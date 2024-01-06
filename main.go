package main

import (
	"fmt"
	"log"
	"os"
	"tea-time/components"
	"tea-time/data"
	"tea-time/template"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
)

type Playlist struct {
	Id        string `db:"id" json:"id"`
	Name      string `db:"name" json:"name"`
	SpotifyId string `db:"SpotifyId" json:"SpotifyId"`
}

func main() {
	app := pocketbase.New()

	// serves static files from the provided public dir (if exists)
	app.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		template.NewTemplateRenderer(e.Router)

		e.Router.GET("/*", apis.StaticDirectoryHandler(os.DirFS("./pb_public"), false))

		e.Router.GET("/hello/:name", func(c echo.Context) error {
			name := c.PathParam("name")
			component := components.Index(name)

			playlists, err := data.GetAllPlaylists(app.DB().Builder)

			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(playlists)

			return template.Html(c, component)
		})

		return nil
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}
