package gir

const (
	CORE_URI = "http://www.gtk.org/introspection/core/1.0"
	C_URI    = "http://www.gtk.org/introspection/c/1.0"
	GLIB_URI = "http://www.gtk.org/introspection/glib/1.0"
)

type Repository struct {
	Namespace   *Namespace
	PackageName string
	CIncludes   []string
	Includes    []*Include
}

type Namespace struct {
	Name                string
	SharedLibraries     []string
	CIdentifierPrefixes []string
	CSymbolPrefixes     []string
	Functions           []*Function
}

type Include struct {
	Name    string
	Version string
}

type Function struct {
	Name        string
	CIdentifier string
	Doc         string
	ReturnValue *ReturnValue
}

type ReturnValue struct {
	Type *Type
	//	transfer-ownership="none"
}

type Type struct {
	Name, CType string
}
