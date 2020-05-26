package rel

import (
	"strings"

	"github.com/arr-ai/wbnf/parser"
	"github.com/go-errors/errors"
)

// BytesExpr is an expression that evaluates to a Byte Array
type BytesExpr struct {
	ExprScanner
	elements []Expr
}

func NewBytesExpr(scanner parser.Scanner, elements ...Expr) Expr {
	bytes := make([]byte, 0, len(elements))
	for _, expr := range elements {
		if byteNum, ok := expr.(Number); ok && isByteNumber(byteNum) {
			bytes = append(bytes, byte(int(byteNum)))
			continue
		}
		return BytesExpr{ExprScanner{scanner}, elements}
	}
	return NewBytes(bytes)
}

func (b BytesExpr) Eval(local Scope) (Value, error) {
	bytes := make([]byte, 0)
	for _, expr := range b.elements {
		value, err := expr.Eval(local)
		if err != nil {
			return nil, wrapContext(err, b)
		}
		switch v := value.(type) {
		case Number:
			if !isByteNumber(v) {
				return nil, wrapContext(errors.Errorf("BytesExpr.Eval: Number does not represent a byte: %v", v), b)
			}
			bytes = append(bytes, byte(v))
		case String:
			if err := b.handleOffset(v); err != nil {
				return nil, err
			}
			bytes = append(bytes, []byte(string(v.s))...)
		case GenericSet:
			if s, isString := AsString(v); isString {
				if err := b.handleOffset(s); err != nil {
					return nil, err
				}
				bytes = append(bytes, []byte(string(s.s))...)
				continue
			}
			return nil, wrapContext(errors.Errorf("BytesExpr.Eval: Set %v is not supported", expr), b)
		default:
			return nil, wrapContext(errors.Errorf("BytesExpr.Eval: %T is not supported", v), b)
		}
	}
	return NewBytes(bytes), nil
}

func (b BytesExpr) handleOffset(s String) error {
	if s.offset != 0 {
		return wrapContext(errors.Errorf("BytesExpr.Eval: offsetted String is not supported: %v", s), b)
	}
	return nil
}

func (b BytesExpr) String() string {
	s := strings.Builder{}
	s.WriteString("<<")
	for i, expr := range b.elements {
		if i > 0 {
			s.WriteString(", ")
		}
		s.WriteString(expr.String())
	}
	s.WriteString(">>")
	return s.String()
}

func isByteNumber(n Number) bool {
	return int(n) >= 0 && int(n) <= 255
}