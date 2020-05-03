package main

import (
	"bitbucket.org/eyediagnosis/ishai_strauss-senior-full-stack-task/db"
	"bitbucket.org/eyediagnosis/ishai_strauss-senior-full-stack-task/env"
	"bitbucket.org/eyediagnosis/ishai_strauss-senior-full-stack-task/handler"
	"bitbucket.org/eyediagnosis/ishai_strauss-senior-full-stack-task/store"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	envInstance, envErr := env.New()
	if envErr != nil {
		panic(envErr)
	}

	dbInstance, dbErr := db.New(envInstance.DB)
	if dbErr != nil {
		panic(dbErr)
	}

	storeInstance := store.New(dbInstance)

	echoInstance := echo.New()
	echoInstance.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	handlerInstance := handler.NewHandlerGroup(storeInstance)
	handlerInstance.HandleRoutes(echoInstance.Group(""))

	echoInstance.Logger.Fatal(echoInstance.Start(":8080"))
}
