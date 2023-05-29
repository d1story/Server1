package main

import (
	"github.com/go-chi/chi/v5"
)

func setUpAdminRouter(api *apiStat) chi.Router {
	adminRouter := chi.NewRouter()
	//apiRouter.Use(middlewareCors)
	adminRouter.Get("/metrics", api.adminMetricFormat)
	return adminRouter
}
