package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"runtime"
)

const src = `
package main

func main() {

}`

func main() {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "", src, 0)
	if err != nil {
		panic(err)
	}

	arch := runtime.GOARCH

	if arch == "arm64" {
		emitForArm64(f)
	}
	if arch == "amd64" {
		emitForAmd64(f)
	}
}

func emitForArm64(f *ast.File) {
	// ...
	for _, decl := range f.Decls {
		switch decl := decl.(type) {
		case *ast.FuncDecl:
			if decl.Name.Name == "main" {
				fmt.Printf(".globl _main\n")
				fmt.Printf("_main:\n")
				fmt.Printf("  stp x29, x30, [sp, #-16]!\n")
				fmt.Printf("  mov x29, sp\n")

				fmt.Printf("  ldp x29, x30, [sp], #16\n")
				fmt.Printf("  ret\n")
			}
		}
	}
}

func emitForAmd64(f *ast.File) {
	// ...
}
