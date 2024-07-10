package app

import (
	"azizdev/controller"
	"azizdev/exception"
	"github.com/julienschmidt/httprouter"
)

func NewRouter(userController controller.UserController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/user", userController.FindAll)
	router.GET("/api/user/:id", userController.FindById)
	router.POST("/api/user", userController.Create)
	router.PUT("/api/user/:id", userController.Update)
	router.DELETE("/api/user/:id", userController.Delete)

	router.PanicHandler = exception.ErrorHandler

	return router
}
