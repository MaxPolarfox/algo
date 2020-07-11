package tries

type SuffixTrie map[byte]SuffixTrie

// Feel free to add to this function.
func NewSuffixTrie() SuffixTrie {
	trie := SuffixTrie{}
	return trie
}

func (trie SuffixTrie) PopulateSuffixTrieFrom(str string) {
	for i := range str {
		node := trie
		for j := i; j < len(str); j++ {
			letter := str[j]
			if _, found := node[letter]; !found {
				node[letter] = NewSuffixTrie()

			}
			node = node[letter]
		}
		node['*'] = nil
	}
}

func (trie SuffixTrie) Contains(str string) bool {
	node := trie
	for i := range str {
		letter := str[i]
		if _, found := node[letter]; !found {
			return false
		}
		node = node[letter]
	}
	_, found := node['*']
	return found
}
