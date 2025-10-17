package database

import (
	"context"
	"github.com/ChanKachan/go_tracky/internal/models"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type GetBoardRequests interface {
	GetBoardInfo(id uuid.UUID) (models.Board, error)
}

type PostBoardRequests interface {
	CreateBoard(id string)
}

type PutBoardRequests interface {
	UpdateBoardInfo(id string)
}

type DeleteBoardRequests interface {
	DeleteBoard(id string)
}

type BoardServise struct {
	Logger *zap.Logger
	Db     *pgxpool.Pool
}

func (b BoardServise) GetBoardInfo(id uuid.UUID) (models.Board, error) {
	b.Logger.Info("Отправляем запрос на информацию о доске", zap.String("id", id.String()))

	var user_id uuid.UUID
	var name, access, color string

	err := b.Db.QueryRow(context.Background(),
		"SELECT user_id, name, access, color FROM boards WHERE id = $1",
		id,
	).Scan(&user_id, &name, &access, &color)

	if err != nil {
		b.Logger.Error("Board not found", zap.String("id", id.String()))
		return models.Board{}, err
	}

	return models.Board{id, user_id, name, access, color}, nil
}
