package internal

type UpsertPostRequest struct {
	Title  string `form:"title"  binding:"required"`
	Detail string `form:"detail"  binding:"required"`
}

type UpsertCommentRequest struct {
	Message string `form:"message"  binding:"required"`
}
