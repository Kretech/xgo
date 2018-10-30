package collection

import (
	"fmt"
	"sort"
	"testing"

	"github.com/Kretech/xgo/random"
	"github.com/Kretech/xgo/test"
)

func newSimpleSkipList() *skipList {
	l := NewSkipList()

	l.Put(4, 6)
	l.Put(10, 1)
	l.Put(2, 8)
	l.Put(6, 4)

	return l
}

func TestBasic(t *testing.T) {

	l := newSimpleSkipList()

	l.Each(func(score int, data interface{}) {
		fmt.Println(score, data)
	})

}

func TestSkipList_RangeByScore(t *testing.T) {
	l := newSimpleSkipList()
	l.RangeByScore(3, 7, func(node *SkipListNode) {
		fmt.Print(node.Score, ",")
	})
}

func TestSkipList_DelByScore(t *testing.T) {
	l := NewSkipList()

	l.DelByScore(4)

	l.EachNode(func(node *SkipListNode) {
		fmt.Print(node.Data, " ")
	})

}

func TestSkipList_Sort(t *testing.T) {
	l := NewSkipList()

	num := 10
	a := make([]int, num)
	for i := 0; i < num; i++ {
		a[i] = random.Int() % 1000000
		l.Put(a[i], a[i])
	}

	i := 0
	l.Each(func(score int, data interface{}) {
		a[i] = data.(int)
		i++
	})

	test.BeTrue(t, sort.IntsAreSorted(a))
}
