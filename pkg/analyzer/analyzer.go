package analyzer

import (
	"go/ast"
	"go/types"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

func NewAnalyzer() *analysis.Analyzer {
	return &analysis.Analyzer{
		Name: "gomutcheck",
		Doc:  "detects mutations of struct fields in value receiver methods",
		Run:  run,
		Requires: []*analysis.Analyzer{
			inspect.Analyzer,
		},
	}
}

func run(pass *analysis.Pass) (interface{}, error) {
	i := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	// We are only interested in function declarations
	nodeFilter := []ast.Node{
		(*ast.FuncDecl)(nil),
	}

	i.Preorder(nodeFilter, func(n ast.Node) {
		fn, ok := n.(*ast.FuncDecl)
		if !ok || fn.Recv == nil {
			return
		}

		// Check if the receiver is a value type (not a pointer)
		recvType := pass.TypesInfo.TypeOf(fn.Recv.List[0].Type)
		if _, isPointer := recvType.(*types.Pointer); isPointer {
			return
		}

		// Traverse the function body to detect assignments to struct fields
		ast.Inspect(fn.Body, func(n ast.Node) bool {
			assign, ok := n.(*ast.AssignStmt)
			if !ok {
				return true
			}

			// Check each left-hand side expression in the assignment
			for _, lhs := range assign.Lhs {
				selExpr, ok := lhs.(*ast.SelectorExpr)
				if !ok {
					continue
				}

				// Ensure the field being assigned belongs to the receiver
				ident, ok := selExpr.X.(*ast.Ident)
				if !ok || ident.Obj != fn.Recv.List[0].Names[0].Obj {
					continue
				}

				pass.Reportf(assign.Pos(), "struct field '%s' is being mutated in value receiver method", selExpr.Sel.Name)
			}
			return true
		})
	})

	return nil, nil
}
