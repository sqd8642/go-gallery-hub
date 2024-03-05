package model

import (
	"context"
	"database/sql"
	"log"
	"time"
)

type Image struct {
	Id             string `json:"id"`
	CreatedAt      string `json:"createdAt"`
	UpdatedAt      string `json:"updatedAt"`
	Url            string `json:"url"`
	Caption        string `json:"caption"`
}

type ImageModel struct {
    DB        *sql.DB
	InfoLog   *log.Logger
	ErrorLog  *log.Logger
}

func (m ImageModel) Insert (image *Image) error {

	query := ''
	args := []interface{}{image.CreatedAt, image.UpdatedAt, image.Url, image.Caption}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return m.DB.QueryRowContext(ctx, query, args ...).Scan(&image.id, &image.CreatedAt, &image.UpdatedAt)
}

func (m MenuModel) Get(id int) (*Menu, error) {
	// Retrieve a specific menu item based on its ID.
	query := `
		SELECT id, created_at, updated_at, caption, url
		FROM image
		WHERE id = $1
		`
	var image Image
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	row := m.DB.QueryRowContext(ctx, query, id)
	err := row.Scan(&image.Id, &image.CreatedAt, &image.UpdatedAt, &image.Caption, &image.Url)
	if err != nil {
		return nil, err
	}
	return &image, nil
}

func (m ImageModel) Update(image *Image) error {
	// Update a specific menu item in the database.
	query := `
		UPDATE image
		SET caption = $1
		WHERE id = $4
		RETURNING updated_at
		`
	args := []interface{}{image.Caption, image.Id}
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	return m.DB.QueryRowContext(ctx, query, args...).Scan(&Image.UpdatedAt)
}

func (m ImageModel) Delete(id int) error {
	// Delete a specific menu item from the database.
	query := `
		DELETE FROM image
		WHERE id = $1
		`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := m.DB.ExecContext(ctx, query, id)
	return err
}