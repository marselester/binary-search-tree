// Package bst implements a Binary Search Tree (BST) using recursive approach which is limited to a stack size.
// BST is a binary tree where key in a node is larger than the keys in all its left children and
// smaller than the keys in right children. New nodes are attached at the bottom of the tree.
//
// The running time depends on the shape of a tree which depends on the order in which keys were inserted.
// Lg N is the best case (tree is balanced), e.g., when keys inserted in a random order, so there are no long paths.
// N in the worst case (N nodes in a search path). The worst case scenario happens when keys are inserted in a sorted order.
package bst

// Tree represents a Binary Search Tree.
type Tree struct {
	root *node
}

// node is a node of a Binary Search Tree.
type node struct {
	// Key is a unique comparable key, e.g., name.
	Key string
	// Value is a value associated with the key, e.g., Bob.
	Value []byte
	// Left is pointer to the left subtree where smaller keys are stored.
	Left *node
	// Right is pointer to the right subtree where larger keys are stored.
	Right *node
}

// Get retrieves a key from the tree.
func (t *Tree) Get(key string) []byte {
	found := search(key, t.root)
	if found == nil {
		return nil
	}
	return found.Value
}

// Set stores the key in the tree. First it looks up the key and if found, updates the value.
// If the key is new, it will be added to the tree.
func (t *Tree) Set(key string, value []byte) {
	if t.root == nil {
		t.root = &node{Key: key, Value: value}
		return
	}
	put(key, value, t.root)
}

// Keys returns all keys sorted in ascending order.
func (t *Tree) Keys() []string {
	return keys(nil, t.root)
}

// search recursively looks up node by key starting from node n.
func search(key string, n *node) *node {
	switch {
	// Search miss.
	case n == nil:
		return nil
	// Search hit.
	case key == n.Key:
		return n
	// Check smaller keys on the left side.
	case key < n.Key:
		return search(key, n.Left)
	// Check larger keys on the right side.
	case key > n.Key:
		return search(key, n.Right)
	}
	return nil
}

// put updates the value of found node which was looked up by key.
// If key is not found, the new node is added to the tree.
// The node which contains the key is returned.
func put(key string, value []byte, n *node) *node {
	switch {
	// Search miss.
	case n == nil:
		return nil
	// Search hit, update the value.
	case key == n.Key:
		n.Value = value
		return n
	// Check smaller keys on the left side.
	// If nothing is found, insert the new node there.
	case key < n.Key:
		if found := put(key, value, n.Left); found != nil {
			return found
		}
		n.Left = &node{Key: key, Value: value}
		return n.Left
	// Check larger keys on the right side.
	// If nothing is found, insert the new node there.
	case key > n.Key:
		if found := put(key, value, n.Right); found != nil {
			return found
		}
		n.Right = &node{Key: key, Value: value}
		return n.Right
	}
	return nil
}

// keys recursively traverses the tree and returns all keys in order.
func keys(kk []string, n *node) []string {
	if n == nil {
		return kk
	}
	kk = keys(kk, n.Left)
	kk = append(kk, n.Key)
	kk = keys(kk, n.Right)
	return kk
}
