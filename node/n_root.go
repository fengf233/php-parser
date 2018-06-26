package node

import (
	"github.com/z7zmey/php-parser/comment"
	"github.com/z7zmey/php-parser/position"
	"github.com/z7zmey/php-parser/walker"
)

// Root node
type Root struct {
	Comments []*comment.Comment
	Position *position.Position
	Stmts    []Node
}

// NewRoot node constructor
func NewRoot(Stmts []Node) *Root {
	return &Root{
		Stmts: Stmts,
	}
}

// SetPosition sets node position
func (n *Root) SetPosition(p *position.Position) {
	n.Position = p
}

// GetPosition returns node positions
func (n *Root) GetPosition() *position.Position {
	return n.Position
}

func (n *Root) AddComments(cc []*comment.Comment, tn comment.TokenName) {
	for _, c := range cc {
		c.SetTokenName(tn)
	}
	n.Comments = append(n.Comments, cc...)
}

func (n *Root) GetComments() []*comment.Comment {
	return n.Comments
}

// Attributes returns node attributes as map
func (n *Root) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *Root) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Stmts != nil {
		v.EnterChildList("Stmts", n)
		for _, nn := range n.Stmts {
			if nn != nil {
				nn.Walk(v)
			}
		}
		v.LeaveChildList("Stmts", n)
	}

	v.LeaveNode(n)
}