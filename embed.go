package bookstore

import (
	"embed"
)

// FrontendSubDir is the name of the subdirectory to make the root of the embedded frontend.
const FrontendSubDir = "frontend"

// Frontend is the embedded file system for frontend assets.
//go:embed frontend
var Frontend embed.FS
