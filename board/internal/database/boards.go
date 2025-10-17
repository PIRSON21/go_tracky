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
	CreateBoard(board *models.Board) error
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

// createId генерирует id с типом uuid
func createId() uuid.UUID {
	return uuid.New()
}

// GetBoardInfo получаете информацию по id доски.

// Если доска не найдена, то вернется структура errorResponse, с номером ошибки, сообщением и кодом ошибки
// Если доска найдена, то возращает всю информацию о доске.
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

	b.Logger.Info("Запрос на поиск успешно выполнен", zap.String("id", id.String()))

	return models.Board{id, user_id, name, access, color}, nil
}

// CreateBoard добавляет доску в базу данных.

// Если данные некорректные, то вернется ошибка добавления доски.
// Если по какой-то причине нельзя добавить данные в базу, то вернется ошибка добавления доски.
func (b BoardServise) CreateBoard(board *models.Board) error {
	b.Logger.Info("Отправляем запрос на создание доски")

	board.Id = createId()

	_, err := b.Db.Exec(context.Background(),
		`INSERT INTO boards (id, user_id, name, access, color) VALUES ($1, $2, $3, $4, $5)`,
		board.Id, board.User_id, board.Name_board, board.Access, board.Color,
	)
	if err != nil {
		b.Logger.Error("Error adding board", zap.String("id", board.Id.String()))
		return err
	}

	b.Logger.Info("Добавление успешно выполнено", zap.String("id", board.Id.String()))
	return nil
}
