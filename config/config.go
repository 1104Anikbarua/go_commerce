package config

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var configurations TSConfig

type TSConfig struct {
	Version     string
	ServiceName string
	HttpPort    int
}

func loadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}
	httpPortStr := os.Getenv("HTTP_PORT")
	if httpPortStr == "" {
		fmt.Println("Port is required!", http.StatusNotFound)
		os.Exit(1)
	}
	httpPortInt, err := strconv.ParseInt(httpPortStr, 10, 64)
	if err != nil {
		fmt.Println("Failed To convert in number")
		os.Exit(1)
	}
	version := os.Getenv("VERSION")
	if version == "" {
		fmt.Println("Version is required!")
		os.Exit(1)
	}
	serviceName := os.Getenv("SERVICE_NAME")
	if serviceName == "" {
		fmt.Println("Service Name is required!")
		os.Exit(1)
	}

	configurations = TSConfig{
		Version:     version,
		ServiceName: serviceName,
		HttpPort:    int(httpPortInt),
	}

}

func GetConfig() TSConfig {
	loadConfig()
	return configurations
}
