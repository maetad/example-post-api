package internal

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewRoute(r *gin.Engine, db *gorm.DB) error {
	h := NewHandler(NewRepository(db))

	r.GET("/posts", h.GetPosts)
	r.POST("/posts", h.CreaePost)
	r.GET("/posts/:id", h.FindPost)
	r.PUT("/posts/:id", h.UpdatePost)
	r.DELETE("/posts/:id", h.DeletePost)
	r.GET("/posts/:id/comments", h.GetComments)
	r.POST("/posts/:id/comments", h.CreateComment)
	r.GET("/posts/:id/comments/:commentId", h.FindComment)
	r.PUT("/posts/:id/comments/:commentId", h.UpdateComment)
	r.DELETE("/posts/:id/comments/:commentId", h.DeleteComment)

	return nil
}
