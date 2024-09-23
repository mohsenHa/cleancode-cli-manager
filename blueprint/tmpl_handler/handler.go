package tmpl_handler

func GetHandlerTmpl() string {
	return `package {{.}}handler

import (
	"clean-code-structure/service/{{.}}service"
	"clean-code-structure/validator/{{.}}validator"
)

type Handler struct {
	{{.}}Validator {{.}}validator.Validator
	{{.}}Service   {{.}}service.Service
}

func New({{.}}Service {{.}}service.Service, {{.}}Validator {{.}}validator.Validator) Handler {
	return Handler{
		{{.}}Service:   {{.}}Service,
		{{.}}Validator: {{.}}Validator,
	}
}
`
}
func GetRoutTmpl() string {
	return `package {{.}}handler

import "github.com/labstack/echo/v4"

func (h Handler) SetRoutes(messageGroup *echo.Group) {

	messageGroup.GET("/sample", h.sample)

}
`
}
func GetSampleTmpl() string {
	return `package {{.}}handler

import (
	"clean-code-structure/param/{{.}}param"
	"clean-code-structure/pkg/httpmsg"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h Handler) sample(c echo.Context) error {
	var req {{.}}param.SampleRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	if fieldErrors, err := h.{{.}}Validator.ValidateSampleRequest(req); err != nil {
		msg, code := httpmsg.Error(err)
		return c.JSON(code, echo.Map{
			"message": msg,
			"errors":  fieldErrors,
		})
	}

	resp, err := h.{{.}}Service.Sample(req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, resp)

}
`
}
