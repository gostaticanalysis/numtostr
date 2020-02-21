package numtostr

import (
	"go/ast"
	"go/types"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var Analyzer = &analysis.Analyzer{
	Name: "numtostr",
	Doc:  Doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

var fmtpkg *types.Package

const Doc = "numstr checks bad conversion number to string"

func init() {
}
func run(pass *analysis.Pass) (interface{}, error) {
	for _, pkg := range pass.Pkg.Imports() {
		if pkg.Path() == "fmt" {
			fmtpkg = pkg
		}
	}
	if fmtpkg == nil {
		return nil, nil
	}

	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	nodeFilter := []ast.Node{
		(*ast.CallExpr)(nil),
	}
	inspect.Preorder(nodeFilter, func(n ast.Node) {
		switch n := n.(type) {
		case *ast.CallExpr:
			// check function is fmt.Sprint
			caller, ok := n.Fun.(*ast.SelectorExpr)
			if !ok {
				return
			}
			if pass.TypesInfo.ObjectOf(caller.Sel).Pkg() != fmtpkg {
				return
			}
			if caller.Sel.Name != "Sprint" {
				return
			}

			// check all arguments are numerical
			numOnly := true
			for _, arg := range n.Args {
				v := pass.TypesInfo.TypeOf(arg)
				switch v.String() {
				case "int", "int8", "int32", "int64":
				case "uint", "uint8", "uint16", "uint32", "uint64":
				case "float32", "float64":
					continue
				default:
					numOnly = false
					break
				}
			}
			if numOnly {
				pass.Reportf(caller.Pos(), "don't use fmt.Sprint to convert number to string. Use strconv.Itoa.")
			}
		}
	})

	return nil, nil
}
