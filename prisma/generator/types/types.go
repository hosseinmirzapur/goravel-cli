package types

import (
	"fmt"

	"github.com/hosseinmirzapur/goravel-cli/prisma/helpers/gocase"
	"github.com/hosseinmirzapur/goravel-cli/prisma/helpers/strcase"
)

// String acts as a builtin string but provides useful casing methods.
type String string

func (s String) String() string {
	return string(s)
}

// GoCase transforms strings into Go-style casing, meaning uppercase including Go casing edge cases.
func (s String) GoCase() string {
	return gocase.To(strcase.ToCamel(string(s)))
}

// GoLowerCase transforms strings into Go-style lowercase casing. It is like GoCase but used for private fields.
func (s String) GoLowerCase() string {
	return gocase.To(strcase.ToLowerCamel(string(s)))
}

// CamelCase transforms strings into camelCase casing. It is often used for json mappings.
func (s String) CamelCase() string {
	return strcase.ToLowerCamel(string(s))
}

// Tag returns the struct tag value of a field.
func (s String) Tag(isRequired bool) string {
	if !isRequired {
		return fmt.Sprintf("`json:\"%s,omitempty\"`", s)
	}
	return fmt.Sprintf("`json:\"%s\"`", s)
}

// PrismaGoCase transforms `relevance` into `Relevance_`
func (s String) PrismaGoCase() string {
	return strcase.ToCamel(string(s)) + "_"
}

// PrismaInternalCase transforms `relevance` into `_relevance`
func (s String) PrismaInternalCase() string {
	return "_" + string(s)
}

// builtin Go types
var builtin = map[string]string{
	"ID":       "string",
	"String":   "string",
	"Boolean":  "bool",
	"Int":      "int",
	"Float":    "float64",
	"DateTime": "DateTime",
	"Json":     "JSON",
	"Bytes":    "Bytes",
	"BigInt":   "BigInt",
}

// Type acts as a builtin string but provides useful methods for type DMMF values.
type Type string

func (t Type) String() string {
	return string(t)
}

// Value returns the native value of a type.
func (t Type) Value() string {
	str := string(t)
	v, ok := builtin[str]
	if ok {
		return v
	}

	return gocase.To(strcase.ToCamel(str))
}

// GoCase transforms strings into Go-style lowercase casing. It is like GoCase but used for private fields.
func (t Type) GoCase() string {
	return gocase.To(strcase.ToCamel(string(t)))
}

// GoLowerCase transforms strings into Go-style lowercase casing. It is like GoCase but used for private fields.
func (t Type) GoLowerCase() string {
	return gocase.To(strcase.ToLowerCamel(string(t)))
}

// CamelCase transforms strings into camelCase casing. It is often used for json mappings.
func (t Type) CamelCase() string {
	return strcase.ToLowerCamel(string(t))
}
