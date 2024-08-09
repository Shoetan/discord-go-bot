package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)


func GetEnv(key string) string{
	err := godotenv.Load(".env")
	
	if err != nil {
		log.Fatal("Could not load environmental variable")
	}

	return os.Getenv(key)

}