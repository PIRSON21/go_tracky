package handlers

import (
	"github.com/ChanKachan/go_tracky/internal/database"
	"github.com/ChanKachan/go_tracky/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

func GetBoardInfoById(logger *zap.Logger, dbpool *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := uuid.Parse(c.Param("id"))
		if err != nil {
			logger.Error("Failed to parse id", zap.String("id", id.String()))
			return
		}

		var getBoardHandler database.GetBoardRequests = &database.BoardServise{Logger: logger, Db: dbpool}

		board, err := getBoardHandler.GetBoardInfo(id)
		if err != nil {
			c.JSON(404, models.ErrorResponse{StatusCode: 404, Code: "BOARD_NOT_FOUND", Message: err.Error()})
			return
		}

		c.JSON(200, board)
	}
}

func PostBoard(logger *zap.Logger, dbpool *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		var PostBoardHandler database.PostBoardRequests = &database.BoardServise{Logger: logger, Db: dbpool}
		board := models.Board{}

		err := c.ShouldBindJSON(&board)
		if err != nil {
			c.JSON(400, models.ErrorResponse{StatusCode: 400, Code: "BOARD_NOT_CORRECT", Message: err.Error()})
			return
		}

		err = PostBoardHandler.CreateBoard(&board)
		if err != nil {
			c.JSON(500, models.ErrorResponse{StatusCode: 500, Code: "BOARD_CREATE_FAILED", Message: err.Error()})
			return
		}

		c.JSON(201, board)
	}
}
