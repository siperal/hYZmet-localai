package localai

import (
	"github.com/labstack/echo/v4"
	"github.com/siperal/hYZmet-localai/core/backend"
	"github.com/siperal/hYZmet-localai/core/config"
	"github.com/siperal/hYZmet-localai/core/http/middleware"
	"github.com/siperal/hYZmet-localai/core/schema"
	"github.com/siperal/hYZmet-localai/pkg/model"
)

// TokenizeEndpoint exposes a REST API to tokenize the content
// @Summary Tokenize the input.
// @Tags tokenize
// @Param request body schema.TokenizeRequest true "Request"
// @Success 200 {object} schema.TokenizeResponse "Response"
// @Router /v1/tokenize [post]
func TokenizeEndpoint(cl *config.ModelConfigLoader, ml *model.ModelLoader, appConfig *config.ApplicationConfig) echo.HandlerFunc {
	return func(c echo.Context) error {
		input, ok := c.Get(middleware.CONTEXT_LOCALS_KEY_LOCALAI_REQUEST).(*schema.TokenizeRequest)
		if !ok || input.Model == "" {
			return echo.ErrBadRequest
		}

		cfg, ok := c.Get(middleware.CONTEXT_LOCALS_KEY_MODEL_CONFIG).(*config.ModelConfig)
		if !ok || cfg == nil {
			return echo.ErrBadRequest
		}

		tokenResponse, err := backend.ModelTokenize(input.Content, ml, *cfg, appConfig)
		if err != nil {
			return err
		}
		return c.JSON(200, tokenResponse)
	}
}
