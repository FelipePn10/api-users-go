package routes

//
//import (
//	"github.com/go-chi/chi/v5"
//	_ "github.com/swaggo/http-swagger/example/go-chi/docs"
//	httpSwagger "github.com/swaggo/http-swagger/v2"
//)
//
//var (
//	docsURL = "http://localhost:8080/docs/doc.json"
//)
//
//// @title 		Swagger Dark Mode
//// @Version 	1.0
//// @securityDefinitions.apikey		ApiKeyAuth
//// @in								header
//// @name							Authorization
//
//func InitDocsRoute(r chi.Router) {
//	r.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL(docsURL),
//		httpSwagger.AfterScript(custom.CustomJS),
//		httpSwagger.DocExpansion("none"),
//		httpSwagger.UIConfig(map[string]string{
//			"defaultModelsExpandDepth": "-1",
//		}),
//	))
//}
