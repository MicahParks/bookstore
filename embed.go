package bookstore

import (
	"embed"
)

// frontend is the embedded file system for frontend assets.
//go:embed frontend
var Frontend embed.FS
