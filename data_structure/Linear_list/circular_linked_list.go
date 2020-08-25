package main

import "fmt"

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
	fmt.Print("circular linked list traversal:", p.data)

}

func createcircularLinkList1(initList []CirElemType) (head circularLinkList) {
	p := &head
	for _, v := range initList {
		q := &circularLinkList{v, nil}
		p = q
		p = p.link
	}
	fmt.Print("circular linked list traversal:", p)
	p.link = &head
	fmt.Print("circular linked list traversal:", p.data)
	return
}

func cirTreavl(cirHead circularLinkList) {
	fmt.Print("\ncircular linked list traversal:")
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

func josephuQuestion(criHead1 *circularLinkList) {

}

func main() {
	cirHead := createCircularHead()
	initList := []CirElemType{3, 4, 2, 1}
	createcircularLinkList(&cirHead, initList)
	cirTreavl(cirHead)
	findNode := findNodeBydata(cirHead)
	fmt.Println("Now, we want to find node which data is 2", findNode)

	//cirHead1 := createCircularHead()
	initOrderList := []CirElemType{1, 2, 3, 4}
	cirHead1 := createcircularLinkList1(initOrderList)
	fmt.Println("Now,we paly josephu game,this is a initialization sequence")
	cirTreavl(cirHead1)
	//josephuQuestion(&criHead1)
}
