package main

// ANSI alphabet
const AlphabetSize = 26

type TrieNode struct {
	children [AlphabetSize]*TrieNode
	isEnd    bool
}

type Trie struct {
	root *TrieNode
}

func InitTrie() *Trie {
	trie := &Trie{root: &TrieNode{}}
	return trie
}

// Insert
func (t *Trie) Insert(word string) {
	p := t.root

	for _, char := range word {
		i := char - 'a'
		if p.children[i] == nil {
			p.children[i] = &TrieNode{}
		}
		p = p.children[i]
	}

	p.isEnd = true
}

// Search
func (t *Trie) Search(word string) bool {
	p := t.root
	var i rune
	for _, char := range word {
		i = char - 'a'
		if p.children[i] == nil {
			return false
		}
		p = p.children[i]
	}

	return p.isEnd
}

// func main() {
// 	trie := InitTrie()
// 	trie.Insert("pinga")
// 	trie.Insert("zapato")
// 	trie.Insert("abrigo")
// 	fmt.Println(trie.root)

// 	fmt.Println(trie.Search("ping"))
// 	fmt.Println(trie.Search("pinga"))
// }
