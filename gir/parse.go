package gir

import (
	myxml "zenhack.net/go/gi/xml"
	"encoding/xml"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"strings"
)

type Parser struct {
	Logger *log.Logger
}

func NewParser(logWriter io.Writer) *Parser {
	if logWriter == nil {
		logWriter = ioutil.Discard
	}
	return &Parser{
		Logger: log.New(logWriter, "GIR Parser:", log.LstdFlags),
	}
}

func (p *Parser) unrecognized(e *myxml.Element) {
	p.Logger.Println("WARNING: Unrecognized Child Element ", e.Name)
}

func (p *Parser) ParseRepository(e *myxml.Element) (repo *Repository, err error) {
	defer func() { err, _ = recover().(error) }()
	requireTag(e, xml.Name{CORE_URI, "repository"})
	repo = &Repository{
		Includes: []*Include{},
	}
	for _, child := range e.Children {
		switch elt := child.(type) {
		case *myxml.Element:
			switch elt.Name {
			case xml.Name{CORE_URI, "include"}:
				inc, err := p.ParseInclude(elt)
				check(err)
				repo.Includes = append(repo.Includes, inc)
			case xml.Name{CORE_URI, "namespace"}:
				ns, err := p.ParseNamespace(elt)
				check(err)
				if repo.Namespace != nil {
					panic(&Duplicate{
						Old: repo.Namespace,
						New: ns,
						In:  e,
					})
				}
				repo.Namespace = ns
			case xml.Name{CORE_URI, "package"}:
				pkg, err := p.ParsePackage(elt)
				check(err)
				if repo.PackageName != "" {
					panic(&Duplicate{
						Old: repo.PackageName,
						New: pkg,
						In:  e,
					})
				}
				repo.PackageName = pkg
			default:
				p.unrecognized(elt)
			}
		}
	}
	return
}

func (p *Parser) ParseNamespace(e *myxml.Element) (ns *Namespace, err error) {
	defer func() { err, _ = recover().(error) }()
	split := func(s string) []string {
		return strings.Split(s, ",")
	}
	attr := func(name xml.Name) string {
		return requireAttr(e.Attrs, name)
	}
	requireTag(e, xml.Name{CORE_URI, "namespace"})
	ns = &Namespace{
		Name:                attr(xml.Name{"", "name"}),
		SharedLibraries:     split(attr(xml.Name{"", "shared-library"})),
		CIdentifierPrefixes: split(attr(xml.Name{C_URI, "identifier-prefixes"})),
		CSymbolPrefixes:     split(attr(xml.Name{C_URI, "symbol-prefixes"})),
		Functions: []*Function{},
	}
	for _, child := range(e.Children) {
		switch elt := child.(type) {
		case *myxml.Element:
			switch elt.Name {
			case xml.Name{CORE_URI, "function"}:
				fn, err := p.ParseFunction(elt)
				check(err)
				ns.Functions = append(ns.Functions, fn)
			default:
				p.unrecognized(elt)
			}
		}
	}
	return
}

func (p *Parser) ParseInclude(elt *myxml.Element) (inc *Include, err error) {
	defer func() { err, _ = recover().(error) }()
	inc = &Include{}
	var nameOk, versionOk bool
	inc.Name, nameOk = elt.Attrs[xml.Name{"", "name"}]
	inc.Version, versionOk = elt.Attrs[xml.Name{"", "version"}]
	if !nameOk {
		panic(errors.New("<include> element has no name attribute."))
	} else if !versionOk {
		panic(errors.New("<include> element has no version attribute."))
	}
	return
}

func (p *Parser) ParsePackage(elt *myxml.Element) (pkg string, err error) {
	defer func() { err, _ = recover().(error) }()
	requireTag(elt, xml.Name{CORE_URI, "package"})
	pkg, ok := elt.Attrs[xml.Name{"", "name"}]
	if !ok {
		panic(errors.New("<package> element has no name attribute."))
	}
	return
}

func (p *Parser) ParseFunction(e *myxml.Element) (fn *Function, err error) {
	defer func() { err, _ = recover().(error) }()
	attr := func(name xml.Name) string {
		return requireAttr(e.Attrs, name)
	}
	fn = &Function{
		Name: attr(xml.Name{"", "name"}),
		CIdentifier: attr(xml.Name{C_URI, "identifier"}),

		// TODO: parse these out of the xml:
		Doc: "",
		ReturnValue: nil,
	}
	return
}
