package collection

import "math/rand"

const MAX_LEVEL = 32
const P = 4

type skipList struct {
	head  *SkipListNode
	tail  *SkipListNode
	size  int
	level int
}

type SkipListNode struct {
	Score int
	Data  interface{}
	prev  *SkipListNode
	// 层深，区间为 [0,level)
	level []*nodeLevelInfo
}

type nodeLevelInfo struct {
	next *SkipListNode
	span int //	暂时没用
}

func NewSkipList() *skipList {
	return &skipList{
		head:  newSkipListNode(-1, nil, MAX_LEVEL),
		tail:  nil,
		size:  0,
		level: 1,
	}
}

func newSkipListNode(score int, data interface{}, level int) *SkipListNode {
	node := &SkipListNode{
		Score: score,
		Data:  data,
		level: make([]*nodeLevelInfo, level),
	}
	for i := range node.level {
		node.level[i] = &nodeLevelInfo{}
	}
	return node
}

func (l *skipList) Put(score int, data interface{}) *SkipListNode {

	var preves [MAX_LEVEL]*SkipListNode

	//	查找 score 对应位置的左节点
	prev := l.head
	for level := l.level - 1; level >= 0; level-- {
		for prev != nil &&
			prev.level[level].next != nil &&
			score > prev.level[level].next.Score {
			prev = prev.level[level].next
		}

		preves[level] = prev
	}

	//	生成level
	level := randomLevel()
	if level > l.level {
		for i := l.level; i < level; i++ {
			preves[i] = l.head
		}
		l.level = level
	}

	//	插入
	node := newSkipListNode(score, data, level)
	for i := 0; i < level; i++ {
		//	.next
		node.level[i].next = preves[i].level[i].next
		preves[i].level[i].next = node
	}

	//	.prev
	node.prev = preves[0]
	if next := node.level[0].next; next != nil {
		next.prev = node
	}

	l.size++
	return node
}

func (l *skipList) DelByScore(score int) {
	preves := make([]*SkipListNode, l.level)

	cur := l.head
	for i := l.level - 1; i >= 0; i-- {
		for cur != nil &&
			cur.level[i].next != nil &&
			cur.level[i].next.Score < score {
			cur = cur.level[i].next
		}
		preves[i] = cur
	}

	if next := cur.level[0].next; next != nil {
		l.DelNode(next, preves)
	}
}

func (l *skipList) DelNode(node *SkipListNode, preves []*SkipListNode) {
	for i := 0; i < len(node.level); i++ {
		preves[i].level[i].next = node.level[i].next
	}

	if next := node.level[0].next; next != nil {
		next.prev = preves[0]
	}

	if l.head.level[l.level-1].next == nil {
		l.level--
	}
}

func (l *skipList) Each(do func(score int, data interface{})) {
	for cur := l.head.level[0].next; cur != nil; cur = cur.level[0].next {
		do(cur.Score, cur.Data)
	}
}

func (l *skipList) EachNode(do func(node *SkipListNode)) {
	for cur := l.head.level[0].next; cur != nil; cur = cur.level[0].next {
		do(cur)
	}
}

func (l *skipList) RangeByScore(left int, right int, do func(node *SkipListNode)) {
	cur := l.head.level[l.level-1].next
	for i := l.level - 1; i >= 0; i-- {
		if cur != nil && cur.Score < left {
			cur = cur.level[i].next
		}
	}

	for cur != nil && cur.Score >= left && cur.Score <= right {
		do(cur)
		cur = cur.level[0].next
	}
}

// 1/p 的概率使其加一层
func randomLevel() int {
	var level = 1
	for rand.Int()&0xFFFF*P < 0xFFFF {
		level += 1
	}
	return level
}
