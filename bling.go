package bling

import (
	retryablehttp "github.com/hashicorp/go-retryablehttp"
)

// Client manages communication with the API.
type Client struct {
	// controller is the internal controller that handles the communication with the API.
	controller *Controller

	// Services used for talking to different parts of the Bling API.
	ExampleService *ExampleService
}

// New creates a new instance of Client with the specified endpoint and secret.
func New(baseURL string, options ...OptionFunc) (*Client, error) {
	c := &Client{}

	// TODO: set envs
	controller, err := NewController(baseURL, append(options, WithErrorHandler(errorHandler))...)
	if err != nil {
		return nil, err
	}

	c.controller = controller
	c.ExampleService = &ExampleService{c: controller}
	return c, nil
}

// Shutdown closes the idle connections.
func (c *Client) Shutdown() error {
	return c.controller.Shutdown()
}

// SetAuth sets the authentication header for the request.
func (c *Client) SetAuth(req *retryablehttp.Request, secret any) error {
	// Bla bla bla
	return nil
}
