package lc

// 146. LRU 缓存（中等）
//
// 设计并实现一个满足 LRU（最近最少使用）约束的数据结构。
//  LRUCache(capacity) 以正整数 capacity 初始化；
//  get(key) 存在返回其值，否则返回 -1；
//  put(key, value) 已存在则更新，不存在则插入，超出容量时逐出最久未使用的键。
//
// 思路：哈希表 + 双向链表
//  - map[key]*node：O(1) 定位节点；
//  - 双向链表（带哨兵头尾）：表头 = 最近使用，表尾 = 最久未使用；
//  - get / put 命中时都把节点摘除后插回头部；
//  - 淘汰时摘掉表尾并同步从 map 删除。
//  get、put 均为 O(1)，空间 O(capacity)。
//
// 骨架：类型和函数都给你，函数体自己填。
// 填的时候注意三个不变式：
//  1) 节点在链表里的位置 = 访问顺序（越靠头越新）
//  2) map 里的每个 key 都在链表上，且 map 和链表同时增删
//  3) 淘汰节点时，链表摘除 + map 删除，两边都要做

// dNode 双向链表节点，带 key 是为了淘汰时能直接从 map 里删掉它
type dNode struct {
	key, val   int
	prev, next *dNode
}

// LRUCache 哈希表 + 带哨兵的双向链表
type LRUCache struct {
	capacity int
	cache    map[int]*dNode
	head     *dNode // 哨兵头，表头（head.next）才是最近使用的节点
	tail     *dNode // 哨兵尾，表尾（tail.prev）是最久未使用的节点
}

// NewLRUCache 初始化：建 map + 串好哨兵（空链表）
func NewLRUCache(capacity int) LRUCache {
	head := &dNode{}
	tail := &dNode{prev: head}
	head.next = tail
	return LRUCache{
		capacity: capacity,
		cache:    make(map[int]*dNode),
		head:     head,
		tail:     tail,
	}
}

// Get 命中：摘除 + 插回头部 后返回值；未命中：返回 -1
func (c *LRUCache) Get(key int) int {
	if node, ok := c.cache[key]; ok {
		node.prev.next = node.next
		node.next.prev = node.prev
		node.next = c.head.next
		c.head.next.prev = node
		node.prev = c.head
		c.head.next = node
		return node.val
	}
	return -1
}

// Put 三种情况：
//  1. key 已存在 → 更新值 + 摘除插回头部
//  2. key 不存在 → 建节点入 map、插到链表头部
//  3. 超出容量 → 摘掉表尾节点并从 map 删掉它的 key
func (c *LRUCache) Put(key, value int) {
	if node, ok := c.cache[key]; ok {
		node.val = value
		node.prev.next = node.next
		node.next.prev = node.prev
		node.next = c.head.next
		c.head.next.prev = node
		node.prev = c.head
		c.head.next = node
		return
	}
	if len(c.cache) == c.capacity {
		delete(c.cache, c.tail.prev.key)
		c.tail.prev.prev.next = c.tail
		c.tail.prev = c.tail.prev.prev
	}
	node := &dNode{key: key, val: value}
	c.cache[key] = node
	node.next = c.head.next
	c.head.next.prev = node
	node.prev = c.head
	c.head.next = node
}
