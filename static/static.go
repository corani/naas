package static

import "embed"

//go:embed www/**
var FS embed.FS
