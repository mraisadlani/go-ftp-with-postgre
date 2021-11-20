package main

import (
	"Training/go-ftp-postgre/config"
	"Training/go-ftp-postgre/docs"
	"Training/go-ftp-postgre/domain"
	"Training/go-ftp-postgre/handler"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func init() {
	config.SetupConfiguration()
}

// @termsOfService http://swagger.io/terms/
// @contact.name Muhammad Rais Adlani
// @contact.url https://gitlab.com/mraisadlani
// @contact.email mraisadlani@gmail.com

// @license.name MIT
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	// programmatically set swagger info
	docs.SwaggerInfo.Title = "Go FTP"
	docs.SwaggerInfo.Description = "Golang upload and read from FTP"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = fmt.Sprintf("%s:%v", config.C.App.Host, config.C.App.Port)
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	logs := config.SetupLog()

	db, err := config.InitDB()
	defer db.Close()

	if err != nil {
		logs.Errorf("Database error : %s", err.Error())
		return
	}

	ping := db.Ping()

	if ping != nil {
		logs.Errorf("Request timeout : %s", err.Error())
		return
	} else {
		logs.Info("Connected Database")
	}

	host := config.C.App.Host
	if host != "localhost" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	BuildHandler(db, logs)
}

func BuildHandler(db *sql.DB, logs *logrus.Logger) {
	productRepo := domain.BuildProductRepository(db)

	r := gin.New()

	r.Use(gin.Recovery())
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := r.Group("/api")

	productHandler := handler.BuildProductHandler(productRepo)

	api.POST("/UploadFile", productHandler.UploadFile)
	api.POST("/ReadFile", productHandler.ReadFile)
	api.POST("/MoveFile", productHandler.MoveFile)
	api.POST("/RenameFile", productHandler.RenameFile)
	api.POST("/DeleteFile", productHandler.DeleteFile)

	logs.Info("Listening on port ", config.C.App.Port)
	err := r.Run(fmt.Sprintf(":%v", config.C.App.Port))

	if err != nil {
		logs.Fatalf("Error listening port server : %v", err)
		return
	}
}