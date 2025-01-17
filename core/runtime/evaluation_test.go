package runtime

import (
	"testing"

	"github.com/AmjedChakhis/GoLang-interpreter/core/lexer"
	"github.com/AmjedChakhis/GoLang-interpreter/core/parser/parserImpl"
	"github.com/AmjedChakhis/GoLang-interpreter/core/types"
)

func TestDefEval(t *testing.T) {
	input := defEval
	for _, test := range input {
		evaluated := getEvaluated(test.input)
		testIntegerObject(t, evaluated, test.expected)
	}
}

func TestIntegerEval(t *testing.T) {
	testData := intTestData
	for _, test := range testData {
		evaluated := getEvaluated(test.input)
		testIntegerObject(t, evaluated, test.expected)
	}
}

func TestBooleanEval(t *testing.T) {
	testData := boolInputData

	for _, test := range testData {
		evaluated := getEvaluated(test.input)
		testBooleanObject(t, evaluated, test.expected)
	}
}

func TestIfElseEval(t *testing.T) {
	input := elseIfEvalTestData
	evaluated := getEvaluated(input)
	intObj, ok := evaluated.(*types.Integer)
	if !ok {
		t.Fatalf("the obj is not of type types.Integer, instead got %T",
			evaluated,
		)
	}
	if intObj.Val != 10 {
		t.Fatalf("the value of the integer object is not valid expected :%d and got %d", 10, intObj.Val)
	}
}

func TestClosure(t *testing.T) {
	input := closuresTests
	evaluated := getEvaluated(input)

	intObj, ok := evaluated.(*types.Integer)
	if !ok {
		t.Fatalf("the evalued object is not of type *types.Integer, instead go %T", evaluated)
	}

	if intObj.Val != 777 {
		t.Fatalf("the value of intObj is no as expected. expected %d instead got %d", 777, intObj.Val)
	}
}

// ------------- TEST HELPERS  --------------
func testBooleanObject(t *testing.T, evaluated types.ObjectJIPL, expected bool) {
	boolObj, ok := evaluated.(*types.Boolean)
	if !ok {
		t.Fatalf("the obj is not of type types.Boolean, instead got %T",
			evaluated,
		)
	}
	if boolObj.Val != expected {
		t.Fatalf("the value of the boolean object is not valid expected :%t and got %t", expected, boolObj.Val)
	}
}
func getEvaluated(input string) types.ObjectJIPL {
	l := lexer.InitLexer(input)
	p := parser.InitParser(l)
	program := p.Parse()
	ctx := types.NewContext()
	ev, _ := Eval(program, ctx)
	return ev
}

func testIntegerObject(t *testing.T, obj types.ObjectJIPL, expected int) {
	intObj, ok := obj.(*types.Integer)
	if !ok {
		t.Fatalf("the obj is not of type types.Integer, instead got %T",
			obj,
		)
	}
	if intObj.Val != expected {
		t.Fatalf("the value of the integer object is not valid expected :%d and got %d", expected, intObj.Val)
	}
}

// --- TESTS DATA ---
var (
	boolInputData = []struct {
		input    string
		expected bool
	}{
		{"true;", true},
		{"false;", false},
		{"1 < 2;", true},
		{"1 > 2;", false},
		{"1 == 2;", false},
		{"1 != 2;", true},
		{"true == true;", true},
		{"true != true;", false},
		{"true == false;", false},
		{"!true;", false},
		{"!false;", true},
	}
	intTestData = []struct {
		input    string
		expected int
	}{
		{"4545;", 4545},
		{"7;", 7},
	}

	defEval = []struct {
		input    string
		expected int
	}{
		{"def var1 = 0; var1;", 0},
		{"def var2 = 1; var2;", 1},
		{"def var3 = 5 * 4; var3;", 20},
		{"def var4 = 1;def var5 =2; var4+var5;", 3},
		{"def var6=1; def var7=7+var6; var7;", 8},
	}

	elseIfEvalTestData = `
	if (10 > 5) {
		10;
	} else {
		5;
	}
	`
	returnEvalTestData = "return 10;5454447;"

	closuresTests = `
	function outer() {
			out("outer function called");
			def a=777;
			function inner() {
					out("inner function called");
					return a;
			}
			return inner;
	}
	def fn = outer();
	fn();
	`
)
