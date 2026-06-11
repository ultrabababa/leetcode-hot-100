package lru_cache

// Node 定义双向链表的节点
type Node struct {
	Key  int
	Val  int
	Prev *Node
	Next *Node
}

// LRUCache 定义缓存结构体
type LRUCache struct {
	capacity int
	cache    map[int]*Node // 哈希表：Key -> 链表节点
	head     *Node         // 伪头节点
	tail     *Node         // 伪尾节点
}

// Constructor 初始化
func Constructor(capacity int) LRUCache {
	l := LRUCache{
		capacity: capacity,
		cache:    make(map[int]*Node),
		head:     &Node{}, // 哨兵节点，方便操作
		tail:     &Node{},
	}
	// 初始化链表： head <-> tail
	l.head.Next = l.tail
	l.tail.Prev = l.head
	return l
}

// Get 获取值
func (lru *LRUCache) Get(key int) int {
	if node, ok := lru.cache[key]; ok {
		// 1. 如果存在，将其移动到头部（表示最近使用）
		lru.moveToHead(node)
		return node.Val
	}
	return -1
}

// Put 写入值
func (lru *LRUCache) Put(key int, value int) {
	if node, ok := lru.cache[key]; ok {
		// 1. 如果 Key 存在，更新值，并移到头部
		node.Val = value
		lru.moveToHead(node)
	} else {
		// 2. 如果 Key 不存在，创建新节点
		newNode := &Node{Key: key, Val: value}
		lru.cache[key] = newNode
		lru.addToHead(newNode) // 放到头部

		// 3. 检查容量，如果超了，删除尾部（最久未使用）
		if len(lru.cache) > lru.capacity {
			removed := lru.removeTail()
			delete(lru.cache, removed.Key) // 别忘了从 map 中删除
		}
	}
}

// ================= 辅助函数 =================

// addToHead 将节点添加到头部（head 之后）
func (lru *LRUCache) addToHead(node *Node) {
	node.Prev = lru.head
	node.Next = lru.head.Next
	lru.head.Next.Prev = node
	lru.head.Next = node
}

// removeNode 从链表中移除节点
func (lru *LRUCache) removeNode(node *Node) {
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
}

// moveToHead 将现有节点移动到头部
func (lru *LRUCache) moveToHead(node *Node) {
	lru.removeNode(node)
	lru.addToHead(node)
}

// removeTail 移除尾部节点（tail 之前），并返回该节点（方便 map 删除 key）
func (lru *LRUCache) removeTail() *Node {
	node := lru.tail.Prev
	lru.removeNode(node)
	return node
}
