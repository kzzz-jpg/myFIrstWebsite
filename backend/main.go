package main

import (
	"fmt"
	"memeboard/handlers"
	"memeboard/models"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func getEnv(key, Default string) string {
	ret := os.Getenv(key)
	if ret == "" {
		return Default
	}
	return ret
}
func main() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Taipei",
		getEnv("DB_HOST", "localhost"),
		getEnv("DB_USER", "memeboard_user"),
		getEnv("DB_PASSWORD", "password"),
		getEnv("DB_NAME", "memeboard"),
		getEnv("DB_PORT", "5432"),
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("connect db failed", err)
		return
	}
	err = db.AutoMigrate(&models.Meme{})
	if err != nil {
		fmt.Println("failed to create form", err)
		return
	}
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "PUT", "POST", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		AllowCredentials: true,
	}))
	api := r.Group("/api")
	{
		memeInfo := handlers.NewMemeInfo(db)
		api.GET("/memeinfo", memeInfo.GetMemeInfo)
		api.POST("/memeinfo", memeInfo.UploadNewMeme)
		api.PUT("/memeinfo", memeInfo.UpdateMemeInfo)
		api.DELETE("/memeinfo/:id", memeInfo.DeleteMeme)
	}
	port := getEnv("PORT", "8080")
	if err := r.Run(":" + port); err != nil {
		fmt.Println("failed to run server", err)
		return
	}
	fmt.Println("server is running at 127.0.0.1:8080")
}
