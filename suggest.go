package main

import (
	"fmt"
	"go/ast"
	"go/build"
	"go/parser"
	"go/token"
	"io/ioutil"
	"path/filepath"
	"strings"
)

func suggestForFiles(gc *gocaser, files []string) {
	fset := token.NewFileSet()
	var decls []decl
	for _, fname := range files {
		f, err := parser.ParseFile(fset, fname, nil, 0)
		if err != nil {
			fmt.Printf("failed parsing: %s.\n", err)
			continue
		}

		ast.Inspect(f, func(n ast.Node) bool {
			switch n := n.(type) {
			case *ast.FuncDecl:
				fd := ast.FuncDecl(*n)
				decls = append(decls, decl{
					name:       fd.Name.Name,
					suggestion: gc.GoCase(fd.Name.Name),
					pos:        fset.Position(fd.Name.Pos()),
					kind:       ast.Fun,
				})
			}
			return true
		})
	}

	print(decls)
}

func SuggestForPkg(pkg string, wordsFiles []string) {
	gc := NewGocaser(wordsFiles)
	files := pkgFiles(pkg)
	if files != nil {
		suggestForFiles(gc, files)
	}
}

//XXX: couldn't find anything in go/ast|build|parser
func pkgFiles(pkgPath string) []string {
	buildCtxt := build.Default
	buildPkg, err := buildCtxt.Import(pkgPath, ".", 0)
	if err != nil {
		fmt.Errorf("%v", err)
		return nil
	}

	var files []string
	for _, name := range buildPkg.GoFiles {
		files = append(files, filepath.Join(buildPkg.Dir, name))
	}

	return files
}

func SuggestForDir(path string, wordsFiles []string) {
	gc := NewGocaser(wordsFiles)
	dfiles, err := ioutil.ReadDir(path)
	if err != nil {
		return
	}

	var files []string
	for _, file := range dfiles {
		if strings.HasSuffix(file.Name(), ".go") {
			files = append(files, filepath.Join(path, file.Name()))
		}
	}
	suggestForFiles(gc, files)
}
