package renderer

import (
	"mainmod/lib/common"
	"mainmod/lib/renderer/ansi"
	"mainmod/lib/renderer/html"
)

var MapRenderer = map[string]common.Renderer{
	"ansi": ansi.Render,
	"html": html.Render,
}
