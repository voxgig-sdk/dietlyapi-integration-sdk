package core

type DietlyapiIntegrationError struct {
	IsDietlyapiIntegrationError bool
	Sdk              string
	Code             string
	Msg              string
	Ctx              *Context
	Result           any
	Spec             any
}

func NewDietlyapiIntegrationError(code string, msg string, ctx *Context) *DietlyapiIntegrationError {
	return &DietlyapiIntegrationError{
		IsDietlyapiIntegrationError: true,
		Sdk:              "DietlyapiIntegration",
		Code:             code,
		Msg:              msg,
		Ctx:              ctx,
	}
}

func (e *DietlyapiIntegrationError) Error() string {
	return e.Msg
}
