package router

import (
	controller "go-fund-science/controller/api/v1"

	"github.com/gin-gonic/gin"
)

func NewRouter(router *gin.Engine) {
	// router.POST("/user/register", controller.UserRegister)
	// router.POST("/user/login", controller.UserLogin)
	// router.Use(middleware.Authorize())
	// router.GET("/user/list", controller.UserList)
	// router.POST("/user/logout", controller.UserLogout)
	// formRouter(router)
	// menuRouter(router)
	// buildRouter(router)
	// ruleRouter(router)
	// projectRouter(router)
	v1 := router.Group("/api/v1")
	{
		persons := v1.Group("/persons")
		{
			personRouter(persons)
		}
	}
}

func personRouter(router *gin.RouterGroup) {
	router.GET("/GetPersonByID", controller.GetPersonByID)
	router.GET("/GetPersonByPersonName", controller.GetPersonByPersonName)
	router.GET("/GetPersonByCertID", controller.GetPersonByCertID)
	// router.POST("/Insert", controller.PersonInsert)
	// router.PUT("/Update", controller.PersonUpdate)
	// router.DELETE("/Delete", controller.PersonDelete)
}

// func projectRouter(router *gin.Engine) {
// 	router.GET("/projects", controller.ProjectList)
// 	router.GET("/project/icon", controller.ProjectIcon)
// 	router.POST("/project", controller.ProjectCreate)
// 	router.PUT("/project", controller.ProjectEdit)
// 	router.DELETE("/project", controller.ProjectDelete)
// 	router.PUT("/project/upgrade", controller.ProjectUpgrade)
// }
