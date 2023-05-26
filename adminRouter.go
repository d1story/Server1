package main

import "github.com/go-chi/chi/v5"

func setUpAdminRouter(api *apiStat) chi.Router {
	adminRouter := chi.NewRouter()
	//apiRouter.Use(middlewareCors)
	htmlFile := file{data: formatAdminMetricPage(api.hits)}
	adminRouter.Get("/metrics", htmlFile.adminMetricHandler)
	return adminRouter
}
