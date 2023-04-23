package route

import (
	"challenge-8/handler"

	"github.com/gin-gonic/gin"
)

func RegisterApi(r *gin.Engine, server handler.HttpServer) {
	api := r.Group("/books") // prefix
	{
		api.POST("", server.CreateBook)       // /Books
		api.GET("/", server.GetAllBook)       // /Books
		api.GET("/:id", server.GetBookByID)   // /Books/:id
		api.PUT("/:id", server.UpdateBook)    // /Books/:id
		api.DELETE("/:id", server.DeleteBook) // /Books/:id
	}
}
