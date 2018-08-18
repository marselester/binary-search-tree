package bst_test

import (
	"fmt"

	"github.com/marselester/binary-search-tree"
)

func Example() {
	tree := bst.Tree{}
	tree.Set("name", []byte("Bob"))
	fmt.Printf("%s", tree.Get("name"))
	// Output: Bob
}
