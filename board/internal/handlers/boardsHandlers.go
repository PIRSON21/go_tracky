package handlers

import (
	"github.com/ChanKachan/go_tracky/config"
	"github.com/ChanKachan/go_tracky/internal/database"
	"github.com/ChanKachan/go_tracky/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

func GetBoardInfoById(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		config.Logger.Error("Failed to parse id", zap.String("id", id.String()))
		return
	}

	var getBoardHandler database.GetBoardRequests = &database.BoardServise{Logger: config.Logger, Db: config.ConnectDatabase()}

	board, err := getBoardHandler.GetBoardInfo(id)
	if err != nil {
		c.JSON(404, models.ErrorResponse{StatusCode: 404, Code: "BOARD_NOT_FOUND", Message: err.Error()})
		return
	}

	c.JSON(200, board)
}
