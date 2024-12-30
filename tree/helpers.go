package tree

import (
	common "github.com/PlayerR9/mygo-data/common"
	"github.com/PlayerR9/mygo-data/tree/internal"
)

// PrependChildren prepends the given slice of *Node to the parent *Node's children. Nil elements are rejected.
//
// Parameters:
//   - parent: The parent *Node to prepend the children to.
//   - nodes: The slice of *Node to prepend to the parent's children.
//
// Returns:
//   - error: An error if the parent is nil, or if any of the nodes in the slice
//     are nil.
//
// Errors:
//   - common.ErrBadParam: If the parent is nil.
//
// Behaviors:
//   - There is a side-effect on the nodes: nil nodes are removed and the slice is
//     trimmed.
func PrependChildren(parent *BaseNode, nodes []*BaseNode) error {
	if parent == nil {
		err := common.NewErrBadParam("parent", "is nil")
		return err
	}

	if len(nodes) == 0 {
		return nil
	}

	_ = internal.RejectNils(&nodes)
	if len(nodes) == 0 {
		return nil
	}

	for i := len(nodes) - 1; i >= 0; i-- {
		node := nodes[i]

		_ = parent.PrependChild(node)
	}

	return nil
}

// AppendChildren appends the given slice of *Node to the parent *Node's children. Nil elements are rejected.
//
// Parameters:
//   - parent: The parent *Node to append the children to.
//   - nodes: The slice of *Node to append to the parent's children.
//
// Returns:
//   - error: An error if the parent is nil, or if any of the nodes in the slice
//     are nil.
//
// Errors:
//   - common.ErrBadParam: If the parent is nil.
//
// Behaviors:
//   - There is a side-effect on the nodes: nil nodes are removed and the slice is
//     trimmed.
func AppendChildren(parent *BaseNode, nodes []*BaseNode) error {
	if parent == nil {
		err := common.NewErrBadParam("parent", "is nil")
		return err
	}

	if len(nodes) == 0 {
		return nil
	}

	_ = internal.RejectNils(&nodes)
	if len(nodes) == 0 {
		return nil
	}

	for _, node := range nodes {
		_ = parent.AppendChild(node)
	}

	return nil
}
