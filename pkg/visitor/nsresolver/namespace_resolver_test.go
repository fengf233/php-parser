package nsresolver_test

import (
	"github.com/z7zmey/php-parser/pkg/visitor/nsresolver"
	"github.com/z7zmey/php-parser/pkg/visitor/traverser"
	"testing"

	"gotest.tools/assert"

	"github.com/z7zmey/php-parser/pkg/ast"
)

func TestResolveStaticCall(t *testing.T) {
	nameAB := &ast.NameName{Parts: []ast.Vertex{&ast.NameNamePart{Value: []byte("A")}, &ast.NameNamePart{Value: []byte("B")}}}
	nameBC := &ast.NameName{Parts: []ast.Vertex{&ast.NameNamePart{Value: []byte("B")}, &ast.NameNamePart{Value: []byte("C")}}}

	stxTree := &ast.StmtStmtList{
		Stmts: []ast.Vertex{
			&ast.StmtUse{
				UseDeclarations: []ast.Vertex{
					&ast.StmtUseDeclaration{
						Use: nameAB,
					},
				},
			},
			&ast.ExprStaticCall{
				Class: nameBC,
				Call:  &ast.Identifier{Value: []byte("foo")},
			},
		},
	}

	expected := map[ast.Vertex]string{
		nameBC: "A\\B\\C",
	}

	nsResolver := nsresolver.NewNamespaceResolver()
	traverser.NewTraverser(nsResolver).Traverse(stxTree)

	assert.DeepEqual(t, expected, nsResolver.ResolvedNames)
}

func TestResolveStaticPropertyFetch(t *testing.T) {
	nameAB := &ast.NameName{Parts: []ast.Vertex{&ast.NameNamePart{Value: []byte("A")}, &ast.NameNamePart{Value: []byte("B")}}}
	nameBC := &ast.NameName{Parts: []ast.Vertex{&ast.NameNamePart{Value: []byte("B")}, &ast.NameNamePart{Value: []byte("C")}}}

	stxTree := &ast.StmtStmtList{
		Stmts: []ast.Vertex{
			&ast.StmtUse{
				UseDeclarations: []ast.Vertex{
					&ast.StmtUseDeclaration{
						Use: nameAB,
					},
				},
			},
			&ast.ExprStaticPropertyFetch{
				Class:    nameBC,
				Property: &ast.Identifier{Value: []byte("foo")},
			},
		},
	}

	expected := map[ast.Vertex]string{
		nameBC: "A\\B\\C",
	}

	nsResolver := nsresolver.NewNamespaceResolver()
	traverser.NewTraverser(nsResolver).Traverse(stxTree)

	assert.DeepEqual(t, expected, nsResolver.ResolvedNames)
}

func TestResolveClassConstFetch(t *testing.T) {
	nameAB := &ast.NameName{Parts: []ast.Vertex{&ast.NameNamePart{Value: []byte("A")}, &ast.NameNamePart{Value: []byte("B")}}}
	nameBC := &ast.NameName{Parts: []ast.Vertex{&ast.NameNamePart{Value: []byte("B")}, &ast.NameNamePart{Value: []byte("C")}}}

	stxTree := &ast.StmtStmtList{
		Stmts: []ast.Vertex{
			&ast.StmtUse{
				UseDeclarations: []ast.Vertex{
					&ast.StmtUseDeclaration{
						Use: nameAB,
					},
				},
			},
			&ast.ExprClassConstFetch{
				Class:        nameBC,
				ConstantName: &ast.Identifier{Value: []byte("FOO")},
			},
		},
	}

	expected := map[ast.Vertex]string{
		nameBC: "A\\B\\C",
	}

	nsResolver := nsresolver.NewNamespaceResolver()
	traverser.NewTraverser(nsResolver).Traverse(stxTree)

	assert.DeepEqual(t, expected, nsResolver.ResolvedNames)
}

