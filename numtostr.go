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
			if isAllNumber(pass, n.Args) {
				pass.Reportf(caller.Pos(), "don't use fmt.Sprint to convert number to string. Use strconv.Itoa.")
			}
		}
	})

	return nil, nil
}

func isAllNumber(pass *analysis.Pass, args []ast.Expr) bool {
	for _, arg := range args {
		v := pass.TypesInfo.TypeOf(arg)
		if !types.Identical(v, types.Typ[types.Int]) {
			return false
		}
	}
	return true
}
