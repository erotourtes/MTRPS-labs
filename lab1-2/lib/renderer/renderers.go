package renderer

import (
	"mainmod/lib/common"
	"mainmod/lib/renderer/ansi"
	"mainmod/lib/renderer/html"
)

const (
	ANSI = "ansi"
	HTML = "html"
)

var MapRenderer = map[string]common.Renderer{
	ANSI: ansi.Render,
	HTML: html.Render,
}
