package implement_trie_prefix_tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrie(t *testing.T) {
	// 模拟 LeetCode 示例 1
	trie := Constructor()

	// 1. insert("apple")
	trie.Insert("apple")

	// 2. search("apple") -> true
	assert.True(t, trie.Search("apple"), "Search 'apple' should return true")

	// 3. search("app") -> false (因为 "app" 只是前缀，还没作为一个完整的词插入过)
	assert.False(t, trie.Search("app"), "Search 'app' should return false")

	// 4. startsWith("app") -> true
	assert.True(t, trie.StartsWith("app"), "StartsWith 'app' should return true")

	// 5. insert("app")
	trie.Insert("app")

	// 6. search("app") -> true (现在 "app" 也是一个完整的词了)
	assert.True(t, trie.Search("app"), "Search 'app' should return true after insertion")

	// 测试额外边界情况
	assert.False(t, trie.Search("hello"), "Search 'hello' should return false")
	assert.False(t, trie.StartsWith("applxe"), "StartsWith 'applxe' should return false")
}
