package controller

import (
	"fmt"
	"go_livecode_persiapan/config"
	"go_livecode_persiapan/delivery/middleware"
	"go_livecode_persiapan/model"
	"go_livecode_persiapan/usecase"
	"go_livecode_persiapan/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TableController struct {
	router *gin.Engine

	ucCrudTable usecase.CrudTableUseCase
}

func(m *TableController) updateTable (c *gin.Context) {
	var table *model.Table

	//pake patch
	//find by id --> update pake id

	id := c.Param("id")

	// var updateMenu map[string]interface{}

	if err := c.BindJSON(&table); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	} else {
		err := m.ucCrudTable.UpdateTable(table, id)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": "Record not found!",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status":  "Update is Success",
			"message": table,
		})

	}

}

func(m *TableController) deleteTable(c *gin.Context) {
	id := c.Param("id")


	err := m.ucCrudTable.DeleteTable(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "Record not found!",
		})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status": fmt.Sprintf("Successfully deleted user: %s", id),
			// "message": newMenu,
		})
}

func(m *TableController) createNewTable(c *gin.Context){
	var newTable *model.Table

	if err := c.BindJSON(&newTable); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	} else {
		err := m.ucCrudTable.CreateTable(newTable)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"status":  "FAILED",
				"message": "Error when creating menu",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": newTable,
		})

	}

}

func NewTableController(router *gin.Engine, ucCrudTable usecase.CrudTableUseCase) *TableController {

	controller := TableController{
		router: router,
		ucCrudTable:ucCrudTable ,
	}

	//tanpa jwt

	router.POST("/table", controller.createNewTable)

	router.DELETE("/table/:id", controller.deleteTable)

	router.PUT("/table/:id", controller.updateTable)

	//nambain JWT

	cfg := config.NewConfigJWT()

	tokenService := utils.NewTokenService(cfg.TokenConfig)

	routerGroup := router.Group("/api")

	routerGroup.POST("/auth/login", func(c *gin.Context){

		var user model.Credential


		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H {
				"message": "cant't bind struct",
			})
			return 
		}

		if user.Username == "enigma" && user.Password == "123" {
			token, err := tokenService.CreateAccessToken(&user)

			if err != nil {
				c.AbortWithStatus(http.StatusUnauthorized)
				return 
			}
			c.JSON(200, gin.H {
				"token": token,
			})
		} else {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

	})

	protectedGroup := routerGroup.Group("/master", middleware.NewTokenValidator(tokenService).RequireToken())

	protectedGroup.POST("/table", controller.createNewTable)

	protectedGroup.DELETE("/table/:id", controller.deleteTable)

	protectedGroup.PUT("/table/:id", controller.updateTable)


	return &controller
}