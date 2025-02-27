package main

import (
	"github.com/SamareshSingh/banking/app"
	"github.com/SamareshSingh/banking/logger"
)

func main() {
	//log.Println("Starting our application...")
	/*
		we need to structure the log in below structured json logging format
		{
			"timestamp": "2025/02/26 21:14:09",
			"msg": "Starting our application..."
		}
	*/
	logger.Info("Starting the application...")
	app.Start()
}
