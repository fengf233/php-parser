package expr

import (
	"github.com/z7zmey/php-parser/comment"
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/position"
	"github.com/z7zmey/php-parser/walker"
)

// Reference node
type Reference struct {
	Comments []*comment.Comment
	Position *position.Position
	Variable node.Node
}

// NewReference node constructor
func NewReference(Variable node.Node) *Reference {
	return &Reference{
		Variable: Variable,
	}
}

// SetPosition sets node position
func (n *Reference) SetPosition(p *position.Position) {
	n.Position = p
}

// GetPosition returns node positions
func (n *Reference) GetPosition() *position.Position {
	return n.Position
}

func (n *Reference) AddComments(cc []*comment.Comment, tn comment.TokenName) {
	for _, c := range cc {
		c.SetTokenName(tn)
	}
	n.Comments = append(n.Comments, cc...)
}

func (n *Reference) GetComments() []*comment.Comment {
	return n.Comments
}

// Attributes returns node attributes as map
func (n *Reference) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *Reference) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Variable != nil {
		v.EnterChildNode("Variable", n)
		n.Variable.Walk(v)
		v.LeaveChildNode("Variable", n)
	}

	v.LeaveNode(n)
}