func TestResolveNew(t *testing.T) {
	nameAB := &ast.NameName{Parts: []ast.Vertex{&ast.NameNamePart{Value: []byte("A")}, &ast.NameNamePart{Value: []byte("B")}}}
	nameBC := &ast.NameName{Parts: []ast.Vertex{&ast.NameNamePart{Value: []byte("B")}, &ast.NameNamePart{Value: []byte("C")}}}

	stxTree := &ast.StmtStmtList{
		Stmts: []ast.Vertex{
			&ast.StmtUse{
				UseDeclarations: []ast.Vertex{
					&ast.StmtUseDeclaration{
						Use: nameAB,
					},
				},
			},
			&ast.ExprNew{
				Class: nameBC,
			},
		},
	}

	expected := map[ast.Vertex]string{
		nameBC: "A\\B\\C",
	}

	nsResolver := nsresolver.NewNamespaceResolver()
	traverser.NewTraverser(nsResolver).Traverse(stxTree)

	assert.DeepEqual(t, expected, nsResolver.ResolvedNames)
}

func TestResolveInstanceOf(t *testing.T) {
	nameAB := &ast.NameName{Parts: []ast.Vertex{&ast.NameNamePart{Value: []byte("A")}, &ast.NameNamePart{Value: []byte("B")}}}
	nameBC := &ast.NameName{Parts: []ast.Vertex{&ast.NameNamePart{Value: []byte("B")}, &ast.NameNamePart{Value: []byte("C")}}}

	stxTree := &ast.StmtStmtList{
		Stmts: []ast.Vertex{
			&ast.StmtUse{
				UseDeclarations: []ast.Vertex{
					&ast.StmtUseDeclaration{
						Use: nameAB,
					},
				},
			},
			&ast.ExprInstanceOf{
				Expr:  &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("foo")}},
				Class: nameBC,
			},
		},
	}

	expected := map[ast.Vertex]string{
		nameBC: "A\\B\\C",
	}

	nsResolver := nsresolver.NewNamespaceResolver()
	traverser.NewTraverser(nsResolver).Traverse(stxTree)

	assert.DeepEqual(t, expected, nsResolver.ResolvedNames)
}

func TestResolveInstanceCatch(t *testing.T) {
	nameAB := &ast.NameName{Parts: []ast.Vertex{&ast.NameNamePart{Value: []byte("A")}, &ast.NameNamePart{Value: []byte("B")}}}
	nameBC := &ast.NameName{Parts: []ast.Vertex{&ast.NameNamePart{Value: []byte("B")}, &ast.NameNamePart{Value: []byte("C")}}}

	nameDE := &ast.NameName{Parts: []ast.Vertex{&ast.NameNamePart{Value: []byte("D")}, &ast.NameNamePart{Value: []byte("E")}}}
	nameF := &ast.NameName{Parts: []ast.Vertex{&ast.NameNamePart{Value: []byte("F")}}}

	stxTree := &ast.StmtStmtList{
		Stmts: []ast.Vertex{
			&ast.StmtUse{
				UseDeclarations: []ast.Vertex{
					&ast.StmtUseDeclaration{
						Use: nameAB,
					},
					&ast.StmtUseDeclaration{
						Use:   nameDE,
						Alias: &ast.Identifier{Value: []byte("F")},
					},
				},
			},
			&ast.StmtTry{
				Stmts: []ast.Vertex{},
				Catches: []ast.Vertex{
					&ast.StmtCatch{
						Types: []ast.Vertex{
							nameBC,
							nameF,
						},
						Var:   &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("foo")}},
						Stmts: []ast.Vertex{},
					},
				},
			},
		},
	}

	expected := map[ast.Vertex]string{
		nameBC: "A\\B\\C",
		nameF:  "D\\E",
	}

	nsResolver := nsresolver.NewNamespaceResolver()
	traverser.NewTraverser(nsResolver).Traverse(stxTree)

	assert.DeepEqual(t, expected, nsResolver.ResolvedNames)
}

func TestResolveFunctionCall(t *testing.T) {
	nameAB := &ast.NameName{Parts: []ast.Vertex{&ast.NameNamePart{Value: []byte("A")}, &ast.NameNamePart{Value: []byte("B")}}}
	nameB := &ast.NameName{Parts: []ast.Vertex{&ast.NameNamePart{Value: []byte("B")}}}

	stxTree := &ast.StmtStmtList{
		Stmts: []ast.Vertex{
			&ast.StmtUse{
				Type: &ast.Identifier{Value: []byte("function")},
				UseDeclarations: []ast.Vertex{
					&ast.StmtUseDeclaration{
						Use: nameAB,
					},
				},
			},
			&ast.ExprFunctionCall{
				Function: nameB,
			},
		},
	}

	expected := map[ast.Vertex]string{
		nameB: "A\\B",
	}

	nsResolver := nsresolver.NewNamespaceResolver()
	traverser.NewTraverser(nsResolver).Traverse(stxTree)

	assert.DeepEqual(t, expected, nsResolver.ResolvedNames)
}

