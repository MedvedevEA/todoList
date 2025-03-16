package store

import (
	"context"
	"embed"
	"log"
	"time"
	"todolist/internal/model"
	"todolist/internal/repository/dto"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
)

type Store struct {
	conn *pgx.Conn
}

func New(databaseConnectString string, embedMigrations embed.FS) (*Store, error) {
	//connect
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var (
		conn  *pgx.Conn
		err   error
		count int
	)
	for {
		count++
		conn, err = pgx.Connect(ctx, databaseConnectString)
		if err != nil {
			log.Printf("Database connect error(%d): %s", count, err.Error())

			if count > 4 {
				return nil, err
			}

			time.Sleep(1 * time.Second)
			continue
		}
		break
	}
	log.Println("Successful connection to the database")

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
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	return s.conn.Close(ctx)
}

func (s *Store) AddTask(req *dto.AddTask) (*model.Task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	query := `INSERT INTO tasks VALUES(default,$1,$2,$3,default,default) RETURNING *`
	task := new(model.Task)
	err := s.conn.QueryRow(ctx, query, req.Title, req.Description, req.Status).Scan(&task.Id, &task.Title, &task.Description, &task.Status, &task.CreatedAt, &task.UpdatedAt)
	return task, err
}
func (s *Store) GetTasks(req *dto.GetTasks) ([]*model.Task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	query := `SELECT * FROM tasks`

	rows, err := s.conn.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tasks := []*model.Task{}
	for rows.Next() {
		task := new(model.Task)
		err := rows.Scan(&task.Id, &task.Title, &task.Description, &task.Status, &task.CreatedAt, &task.UpdatedAt)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}
func (s *Store) UpdateTask(req *dto.UpdateTask) (*model.Task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	query := `UPDATE tasks SET title=$2, description=$3, status=$4, updated_at=now() WHERE id=$1 RETURNING *`
	task := new(model.Task)
	err := s.conn.QueryRow(ctx, query, req.Id, req.Title, req.Description, req.Status).Scan(&task.Id, &task.Title, &task.Description, &task.Status, &task.CreatedAt, &task.UpdatedAt)
	return task, err
}
func (s *Store) RemoveTask(req *dto.RemoveTask) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	query := `DELETE FROM tasks	WHERE id=$1 RETURNING id`
	taskId := new(int)
	return s.conn.QueryRow(ctx, query, req.Id).Scan(taskId)

}
