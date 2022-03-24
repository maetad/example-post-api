package internal

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	repo RepositoryInterface
}

func NewHandler(repo RepositoryInterface) *Handler {
	return &Handler{repo}
}

func (h *Handler) GetPosts(ctx *gin.Context) {
	p := h.repo.GetPosts()
	ctx.JSON(http.StatusOK, p)
}

func (h *Handler) CreaePost(ctx *gin.Context) {
	var req UpsertPostRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
		return
	}

	p, err := h.repo.CreatePost(req)

	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusCreated, p)
}

func (h *Handler) FindPost(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	p, err := h.repo.FindPost(uint(id))

	if err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	ctx.JSON(http.StatusOK, p)
}

func (h *Handler) UpdatePost(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	p, err := h.repo.FindPost(uint(id))

	if err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	var req UpsertPostRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
		return
	}

	err = h.repo.UpdatePost(p, req)

	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, p)
}

func (h *Handler) DeletePost(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	p, err := h.repo.FindPost(uint(id))

	if err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	if err = h.repo.DeletePost(p); err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.AbortWithStatus(http.StatusNoContent)
}

func (h *Handler) GetComments(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	p, err := h.repo.FindPost(uint(id))

	if err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	c := h.repo.GetComments(p)
	ctx.JSON(http.StatusOK, c)
}

func (h *Handler) CreateComment(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	p, err := h.repo.FindPost(uint(id))

	if err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	var req UpsertCommentRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
		return
	}

	c, err := h.repo.CreateComment(p, req)

	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusCreated, c)
}

func (h *Handler) FindComment(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	p, err := h.repo.FindPost(uint(id))

	if err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	cid, err := strconv.ParseUint(ctx.Param("commentId"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	c, err := h.repo.FindComment(p, uint(cid))

	if err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	ctx.JSON(http.StatusOK, c)
}

func (h *Handler) UpdateComment(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	p, err := h.repo.FindPost(uint(id))

	if err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	cid, err := strconv.ParseUint(ctx.Param("commentId"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	c, err := h.repo.FindComment(p, uint(cid))

	if err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	var req UpsertCommentRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
		return
	}

	err = h.repo.UpdateComment(c, req)

	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, c)
}

func (h *Handler) DeleteComment(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	p, err := h.repo.FindPost(uint(id))

	if err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	cid, err := strconv.ParseUint(ctx.Param("commentId"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	c, err := h.repo.FindComment(p, uint(cid))

	if err != nil {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	if err = h.repo.DeleteComment(c); err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.AbortWithStatus(http.StatusNoContent)
}
