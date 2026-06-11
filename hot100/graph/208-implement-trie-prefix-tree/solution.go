package implement_trie_prefix_tree

// Trie 定义前缀树节点
type Trie struct {
	children [26]*Trie // 指向 26 个可能的小写字母子节点
	isEnd    bool      // 标记该节点是否是一个单词的结尾
}

// Constructor 初始化前缀树对象
func Constructor() Trie {
	return Trie{}
}

// Insert 向前缀树中插入字符串 word
func (this *Trie) Insert(word string) {
	node := this // 从根节点开始
	for _, ch := range word {
		index := ch - 'a' // 计算字母对应的数组索引 (0-25)

		// 如果对应字母的子节点不存在，就新建一个
		if node.children[index] == nil {
			node.children[index] = &Trie{}
		}
		// 指针向下移动到子节点
		node = node.children[index]
	}
	// 单词的最后一个字符对应的节点，打上“单词结尾”的标记
	node.isEnd = true
}

// Search 如果字符串 word 在前缀树中，返回 true
func (this *Trie) Search(word string) bool {
	node := this
	for _, ch := range word {
		index := ch - 'a'
		// 如果路断了，说明树里根本没这个词
		if node.children[index] == nil {
			return false
		}
		node = node.children[index]
	}
	// 单词找完了，但还要检查这里是不是一个完整单词的结尾
	return node.isEnd
}

// StartsWith 如果之前已经插入的字符串 word 的前缀之一为 prefix ，返回 true
func (this *Trie) StartsWith(prefix string) bool {
	node := this
	for _, ch := range prefix {
		index := ch - 'a'
		// 如果路断了，说明连这个前缀都没有
		if node.children[index] == nil {
			return false
		}
		node = node.children[index]
	}
	// 只要能顺利走完 prefix 的所有字符，说明这个前缀存在！不需要管 isEnd
	return true
}
