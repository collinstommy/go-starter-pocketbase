package main

import (
	"log"
	"os"
	"tea-time/components"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
)

func main() {
	app := pocketbase.New()

	app.OnServe().BindFunc(func(se *core.ServeEvent) error {
		// Serve static files
		se.Router.GET("/{path...}", apis.Static(os.DirFS("./pb_public"), false))

		// Other routes
		se.Router.GET("/hello/{name}", helloRoute)

		return se.Next()
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}

func helloRoute(e *core.RequestEvent) error {
	name := e.Request.PathValue("name")

	component := components.Index(name)

	e.Response.Header().Set("Content-Type", "text/html; charset=utf-8")

	return component.Render(e.Request.Context(), e.Response)
}
