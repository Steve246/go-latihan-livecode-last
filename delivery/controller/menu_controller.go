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

type MenuController struct {
	router *gin.Engine

	ucCrudMenu usecase.CrudMenuUseCase

}

func (m *MenuController) updateMenu (c *gin.Context){
	var menu *model.Menu

	//pake patch
	//find by id --> update pake id

	id := c.Param("id")

	// var updateMenu map[string]interface{}

	if err := c.BindJSON(&menu); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	} else {
		err := m.ucCrudMenu.UpdateMenu(menu, id)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": "Record not found!",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status":  "Update is Success",
			"message": menu,
		})

	}


}

func(m *MenuController) deleteMenu (c *gin.Context){

	id := c.Param("id")


	err := m.ucCrudMenu.DeleteMenu(id)
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



func(m *MenuController) createNewMenu(c *gin.Context){
	var newMenu *model.Menu

	if err := c.BindJSON(&newMenu); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	} else {
		err := m.ucCrudMenu.CreateMenu(newMenu)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"status":  "FAILED",
				"message": "Error when creating menu",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": newMenu,
		})

	}

}


func NewMenuController(router *gin.Engine, ucCrudMenu usecase.CrudMenuUseCase) *MenuController {

	controller := MenuController{
		router: router,
		ucCrudMenu: ucCrudMenu,
	}

	//nambain JWT

	cfg := config.NewConfigJWT()

	tokenService := utils.NewTokenService(cfg.TokenConfig)

	routerGroup := router.Group("/api")

	protectedGroup := routerGroup.Group("/master", middleware.NewTokenValidator(tokenService).RequireToken())
	

	protectedGroup.POST("/menu", controller.createNewMenu)

	protectedGroup.DELETE("/menu/:id", controller.deleteMenu)

	protectedGroup.PUT("/menu/:id", controller.updateMenu)


	return &controller
}