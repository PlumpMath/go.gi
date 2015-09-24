package gir

import (
	"io"
)

func (r *Repository) GenGo(w io.Writer) (err error) {
	return girTemplate.Execute(w, r)
}
