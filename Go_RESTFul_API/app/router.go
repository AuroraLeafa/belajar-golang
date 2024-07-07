package app

import (
	"Go_RESTFul_API/controller"
	"Go_RESTFul_API/exception"
	"github.com/julienschmidt/httprouter"
)

func NewRouter(controller controller.CategoryController) *httprouter.Router {
	router := httprouter.New()
	//router.GET("/api/categories", categoryController.FindAll)
	router.GET("/api/categories", controller.FindAll)
	router.GET("/api/categories/:categoryId", controller.FindById)
	router.POST("/api/categories", controller.Create)
	router.PUT("/api/categories/:categoryId", controller.Update)
	router.DELETE("/api/categories/:categoryId", controller.Delete)

	router.PanicHandler = exception.ErrorHandler
	return router
}
