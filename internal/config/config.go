package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

func LoadFromEnv() StorageConfig {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file")
	}

	host := os.Getenv("HOST")
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	dbname := os.Getenv("NAME")
	port := os.Getenv("PORT")

	fmt.Printf("Host: %s\n", host)
	fmt.Printf("User: %s\n", user)
	fmt.Printf("Password: %s\n", password)
	fmt.Printf("DBName: %s\n", dbname)
	fmt.Printf("Port: %s\n", port)

	return StorageConfig{
		Host:     host,
		User:     user,
		Password: password,
		Dbname:   dbname,
		Port:     port,
	}
}

type StorageConfig struct {
	Host     string `json:"host"`
	User     string `json:"user"`
	Password string `json:"password"`
	Dbname   string `json:"dbname"`
	Port     string `json:"port"`
}

func (c StorageConfig) ConnectionString() string {
	connectionString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		c.Host, c.User, c.Password, c.Dbname, c.Port)
	return connectionString
}
