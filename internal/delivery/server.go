package delivery

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"shorthack_backend/internal/delivery/handlers"
	"shorthack_backend/internal/repository/student"
	"shorthack_backend/internal/service"

	"shorthack_backend/internal/delivery/docs"

	"shorthack_backend/pkg/log"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Start(db *sqlx.DB, logger *log.Log) {
	r := gin.Default()

	docs.SwaggerInfo.BasePath = "/"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// todo: имплементировать cors middleware
	//middlewareStruct := middleware.InitMiddleware(logger, jwtUtils, session)

	//r.Use(middlewareStruct.CORSMiddleware())

	//routers.InitRouting(r, db, logger, middlewareStruct, jwtUtils, session, tracer)

	publicRouter := r.Group("/student")

	userRepo := student.NewStudentRepo(db)

	publicService := service.InitStudentService(userRepo)
	publicHandler := handlers.InitPublicHandler(publicService)

	publicRouter.POST("/create", publicHandler.CreateUser)
	//publicRouter.POST("/login", publicHandler.LoginUser)
	//
	//publicRouter.POST("/refresh", publicHandler.Refresh)

	if err := r.Run("0.0.0.0:8080"); err != nil {
		panic(fmt.Sprintf("error running client: %v", err.Error()))
	}
}
