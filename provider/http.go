package provider

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

type APIStyle string

const (
	OpenAI    APIStyle = "openai"
	Anthropic APIStyle = "anthropic"
	Gemini    APIStyle = "gemini"
)

type HTTPModel struct {
	Name     string
	Style    APIStyle
	Endpoint string
	APIKey   string
	Model    string
	Client   *http.Client
}

func (p *HTTPModel) Execute(ctx context.Context, request Request) (Response, error) {
	if p.Endpoint == "" || p.Model == "" {
		return Response{}, fmt.Errorf("provider %s is not configured", p.Name)
	}
	body := p.requestBody(request)
