package handlers

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/goesbams/mini-books-library/backend/entities"
	"github.com/goesbams/mini-books-library/backend/services"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type UrlHandler struct {
	service services.UrlServiceInterface
}

func NewUrlHandler(service services.UrlServiceInterface) *UrlHandler {
	return &UrlHandler{
		service: service,
	}
}

// ProcessURL processes URL cleanup and redirection
// @Summary Process URL cleanup/redirection
// @Description Clean or redirect a given URL based on operation type (canonical, redirection, all)
// @Tags urls
// @Accept json
// @Produce json
// @Param request body entities.URLRequest true "URL and Operation"
// @Success 200 {object} entities.URLResponse
// @Failure 400 {object} map[string]interface{}
// @Router /urls/process [post]
func (h *UrlHandler) ProcessUrl(c echo.Context) error {
	var req entities.URLRequest
	if err := c.Bind(&req); err != nil {
		logrus.WithError(err).Error("failed to bind url request")
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error":   "bad_request",
			"message": "invalid input format",
		})
	}

	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error":   "validation_error",
			"message": err.Error(),
		})
	}

	processed, err := h.service.ProcessUrl(req.URL, req.Operation)
	if err != nil {
		logrus.WithError(err).Error("failed to process url")
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error":   "processing_error",
			"message": err.Error(),
		})
	}

	logrus.Infof("processed url operation:%s result:%s", req.Operation, processed)
	return c.JSON(http.StatusOK, entities.URLResponse{ProcessedURL: processed})
}