func TestResolveConstFetch(t *testing.T) {
	nameAB := &ast.NameName{Parts: []ast.Vertex{&ast.NameNamePart{Value: []byte("A")}, &ast.NameNamePart{Value: []byte("B")}}}
	nameB := &ast.NameName{Parts: []ast.Vertex{&ast.NameNamePart{Value: []byte("B")}}}

	stxTree := &ast.StmtStmtList{
		Stmts: []ast.Vertex{
			&ast.StmtUse{
				Type: &ast.Identifier{Value: []byte("const")},
				UseDeclarations: []ast.Vertex{
					&ast.StmtUseDeclaration{
						Use: nameAB,
					},
				},
			},
			&ast.ExprConstFetch{
				Const: nameB,
			},
		},
	}

	expected := map[ast.Vertex]string{
		nameB: "A\\B",
	}

	nsResolver := nsresolver.NewNamespaceResolver()
	traverser.NewTraverser(nsResolver).Traverse(stxTree)

	assert.DeepEqual(t, expected, nsResolver.ResolvedNames)
}

func TestResolveGroupUse(t *testing.T) {
	nameAB := &ast.NameName{Parts: []ast.Vertex{&ast.NameNamePart{Value: []byte("A")}, &ast.NameNamePart{Value: []byte("B")}}}
	nameBD := &ast.NameName{Parts: []ast.Vertex{&ast.NameNamePart{Value: []byte("B")}, &ast.NameNamePart{Value: []byte("D")}}}
	nameE := &ast.NameName{Parts: []ast.Vertex{&ast.NameNamePart{Value: []byte("E")}}}
	nameC := &ast.NameName{Parts: []ast.Vertex{&ast.NameNamePart{Value: []byte("C")}}}
	nameF := &ast.NameName{Parts: []ast.Vertex{&ast.NameNamePart{Value: []byte("F")}}}

	stxTree := &ast.StmtStmtList{
		Stmts: []ast.Vertex{
			&ast.StmtGroupUse{
				Prefix: nameAB,
				UseDeclarations: []ast.Vertex{
					&ast.StmtUseDeclaration{
						Type: &ast.Identifier{Value: []byte("Function")},
						Use:  nameF,
					},
					&ast.StmtUseDeclaration{
						Type: &ast.Identifier{Value: []byte("const")},
						Use:  nameC,
					},
				},
			},
			&ast.StmtGroupUse{
				Prefix: nameBD,
				Type:   &ast.Identifier{Value: []byte("Function")},
				UseDeclarations: []ast.Vertex{
					&ast.StmtUseDeclaration{
						Use: nameE,
					},
				},
			},
			&ast.ExprConstFetch{
				Const: nameC,
			},
			&ast.ExprFunctionCall{
				Function: nameF,
			},
			&ast.ExprFunctionCall{
				Function: nameE,
			},
		},
	}

	expected := map[ast.Vertex]string{
		nameC: "A\\B\\C",
		nameF: "A\\B\\F",
		nameE: "B\\D\\E",
	}

	nsResolver := nsresolver.NewNamespaceResolver()
	traverser.NewTraverser(nsResolver).Traverse(stxTree)

	assert.DeepEqual(t, expected, nsResolver.ResolvedNames)
}

