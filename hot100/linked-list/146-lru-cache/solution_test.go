package lru_cache

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLRUCache(t *testing.T) {
	// 按照题目示例进行测试
	// LRUCache lRUCache = new LRUCache(2);
	lru := Constructor(2)

	// lRUCache.put(1, 1); // 缓存是 {1=1}
	lru.Put(1, 1)

	// lRUCache.put(2, 2); // 缓存是 {1=1, 2=2}
	lru.Put(2, 2)

	// lRUCache.get(1);    // 返回 1
	// 因为访问了 1，此时 1 变成最新的，2 变成最旧的 -> {2=2, 1=1}
	assert.Equal(t, 1, lru.Get(1))

	// lRUCache.put(3, 3); // 该操作会使得关键字 2 作废，缓存是 {1=1, 3=3}
	// 容量满了(2)，插入 3，需要淘汰最旧的 2
	lru.Put(3, 3)

	// lRUCache.get(2);    // 返回 -1 (未找到)
	assert.Equal(t, -1, lru.Get(2))

	// lRUCache.put(4, 4); // 该操作会使得关键字 1 作废，缓存是 {4=4, 3=3}
	// 此时缓存里是 {1=1, 3=3} (1是旧的, 3是新的)。插入 4，淘汰 1
	lru.Put(4, 4)

	// lRUCache.get(1);    // 返回 -1 (未找到)
	assert.Equal(t, -1, lru.Get(1))

	// lRUCache.get(3);    // 返回 3
	assert.Equal(t, 3, lru.Get(3))

	// lRUCache.get(4);    // 返回 4
	assert.Equal(t, 4, lru.Get(4))
}
