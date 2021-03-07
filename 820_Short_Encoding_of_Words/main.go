package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	words := []string{"time", "me", "bell"}
	fmt.Println(minimumLengthEncoding(words))
}

// Method 1 - brute force
func minimumLengthEncoding_(words []string) int {
	res, m := 0, map[string]bool{}
	for _, w := range words {
		m[w] = true
	}
	for w := range m {
		for i := 1; i < len(w); i++ {
			delete(m, w[i:])
		}
	}
	for w := range m {
		res += len(w) + 1
	}
	return res
}

// Trie
// Reference - https://en.wikipedia.org/wiki/Trie
type node struct {
	value byte
	sub   []*node
}

func (t *node) has(b byte) (*node, bool) {
	if t == nil {
		return nil, false
	}
	for i := range t.sub {
		if t.sub[i] != nil && t.sub[i].value == b {
			return t.sub[i], true
		}
	}
	return nil, false
}

func (t *node) isLeaf() bool {
	if t == nil {
		return false
	}
	return len(t.sub) == 0
}

func (t *node) add(s []byte) {
	now := t
	for i := len(s) - 1; i > -1; i-- {
		if v, ok := now.has(s[i]); ok {
			now = v
			continue
		}
		temp := new(node)
		temp.value = s[i]
		now.sub = append(now.sub, temp)
		now = temp
	}
}

func (t *node) endNodeOf(s []byte) *node {
	fmt.Println(string(s))
	now := t
	for i := len(s) - 1; i > -1; i-- {
		if v, ok := now.has(s[i]); ok {
			now = v
			continue
		}
		return nil
	}
	return now
}

func minimumLengthEncoding(words []string) int {
	res, tree, m := 0, new(node), make(map[string]bool)
	for i := range words {
		if !m[words[i]] {
			tree.add([]byte(words[i]))
			m[words[i]] = true
		}
	}
	for s := range m {
		if tree.endNodeOf([]byte(s)).isLeaf() {
			res += len(s)
			res++
		}
	}
	fmt.Println(PrettyPrint(tree.sub))
	return res
}

// PrettyPrint to print struct in a readable way
func PrettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}
