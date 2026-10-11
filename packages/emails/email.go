package emails

import (
	"embed"
)

//go:embed dist
var Templates embed.FS

type (
	PresenterConfig struct {
		APIBaseURL string
	}
)

func DefaultPresenterConfig(baseURL string) PresenterConfig {
	return PresenterConfig{
		APIBaseURL: "",
	}
}

func renderEmail() {}
