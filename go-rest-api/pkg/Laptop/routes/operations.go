package routes

import (
	"be-training/go-rest-api/bootstrap"
	"be-training/go-rest-api/pkg/Laptop/handler"
	"time"

	"github.com/gin-gonic/gin"
)

func CRUDoperations(env *bootstrap.Env, timeout time.Duration, group *gin.RouterGroup) {


	//Create List godoc

	// @Summary      Create List
	// @Description  Create List
	// @Tags         Create List
	// @Accept       json
	// @Produce      json
	// @Param        request body LaptopStruct true "Create List"
	// @Success      200  {object}  string
	// @Failure      400  {object}  httputil.HTTPError
	// @Failure      404  {object}  httputil.HTTPError
	// @Failure      500  {object}  httputil.HTTPError
	// @Router       /create [post]
	// @Security     ApiKeyAuth

	group.POST("/",handler.CreateList)



	//Delete by id godoc
	// @Summary      Delete by id
	// @Description  Delete by id
	// @Tags         Delete by id
	// @Accept       json
	// @Produce      json
	// @Param        request body IdResponse true "Delete by id"
	// @Success      200  {object}  string
	// @Failure      400  {object}  httputil.HTTPError
	// @Failure      404  {object}  httputil.HTTPError
	// @Failure      500  {object}  httputil.HTTPError
	// @Router       /delete [delete]
	// @Security     ApiKeyAuth
	
	group.GET("/",handler.GetList)



	//Get list godoc
	// @Summary      Get list
	// @Description  Get list
	// @Tags         Get list
	// @Accept       json
	// @Produce      json
	// @Param        request body IdResponse true "Get list"
	// @Success      200  {object}  string
	// @Failure      400  {object}  httputil.HTTPError
	// @Failure      404  {object}  httputil.HTTPError
	// @Failure      500  {object}  httputil.HTTPError

	group.DELETE("/delete",handler.DeleteById)


	//Update list godoc
	// @Summary      Update list
	// @Description  Update list
	// @Tags         Update list
	// @Accept       json
	// @Produce      json
	// @Param        request body LaptopStruct true "Update list"
	// @Success      200  {object}  string
	// @Failure      400  {object}  httputil.HTTPError
	// @Failure      404  {object}  httputil.HTTPError
	// @Failure      500  {object}  httputil.HTTPError
	// @Router       /update [put]
	// @Security     ApiKeyAuth
	group.PUT("/update")


	//Heartbeat godoc
	// @Summary      Heartbeat
	// @Description 
	group.GET("/heartbeat",handler.Heartbeat)


}
