package repository

import (
	"context"

	"github.com/ericoliveiras/alert-bot-go/models"
	"github.com/jmoiron/sqlx"
)

type IStreamRepository interface {
	Create(ctx context.Context, stream *models.Stream) (*models.Stream, error)
	GetByStreamName(ctx context.Context, name string) (*models.Stream, error)
	GetByStreamID(ctx context.Context, id int) (*models.Stream, error)
	GetAllStreams(ctx context.Context) ([]models.Stream, error)
	UpdateStreamIsLive(ctx context.Context, id int, isLive bool) error
}

type StreamRepository struct {
	DB *sqlx.DB
}

func NewStreamRepository(db *sqlx.DB) *StreamRepository {
	return &StreamRepository{DB: db}
}

func (sr *StreamRepository) Create(ctx context.Context, stream *models.Stream) (*models.Stream, error) {
	query := `
		INSERT INTO streams 
			(name, image_url, is_live, created_at, updated_at) 
		VALUES 
			($1, $2, $3, $4, $5)
		RETURNING id
	`

	var id int
	err := sr.DB.QueryRowContext(
		ctx, query,
		stream.Name, stream.ImageUrl, stream.IsLive, stream.CreatedAt, stream.UpdatedAt,
	).Scan(&id)

	if err != nil {
		return nil, err
	}

	stream.ID = id

	return stream, nil
}

func (sr *StreamRepository) GetByStreamName(ctx context.Context, name string) (*models.Stream, error) {
	var stream models.Stream

	query := "SELECT * FROM streams WHERE name = $1"
	err := sr.DB.GetContext(ctx, &stream, query, name)
	if err != nil {
		return &models.Stream{}, err
	}

	return &stream, nil
}

func (sr *StreamRepository) GetByStreamID(ctx context.Context, id int) (*models.Stream, error) {
	var stream models.Stream

	query := "SELECT * FROM streams WHERE id = $1"
	err := sr.DB.GetContext(ctx, &stream, query, id)
	if err != nil {
		return &models.Stream{}, err
	}

	return &stream, nil
}

func (sr *StreamRepository) GetAllStreams(ctx context.Context) ([]models.Stream, error) {
	var streams []models.Stream

	query := "SELECT id, name, image_url, is_live, created_at, updated_at FROM streams"
	err := sr.DB.SelectContext(ctx, &streams, query)
	if err != nil {
		return nil, err
	}

	return streams, nil
}

func (sr *StreamRepository) UpdateStreamIsLive(ctx context.Context, id int, isLive bool) error {
	query := "UPDATE streams SET is_live = $1 WHERE id = $2"

	_, err := sr.DB.ExecContext(ctx, query, isLive, id)
	if err != nil {
		return err
	}

	return nil
}
