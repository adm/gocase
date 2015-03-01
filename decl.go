package main

import (
	"fmt"
	"go/ast"
	"go/token"
)

type decl struct {
	name       string
	suggestion string
	pos        token.Position
	kind       ast.ObjKind
	freq       int32
}

func (d decl) String() string {
	return fmt.Sprintf("%s:#%d,%s,%s,%s", d.pos.Filename, d.pos.Offset, d.kind.String(), d.name, d.suggestion)
}

func print(decls []decl) {
	fmt.Printf("#offset position in file, kind, identifier, suggested\n")
	for _, s := range decls {
		comment := ""
		if s.name == s.suggestion {
			comment = "#"
		}
		fmt.Printf("%s %s\n", comment, s.String())
	}
}
