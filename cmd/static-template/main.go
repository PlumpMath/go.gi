package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
)

var (
	name = flag.String("name", "", "Name of the package to generate.")
)

func check(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error: ", err)
		os.Exit(1)
	}
}

func main() {
	flag.Parse()
	if *name == "" {
		fmt.Fprintln(os.Stderr, "Please supply a package name.")
		flag.PrintDefaults()
		os.Exit(1)
	}
	buf := &bytes.Buffer{}
	_, err := io.Copy(buf, os.Stdin)
	check(err)
	fmt.Println("package ", *name)
	fmt.Printf("const templateText = %q\n", buf.String())
}
