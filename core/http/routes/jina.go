package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/siperal/hYZmet-localai/core/config"
	"github.com/siperal/hYZmet-localai/core/http/endpoints/jina"
	"github.com/siperal/hYZmet-localai/core/http/middleware"
	"github.com/siperal/hYZmet-localai/core/schema"

	"github.com/siperal/hYZmet-localai/pkg/model"
)

func RegisterJINARoutes(app *echo.Echo,
	re *middleware.RequestExtractor,
	cl *config.ModelConfigLoader,
	ml *model.ModelLoader,
	appConfig *config.ApplicationConfig) {

	// POST endpoint to mimic the reranking
	rerankHandler := jina.JINARerankEndpoint(cl, ml, appConfig)
	app.POST("/v1/rerank",
		rerankHandler,
		middleware.ExposeNodeHeader(appConfig),
		re.BuildFilteredFirstAvailableDefaultModel(config.BuildUsecaseFilterFn(config.FLAG_RERANK)),
		re.SetModelAndConfig(func() schema.LocalAIRequest { return new(schema.JINARerankRequest) }))
}
