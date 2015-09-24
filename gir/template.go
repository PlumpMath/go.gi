package gir

import (
	"text/template"
	"strings"
)

var (
	girTemplate = template.Must(template.New("GIR").Funcs(map[string]interface{} {
		"ToLower": strings.ToLower,
	}).Parse(templateText))
)
