package problem_218

import (
	"math/rand"
	"sort"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type Node struct {
	x      int
	height int
	isAdd  bool
}

func getSkyline(buildings [][]int) [][]int {
	nodes := make([]Node, 2*len(buildings))
	for i := range buildings {
		nodes[i<<1] = Node{buildings[i][0], buildings[i][2], true}
		nodes[i<<1|1] = Node{buildings[i][1], buildings[i][2], false}
	}
	sort.Slice(nodes, func(i, j int) bool {
		return nodes[i].x < nodes[j].x
	})
	heightTimesMap := NewSkipListMap()
	xMaxHeightMap := NewSkipListMap()
	for _, node := range nodes {
		x := toIntPointer(node.x)
		height := toIntPointer(node.height)
		if node.isAdd {
			heightTimesMap.Put(height, heightTimesMap.Get(height)+1)
		} else {
			if heightTimesMap.Get(height) == 1 {
				heightTimesMap.Remove(height)
			} else {
				heightTimesMap.Put(height, heightTimesMap.Get(height)-1)
			}
		}
		if heightTimesMap.IsEmpty() {
			xMaxHeightMap.Put(x, 0)
		} else {
			xMaxHeightMap.Put(x, *heightTimesMap.LastKey())
		}
	}
	ans := make([][]int, 0)
	for _, kv := range xMaxHeightMap.GetKeyValueSet() {
		x, height := kv[0], kv[1]
		if len(ans) == 0 || ans[len(ans)-1][1] != height {
			ans = append(ans, []int{x, height})
		}
	}
	return ans
}

type SkipListNode struct {
	key       *int
	val       int
	nextNodes []*SkipListNode
}

func (n *SkipListNode) isKeyLess(otherKey *int) bool {
	return otherKey != nil && (n.key == nil || *n.key < *otherKey)
}

func (n *SkipListNode) isKeyEqual(otherKey *int) bool {
	return (n.key == nil && otherKey == nil) ||
		(n.key != nil && otherKey != nil && *n.key == *otherKey)
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

func (slm *SkipListMap) Put(key *int, val int) {
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

func (slm *SkipListMap) Get(key *int) int {
	if key == nil {
		return 0
	}
	node := slm.mostRightLessNodeInZeroLevel(key).nextNodes[0]
	if node == nil || !node.isKeyEqual(key) {
		return 0
	}
	return node.val
}

func (slm *SkipListMap) ContainsKey(key *int) bool {
	if key == nil {
		return false
	}
	node := slm.mostRightLessNodeInZeroLevel(key).nextNodes[0]
	return node != nil && node.isKeyEqual(key)
}

func (slm *SkipListMap) Remove(key *int) {
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

func (slm *SkipListMap) LastKey() *int {
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

func (slm *SkipListMap) GetKeyValueSet() [][]int {
	ans := make([][]int, 0, slm.size)
	cur := slm.head.nextNodes[0]
	for cur != nil {
		ans = append(ans, []int{*cur.key, cur.val})
		cur = cur.nextNodes[0]
	}
	return ans
}

func (slm *SkipListMap) IsEmpty() bool {
	return slm.size == 0
}

// 从最高层开始一直找下去，找到第0层<key的最右节点
func (slm *SkipListMap) mostRightLessNodeInZeroLevel(key *int) *SkipListNode {
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
func (slm *SkipListMap) mostRightLessonNodeInLevel(key *int, cur *SkipListNode, level int) *SkipListNode {
	for next := cur.nextNodes[level]; next != nil && next.isKeyLess(key); {
		cur = next
		next = cur.nextNodes[level]
	}
	return cur
}

func toIntPointer(i int) *int {
	return &i
}
