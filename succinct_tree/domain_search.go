package succinct_tree

import (
	"unicode/utf8"
)

/**
  Look-up a word suffix in the trie. Returns true if and only if the suffix exists
  in the trie as full word and the cursor world
*/
func (f *FrozenTrie) LookupSuffix(word string) bool {
	node := f.GetRoot()
	for i, w := 0, 0; i < len(word); i += w {
		runeValue, width := utf8.DecodeRuneInString(word[i:])
		w = width
		var child FrozenTrieNode
		var j uint = 0
		for ; j < node.GetChildCount(); j++ {
			child = node.GetChild(j)
			if child.Letter() == string(runeValue) {
				break
			}
		}

		if j == node.GetChildCount() {
			// if we cant find more matching children,
			// and the next rune is . then the word is a subdomain of a domin in the tree
			return node.Final() && runeValue == '.'
		}
		node = child
	}

	return node.Final()
}
