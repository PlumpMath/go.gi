package gir

import (
	myxml "../xml"
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

type BadElementType struct {
	Expected []xml.Name
	Actual   *myxml.Element
}

func (e *BadElementType) Error() string {
	buf := &bytes.Buffer{}
	fmt.Fprintln(buf, "Bad element type. Expected one of:")
	for i := range e.Expected {
		fmt.Fprintf(buf, "  <%q:%s>\n", e.Expected[i].Space, e.Expected[i].Local)
	}
	xmlname := e.Actual.Name
	fmt.Fprintf(buf, "but got <%q:%s>\n", xmlname.Space, xmlname.Local)
	return buf.String()
}

func requireTag(actual *myxml.Element, expected ...xml.Name) {
	for i := range expected {
		if actual.Name == expected[i] {
			return
		}
	}
	panic(&BadElementType{
		Expected: expected,
		Actual:   actual,
	})
}

func requireAttr(attrs map[xml.Name]string, name xml.Name) string {
	attr, ok := attrs[name]
	if !ok {
		panic(errors.New(fmt.Sprintf("No such attribute %v\n", name)))
	}
	return attr
}

type Duplicate struct {
	Old, New, In interface{}
}

func (e *Duplicate) Error() string {
	return fmt.Sprintln(
		"Illegal duplicate value. Element", e.New,
		"In element", e.In,
		"would replace old value", e.Old, ".",
	)
}
