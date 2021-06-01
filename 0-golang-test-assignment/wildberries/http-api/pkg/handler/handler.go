package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/p-12s/own-golang-manual/0-golang-test-assignment/wildberries/http-api/pkg/service"
	//"github.com/p-12s/own-golang-manual/0-golang-test-assignment/wildberries/http-api/pkg/repository"
	//"github.com/swaggo/gin-swagger"
	//"github.com/swaggo/gin-swagger/swaggerFiles"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	// router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	/*auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}*/

	api := router.Group("/api", nil)//, h.userIdentity
	{
		/*lists := api.Group("/lists")
		{
			lists.POST("/", h.createList)
			lists.GET("/", h.getAllLists)
			lists.GET("/:id", h.getListById)
			lists.PUT("/:id", h.updateList)
			lists.DELETE("/:id", h.deleteList)

			items := lists.Group(":id/items")
			{
				items.POST("/", h.createItem)
				items.GET("/", h.getAllItems)
			}
		}

		items := api.Group("items")
		{
			items.GET("/:id", h.getItemById)
			items.PUT("/:id", h.updateItem)
			items.DELETE("/:id", h.deleteItem)
		}*/
	}
	_ = api
	return router
}