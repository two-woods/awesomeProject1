package main

import "fmt"

type doubleElemType int
type doubleLinkList struct {
	data         doubleElemType
	llink, rlink *doubleLinkList
}

func createDoubleHead() (head doubleLinkList) {
	head = doubleLinkList{0, nil, nil}
	head.rlink = &head
	return
}

func insert() {
}

func delete() {

}
func createDoubleList(head *doubleLinkList, initList []int) {
	p := head
	for _, v := range initList {
		q := &doubleLinkList{doubleElemType(v), nil, nil}
		p.rlink = q
		q.llink = p
		p = q
	}
	fmt.Println("2312312", head.rlink.data)
}

func main() {
	head := createDoubleHead()
	initList := []int{1, 2, 3, 4}
	createDoubleList(&head, initList)

	//x =
	//insert(head,x,item)
}
