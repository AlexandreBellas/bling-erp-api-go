package bling

import "context"

// ExampleService handles communication with the example related methods of the API.
type ExampleService struct {
	c *Controller
}

// Example represents an example response from the API.
type Example struct {
	ID int `json:"id"`
}

// GetExample returns an example.
func (s *ExampleService) GetExample(ctx context.Context) (*Example, error) {
	req, err := s.c.NewRequest(ctx, "GET", "/example", nil)
	if err != nil {
		return nil, err
	}

	example := new(Example)
	_, err = s.c.Do(req, example)
	return example, err
}
