package bst

import (
	"bytes"
	"testing"
)

func equal(s1, s2 []string) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}

func abcTree() *Tree {
	r := node{
		Key:   "S",
		Value: []byte("sea"),
		Left: &node{
			Key:   "E",
			Value: []byte("ear"),
			Left: &node{
				Key:   "A",
				Value: []byte("apple"),
				Right: &node{
					Key:   "C",
					Value: []byte("cat"),
				},
			},
			Right: &node{
				Key:   "R",
				Value: []byte("rock"),
				Left: &node{
					Key:   "H",
					Value: []byte("house"),
					Right: &node{
						Key:   "M",
						Value: []byte("mouse"),
					},
				},
			},
		},
		Right: &node{
			Key: "X",
		},
	}
	return &Tree{root: &r}
}

func TestSearch(t *testing.T) {
	tree := abcTree()

	tt := []struct {
		name string
		key  string
		n    *node
		want *node
	}{
		{
			name: "blank tree",
			key:  "fizz",
			n:    nil,
			want: nil,
		},
		{
			name: "S to S",
			key:  "S",
			n:    tree.root,
			want: tree.root,
		},
		{
			name: "S to E",
			key:  "E",
			n:    tree.root,
			want: tree.root.Left,
		},
		{
			name: "S to A",
			key:  "A",
			n:    tree.root,
			want: tree.root.Left.Left,
		},
		{
			name: "S to C",
			key:  "C",
			n:    tree.root,
			want: tree.root.Left.Left.Right,
		},
		{
			name: "E to C",
			key:  "C",
			n:    tree.root.Left,
			want: tree.root.Left.Left.Right,
		},
		{
			name: "S to R",
			key:  "R",
			n:    tree.root,
			want: tree.root.Left.Right,
		},
		{
			name: "S to H",
			key:  "H",
			n:    tree.root,
			want: tree.root.Left.Right.Left,
		},
		{
			name: "S to M",
			key:  "M",
			n:    tree.root,
			want: tree.root.Left.Right.Left.Right,
		},
		{
			name: "S to X",
			key:  "X",
			n:    tree.root,
			want: tree.root.Right,
		},
		{
			name: "X to X",
			key:  "X",
			n:    tree.root.Right,
			want: tree.root.Right,
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got := search(tc.key, tc.n)
			if got != tc.want {
				t.Errorf("search(%q, %+v) got %+v, want %+v", tc.key, tc.n, got, tc.want)
			}
		})
	}
}

func TestTree_Get(t *testing.T) {
	tree := abcTree()

	tt := []struct {
		key  string
		want []byte
	}{
		{"S", []byte("sea")},
		{"R", []byte("rock")},
		{"unknown", nil},
	}
	for _, tc := range tt {
		t.Run(tc.key, func(t *testing.T) {
			got := tree.Get(tc.key)
			if !bytes.Equal(got, tc.want) {
				t.Errorf("Get(%q) got %q, want %q", tc.key, got, tc.want)
			}
		})
	}

	blank := Tree{}
	key := "missing"
	got := blank.Get(key)
	if got != nil {
		t.Errorf("Get(%q) got %+v from blank tree, want nil", key, got)
	}
}

func TestPut_update(t *testing.T) {
	tree := abcTree()

	tt := []struct {
		name  string
		key   string
		value []byte
		n     *node
		want  *node
	}{
		{
			name:  "blank tree",
			key:   "S",
			value: []byte("sea"),
			n:     nil,
			want:  nil,
		},
		{
			name:  "update S from S",
			key:   "S",
			value: []byte("sea"),
			n:     tree.root,
			want:  tree.root,
		},
		{
			name:  "update C from S",
			key:   "C",
			value: []byte("cat"),
			n:     tree.root,
			want:  tree.root.Left.Left.Right,
		},
		{
			name:  "update C from A",
			key:   "C",
			value: []byte("cat"),
			n:     tree.root.Left.Left,
			want:  tree.root.Left.Left.Right,
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got := put(tc.key, tc.value, tc.n)
			if got != tc.want {
				t.Errorf("put(%q, %q, n) got node %+v, want %+v", tc.key, tc.value, got, tc.want)
			}
			if tc.want != nil && !bytes.Equal(got.Value, tc.value) {
				t.Errorf("put(%q, %q, n) got value %q, want %q", tc.key, tc.value, got.Value, tc.value)
			}
		})
	}
}

func TestPut_insert(t *testing.T) {
	tree := abcTree()

	got := put("Y", []byte("yolo"), tree.root)
	if got == nil {
		t.Fatal(got)
	}
	if got != tree.root.Right.Right {
		t.Errorf("put(Y, yolo, root) got node %+v, want %+v", got, tree.root.Right.Right)
	}

	got = put("W", []byte("woo"), tree.root)
	if got == nil {
		t.Fatal(got)
	}
	if got != tree.root.Right.Left {
		t.Errorf("put(W, woo, root) got node %+v, want %+v", got, tree.root.Right.Left)
	}
}

func TestTree_Set(t *testing.T) {
	tree := Tree{}

	key := "name"
	value := []byte("Bob")
	tree.Set(key, value)
	if tree.root == nil {
		t.Errorf("Set(%q, %q) root is nil", key, value)
	}
	if tree.root.Key != key {
		t.Errorf("Set(%q, %q) wrong root key, got %q", key, value, tree.root.Key)
	}
	if !bytes.Equal(tree.root.Value, value) {
		t.Errorf("Set(%q, %q) wrong root value, got %q", key, value, tree.root.Value)
	}

	key = "planet"
	value = []byte("Earth")
	tree.Set(key, value)
	if tree.root.Right == nil {
		t.Errorf("Set(%q, %q) right child is nil", key, value)
	}
	if tree.root.Right.Key != key {
		t.Errorf("Set(%q, %q) wrong right child key, got %q", key, value, tree.root.Right.Key)
	}
	if !bytes.Equal(tree.root.Right.Value, value) {
		t.Errorf("Set(%q, %q) wrong right child value, got %q", key, value, tree.root.Right.Value)
	}
}

func TestKeys(t *testing.T) {
	tree := abcTree()
	kk := keys(nil, tree.root)
	want := []string{"A", "C", "E", "H", "M", "R", "S", "X"}
	if !equal(kk, want) {
		t.Errorf("keys(nil, root) got %v, want %v", kk, want)
	}
}

func TestTree_Keys(t *testing.T) {
	tree := &Tree{}
	kk := tree.Keys()
	if kk != nil {
		t.Errorf("Keys() got %v, want nil", kk)
	}

	tree = abcTree()
	kk = tree.Keys()
	want := []string{"A", "C", "E", "H", "M", "R", "S", "X"}
	if !equal(kk, want) {
		t.Errorf("Keys() got %v, want %v", kk, want)
	}
}
