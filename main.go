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
	var app = infrastructure.NewConfig().
		Name(os.Getenv("APP_NAME")).
		ContextTimeout(10 * time.Second).
		Logger(log.InstanceLogrusLogger).
		Validator(validation.InstanceGoPlayground).
		// DbSQL(database.InstancePostgres).
		DbNoSQL(database.InstanceMongoDB)

	app.WebServerPort(os.Getenv("APP_PORT")).
		WebServer(router.InstanceGorillaMux).
		Start()
}
