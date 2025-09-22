package services

import (
	"fmt"
	"net/url"
	"strings"
)

type UrlServiceInterface interface {
	ProcessUrl(rawURL, operation string) (string, error)
}

type UrlService struct{}

func NewURLService() UrlServiceInterface {
	return &UrlService{}
}

func (s *UrlService) ProcessUrl(rawURL, operation string) (string, error) {
	u, err := url.Parse(rawURL)
	if err != nil {
		return "", fmt.Errorf("invalid url: %w", err)
	}

	// Canonical: remove query params, trailing slash
	canonical := func(u *url.URL) string {
		u.RawQuery = ""
		cleaned := strings.TrimSuffix(u.String(), "/")
		return cleaned
	}

	// Redirection: enforce www.byfood.com, lowercase
	redirect := func(u *url.URL) string {
		u.Host = "www.byfood.com"
		return strings.ToLower(u.String())
	}

	switch operation {
	case "canonical":
		return canonical(u), nil
	case "redirection":
		return redirect(u), nil
	case "all":
		cleaned := canonical(u)
		parsed, _ := url.Parse(cleaned)
		return redirect(parsed), nil
	default:
		return "", fmt.Errorf("unsupported operation: %s", operation)
	}
}
