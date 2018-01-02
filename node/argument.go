package node

type Argument struct {
	name       string
	attributes map[string]interface{}
	position   *Position
	expr       Node
	variadic   bool
}

func NewArgument(expression Node, variadic bool) Node {
	return Argument{
		"Argument",
		map[string]interface{}{
			"variadic": variadic,
		},
		nil,
		expression,
		variadic,
	}
}

func (n Argument) Name() string {
	return "Argument"
}

func (n Argument) Attributes() map[string]interface{} {
	return n.attributes
}

func (n Argument) Attribute(key string) interface{} {
	return n.attributes[key]
}

func (n Argument) SetAttribute(key string, value interface{}) {
	n.attributes[key] = value
}

func (n Argument) Position() *Position {
	return n.position
}

func (n Argument) SetPosition(p *Position) Node {
	n.position = p
	return n
}

func (n Argument) Walk(v Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.expr != nil {
		vv := v.GetChildrenVisitor("expr")
		n.expr.Walk(vv)
	}

	v.LeaveNode(n)
}