func TestResolveTraitUse(t *testing.T) {
	nameAB := &ast.NameName{Parts: []ast.Vertex{&ast.NameNamePart{Value: []byte("A")}, &ast.NameNamePart{Value: []byte("B")}}}
	nameB := &ast.NameName{Parts: []ast.Vertex{&ast.NameNamePart{Value: []byte("B")}}}
	nameD := &ast.NameName{Parts: []ast.Vertex{&ast.NameNamePart{Value: []byte("D")}}}

	fullyQualifiedNameB := &ast.NameFullyQualified{Parts: []ast.Vertex{&ast.NameNamePart{Value: []byte("B")}}}
	fullyQualifiedNameBC := &ast.NameFullyQualified{Parts: []ast.Vertex{&ast.NameNamePart{Value: []byte("B")}, &ast.NameNamePart{Value: []byte("C")}}}
	relativeNameB := &ast.NameRelative{Parts: []ast.Vertex{&ast.NameNamePart{Value: []byte("B")}}}
	relativeNameBC := &ast.NameRelative{Parts: []ast.Vertex{&ast.NameNamePart{Value: []byte("B")}, &ast.NameNamePart{Value: []byte("C")}}}

	stxTree := &ast.StmtStmtList{
		Stmts: []ast.Vertex{
			&ast.StmtUse{
				UseDeclarations: []ast.Vertex{
					&ast.StmtUseDeclaration{
						Use: nameAB,
					},
				},
			},
			&ast.StmtTraitUse{
				Traits: []ast.Vertex{
					nameB,
					relativeNameB,
				},
				Adaptations: []ast.Vertex{
					&ast.StmtTraitUsePrecedence{
						Trait:     fullyQualifiedNameB,
						Method:    &ast.Identifier{Value: []byte("foo")},
						Insteadof: []ast.Vertex{fullyQualifiedNameBC},
					},
					&ast.StmtTraitUseAlias{
						Trait:  relativeNameBC,
						Method: &ast.Identifier{Value: []byte("foo")},
						Alias:  &ast.Identifier{Value: []byte("bar")},
					},
				},
			},
			&ast.StmtTraitUse{
				Traits: []ast.Vertex{
					nameD,
				},
			},
		},
	}

	expected := map[ast.Vertex]string{
		nameB:                "A\\B",
		nameD:                "D",
		relativeNameB:        "B",
		fullyQualifiedNameB:  "B",
		fullyQualifiedNameBC: "B\\C",
		relativeNameBC:       "B\\C",
	}

	nsResolver := nsresolver.NewNamespaceResolver()
	traverser.NewTraverser(nsResolver).Traverse(stxTree)

	assert.DeepEqual(t, expected, nsResolver.ResolvedNames)
}

func TestResolveClassName(t *testing.T) {
	nameAB := &ast.NameName{Parts: []ast.Vertex{&ast.NameNamePart{Value: []byte("A")}, &ast.NameNamePart{Value: []byte("B")}}}
	nameBC := &ast.NameName{Parts: []ast.Vertex{&ast.NameNamePart{Value: []byte("B")}, &ast.NameNamePart{Value: []byte("C")}}}

	class := &ast.StmtClass{
		ClassName: &ast.Identifier{Value: []byte("A")},
		Extends:   nameAB,
		Implements: []ast.Vertex{
			nameBC,
		},
	}

	stxTree := &ast.StmtStmtList{
		Stmts: []ast.Vertex{
			class,
		},
	}

	expected := map[ast.Vertex]string{
		class:  "A",
		nameAB: "A\\B",
		nameBC: "B\\C",
	}

	nsResolver := nsresolver.NewNamespaceResolver()
	traverser.NewTraverser(nsResolver).Traverse(stxTree)

	assert.DeepEqual(t, expected, nsResolver.ResolvedNames)
}

func TestResolveInterfaceName(t *testing.T) {
	nameAB := &ast.NameName{Parts: []ast.Vertex{&ast.NameNamePart{Value: []byte("A")}, &ast.NameNamePart{Value: []byte("B")}}}
	nameBC := &ast.NameName{Parts: []ast.Vertex{&ast.NameNamePart{Value: []byte("B")}, &ast.NameNamePart{Value: []byte("C")}}}

	interfaceNode := &ast.StmtInterface{
		InterfaceName: &ast.Identifier{Value: []byte("A")},
		Extends: []ast.Vertex{
			nameAB,
			nameBC,
		},
	}

	stxTree := &ast.StmtStmtList{
		Stmts: []ast.Vertex{
			interfaceNode,
		},
	}

	expected := map[ast.Vertex]string{
		interfaceNode: "A",
		nameAB:        "A\\B",
		nameBC:        "B\\C",
	}

	nsResolver := nsresolver.NewNamespaceResolver()
	traverser.NewTraverser(nsResolver).Traverse(stxTree)

	assert.DeepEqual(t, expected, nsResolver.ResolvedNames)
}

func TestResolveTraitName(t *testing.T) {
	traitNode := &ast.StmtTrait{
		TraitName: &ast.Identifier{Value: []byte("A")},
		Stmts:     []ast.Vertex{},
	}

	stxTree := &ast.StmtStmtList{
		Stmts: []ast.Vertex{
			traitNode,
		},
	}

	expected := map[ast.Vertex]string{
		traitNode: "A",
	}

	nsResolver := nsresolver.NewNamespaceResolver()
	traverser.NewTraverser(nsResolver).Traverse(stxTree)

	assert.DeepEqual(t, expected, nsResolver.ResolvedNames)
}

