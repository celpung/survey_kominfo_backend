package main

import (
	"fmt"
	"log"
	"os"
	"time"

	mysql_configs "github.com/celpung/gocleanarch/configs/database/mysql"
	"github.com/celpung/gocleanarch/configs/environment"
	survey_answer_router "github.com/celpung/gocleanarch/domain/answer/delivery/route"
	survey_category_router "github.com/celpung/gocleanarch/domain/category/delivery/router"
	survey_question_router "github.com/celpung/gocleanarch/domain/question/delivery/route"
	survey_router "github.com/celpung/gocleanarch/domain/survey/delivery/router"
	uploader_router "github.com/celpung/gocleanarch/domain/uploader/route"
	user_router "github.com/celpung/gocleanarch/domain/user/delivery/gin/router"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// load .env
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables from system")
	}

	// Connect to the database and auto migrate
	mysql_configs.CreateDatabaseIfNotExists()
	mysql_configs.ConnectDatabase()
	mysql_configs.AutoMigrage()

	//setup gin
	r := gin.Default()

	// setup mode
	mode := os.Getenv("MODE")

	if mode == "debug" {
		gin.SetMode(gin.DebugMode)
	} else if mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		fmt.Println("-------------------------------------------------")
		fmt.Println("Please set the mode debug/release on environment!")
		fmt.Println("Example : [MODE: debug] or [MODE: release]")
		fmt.Println("-------------------------------------------------")
	}

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:5173",
			"http://localhost:3000",
			"http://localhost",
			"http://103.116.168.244",
			"http://103.116.168.244/survey",
			"http://surveyflex.medan.go.id",
			"https://surveyflex.medan.go.id",
		},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// setup router
	api := r.Group("/api")

	user_router.Router(api)
	survey_router.Router(api)
	survey_category_router.Router(api)
	survey_question_router.Router(api)
	survey_answer_router.Router(api)
	uploader_router.Router(api)

	// Serve static files
	r.GET("/", func(c *gin.Context) {
		c.File("./public")
	})

	r.Static("/images", "./public/images")
	r.Static("/files", "./public/files")

	// Start the server
	r.Run(fmt.Sprintf(":%s", environment.Env.PORT))
}
