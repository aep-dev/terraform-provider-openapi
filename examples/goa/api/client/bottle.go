// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "cellar": bottle Resource Client
//
// Command:
// $ goagen
// --design=github.com/dikhan/terraform-provider-openapi/examples/goa/api/design
// --out=$(GOPATH)/src/github.com/dikhan/terraform-provider-openapi/examples/goa/api
// --version=v1.3.1

package client

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"net/url"
)

// CreateBottlePath computes a request path to the create action of bottle.
func CreateBottlePath() string {

	return fmt.Sprintf("/bottles/")
}

// creates a bottle
func (c *Client) CreateBottle(ctx context.Context, path string, payload *BottlePayload, contentType string) (*http.Response, error) {
	req, err := c.NewCreateBottleRequest(ctx, path, payload, contentType)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewCreateBottleRequest create the request corresponding to the create action endpoint of the bottle resource.
func (c *Client) NewCreateBottleRequest(ctx context.Context, path string, payload *BottlePayload, contentType string) (*http.Request, error) {
	var body bytes.Buffer
	if contentType == "" {
		contentType = "*/*" // Use default encoder
	}
	err := c.Encoder.Encode(payload, &body, contentType)
	if err != nil {
		return nil, fmt.Errorf("failed to encode body: %s", err)
	}
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("POST", u.String(), &body)
	if err != nil {
		return nil, err
	}
	header := req.Header
	if contentType == "*/*" {
		header.Set("Content-Type", "application/json")
	} else {
		header.Set("Content-Type", contentType)
	}
	return req, nil
}

// ShowBottlePath computes a request path to the show action of bottle.
func ShowBottlePath(id string) string {
	param0 := id

	return fmt.Sprintf("/bottles/%s", param0)
}

// shows a bottle
func (c *Client) ShowBottle(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewShowBottleRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewShowBottleRequest create the request corresponding to the show action endpoint of the bottle resource.
func (c *Client) NewShowBottleRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}