package main

import (
	"fmt"
	"log"
	"os"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
	"http_api/modules/item/transport/handler"
	"http_api/middleware"
	"http_api/common"
)

type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	DBSSLMode  string
	DBTimezone string
}

func LoadConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Failed to load .env file")
	}

	return &Config{
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
		DBSSLMode:  os.Getenv("DB_SSLMODE"),
		DBTimezone: os.Getenv("DB_TIMEZONE"),
	}
}

func main() {
	config := LoadConfig()

	dns := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s TimeZone=%s",
		config.DBHost,
		config.DBPort,
		config.DBUser,
		config.DBPassword,
		config.DBName,
		config.DBSSLMode,
		config.DBTimezone,
	)

	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		common.ErrInternalServer(err)
	}

	// gin.SetMode(gin.ReleaseMode)
	router := gin.Default() // Neu dung Recovery() rieng thi nen chon gin.New()
	router.Use(middleware.Recovery())

	router.GET("/ping", func(c *gin.Context) {

		go func() {
			defer common.Recovery()
			fmt.Println([]int{}[0])
		}()
		
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	v1 := router.Group("/v1")
	{
		v1.POST("/parameters", handler.CreateParameter(db))
		v1.GET("/parameters/:id", handler.GetParameterByID(db))
		v1.PUT("/parameters/:id", handler.UpdateParameterByID(db))
		v1.DELETE("/parameters/:id", handler.DeleteParameter(db))
		v1.GET("/parameters", handler.ListParameters(db))
	}

	router.Run()
}
