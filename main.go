package main

import (
	"log"
	"os"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"github.com/ryuuzake/htmx-list-detail-view/controller"
	"github.com/ryuuzake/htmx-list-detail-view/middleware"
)

func main() {
	app := pocketbase.New()

	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.Use(middleware.LoadAuthContextFromCookie(app))

		projectHandler := controller.NewProjectHandler(app)
		e.Router.GET("/", projectHandler.GetProjects)
		e.Router.GET("/project/:id", projectHandler.GetProjectByParentId)

		e.Router.GET("/*", apis.StaticDirectoryHandler(os.DirFS("./public"), false))
		return nil
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
