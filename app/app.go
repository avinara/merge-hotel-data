package app

import (
	"fmt"
	"net/http"

	"github.com/merge-hotel-data/config"
	"github.com/merge-hotel-data/controllers"
	"github.com/merge-hotel-data/errors"
	"github.com/merge-hotel-data/routes"
	"github.com/merge-hotel-data/services"
)

type App struct {
	Config config.Config
	Router routes.RouterInterface
}

func NewApp() *App {
	config, err := config.LoadConfig("config.json")
	if err != nil {
		err2 := errors.LoadingConfigurationFileError()
		panic(err2)
	}

	return &App{
		Config: *config,
	}
}

func (app *App) Init() {
	mergeHotelDataService := services.NewMergeHotelDataService(app.Config)
	supplierService := services.NewSupplierService(app.Config)
	mergeHotelDataController := controllers.NewMergeHotelDataController(mergeHotelDataService, supplierService)
	router := routes.NewRouter()
	app.Router = router
	router.InitRoutes(mergeHotelDataController)
}

func (app *App) Run() {
	fmt.Println("Starting server on port 8080")
	fmt.Println("############## Server Started ##############")
	http.ListenAndServe(":8080", app.Router.GetMux())

}
