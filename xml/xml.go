package xml

import (
	stdxml "encoding/xml"
)

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
