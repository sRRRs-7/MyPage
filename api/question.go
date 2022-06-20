package api

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	db "github.com/sRRRs-7/MyPage/db/sqlc"
)

type ResponseQ struct {
	ID        int64         `json:"id"`
	AnswerID  int64 `json:"answer_id"`
	Text      string        `json:"text"`
	CreatedAt time.Time     `json:"created_at"`
}

// createQ response
func NewResponseQ(question db.Question) ResponseQ {
	return ResponseQ {
		ID: question.ID,
		AnswerID: question.AnswerID,
		Text: question.Text,
		CreatedAt: question.CreatedAt,
	}
}

type CreateQRequest struct {
	Text string `json:"text" binding:"required,max=1000"`
}

func (server *Server) createQ (ctx *gin.Context) {
	var req CreateQRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	q, err := server.store.CreateQuestion(ctx, req.Text)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	resp := NewResponseQ(q)
	ctx.JSON(http.StatusOK, resp)
}

type GetQListRequest struct {
	PageID int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) getQList (ctx *gin.Context) {
	var req GetBlogListRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListQuestionParams {
		Limit: req.PageSize,
		Offset: (req.PageID -1) * req.PageSize,
	}

	qList, err := server.store.ListQuestion(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, qList)
}

type GetQRequest struct {
	ID int64 `uri:"id" binding:"required"`
}

func (server *Server) getQ (ctx *gin.Context) {
	var req GetQRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	q, err := server.store.GetQuestion(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, q)
}

type UpdateQRequestID struct {
	ID int64 `uri:"id" binding:"required"`
}

type UpdateQRequest struct {
	Text string `json:"text" binding:"required,max=1000"`
}

func (server *Server) updateQ (ctx *gin.Context) {
	var reqID UpdateQRequestID
	var req UpdateQRequest
	if err := ctx.ShouldBindUri(&reqID); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdateQuestionParams {
		ID: reqID.ID,
		Text: req.Text,
	}

	q, err := server.store.UpdateQuestion(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, q)
}

type DeleteQRequest struct {
	ID int64 `uri:"id" binding:"required"`
}

func (server *Server) deleteQ (ctx *gin.Context) {
	var req DeleteQRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err := server.store.DeleteQuestion(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "executed delete"})
}