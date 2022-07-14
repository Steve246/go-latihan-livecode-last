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

type TransTypeController struct {
	router *gin.Engine

	ucCrudTransType usecase.CrudTransTypeUseCase
}

func (m *TransTypeController) updateTransType (c *gin.Context){
	var trans *model.Trans_Type

	//pake patch
	//find by id --> update pake id

	id := c.Param("id")

	// var updateMenu map[string]interface{}

	if err := c.BindJSON(&trans); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	} else {
		err := m.ucCrudTransType.UpdateTransType(trans, id)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": "Record not found!",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status":  "Update is Success",
			"message": trans,
		})

	}


}

func(m *TransTypeController) deleteTransType (c *gin.Context){

	id := c.Param("id")


	err := m.ucCrudTransType.DeleteTransType(id)
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



func(m *TransTypeController) createNewTransType(c *gin.Context){
	var newTransType *model.Trans_Type

	if err := c.BindJSON(&newTransType); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	} else {
		err := m.ucCrudTransType.CreateTransType(newTransType)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"status":  "FAILED",
				"message": "Error when creating trans",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": newTransType,
		})

	}

}


func NewTransTypeController(router *gin.Engine,ucCrudTransType usecase.CrudTransTypeUseCase) *TransTypeController {

	controller := TransTypeController{
		router: router,
		ucCrudTransType: ucCrudTransType,
	}

	//tanpa jwt

	router.POST("/transType", controller.createNewTransType)

	router.DELETE("/transType/:id", controller.deleteTransType)

	router.PUT("/transType/:id", controller.updateTransType)

	//nambain JWT

	cfg := config.NewConfigJWT()

	tokenService := utils.NewTokenService(cfg.TokenConfig)

	routerGroup := router.Group("/api")

	routerGroup.POST("/auth/loginTransType", func(c *gin.Context){

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
	

	protectedGroup.POST("/transType", controller.createNewTransType)

	protectedGroup.DELETE("/transType/:id", controller.deleteTransType)

	protectedGroup.PUT("/transType/:id", controller.updateTransType)


	return &controller
}