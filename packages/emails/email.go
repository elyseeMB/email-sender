package emails

import (
	"embed"
)

//go:embed dist
var Templates embed.FS
