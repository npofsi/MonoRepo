//use this to embed static files into the go binary.

package webpages

import (
	"embed"
)

//go:embed build/**/*
var WebpageFS embed.FS
