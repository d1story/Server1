package main

import "github.com/go-chi/chi/v5"

func setUpApiRouter(api *apiStat) chi.Router {
	apiRouter := chi.NewRouter()
	//apiRouter.Use(middlewareCors)
	apiRouter.Get("/healthz", healthHandler)
	apiRouter.Get("/metrics", api.displayStats)
	apiRouter.Get("/profane", checkProfane)
	return apiRouter
}