func TestResolveFunctionName(t *testing.T) {
	nameAB := &ast.NameName{Parts: []ast.Vertex{&ast.NameNamePart{Value: []byte("A")}, &ast.NameNamePart{Value: []byte("B")}}}
	nameBC := &ast.NameName{Parts: []ast.Vertex{&ast.NameNamePart{Value: []byte("B")}, &ast.NameNamePart{Value: []byte("C")}}}

	functionNode := &ast.StmtFunction{
		FunctionName: &ast.Identifier{Value: []byte("A")},
		Params: []ast.Vertex{
			&ast.Parameter{
				Type: nameAB,
				Var:  &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("foo")}},
			},
		},
		ReturnType: &ast.Nullable{Expr: nameBC},
		Stmts:      []ast.Vertex{},
	}

	stxTree := &ast.StmtStmtList{
		Stmts: []ast.Vertex{
			functionNode,
		},
	}

	expected := map[ast.Vertex]string{
		functionNode: "A",
		nameAB:       "A\\B",
		nameBC:       "B\\C",
	}

	nsResolver := nsresolver.NewNamespaceResolver()
	traverser.NewTraverser(nsResolver).Traverse(stxTree)

	assert.DeepEqual(t, expected, nsResolver.ResolvedNames)
}

func TestResolveMethodName(t *testing.T) {
	nameAB := &ast.NameName{Parts: []ast.Vertex{&ast.NameNamePart{Value: []byte("A")}, &ast.NameNamePart{Value: []byte("B")}}}
	nameBC := &ast.NameName{Parts: []ast.Vertex{&ast.NameNamePart{Value: []byte("B")}, &ast.NameNamePart{Value: []byte("C")}}}

	methodNode := &ast.StmtClassMethod{
		MethodName: &ast.Identifier{Value: []byte("A")},
		Params: []ast.Vertex{
			&ast.Parameter{
				Type: nameAB,
				Var:  &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("foo")}},
			},
		},
		ReturnType: &ast.Nullable{Expr: nameBC},
		Stmt: &ast.StmtStmtList{
			Stmts: []ast.Vertex{},
		},
	}

	expected := map[ast.Vertex]string{
		nameAB: "A\\B",
		nameBC: "B\\C",
	}

	nsResolver := nsresolver.NewNamespaceResolver()
	traverser.NewTraverser(nsResolver).Traverse(methodNode)

	assert.DeepEqual(t, expected, nsResolver.ResolvedNames)
}

func TestResolveClosureName(t *testing.T) {
	nameAB := &ast.NameName{Parts: []ast.Vertex{&ast.NameNamePart{Value: []byte("A")}, &ast.NameNamePart{Value: []byte("B")}}}
	nameBC := &ast.NameName{Parts: []ast.Vertex{&ast.NameNamePart{Value: []byte("B")}, &ast.NameNamePart{Value: []byte("C")}}}

	closureNode := &ast.ExprClosure{
		Params: []ast.Vertex{
			&ast.Parameter{
				Type: nameAB,
				Var:  &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("foo")}},
			},
		},
		ReturnType: &ast.Nullable{Expr: nameBC},
		Stmts:      []ast.Vertex{},
	}

	expected := map[ast.Vertex]string{
		nameAB: "A\\B",
		nameBC: "B\\C",
	}

	nsResolver := nsresolver.NewNamespaceResolver()
	traverser.NewTraverser(nsResolver).Traverse(closureNode)

	assert.DeepEqual(t, expected, nsResolver.ResolvedNames)
}

func TestResolveConstantsName(t *testing.T) {
	nameAB := &ast.NameName{Parts: []ast.Vertex{&ast.NameNamePart{Value: []byte("A")}, &ast.NameNamePart{Value: []byte("B")}}}

	constantB := &ast.StmtConstant{
		Name: &ast.Identifier{Value: []byte("B")},
		Expr: &ast.ScalarLnumber{Value: []byte("1")},
	}
	constantC := &ast.StmtConstant{
		Name: &ast.Identifier{Value: []byte("C")},
		Expr: &ast.ScalarLnumber{Value: []byte("1")},
	}

	stxTree := &ast.StmtStmtList{
		Stmts: []ast.Vertex{
			&ast.StmtNamespace{
				Name: nameAB,
			},
			&ast.StmtConstList{
				Consts: []ast.Vertex{
					constantB,
					constantC,
				},
			},
		},
	}

	expected := map[ast.Vertex]string{
		constantB: "A\\B\\B",
		constantC: "A\\B\\C",
	}

	nsResolver := nsresolver.NewNamespaceResolver()
	traverser.NewTraverser(nsResolver).Traverse(stxTree)

	assert.DeepEqual(t, expected, nsResolver.ResolvedNames)
}

