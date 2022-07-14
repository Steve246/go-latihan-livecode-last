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

type DiscountController struct {
	router *gin.Engine

	ucCrudDiscount usecase.CrudDiscountUseCase
}

func (m *DiscountController) updateDiscount (c *gin.Context){
	var disc *model.Discount

	//pake patch
	//find by id --> update pake id

	id := c.Param("id")

	// var updateMenu map[string]interface{}

	if err := c.BindJSON(&disc); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	} else {
		err := m.ucCrudDiscount.UpdateDiscount(disc, id)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": "Record not found!",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status":  "Update is Success",
			"message": disc,
		})

	}


}

func(m *DiscountController) deleteDiscount (c *gin.Context){

	id := c.Param("id")


	err := m.ucCrudDiscount.DeleteDiscount(id)
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



func(m *DiscountController) createNewDiscount(c *gin.Context){
	var newDiscount *model.Discount

	if err := c.BindJSON(&newDiscount); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	} else {
		err := m.ucCrudDiscount.CreateDiscount(newDiscount)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"status":  "FAILED",
				"message": "Error when creating trans",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": newDiscount,
		})

	}

}


func NewDiscountController(router *gin.Engine,ucCrudDiscount usecase.CrudDiscountUseCase)*DiscountController {

	controller := DiscountController{
		router: router,
		ucCrudDiscount: ucCrudDiscount,
	}

	//tanpa jwt

	router.POST("/discount", controller.createNewDiscount)

	router.DELETE("/discount/:id", controller.deleteDiscount)

	router.PUT("/discount/:id", controller.updateDiscount)

	//nambain JWT

	cfg := config.NewConfigJWT()

	tokenService := utils.NewTokenService(cfg.TokenConfig)

	routerGroup := router.Group("/api")

	routerGroup.POST("/auth/loginDiscount", func(c *gin.Context){

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
	

	protectedGroup.POST("/discount", controller.createNewDiscount)

	protectedGroup.DELETE("/discount/:id", controller.deleteDiscount)

	protectedGroup.PUT("/discount/:id", controller.updateDiscount)


	return &controller
}