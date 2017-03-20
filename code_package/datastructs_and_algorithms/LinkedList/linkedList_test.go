package LinkedList

import "testing"

func TestLinkedList(t *testing.T) {
	mylist := Node{Data: "start"}
	for index,data := range []string{"docker", "chongzhi", "google", "dota", "test", "d", "github"}{
		mylist.Insert(index +1, data)
	}

	dota := mylist.FindNth(5)
	github := mylist.FindNth(8)

	if "dota" != dota.(string) || "github" != github.(string){
		t.Error(`linkedlist failed`)
	}
}

func TestDelete(t *testing.T) {
	mylist := Node{Data: "start"}
	for index,data := range []string{"docker", "chongzhi", "google", "dota", "test", "d", "github"}{
		mylist.Insert(index +1, data)
	}
	// 边界值不能删除
	_, err1 := mylist.Delete(9)
	_, err2 := mylist.Delete(0)

	if err1 == nil || err2 == nil{
		t.Error(`linkedlist failed`)
	}

	// 能正确删除并返回结果，mylist被修改
	data1, _ := mylist.Delete(8)
	data2, _ := mylist.Delete(2)
	_, err3 := mylist.Delete(7)

	if data1.(string) != "github" || data2.(string)  != "docker"|| err3 == nil{
		t.Error(`linkedlist failed`)
	}
}


