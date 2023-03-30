package app

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"gitlab.com/meta-node/client/config"
	"gitlab.com/meta-node/client/server/core/router"
)

type App struct {
	config *config.Config
	router *mux.Router
}

func InitApp(
	config *config.Config,
	db *sqlx.DB,
	hub *router.Hub,
) *App {
	server := router.Server{}
	serverapp := server.Init(config)
	serverapp.Start(router.GetPORT())

	router := router.InitRouter(serverapp, db, hub)
	return &App{
		config: config,
		router: router,
	}
}

func (app *App) Run() {
	http.ListenAndServe(":2001", app.router)

	fmt.Println("Server is running: http://localhost:2001")
}
