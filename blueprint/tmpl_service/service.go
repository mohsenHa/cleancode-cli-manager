package tmpl_service

func GetSampleTmpl() string {
	return `package {{.}}service

import (
	"clean-code-structure/param/{{.}}param"
)

func (s Service) Sample(req {{.}}param.SampleRequest) ({{.}}param.SampleResponse, error) {
	return {{.}}param.SampleResponse{Message: "everything is good!"}, nil
}
`
}

func GetServiceTmpl() string {
	return `package {{.}}service

type Service struct {
}

func New() Service {
	return Service{}
}
`
}
