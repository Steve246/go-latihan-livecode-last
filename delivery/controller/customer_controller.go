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

type CustomerController struct {
	router *gin.Engine

	ucCrudCustomer usecase.CrudCustomerUseCase
}

func (m *CustomerController) updateCustomer (c *gin.Context){
	var cust *model.Customer

	//pake patch
	//find by id --> update pake id

	id := c.Param("id")

	// var updateMenu map[string]interface{}

	if err := c.BindJSON(&cust); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	} else {
		err := m.ucCrudCustomer.UpdateCustomer(cust, id)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": "Record not found!",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status":  "Update is Success",
			"message": cust,
		})

	}


}

func(m *CustomerController) deleteCustomer (c *gin.Context){

	id := c.Param("id")


	err := m.ucCrudCustomer.DeleteCustomer(id)
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



func(m *CustomerController) createNewCustomer(c *gin.Context){
	var newCustomer *model.Customer

	if err := c.BindJSON(&newCustomer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	} else {
		err := m.ucCrudCustomer.CreateCustomer(newCustomer)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"status":  "FAILED",
				"message": "Error when creating trans",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": newCustomer,
		})

	}

}



func NewCustomerController(router *gin.Engine,ucCrudCustomer usecase.CrudCustomerUseCase)*CustomerController {

	controller := CustomerController{
		router: router,
		ucCrudCustomer: ucCrudCustomer,
	}

	//tanpa jwt

	router.POST("/customer", controller.createNewCustomer)

	router.DELETE("/customer/:id", controller.deleteCustomer)

	router.PUT("/customer/:id", controller.updateCustomer)

	router.PUT("/customerRegister/:id", controller.updateCustomer)

	//nambain JWT

	cfg := config.NewConfigJWT()

	tokenService := utils.NewTokenService(cfg.TokenConfig)

	routerGroup := router.Group("/api")

	routerGroup.POST("/auth/loginCustomer", func(c *gin.Context){

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
	

	protectedGroup.POST("/customer", controller.createNewCustomer)

	protectedGroup.DELETE("/customer/:id", controller.deleteCustomer)

	protectedGroup.PUT("/customer/:id", controller.updateCustomer)

	protectedGroup.PUT("/customerRegister/:id", controller.updateCustomer)


	return &controller
}