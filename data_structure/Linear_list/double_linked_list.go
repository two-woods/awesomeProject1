package main

import (
	"fmt"
)

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

func insertByData(head *doubleLinkList, x doubleElemType, item doubleElemType) (s string) {
	s = "insert success"
	p := head.rlink
	for p != nil && p.data != x {
		p = p.rlink
	}
	if p == nil {
		s = fmt.Sprintf("no have node which data is %d", x)
		return
	}
	q := &doubleLinkList{0, p, p.rlink}
	p.rlink.llink = q
	p.rlink = q
	return s
}

func delete() {

}
func createDoubleList(head *doubleLinkList, initList []int) {
	p := head
	for _, v := range initList {
		//q := &doubleLinkList{doubleElemType(v), nil, nil}
		//p.rlink = q
		//q.llink = p
		//p = q
		q := doubleLinkList{doubleElemType(v), nil, nil}
		p.rlink = &q
		q.llink = p
		p = &q
	}

}

func treavlList(head doubleLinkList) {

	//if head.rlink == &head{                    //此处要用main.doubleLinkList值类型来比较，而非指针类型
	if *head.rlink == head {
		fmt.Println("no node")
		return
	}
	fmt.Print("treavl : ")
	r := &head
	p := head.rlink
	for p != nil {
		fmt.Print(p.data, ",")
		r = p
		p = p.rlink
	}
	fmt.Print("reverse traveral:")
	//fmt.Println("0-0-0-0--0-0-",r.llink)
	for *r != head {
		fmt.Print(r.data, ",")
		r = r.llink
	}

}
func main() {
	head := createDoubleHead()
	initList := []int{1, 2, 3, 4}
	//initList := []int{}
	createDoubleList(&head, initList)
	treavlList(head)

	fmt.Print("\nNow ,we insert data after 2,")
	insertByData(&head, 2, 0)
	treavlList(head)

	//x =
	//insert(head,x,item)
}
