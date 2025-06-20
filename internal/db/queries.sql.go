// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0
// source: queries.sql

package db

import (
	"context"
)

const createCategory = `-- name: CreateCategory :exec
INSERT INTO categories (id, name) VALUES(?, ?)
`

type CreateCategoryParams struct {
	ID   int32
	Name string
}

func (q *Queries) CreateCategory(ctx context.Context, arg CreateCategoryParams) error {
	_, err := q.db.ExecContext(ctx, createCategory, arg.ID, arg.Name)
	return err
}

const createCourse = `-- name: CreateCourse :exec
INSERT INTO courses (id, name, category_id) VALUES (?, ?, ?)
`

type CreateCourseParams struct {
	ID         int32
	Name       string
	CategoryID int32
}

func (q *Queries) CreateCourse(ctx context.Context, arg CreateCourseParams) error {
	_, err := q.db.ExecContext(ctx, createCourse, arg.ID, arg.Name, arg.CategoryID)
	return err
}