func TestResolveNamespaces(t *testing.T) {
	namespaceAB := &ast.NameName{Parts: []ast.Vertex{&ast.NameNamePart{Value: []byte("A")}, &ast.NameNamePart{Value: []byte("B")}}}
	namespaceCD := &ast.NameName{Parts: []ast.Vertex{&ast.NameNamePart{Value: []byte("C")}, &ast.NameNamePart{Value: []byte("D")}}}

	nameAC := &ast.NameName{Parts: []ast.Vertex{&ast.NameNamePart{Value: []byte("A")}, &ast.NameNamePart{Value: []byte("C")}}}
	nameCF := &ast.NameName{Parts: []ast.Vertex{&ast.NameNamePart{Value: []byte("C")}, &ast.NameNamePart{Value: []byte("F")}}}
	nameFG := &ast.NameName{Parts: []ast.Vertex{&ast.NameNamePart{Value: []byte("F")}, &ast.NameNamePart{Value: []byte("G")}}}
	relativeNameCE := &ast.NameRelative{Parts: []ast.Vertex{&ast.NameNamePart{Value: []byte("C")}, &ast.NameNamePart{Value: []byte("E")}}}

	constantB := &ast.StmtConstant{
		Name: &ast.Identifier{Value: []byte("B")},
		Expr: &ast.ScalarLnumber{Value: []byte("1")},
	}
	constantC := &ast.StmtConstant{
		Name: &ast.Identifier{Value: []byte("C")},
		Expr: &ast.ScalarLnumber{Value: []byte("1")},
	}

	stxTree := &ast.StmtStmtList{
		Stmts: []ast.Vertex{
			&ast.StmtNamespace{
				Name: namespaceAB,
			},
			&ast.StmtConstList{
				Consts: []ast.Vertex{
					constantB,
					constantC,
				},
			},
			&ast.ExprStaticCall{
				Class: nameFG,
				Call:  &ast.Identifier{Value: []byte("foo")},
			},
			&ast.StmtNamespace{
				Stmts: []ast.Vertex{},
			},
			&ast.StmtNamespace{
				Name: namespaceCD,
				Stmts: []ast.Vertex{
					&ast.StmtUse{
						UseDeclarations: []ast.Vertex{
							&ast.StmtUseDeclaration{
								Use: nameAC,
							},
						},
					},
					&ast.ExprStaticCall{
						Class: relativeNameCE,
						Call:  &ast.Identifier{Value: []byte("foo")},
					},
					&ast.ExprStaticCall{
						Class: nameCF,
						Call:  &ast.Identifier{Value: []byte("foo")},
					},
				},
			},
		},
	}

	expected := map[ast.Vertex]string{
		constantB:      "A\\B\\B",
		constantC:      "A\\B\\C",
		nameFG:         "A\\B\\F\\G",
		relativeNameCE: "C\\D\\C\\E",
		nameCF:         "A\\C\\F",
	}

	nsResolver := nsresolver.NewNamespaceResolver()
	traverser.NewTraverser(nsResolver).Traverse(stxTree)

	assert.DeepEqual(t, expected, nsResolver.ResolvedNames)
}

func TestResolveStaticCallDinamicClassName(t *testing.T) {
	stxTree := &ast.StmtStmtList{
		Stmts: []ast.Vertex{
			&ast.ExprStaticCall{
				Class: &ast.ExprVariable{VarName: &ast.Identifier{Value: []byte("foo")}},
				Call:  &ast.Identifier{Value: []byte("foo")},
			},
		},
	}

	expected := map[ast.Vertex]string{}

	nsResolver := nsresolver.NewNamespaceResolver()
	traverser.NewTraverser(nsResolver).Traverse(stxTree)

	assert.DeepEqual(t, expected, nsResolver.ResolvedNames)
}

