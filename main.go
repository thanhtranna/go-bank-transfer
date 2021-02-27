package main

import (
	"os"
	"time"

	"go-bank-transfer/infrastructure"
	"go-bank-transfer/infrastructure/database"
	"go-bank-transfer/infrastructure/log"
	"go-bank-transfer/infrastructure/router"
	"go-bank-transfer/infrastructure/validation"
)

func main() {
	// Set Environment Variables
	// os.Setenv("APP_NAME", "go-bank-transfer")
	// os.Setenv("APP_PORT", "8001")

	// os.Setenv("GO111MODULE", "on")
	// os.Setenv("CGO_ENABLED", "0")
	// os.Setenv("GOOS", "linux")
	// os.Setenv("GOARCH", "amd64")

	// os.Setenv("MONGODB_HOST", "localhost")
	// os.Setenv("MONGODB_PORT", "27017")
	// os.Setenv("MONGODB_DATABASE", "bank")
	// os.Setenv("MONGODB_ROOT_USER", "")
	// os.Setenv("MONGODB_ROOT_PASSWORD", "")

	// os.Setenv("POSTGRES_HOST", "localhost")
	// os.Setenv("POSTGRES_PORT", "5432")
	// os.Setenv("POSTGRES_USER", "postgres")
	// os.Setenv("POSTGRES_PASSWORD", "")
	// os.Setenv("POSTGRES_DATABASE", "bank")
	// os.Setenv("POSTGRES_DRIVER", "postgres")

	var app = infrastructure.NewConfig().
		Name(os.Getenv("APP_NAME")).
		ContextTimeout(10 * time.Second).
		Logger(log.InstanceLogrusLogger).
		Validator(validation.InstanceGoPlayground).
		DbSQL(database.InstancePostgres).
		DbNoSQL(database.InstanceMongoDB)

	app.WebServerPort(os.Getenv("APP_PORT")).
		WebServer(router.InstanceGorillaMux).
		Start()
}
