package controllers

// ControllerResult is returned from controllers that then need to be mapped to http responses
type ControllerResult struct {
	Success      bool
	Code         int
	Error        error
	ErrorMessage string
	Data         interface{}
}
