package validate

import (
	"github.com/booleangate/sqlc/internal/sql/ast"
	"github.com/booleangate/sqlc/internal/sql/astutils"
	"github.com/booleangate/sqlc/internal/sql/named"
	"github.com/booleangate/sqlc/internal/sql/sqlerr"
)

// A query can use one (and only one) of the following formats:
// - positional parameters           $1
// - named parameter operator        @param
// - named parameter function calls  sqlc.arg(param)
func ParamStyle(n ast.Node) error {
	namedFunc := astutils.Search(n, named.IsParamFunc)
	for _, f := range namedFunc.Items {
		fc, ok := f.(*ast.FuncCall)
		if ok {
			switch val := fc.Args.Items[0].(type) {
			case *ast.FuncCall:
				return &sqlerr.Error{
					Code:     "", // TODO: Pick a new error code
					Message:  "Invalid argument to sqlc.arg()",
					Location: val.Location,
				}
			case *ast.ParamRef:
				return &sqlerr.Error{
					Code:     "", // TODO: Pick a new error code
					Message:  "Invalid argument to sqlc.arg()",
					Location: val.Location,
				}
			case *ast.A_Const, *ast.ColumnRef:
			default:
				return &sqlerr.Error{
					Code:    "", // TODO: Pick a new error code
					Message: "Invalid argument to sqlc.arg()",
				}

			}
		}
	}
	return nil
}
