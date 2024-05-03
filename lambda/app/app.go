package app

import (
	"lambda-func/api"
	"lambda-func/database"
)

type App struct {
	ApiHandler api.ApiHandler
}

func NewApp() App {
	// we initialize our DB store
	db := database.NewDynamoDBClient()
	apiHandler := api.NewApiHandler(db)

	return App{
		ApiHandler: apiHandler,
	}
}
