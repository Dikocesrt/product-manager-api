package config

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
	Redis *redis.Client
)

func Init() {
    // Load environment variables from .env file
    LoadEnv()

    // Initialize database connection
    InitDB()

    // Initialize Redis connection
    InitRedis()
}

func LoadEnv() {
    if err := godotenv.Load(); err != nil {
        log.Fatal("Error loading .env file")
    }
}

func InitDB() {
    var err error
    dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
        os.Getenv("DB_HOST"),
        os.Getenv("DB_USER"),
        os.Getenv("DB_PASSWORD"),
        os.Getenv("DB_NAME"),
        os.Getenv("DB_PORT"),
    )

    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database: ", err)
    }

    log.Println("Database connection established")
}

func InitRedis() {
    ctx := context.Background()
    Redis = redis.NewClient(&redis.Options{
        Addr:     os.Getenv("REDIS_ADDR"),     // Redis address, default is localhost:6379
        Password: os.Getenv("REDIS_PASSWORD"), // Redis password, default is empty
        DB:       0,                           // Default DB
    })

    _, err := Redis.Ping(ctx).Result()
    if err != nil {
        log.Fatal("Failed to connect to Redis: ", err)
    }

    log.Println("Redis connection established")
}