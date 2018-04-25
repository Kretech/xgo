package collection

import (
	"fmt"
	"testing"
)

func newList() *skipList {
	l := NewSkipList()

	l.Put(4, 6)
	l.Put(10, 1)
	l.Put(2, 8)
	l.Put(6, 4)

	return l
}

func TestBasic(t *testing.T) {

	l := newList()

	l.Each(func(score int, data interface{}) {
		fmt.Println(score, data)
	})

}

func TestSkipList_RangeByScore(t *testing.T) {
	l := newList()
	l.RangeByScore(3, 7, func(node *SkipListNode) {
		fmt.Print(node.Score, ",")
	})
}

func TestSkipList_DelByScore(t *testing.T) {
	l := newList()

	l.DelByScore(4)

	l.EachNode(func(node *SkipListNode) {
		fmt.Print(node.Data, " ")
	})

}
