// Package xml emplements an element type that can be used to ummarshal
// arbitrary xml, dealing with it in the form of a tree. This can often be
// easier than defining specialized types that implement UnmarshalXML
// directly.
package xml // import "zenhack.net/go/gi/xml"

import (
	stdxml "encoding/xml"
)

// An xml element
type Element struct {
	Name     stdxml.Name
	Attrs    map[stdxml.Name]string
	Children []interface{}
}

func newElement() *Element {
	return &Element{
		Attrs:    make(map[stdxml.Name]string),
		Children: []interface{}{},
	}
}

func (e *Element) UnmarshalXML(d *stdxml.Decoder, start stdxml.StartElement) error {
	*e = *newElement()
	e.Name = start.Name
	for _, attr := range start.Attr {
		e.Attrs[attr.Name] = attr.Value
	}
	for {
		t, err := d.Token()
		if err != nil {
			return err
		}
		switch child := t.(type) {
		case stdxml.EndElement:
			return nil
		case stdxml.StartElement:
			elt := newElement()
			err = elt.UnmarshalXML(d, child)
			if err != nil {
				return err
			}
			e.Children = append(e.Children, elt)
		default:
			e.Children = append(e.Children, child)
		}
	}
}
