package store

import (
	"context"
	"database/sql"
	"embed"
	"errors"
	"todolist/internal/model"
	"todolist/internal/repository/dto"
	"todolist/internal/servererrors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
)

type Store struct {
	conn *pgx.Conn
}

func New(databaseConnectString string, embedMigrations embed.FS) (*Store, error) {
	//connect
	ctx := context.TODO()
	conn, err := pgx.Connect(ctx, databaseConnectString)
	if err != nil {
		return nil, err
	}
	//migration
	goose.SetBaseFS(embedMigrations)
	config := conn.Config()
	db := stdlib.OpenDB(*config)
	if err := goose.SetDialect("postgres"); err != nil {
		return nil, err
	}
	if err := goose.Up(db, "migrations"); err != nil {
		return nil, err
	}

	return &Store{
		conn: conn,
	}, nil
}
func (s *Store) Close() error {
	ctx := context.Background()
	return s.conn.Close(ctx)
}

func (s *Store) AddTask(req *dto.AddTask) (*model.Task, error) {
	ctx := context.TODO()
	query := `INSERT INTO tasks VALUE(default,$1,$2,$3,default,default) RETURNING *`
	task := new(model.Task)
	err := s.conn.QueryRow(ctx, query, req.Title, req.Description, req.Status).Scan(&task.Id, &task.Title, &task.Description, &task.Status, &task.Updated_at, &task.Updated_at)
	return task, err
}
func (s *Store) GetTasks(req *dto.GetTasks) ([]*model.Task, error) {
	ctx := context.TODO()
	query := `SELECT * FROM tasks`

	rows, err := s.conn.Query(ctx, query)
	if err != nil {
		return nil, servererrors.ErrorInternal
	}
	defer rows.Close()

	tasks := []*model.Task{}
	for rows.Next() {
		task := new(model.Task)
		err := rows.Scan(&task.Id, &task.Title, &task.Description, &task.Status, &task.Created_at, &task.Updated_at)
		if err != nil {
			return nil, servererrors.ErrorInternal
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}
func (s *Store) UpdateTask(req *dto.UpdateTask) (*model.Task, error) {
	ctx := context.TODO()
	query := `UPDATE tasks SET title=$2, description=$3, status=$4 WHERE id=$1 RETURNING *`
	task := new(model.Task)
	err := s.conn.QueryRow(ctx, query, req.Id, req.Title, req.Description, req.Status).Scan(&task.Id, &task.Title, &task.Description, &task.Status, &task.Updated_at, &task.Updated_at)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, servererrors.ErrorRecordNotFound
	}
	if err != nil {
		return nil, servererrors.ErrorInternal
	}
	return task, err
}
func (s *Store) RemoveTask(req *dto.RemoveTask) error {
	ctx := context.TODO()
	query := `DELETE FROM tasks	WHERE id=$1 RETURNING id`
	row := s.conn.QueryRow(ctx, query, req.Id)
	var taskId int
	err := row.Scan(&taskId)
	if errors.Is(err, sql.ErrNoRows) {
		return servererrors.ErrorRecordNotFound
	}
	if err != nil {
		return servererrors.ErrorInternal
	}
	return nil
}
