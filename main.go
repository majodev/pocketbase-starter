package main

import (
	"embed"
	"io/fs"
	"log"
	"os"
	"strings"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/plugins/migratecmd"
	// _ "github.com/majodev/pocketbase-starter/migrations" // TODO: uncomment as soon as you have ./migrations/*.go files
)

// embeddedFS holds our static content within public, embedded into the binary
// no need to transfer the public folder into the Docker image or with the app binary
//
//go:embed public/*
var embeddedFS embed.FS

func main() {

	app := pocketbase.New()

	// loosely check if it was executed using "go run"
	isGoRun := strings.HasPrefix(os.Args[0], os.TempDir())

	migratecmd.MustRegister(app, app.RootCmd, migratecmd.Config{
		// enable auto creation of migration files when making collection changes in the Admin UI
		// (the isGoRun check is to enable it only during development)
		Automigrate: isGoRun,
	})

	if err := bindAppHooks(app); err != nil {
		log.Fatal(err)
	}

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}

// Add pocketbase hooks as needed here...
// https://pocketbase.io/docs/go-overview/
// https://pocketbase.io/docs/go-event-hooks/
// https://pocketbase.io/v023upgrade/go/
func bindAppHooks(app core.App) error {

	// extract the embedded public filesystem
	public, err := fs.Sub(embeddedFS, "public")

	if err != nil {
		return err
	}

	app.OnServe().BindFunc(func(se *core.ServeEvent) error {
		// serves static files from the provided embedded public
		se.Router.GET("/{path...}", apis.Static(public, true))

		return se.Next()
	})

	return nil
}