func TestDoNotResolveReservedConstants(t *testing.T) {
	namespaceName := &ast.NameName{Parts: []ast.Vertex{&ast.NameNamePart{Value: []byte("Foo")}}}

	constantTrue := &ast.NameName{
		Parts: []ast.Vertex{
			&ast.NameNamePart{Value: []byte("True")},
		},
	}

	constantFalse := &ast.NameName{
		Parts: []ast.Vertex{
			&ast.NameNamePart{Value: []byte("False")},
		},
	}

	constantNull := &ast.NameName{
		Parts: []ast.Vertex{
			&ast.NameNamePart{Value: []byte("NULL")},
		},
	}

	stxTree := &ast.StmtStmtList{
		Stmts: []ast.Vertex{
			&ast.StmtNamespace{
				Name: namespaceName,
			},
			&ast.StmtExpression{
				Expr: &ast.ExprConstFetch{
					Const: constantTrue,
				},
			},
			&ast.StmtExpression{
				Expr: &ast.ExprConstFetch{
					Const: constantFalse,
				},
			},
			&ast.StmtExpression{
				Expr: &ast.ExprConstFetch{
					Const: constantNull,
				},
			},
		},
	}

	expected := map[ast.Vertex]string{
		constantTrue:  "true",
		constantFalse: "false",
		constantNull:  "null",
	}

	nsResolver := nsresolver.NewNamespaceResolver()
	traverser.NewTraverser(nsResolver).Traverse(stxTree)

	assert.DeepEqual(t, expected, nsResolver.ResolvedNames)
}

func TestDoNotResolveReservedNames(t *testing.T) {

	nameInt := &ast.NameName{
		Parts: []ast.Vertex{
			&ast.NameNamePart{Value: []byte("int")},
		},
	}

	nameFloat := &ast.NameName{
		Parts: []ast.Vertex{
			&ast.NameNamePart{Value: []byte("float")},
		},
	}

	nameBool := &ast.NameName{
		Parts: []ast.Vertex{
			&ast.NameNamePart{Value: []byte("bool")},
		},
	}

	nameString := &ast.NameName{
		Parts: []ast.Vertex{
			&ast.NameNamePart{Value: []byte("string")},
		},
	}

	nameVoid := &ast.NameName{
		Parts: []ast.Vertex{
			&ast.NameNamePart{Value: []byte("void")},
		},
	}

	nameIterable := &ast.NameName{
		Parts: []ast.Vertex{
			&ast.NameNamePart{Value: []byte("iterable")},
		},
	}

	nameObject := &ast.NameName{
		Parts: []ast.Vertex{
			&ast.NameNamePart{Value: []byte("object")},
		},
	}

	function := &ast.StmtFunction{
		FunctionName: &ast.Identifier{Value: []byte("bar")},
		Params: []ast.Vertex{
			&ast.Parameter{
				Type: nameInt,
				Var: &ast.ExprVariable{
					VarName: &ast.Identifier{Value: []byte("Int")},
				},
			},
			&ast.Parameter{
				Type: nameFloat,
				Var: &ast.ExprVariable{
					VarName: &ast.Identifier{Value: []byte("Float")},
				},
			},
			&ast.Parameter{
				Type: nameBool,
				Var: &ast.ExprVariable{
					VarName: &ast.Identifier{Value: []byte("Bool")},
				},
			},
			&ast.Parameter{
				Type: nameString,
				Var: &ast.ExprVariable{
					VarName: &ast.Identifier{Value: []byte("String")},
				},
			},
			&ast.Parameter{
				Type: nameVoid,
				Var: &ast.ExprVariable{
					VarName: &ast.Identifier{Value: []byte("Void")},
				},
			},
			&ast.Parameter{
				Type: nameIterable,
				Var: &ast.ExprVariable{
					VarName: &ast.Identifier{Value: []byte("Iterable")},
				},
			},
			&ast.Parameter{
				Type: nameObject,
				Var: &ast.ExprVariable{
					VarName: &ast.Identifier{Value: []byte("Object")},
				},
			},
		},
	}

	stxTree := &ast.StmtStmtList{
		Stmts: []ast.Vertex{
			&ast.StmtNamespace{
				Name: &ast.NameName{
					Parts: []ast.Vertex{
						&ast.NameNamePart{Value: []byte("Foo")},
					},
				},
			},
			function,
		},
	}

	expected := map[ast.Vertex]string{
		function:     "Foo\\bar",
		nameInt:      "int",
		nameFloat:    "float",
		nameBool:     "bool",
		nameString:   "string",
		nameVoid:     "void",
		nameIterable: "iterable",
		nameObject:   "object",
	}

	nsResolver := nsresolver.NewNamespaceResolver()
	traverser.NewTraverser(nsResolver).Traverse(stxTree)

	assert.DeepEqual(t, expected, nsResolver.ResolvedNames)
}

