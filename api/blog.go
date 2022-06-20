package api

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	db "github.com/sRRRs-7/MyPage/db/sqlc"
)

type Response struct {
	ID int64 `json:"id" binding:"required"`
	Title string `json:"title" binding:"required,max=100"`
	Text  string `json:"text" binding:"required,max=1000"`
	Image []byte `json:"image"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// createBlog response
func NewBlogResponse(blog db.Blog) Response {
	return Response {
		ID: blog.ID,
		Title: blog.Title,
		Text: blog.Text,
		Image: blog.Image,
		CreatedAt: blog.CreatedAt,
		UpdatedAt: blog.UpdatedAt,
	}
}

type CreateBlogRequest struct {
	Title string `json:"title" binding:"required,max=100"`
	Text  string `json:"text" binding:"required,max=1000"`
	Image []byte `json:"image"`
}

func (server *Server) createBlog (ctx *gin.Context) {
	var req CreateBlogRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateBlogParams{
		Title: req.Title,
		Text: req.Text,
		Image: req.Image,
	}

	blog, err := server.store.CreateBlog(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	resp := NewBlogResponse(blog)
	ctx.JSON(http.StatusOK, resp)
}

type GetBlogListRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) getBlogList (ctx *gin.Context) {
	var req GetBlogListRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListBlogParams {
		Limit: req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	blogs, err := server.store.ListBlog(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, blogs)
}

type GetBlogRequest struct {
	ID int64 `uri:"id" binding:"required"`
}

func (server *Server) getBlog (ctx *gin.Context) {
	var req GetBlogRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	blog, err := server.store.GetBlog(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, blog)
}

type UpdateBlogRequestID struct {
	ID    int64  `uri:"id"`
}

type UpdateBlogRequest struct {
	Title string `json:"title" binding:"required,max=100"`
	Text  string `json:"text" binding:"required,max=1000"`
	Image []byte `json:"image"`
}

func (server *Server) updateBlog (ctx *gin.Context) {
	var reqID UpdateBlogRequestID
	var req UpdateBlogRequest
	if err := ctx.ShouldBindUri(&reqID); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdateBlogParams {
		ID: reqID.ID,
		Title: req.Title,
		Text: req.Text,
		Image: req.Image,
	}

	blog, err := server.store.UpdateBlog(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, blog)
}

type DeleteBlogRequest struct {
	ID int64 `uri:"id" binding:"required"`
}

func (server *Server) deleteBlog (ctx *gin.Context) {
	var req DeleteBlogRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err := server.store.DeleteBlog(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "executed delete"})
}



