package controller

import (
	"fmt"
	"go_livecode_persiapan/model"
	"go_livecode_persiapan/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MenuPriceController struct {
	router *gin.Engine

	ucCrudMenuPrice usecase.CrudMenuPriceUseCase
}

func (m *MenuPriceController) updateMenuPrice (c *gin.Context){
	var menu *model.Menu_Price

	//pake patch
	//find by id --> update pake id

	id := c.Param("id")

	// var updateMenu map[string]interface{}

	if err := c.BindJSON(&menu); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	} else {
		err := m.ucCrudMenuPrice.UpdateMenu(menu, id)
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

func(m *MenuPriceController) deleteMenuPrice (c *gin.Context){

	id := c.Param("id")


	err := m.ucCrudMenuPrice.DeleteMenu(id)
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

func(m *MenuPriceController) createNewMenuPrice(c *gin.Context){
	var newMenu *model.Menu_Price

	if err := c.BindJSON(&newMenu); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	} else {
		err := m.ucCrudMenuPrice.CreateMenu(newMenu)
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


func NewMenuPriceController(router *gin.Engine, ucCrudMenuPrice usecase.CrudMenuPriceUseCase) *MenuPriceController {

	controller := MenuPriceController{
		router: router,
		ucCrudMenuPrice:ucCrudMenuPrice,
	}

	router.POST("/menuPrice", controller.createNewMenuPrice)

	router.DELETE("/menuPrice/:id", controller.deleteMenuPrice)

	router.PUT("/menuPrice/:id", controller.updateMenuPrice)


	return &controller
}