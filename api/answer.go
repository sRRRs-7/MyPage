package api

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	db "github.com/sRRRs-7/MyPage/db/sqlc"
)

type ResponseA struct {
	ID        int64         `json:"id"`
	AnswerID  int32  `json:"answer_id"`
	Text      string        `json:"text"`
	CreatedAt time.Time     `json:"created_at"`
}

// createQ response
func NewResponseA(answer db.Answer) ResponseA {
	return ResponseA {
		ID: answer.ID,
		AnswerID: answer.AnswerID,
		Text: answer.Text,
		CreatedAt: answer.CreatedAt,
	}
}

type CreateARequest struct {
	Text string `json:"text" binding:"required,max=1000"`
	AnswerID int32 `json:"answer_id" binding:"required"`
}

// go channel

// func (server *Server) GetQuestion(ctx *gin.Context, req CreateARequest) <-chan db.Question {
// 	r := make(chan db.Question)
// 	go func() {
// 		defer close(r)
// 		q, err := server.store.GetQuestion(ctx, int64(req.AnswerID))
// 		if err != nil {
// 			ctx.JSON(http.StatusInternalServerError, errorResponse(err))
// 			return
// 		}
// 		r <- q
// 	} ()
// 	return r
// }

func (server *Server) createA (ctx *gin.Context) {
	var req CreateARequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	//q := <-server.GetQuestion(ctx, req)

	q, err := server.store.GetQuestion(ctx, int64(req.AnswerID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	arg := db.CreateAnswerParams {
		Text: req.Text,
		AnswerID: int32(q.AnswerID),
	}

	a, err := server.store.CreateAnswer(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, a)
}

type GetAListRequest struct {
	PageID int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) getAList (ctx *gin.Context) {
	var req GetAListRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListAnswerParams {
		Limit: req.PageSize,
		Offset: (req.PageID-1) * req.PageSize,
	}

	aList, err := server.store.ListAnswer(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, aList)
}

type GetARequest struct {
	ID int64 `uri:"id" binding:"required"`
}

func (server *Server) getA (ctx *gin.Context) {
	var req GetARequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	a, err := server.store.GetAnswer(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, a)
}

type UpdateARequestID struct {
	ID   int64  `uri:"id" binding:"required"`
}

type UpdateARequest struct {
	Text string `json:"text" binding:"required,max=1000"`
}

func (server *Server) updateA (ctx *gin.Context) {
	var reqID UpdateARequestID
	var req UpdateARequest
	if err := ctx.ShouldBindUri(&reqID); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdateAnswerParams {
		ID: reqID.ID,
		Text: req.Text,
	}

	a, err := server.store.UpdateAnswer(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, a)
}

type DeleteARequest struct {
	ID int64 `uri:"id" binding:"required"`
}

func (server *Server) deleteA (ctx *gin.Context) {
	var req DeleteARequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err := server.store.DeleteAnswer(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "executed delete"})
}