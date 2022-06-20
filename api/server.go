package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	db "github.com/sRRRs-7/MyPage/db/sqlc"
	"github.com/sRRRs-7/MyPage/token"
	"github.com/sRRRs-7/MyPage/utils"
)

// Server services HTTP routers
type Server struct {
	config     utils.Config
	store      db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

func NewServer(config utils.Config, store db.Store) (*Server, error) {
	// PASETO -> JWT : NewPasetoMaker -> NewJWTMaker
	tokenMaker, err := token.NewPasetoMaker(config.TOKEN_SYMMETRIC_KEY)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	server.setupRouter()

	return server, nil
}


func (server *Server) setupRouter() {
	router := gin.Default()

	blogRouter := router.Group("/blog")//.Use(authMiddleware(server.tokenMaker))
	blogRouter.POST("/create", server.createBlog)
	blogRouter.GET("/list", server.getBlogList)
	blogRouter.GET("/get/:id", server.getBlog)
	blogRouter.PUT("/update/:id", server.updateBlog)
	blogRouter.DELETE("/delete/:id", server.deleteBlog)

	qRouter := router.Group("/q")//.Use(authMiddleware(server.tokenMaker))
	qRouter.POST("/create", server.createQ)
	qRouter.GET("/list", server.getQList)
	qRouter.GET("/get/:id", server.getQ)
	qRouter.PUT("/update/:id", server.updateQ)
	qRouter.DELETE("/delete/:id", server.deleteQ)

	aRouter := router.Group("/a")//.Use(authMiddleware(server.tokenMaker))
	aRouter.POST("/create", server.createA)
	aRouter.GET("/list", server.getAList)
	aRouter.GET("/get/:id", server.getA)
	aRouter.PUT("/update/:id", server.updateA)
	aRouter.DELETE("/delete/:id", server.deleteA)

	server.router = router
}

// Start runs the HTTP server on a specific address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}