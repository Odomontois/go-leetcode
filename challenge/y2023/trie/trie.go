package trie

type Trie struct {
	children [26]*Trie
	end      bool
}

func Constructor() Trie {
	return Trie{}
}

func (trie *Trie) Insert(word string) {
	for i := 0; i < len(word); i++ {
		l := word[i] - 'a'
		if trie.children[l] == nil {
			trie.children[l] = &Trie{}
		}
		trie = trie.children[l]
	}
	trie.end = true
}

func (trie *Trie) searchPrefix(word string) *Trie {
	for i := 0; i < len(word); i++ {
		l := word[i] - 'a'
		if trie.children[l] == nil {
			return nil
		}
		trie = trie.children[l]
	}
	return trie
}

func (trie *Trie) Search(word string) bool {
	trie = trie.searchPrefix(word)
	return trie != nil && trie.end
}

func (trie *Trie) StartsWith(prefix string) bool {
	return trie.searchPrefix(prefix) != nil
}
