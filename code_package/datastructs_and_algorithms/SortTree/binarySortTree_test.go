package SortTree

import (
	"testing"
	"math/rand"
	"golang-bibble/code_package/utils"
)

func TestSort(t *testing.T)  {
	tree := NewSortTree(22)
	array := []int{1,19,5,32,8, 66, 14, 7, 25, 82}
	for _, i:= range array{
		tree.Insert(i)
	}
	pre := []int{22, 1, 19, 5, 8, 7, 14, 32, 25, 66, 82}
	if !utils.CompareSlice(pre, *tree.Order(PreOrder)){
		t.Error(`PreOrder failed`)
	}
	in := []int{1, 5, 7, 8, 14, 19, 22, 25, 32, 66, 82}
	if !utils.CompareSlice(in, *tree.Order(InOrder)){
		t.Error(`InOrder failed`)
	}
	post := []int{7, 14, 8, 5, 19, 1, 25, 82, 66, 32, 22}
	if !utils.CompareSlice(post, *tree.Order(PostOrder)){
		t.Error(`PostOrder failed`)
	}

	node, times := tree.Find(14)
	if node.data != 14 || times != 6{
		t.Error(`Find value failed`)
	}
}

func BenchmarkInsertFind(b *testing.B)  {
	benchmarkInsertFind(b, b.N)
}
func BenchmarkInsertFind_100W(b *testing.B)  {
	benchmarkInsertFind(b, 100000)
}
func benchmarkInsertFind(b *testing.B, size int)  {
	tree := NewSortTree(rand.Int())
	for i:=0;i<size; i++{
		tree.Insert(rand.Int())
	}

	//tree.Find(rand.Int())
}

