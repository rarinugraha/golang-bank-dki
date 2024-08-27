package main

import (
	"go-bank-dki/config"
	"go-bank-dki/database"
	"go-bank-dki/handlers"
	"go-bank-dki/logger"
	"go-bank-dki/middleware"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

func main() {
	config.LoadEnv()

	db, err := gorm.Open("postgres", config.GetDBConnectionString())
	if err != nil {
		panic("failed to connect to database")
	}
	defer db.Close()

	database.Migrate(db)
	database.SeedUsers(db)

	logger.InitLogger()

	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	r.Use(middleware.RequestLogger())

	r.POST("/login", handlers.Login)

	protected := r.Group("/", middleware.AuthMiddleware())
	{
		protected.POST("/stocks", handlers.CreateStock)
		protected.GET("/stocks", handlers.ListStock)
		protected.GET("/stocks/:id", handlers.DetailStock)
		protected.PUT("/stocks/:id", handlers.UpdateStock)
		protected.DELETE("/stocks/:id", handlers.DeleteStock)
	}

	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to run server: ", err)
	}
}