func TestDoNotResolveReservedSpecialNames(t *testing.T) {

	nameSelf := &ast.NameName{
		Parts: []ast.Vertex{
			&ast.NameNamePart{Value: []byte("Self")},
		},
	}

	nameStatic := &ast.NameName{
		Parts: []ast.Vertex{
			&ast.NameNamePart{Value: []byte("Static")},
		},
	}

	nameParent := &ast.NameName{
		Parts: []ast.Vertex{
			&ast.NameNamePart{Value: []byte("Parent")},
		},
	}

	cls := &ast.StmtClass{
		ClassName: &ast.Identifier{Value: []byte("Bar")},
		Stmts: []ast.Vertex{
			&ast.StmtExpression{
				Expr: &ast.ExprStaticCall{
					Class: nameSelf,
					Call:  &ast.Identifier{Value: []byte("func")},
				},
			},
			&ast.StmtExpression{
				Expr: &ast.ExprStaticCall{
					Class: nameStatic,
					Call:  &ast.Identifier{Value: []byte("func")},
				},
			},
			&ast.StmtExpression{
				Expr: &ast.ExprStaticCall{
					Class: nameParent,
					Call:  &ast.Identifier{Value: []byte("func")},
				},
			},
		},
	}

	stxTree := &ast.StmtStmtList{
		Stmts: []ast.Vertex{
			&ast.StmtNamespace{
				Name: &ast.NameName{
					Parts: []ast.Vertex{
						&ast.NameNamePart{Value: []byte("Foo")},
					},
				},
			},
			cls,
		},
	}

	expected := map[ast.Vertex]string{
		cls:        "Foo\\Bar",
		nameSelf:   "self",
		nameStatic: "static",
		nameParent: "parent",
	}

	nsResolver := nsresolver.NewNamespaceResolver()
	traverser.NewTraverser(nsResolver).Traverse(stxTree)

	assert.DeepEqual(t, expected, nsResolver.ResolvedNames)
}
func TestResolvePropertyTypeName(t *testing.T) {
	nameSimple := &ast.NameName{Parts: []ast.Vertex{&ast.NameNamePart{Value: []byte("A")}, &ast.NameNamePart{Value: []byte("B")}}}
	nameRelative := &ast.NameRelative{Parts: []ast.Vertex{&ast.NameNamePart{Value: []byte("A")}, &ast.NameNamePart{Value: []byte("B")}}}
	nameFullyQualified := &ast.NameFullyQualified{Parts: []ast.Vertex{&ast.NameNamePart{Value: []byte("A")}, &ast.NameNamePart{Value: []byte("B")}}}

	propertyNodeSimple := &ast.StmtPropertyList{
		Type: nameSimple,
	}

	propertyNodeRelative := &ast.StmtPropertyList{
		Type: nameRelative,
	}

	propertyNodeFullyQualified := &ast.StmtPropertyList{
		Type: nameFullyQualified,
	}

	classNode := &ast.StmtClass{
		ClassName: &ast.Identifier{Value: []byte("Bar")},
		Stmts: []ast.Vertex{
			propertyNodeSimple,
			propertyNodeRelative,
			propertyNodeFullyQualified,
		},
	}

	stmts := &ast.StmtStmtList{
		Stmts: []ast.Vertex{
			&ast.StmtNamespace{
				Name: &ast.NameName{
					Parts: []ast.Vertex{
						&ast.NameNamePart{Value: []byte("Foo")},
					},
				},
			},
			classNode,
		},
	}

	expected := map[ast.Vertex]string{
		nameSimple:         "Foo\\A\\B",
		nameRelative:       "Foo\\A\\B",
		nameFullyQualified: "A\\B",
		classNode:          "Foo\\Bar",
	}

	nsResolver := nsresolver.NewNamespaceResolver()
	traverser.NewTraverser(nsResolver).Traverse(stmts)

	assert.DeepEqual(t, expected, nsResolver.ResolvedNames)
}