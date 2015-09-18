package gir

import (
	"io"
	"strings"
)

type stopWriter struct {
	w   io.Writer
	err *error
}

func (w *stopWriter) Write(p []byte) (n int, err error) {
	if *w.err != nil {
		return 0, *w.err
	}
	n, *w.err = w.w.Write(p)
	return n, *w.err
}

func (r *Repository) GenGo(w io.Writer) (err error) {
	w = &stopWriter{w, &err}
	out := []string{
		`package `, strings.ToLower(r.Namespace.Name), "\n",
		"// #cgo CFLAGS: `pkg-config --cflags ", r.PackageName, "`", "\n",
		"// #cgo LDFLAGS: `pkg-config --libs ", r.PackageName, "`", "\n",
		`import "C"`, "\n",
	}
	for _, s := range out {
		io.WriteString(w, s)
	}
	return
}
