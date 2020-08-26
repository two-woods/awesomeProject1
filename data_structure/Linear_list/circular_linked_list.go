package main

import (
	"fmt"
)

type CirElemType int
type circularLinkList struct {
	data CirElemType
	link *circularLinkList
}

func createCircularHead() (head circularLinkList) {
	head = circularLinkList{0, nil}
	head.link = &head
	return head
}

func createcircularLinkList(head *circularLinkList, initList []CirElemType) {
	p := head
	for _, v := range initList {
		q := &circularLinkList{v, nil}
		p.link = q
		p = p.link
	}
	p.link = head
}

//	initOrderList := []CirElemType{1, 2, 3, 4}
func createcircularLinkList1(n int) (head *circularLinkList) {
	//p := &head
	//r:= &head
	//for _, v := range initList {
	//	q := circularLinkList{}
	//	p.data = v
	//	p.link = &q
	//	r = p                                      //for遍历完，为最后一个节点
	//	p = p.link
	//}
	//r.link = &head
	//return
	r := &circularLinkList{}
	//var r *circularLinkList                                                     //此处写法，会使得下面使用的r.link报警告Accessing field 'r.link' may lead to nil ...,因为你没给他赋初始值
	for i := 1; i <= n; i++ {
		p := circularLinkList{CirElemType(i), nil}
		if head == nil {
			head = &p //p指向变 head也变 ;不同于head = p，p的指向变，head不变
		} else {
			r.link = &p
		}
		r = &p //r永远指向最后一个节点
	}
	if head == nil {
		head = &circularLinkList{}
		head.link = head
	}
	r.link = head
	return
}

//have head node
func cirTreavl(cirHead circularLinkList) {
	fmt.Print("\ncircular linked list traversal:")
	p := cirHead.link
	fmt.Print("\ncircular linked list traversal:", p)

	for *p != cirHead {
		fmt.Print(p.data, ",")
		p = p.link
	}
	fmt.Print("\n")
}

//haven't head node
func cirTreavl1(cirHead circularLinkList) {
	//fmt.Print("circular linked list traversal:")

	fmt.Print(cirHead.data, ",")
	p := cirHead.link
	for *p != cirHead {
		fmt.Print(p.data, ",")
		p = p.link
	}
	fmt.Print("\n")
}

func findNodeBydata(cirHead circularLinkList) (resultNode circularLinkList) {
	p := cirHead.link
	//for p != &cirHead{							p变量容器地址一定,两个地址永远不一样,所以此处一直为ture
	for *p != cirHead { //比较两个节点是否一样来判断
		if p.data == 2 {
			resultNode = *p // *p表示circularLinkList类型，而p用*circularLinkList类型的变量来赋值
			return
		}
		p = p.link
	}
	return
}

func josephuQuestion(n int, k int, m int) {
	fmt.Println("This is initial sequence")
	criHead1 := createcircularLinkList1(n)
	cirTreavl1(*criHead1)
	r := &circularLinkList{}
	p := criHead1
	for i := 1; i < k; i++ {
		r = p
		p = p.link
	}
	for p.link != p {
		for i := 1; i < m; i++ { //i<m每次判定的是i+1,而每次循环内用的i，所以到报数m-1时退出，此时r指向m-1，p指向第m人
			r = p
			p = p.link
		}
		r.link = p.link
		fmt.Println("out", p.data)
		p = r.link
	}
	fmt.Println("Finally，The last one left,", p.data)

}

func main() {
	cirHead := createCircularHead()
	initList := []CirElemType{3, 4, 2, 1}
	createcircularLinkList(&cirHead, initList)
	cirTreavl(cirHead)
	findNode := findNodeBydata(cirHead)
	fmt.Println("Now, we want to find node which data is 2", findNode)

	//cirHead1 := createCircularHead()
	fmt.Println("Now,we paly josephu game,this is a initialization sequence")
	n, m, k := 8, 4, 3
	josephuQuestion(n, k, m)
}
