# Binary Search Tree

[![Documentation](https://godoc.org/github.com/marselester/binary-search-tree?status.svg)](https://godoc.org/github.com/marselester/binary-search-tree)
[![Go Report Card](https://goreportcard.com/badge/github.com/marselester/binary-search-tree)](https://goreportcard.com/report/github.com/marselester/binary-search-tree)

Binary Search Tree (BST) is a tree structure where key in a node is larger than the keys in all its **left** children
and smaller than the keys in **right** children. New nodes are attached at the bottom of the tree.

BST can provide fast search/inserts, and you can get all the keys in order:

1. get all left children (smaller)
2. get the current key
3. get all right children (larger).

Those properties can be leveraged in [HastyDB](https://github.com/marselester/hastydb) to implement a memtable,
though in practise it's better to use a self-balancing tree such as Red-Black tree.

The running time depends on the shape of a tree which depends on the order in which keys were inserted.
Lg N is the best case (tree is balanced), e.g., when keys inserted in a random order, so there are no long paths.
N in the worst case (N nodes in a search path). The worst case scenario happens when keys are inserted in a sorted order.

## Usage Example

```go
package main

import (
	"fmt"

	"github.com/marselester/binary-search-tree"
)

func main() {
	tree := bst.Tree{}
	tree.Set("name", []byte("Bob"))
	fmt.Printf("%s", tree.Get("name"))
}
```
