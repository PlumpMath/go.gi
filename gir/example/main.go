package main

import (
	".."
	myxml "../../xml"
	"encoding/xml"
	"fmt"
	"os"
)

func main() {
	dec := xml.NewDecoder(os.Stdin)
	elt := &myxml.Element{}
	err := dec.Decode(&elt)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error: ", err)
		return
	}
	p := gir.NewParser(os.Stderr)
	repo, err := p.ParseRepository(elt)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error: ", err)
		return
	}
	err = repo.GenGo(os.Stdout)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error: ", err)
		return
	}
}
