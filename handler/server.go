package handler

import (
	"backend/json-server/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ServerHandler struct {
	serverService *service.ServerService
}

func NewServiceHandler(serverService *service.ServerService) *ServerHandler {
	return &ServerHandler{serverService}
}
func (s *ServerHandler) InsertJSON(c *gin.Context) {
	var body map[string]any
	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	id, err := s.serverService.InsertNewJSON(c, body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success add json:" + id})
}

func (s *ServerHandler) FindByID(c *gin.Context) {
	id := c.Param("id")
	getJSON, err := s.serverService.GetJSON(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, getJSON.Raw)
}
