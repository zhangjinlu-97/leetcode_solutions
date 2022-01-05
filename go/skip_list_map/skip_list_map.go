package skip_list_map

import (
	"math/rand"
)

// Go 基于跳跃表实现的有序表

type Comparable interface {
	Less(c Comparable) bool
	Equals(c Comparable) bool
}

type SkipListNode struct {
	key       Comparable
	val       interface{}
	nextNodes []*SkipListNode
}

func (n *SkipListNode) isKeyLess(otherKey Comparable) bool {
	return otherKey != nil && (n.key == nil || n.key.Less(otherKey))
}

func (n *SkipListNode) isKeyEqual(otherKey Comparable) bool {
	return (n.key == nil && otherKey == nil) ||
		(n.key != nil && otherKey != nil && n.key.Equals(otherKey))
}

type SkipListMap struct {
	head     *SkipListNode
	size     int
	maxLevel int
}

func NewSkipListMap() *SkipListMap {
	return &SkipListMap{
		head:     &SkipListNode{nil, 0, make([]*SkipListNode, 1)},
		size:     0,
		maxLevel: 0,
	}
}

func (slm *SkipListMap) Put(key Comparable, val interface{}) {
	if key == nil {
		return
	}
	less := slm.mostRightLessNodeInZeroLevel(key)
	find := less.nextNodes[0]
	if find != nil && find.isKeyEqual(key) {
		find.val = val
		return
	}
	// key不存在，则新增
	slm.size++
	newNodeLevel := 0
	for rand.Float64() < 0.5 {
		newNodeLevel++
	}
	// 新节点层数大于maxLevel，则把head的nextNodes扩充至newNodeLevel
	if newNodeLevel > slm.maxLevel {
		slm.maxLevel = newNodeLevel
		newNextNodes := make([]*SkipListNode, newNodeLevel+1)
		for i := range slm.head.nextNodes {
			newNextNodes[i] = slm.head.nextNodes[i]
		}
		slm.head.nextNodes = newNextNodes
	}
	// 新建一个节点，层数为newNodeLevel
	newNode := &SkipListNode{
		key:       key,
		val:       val,
		nextNodes: make([]*SkipListNode, newNodeLevel+1),
	}
	for level, pre := slm.maxLevel, slm.head; level >= 0; level-- {
		// level层中，找到最右的 < key的节点，并连接好
		pre = slm.mostRightLessonNodeInLevel(key, pre, level)
		if level <= newNodeLevel {
			newNode.nextNodes[level] = pre.nextNodes[level]
			pre.nextNodes[level] = newNode
		}
	}
}

func (slm *SkipListMap) Get(key Comparable) interface{} {
	if key == nil {
		return 0
	}
	node := slm.mostRightLessNodeInZeroLevel(key).nextNodes[0]
	if node == nil || !node.isKeyEqual(key) {
		return 0
	}
	return node.val
}

func (slm *SkipListMap) ContainsKey(key Comparable) bool {
	if key == nil {
		return false
	}
	node := slm.mostRightLessNodeInZeroLevel(key).nextNodes[0]
	return node != nil && node.isKeyEqual(key)
}

func (slm *SkipListMap) Remove(key Comparable) {
	if !slm.ContainsKey(key) {
		return
	}
	slm.size--
	pre := slm.head
	for level := slm.maxLevel; level >= 0; level-- {
		pre = slm.mostRightLessonNodeInLevel(key, pre, level)
		next := pre.nextNodes[level]
		if next != nil && next.isKeyEqual(key) {
			pre.nextNodes[level] = next.nextNodes[level]
		}
		if level != 0 && pre == slm.head && pre.nextNodes[level] == nil {
			slm.head.nextNodes = slm.head.nextNodes[:slm.maxLevel]
			slm.maxLevel--
		}
	}
}

func (slm *SkipListMap) CeilingKey(key Comparable) Comparable {
	if key == nil {
		return nil
	}
	node := slm.mostRightLessNodeInZeroLevel(key).nextNodes[0]
	if node == nil {
		return nil
	}
	return node.key
}

func (slm *SkipListMap) FloorKey(key Comparable) Comparable {
	if key == nil {
		return nil
	}
	less := slm.mostRightLessNodeInZeroLevel(key)
	next := less.nextNodes[0]
	if next != nil && next.isKeyEqual(key) {
		return next.key
	}
	return less.key
}

func (slm *SkipListMap) FirstKey() Comparable {
	first := slm.head.nextNodes[0]
	if first != nil {
		return first.key
	}
	return nil
}

func (slm *SkipListMap) LastKey() Comparable {
	cur := slm.head
	for level := slm.maxLevel; level >= 0; level-- {
		next := cur.nextNodes[level]
		for next != nil {
			cur = next
			next = cur.nextNodes[level]
		}
	}
	return cur.key
}

func (slm *SkipListMap) IsEmpty() bool {
	return slm.size == 0
}

// 从最高层开始一直找下去，找到第0层<key的最右节点
func (slm *SkipListMap) mostRightLessNodeInZeroLevel(key Comparable) *SkipListNode {
	if key == nil {
		return nil
	}
	cur := slm.head
	for level := slm.maxLevel; level >= 0; level-- {
		cur = slm.mostRightLessonNodeInLevel(key, cur, level)
	}
	return cur
}

// 在level层里，找到<key的最后一个节点并返回
func (slm *SkipListMap) mostRightLessonNodeInLevel(key Comparable, cur *SkipListNode, level int) *SkipListNode {
	for next := cur.nextNodes[level]; next != nil && next.isKeyLess(key); {
		cur = next
		next = cur.nextNodes[level]
	}
	return cur
}

func toIntPointer(i int) *int {
	return &i
}
