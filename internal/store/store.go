package store

import (
	"context"
	"embed"
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
	ctx := context.Background()
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
	return nil, nil
}
func (s *Store) GetTasks(req *dto.GetTasks) ([]*model.Task, error) {
	return nil, nil
}
func (s *Store) UpdateTask(req *dto.UpdateTask) (*model.Task, error) {
	return nil, nil
}
func (s *Store) RemoveTask(req *dto.RemoveTask) error {
	return nil
}
