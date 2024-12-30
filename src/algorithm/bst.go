package algorithm

import (
	"errors"
	"sync"

	err_msg "github.com/elecbug/go-graphtric/err/bst"
)

type node struct {
	key   string
	left  *node
	right *node
}

type BST struct {
	root *node
	mu   sync.Mutex
}

func (bst *BST) Insert(key string) error {
	bst.mu.Lock()
	defer bst.mu.Unlock()

	if bst.root == nil {
		bst.root = &node{key: key}

		return nil
	}

	temp := bst.root

	for {
		if temp.key == key {
			return errors.New(err_msg.KeyAlreadyExist(key))
		} else if temp.key > key {
			if temp.left == nil {
				temp.left = &node{key: key}

				return nil
			} else {
				temp = temp.left
			}
		} else {
			if temp.right == nil {
				temp.right = &node{key: key}

				return nil
			} else {
				temp = temp.right
			}
		}
	}
}

func (bst *BST) Delete(key string) error {
	bst.mu.Lock()
	defer bst.mu.Unlock()

	var err error

	bst.root, err = deleteNode(bst.root, key)

	return err
}

func (bst *BST) Balance() {
	bst.mu.Lock()
	defer bst.mu.Unlock()

	keys := []string{}

	inOrderTraversalToSlice(bst.root, &keys)
	bst.root = buildBalancedTree(keys, 0, len(keys)-1)
}

func findMin(n *node) *node {
	for n.left != nil {
		n = n.left
	}

	return n
}

func deleteNode(current *node, key string) (*node, error) {
	if current == nil {
		return nil, errors.New(err_msg.KeyDoNotExist(key))
	}

	if key < current.key {
		current.left, _ = deleteNode(current.left, key)
	} else if key > current.key {
		current.right, _ = deleteNode(current.right, key)
	} else {
		if current.left == nil && current.right == nil {
			return nil, nil
		} else if current.left == nil {
			return current.right, nil
		} else if current.right == nil {
			return current.left, nil
		} else {
			successor := findMin(current.right)
			current.key = successor.key
			current.right, _ = deleteNode(current.right, successor.key)
		}
	}

	return current, nil
}

func inOrderTraversalToSlice(n *node, keys *[]string) {
	if n == nil {
		return
	}

	inOrderTraversalToSlice(n.left, keys)
	*keys = append(*keys, n.key)

	inOrderTraversalToSlice(n.right, keys)
}

func buildBalancedTree(keys []string, low, high int) *node {
	if low > high {
		return nil
	}

	mid := (low + high) / 2
	root := &node{key: keys[mid]}
	root.left = buildBalancedTree(keys, low, mid-1)
	root.right = buildBalancedTree(keys, mid+1, high)

	return root
}
