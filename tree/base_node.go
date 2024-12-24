package tree

import (
	"strconv"
	"strings"

	"github.com/PlayerR9/mygo-data/errors"
)

// BaseNode is a node in the tree.
type BaseNode struct {
	// Parent, NextSibling, PrevSibling, FirstChild, and LastChild are pointers to
	// other nodes in the tree.
	Parent, NextSibling, PrevSibling, FirstChild, LastChild *BaseNode

	// Type is the type of the node.
	Type string

	// Data is the data associated with the node.
	Data string
}

// String implements Node.
func (n BaseNode) String() string {
	var builder strings.Builder

	_, _ = builder.WriteString("Node[")
	_, _ = builder.WriteString(n.Type)

	if n.Data != "" {
		_, _ = builder.WriteString(" (")
		_, _ = builder.WriteString(strconv.Quote(n.Data))
		_, _ = builder.WriteRune(')')
	}

	_, _ = builder.WriteRune(']')

	str := builder.String()
	return str
}

// NewBaseNode creates a new Node with the specified type and data.
//
// Parameters:
//   - type_: The type of the node.
//   - data: The data associated with the node.
//
// Returns:
//   - *BaseNode: A pointer to the newly created Node. Never returns nil.
func NewBaseNode(type_, data string) *BaseNode {
	n := &BaseNode{
		Type: type_,
		Data: data,
	}

	return n
}

// PrependChild prepends the given child to the node's children.
//
// Parameters:
//   - child: The child to be prepended.
//
// Returns:
//   - error: An error if the receiver is nil.
//
// Errors:
//   - common.ErrNilReceiver: If the receiver is nil.
func (n *BaseNode) PrependChild(child *BaseNode) error {
	if n == nil {
		return errors.ErrNilReceiver
	}

	if child == nil {
		return nil
	}

	child.Parent = n

	child.PrevSibling = nil

	if n.FirstChild == nil {
		child.NextSibling = nil

		n.LastChild = child
	} else {
		child.NextSibling = n.FirstChild
		n.FirstChild.PrevSibling = child
	}

	n.FirstChild = child

	return nil
}

// AppendChild appends the given child to the node's children.
//
// Parameters:
//   - child: The child to be appended.
//
// Returns:
//   - error: An error if the receiver is nil.
//
// Errors:
//   - common.ErrNilReceiver: If the receiver is nil.
func (n *BaseNode) AppendChild(child *BaseNode) error {
	if n == nil {
		return errors.ErrNilReceiver
	}

	if child == nil {
		return nil
	}

	child.Parent = n

	child.NextSibling = nil

	if n.LastChild == nil {
		child.PrevSibling = nil

		n.FirstChild = child
	} else {
		n.LastChild.NextSibling = child
		child.PrevSibling = n.LastChild
	}

	n.LastChild = child

	return nil
}

// Children returns a slice of pointers to the node's children.
//
// Returns:
//   - []*BaseNode: A slice containing pointers to the children of the node.
//     If the node has no children, returns nil.
func (n BaseNode) Children() []*BaseNode {
	if n.FirstChild == nil {
		return nil
	}

	var children []*BaseNode

	for child := n.FirstChild; child != nil; child = child.NextSibling {
		children = append(children, child)
	}

	return children
}
