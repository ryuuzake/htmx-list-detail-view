package main

import (
	"log"
	"os"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/plugins/migratecmd"
	"github.com/ryuuzake/htmx-list-detail-view/controller"
	"github.com/ryuuzake/htmx-list-detail-view/middleware"

	_ "github.com/ryuuzake/htmx-list-detail-view/migrations"
)

func main() {
	app := pocketbase.New()

	migratecmd.MustRegister(app, app.RootCmd, migratecmd.Config{
		Automigrate: true,
	})

	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.Use(middleware.LoadAuthContextFromCookie(app))

		projectHandler := controller.NewProjectHandler(app)
		e.Router.GET("/", projectHandler.GetProjects)
		e.Router.GET("/project/:id", projectHandler.GetProjectById)

		e.Router.GET("/*", apis.StaticDirectoryHandler(os.DirFS("./public"), false))
		return nil
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
