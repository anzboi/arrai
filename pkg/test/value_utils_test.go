package test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/arr-ai/arrai/pkg/arraictx"
	"github.com/arr-ai/arrai/rel"
	"github.com/arr-ai/arrai/syntax"
)

func TestIsLiteralTrue(t *testing.T) { //nolint:dupl
	t.Parallel()

	require.True(t, isLiteralTrue(eval("true")))
	require.True(t, isLiteralTrue(eval("{()}")))

	require.False(t, isLiteralTrue(eval("false")))
	require.False(t, isLiteralTrue(eval("1")))
	require.False(t, isLiteralTrue(eval("0")))
	require.False(t, isLiteralTrue(eval("()")))
	require.False(t, isLiteralTrue(eval("{true}")))
	require.False(t, isLiteralTrue(eval("{false}")))
	require.False(t, isLiteralTrue(eval("{(1)}")))
	require.False(t, isLiteralTrue(eval("{(),1}")))
	require.False(t, isLiteralTrue(eval("[true]")))
	require.False(t, isLiteralTrue(eval("(val: true)")))
}

func TestIsLiteralFalse(t *testing.T) { //nolint:dupl
	t.Parallel()

	require.True(t, isLiteralFalse(eval("false")))
	require.True(t, isLiteralFalse(eval("{}")))

	require.False(t, isLiteralFalse(eval("true")))
	require.False(t, isLiteralFalse(eval("1")))
	require.False(t, isLiteralFalse(eval("0")))
	require.False(t, isLiteralFalse(eval("()")))
	require.False(t, isLiteralFalse(eval("{true}")))
	require.False(t, isLiteralFalse(eval("{false}")))
	require.False(t, isLiteralFalse(eval("{(1)}")))
	require.False(t, isLiteralFalse(eval("{(),1}")))
	require.False(t, isLiteralFalse(eval("[true]")))
	require.False(t, isLiteralFalse(eval("(val: true)")))
}

func TestForeachLeaf(t *testing.T) {
	t.Parallel()

	// No root
	require.Equal(t,
		leavesShouldBe{"<root>": "true"},
		forInput("true"))
	require.Equal(t,
		leavesShouldBe{"<root>": "false"},
		forInput("false"))
	require.Equal(t,
		leavesShouldBe{"<root>": "42"},
		forInput("42"))

	// Tuple root
	require.Equal(t,
		leavesShouldBe{},
		forInput("()"))
	require.Equal(t,
		leavesShouldBe{"a": "true", "b": "false"},
		forInput("(a: true, b: false)"))
	require.Equal(t,
		leavesShouldBe{"a.b.c": "true", "d.e.f": "false"},
		forInput("(a: (b: (c: true)), d: (e: (f: false)))"))
	require.Equal(t,
		leavesShouldBe{"a(0)": "0", "a(1)": "1", "a(2)": "'2'"},
		forInput("(a: [0, 1, '2'])"))
	require.Equal(t,
		leavesShouldBe{"a(0)(0)": "0", "a(0)(1)": "1", "a(1)(0)(0)": "100", "a(1)(1)": "11"},
		forInput("(a: [[0, 1], [[100], 11]])"))
	require.Equal(t,
		leavesShouldBe{"a('k0')": "0", "a('k1')": "1", "a(2)": "2", "a((x: 3))": "3"},
		forInput("(a: {'k0': 0, 'k1': 1, 2: 2, (x: 3):3})"))

	// Array root
	require.Equal(t,
		leavesShouldBe{"<root>": "false"}, // An unfortunate side effect of everything being a set
		forInput("[]"))
	require.Equal(t,
		leavesShouldBe{"(0)": "true", "(1)": "false", "(2).a": "1", "(2).b": "'2'", "(2).c('three')": "3", "(2).c(4)": "'4'"},
		forInput("[true, false, (a: 1, b: '2', c: { 'three': 3, 4: '4' })]"))

	// Dictionary root
	require.Equal(t,
		leavesShouldBe{"(0)": "true", "(1)": "false", "(2).a": "1", "(2).b": "'2'", "(2).c('three')": "3", "(2).c(4)": "'4'"},
		forInput("{0: true, 1: false, 2: (a: 1, b: '2', c: { 'three': 3, 4: '4' })}"))
}

type leavesShouldBe map[string]string

// Runs ForeachLeaf on the result of the arrai source. Returns a dictionary of results (node path -> leaf value) or
// or nil if it failed to parse (instead of err, so it can be inlined)
func forInput(source string) leavesShouldBe {
	leaves := make(leavesShouldBe)
	tree, err := evalErr(source)
	if err != nil {
		return leavesShouldBe{"PARSING ERROR": err.Error()}
	}

	ForeachLeaf(tree, "", func(val rel.Value, path string) {
		valStr := val.String()
		if valStr == "{}" {
			valStr = "false"
		} else if _, ok := val.(rel.String); ok {
			valStr = "'" + valStr + "'"
		}
		leaves[path] = valStr
	})

	return leaves
}

func eval(source string) rel.Value {
	value, _ := evalErr(source) //nolint:errcheck
	return value
}

// Evaluates the arrai source and returns the result, or nil if it failed to parse (instead of err so it can be inlined)
func evalErr(source string) (rel.Value, error) {
	ctx := arraictx.InitRunCtx(context.Background())
	value, err := syntax.EvaluateExpr(ctx, "", source)
	if err != nil {
		return nil, err
	}
	return value, nil
}
