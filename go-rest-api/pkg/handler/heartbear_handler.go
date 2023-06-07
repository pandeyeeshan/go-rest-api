package handler

import (
	model "be-training/go-rest-api/pkg/model"
	"net/http"

	"github.com/gin-gonic/gin"
)



func Heartbeat(c *gin.Context) {
	c.JSON(http.StatusOK, model.HeartbeatResponse{Status: "OK",Code:200})
  }