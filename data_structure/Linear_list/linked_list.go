package main

import (
	"fmt"
	"math/rand"
	"reflect"
)

type ElemType int

type linkList struct {
	data ElemType
	link *linkList
}

func generateRandomNumber() (n ElemType) {
	n = ElemType(rand.Intn(10))
	return
}

//自定义结构体判空,Custom structure nulls
func (head linkList) isEmpty() bool {
	return reflect.DeepEqual(head, linkList{})
}

///////////////////////////返回一个指针
//func createlinkList(n int) (head *linkList) {
//	var r *linkhead
//	for i :=1;i<=n;i++{
//		a := generateRandomNumber()
//		p := &linkList{a,nil}
//		//判断结构体为空不用nil,It is not use 'nil' to judge the structure as empty
//		//使用head == (linkList{})，或自定义结构体判空(Custom structure nulls)
//		//fmt.Println("isempty",head.isEmpty())
//		if head == nil{
//			head = p
//			r = head
//		}else{
//			r.link = p
//			r = p
//		}
//	}
//	return &*head
//}
//treavlhead(head)

/******************
*main.linkhead（指针）  转化为main.linkhead（指针变量） 加*
*************************/
//返回一个链表地址 对比上面的
func createLinkedhead(n int) (head linkList) {
	r := &head
	for i := 1; i <= n; i++ {
		a := generateRandomNumber()
		p := &linkList{a, nil} //reflect.TypeOf(p)   *main.linkList

		//判断结构体为空不用nil,It is not use 'nil' to judge the structure as empty
		//使用head == (linkhead{})，或自定义结构体判空(Custom structure nulls)
		//fmt.Println("isempty",head.isEmpty())
		if head == (linkList{}) {
			head.link = p
			r = p
		} else {
			r.link = p
			r = p
		}
	}
	return
}

func createLinkedhead1(insertList []ElemType) (head linkList) {
	r := &head
	for _, v := range insertList {
		q := &linkList{v, nil}
		if isEmpty(head) {
			head.link = q
		} else {
			r.link = q
		}
		r = q
	}
	return
}

func treavlhead(head linkList) {
	fmt.Print("head traversal:")
	p := head.link
	for p != nil {
		fmt.Print(p.data, ",")
		p = p.link
	}
	fmt.Println("\n")

}

// 非递归  non-recursive
func lengthLinkedhead(head linkList) (length int) {
	var p *linkList = &head
	for p.link != nil {
		length++
		p = p.link
	}
	return
}

//panic: runtime error: invalid memory address or nil pointer dereference ,指针指向nil会报错
//递归
func lengthLinkedhead1(head linkList) (length int) {
	var p *linkList = &head
	if p.link == nil {
		return 0
	} else {
		return 1 + lengthLinkedhead1(*head.link)
	}
}

func isEmpty(head linkList) (isEmpty bool) {
	return head.link == nil
}

func findElement(head linkList, item ElemType) (element *linkList) {
	var p *linkList = head.link
	for p != nil && p.data != item {
		p = p.link
	}
	return p

}

func aheadInsert(head *linkList, item ElemType) {
	fmt.Println("head123123:", &head)
	p := &linkList{item, nil}
	p.link = head.link
	head.link = p
}

func rearInsert(head *linkList, item ElemType) {
	p := head.link
	q := &linkList{item, nil}
	for p.link != nil { //此处 非p!=nil,p.link !=nil刚好在最后一个节点退出for,进而p退出时指向最后一个节点
		p = p.link
	}
	p.link = q
}

func beforeQpointernodeInsert(head *linkList, q *linkList, item ElemType) {
	p := &linkList{item, nil}
	if head.link == nil {
		head.link = p
	} else {
		p.link = q.link
		q.link = p
	}
}

func beforeInodeInsert(head *linkList, i int, item ElemType) {
	var n int
	q := head.link
	p := &linkList{item, nil}
	for q != nil { //此处 非p.link !=nil,p!=nil刚好在最后一个节点的下一个空节点不进行n++
		n++
		if n == i {
			p.link = q.link
			q.link = p
		}
		q = q.link
	}
}

func orderListInsert(head *linkList, item ElemType) {
	var z *linkList //此处不用var z linkList，要用指针，否则z新变量，新的地址空间，head并不会与其连接上
	//只有用指针变量才能修改链表指向的节点的data域以及指针域，对比引用传递和值传递
	p := head.link
	q := &linkList{item, nil}
	for p != nil {
		if p.data <= item {
		} else {
			q.link = p
			z.link = q //var z *linkList
			break
		}
		z = p //z = *p,配合上边的注释
		p = p.link
	}
}

func delNodeqPoint(head *linkList, q *linkList) {
	/*if head == q{         无头结点，要讨论此情况
		head = q.link
	}*/
	p := head
	for p.link != nil && p.link != q {
		p = p.link
	}
	p.link = q.link
}

func delAllDataIsTwoNode(head *linkList, item ElemType) {
	//if head.data == item{  无头结点，要讨论此情况
	//	head = head.link
	//}
	r, p := head, head.link
	for p != nil {
		if p.data == item {
			r.link = p.link
		}
		r = p
		p = p.link
	}
}

///////////////////////////////////////////////////////////
func reverseLinkList(head *linkList) {
	p := head.link
	//var q linkList = linkList{}
	var q *linkList = nil //不用var  q listList 否则后边会多一个data为0的节点,
	//必须使用*linkList来操控所指向地址的指针域
	fmt.Println("11111", q)
	/*先记住前节点，然后指向q去指向当前遍历节点p，p就可以link进行转移
	，再操作q.link（即操作当前p的指针）,指向r对应指向的新链表的头部*/
	for p != nil {
		r := q
		q = p
		p = p.link
		q.link = r
	}
	head.link = q
}

//三元运算符工具函数
func If(condition bool, trueVal, falseVal *linkList) *linkList {
	if condition {
		//fmt.Println("\nNumber of nodes1:",trueVal)
		return trueVal
	}
	//fmt.Println("\nNumber of nodes2:",falseVal)

	return falseVal
}

func mergeOrederedLinkList(head1 linkList, head2 linkList) (head3 linkList) {
	p := head1.link
	q := head2.link
	r := &head3 //r :=head3 则生成一个新的节点，无法操作head3,
	for p != nil && q != nil {
		if p.data <= q.data {
			r.link = p
			r = p //由r :=head3 此处则要写r=*p，从而生成新的变量而无法操作p的指针域
			p = p.link

		} else {
			r.link = q
			r = q //由r :=head3 此处则要写r=*q，从而生成新的变量而无法操作q的指针域,导致最后值操作了r未操作head3
			q = q.link
		}
	}
	r.link = If(p != nil, p, q)
	return head3
}

func copyLinkList(head3 *linkList) (head4 *linkList) {
	if head3 == nil {
		return nil
	} else {
		head4 = &linkList{head3.data, nil}
		head4.link = copyLinkList(head3.link)
	}
	return
}

func main() {
	head := createLinkedhead(3) //main.linkList
	fmt.Println("11111", &head)
	treavlhead(head)
	n := lengthLinkedhead(head)
	fmt.Println("\nNumber of nodes:", n)

	n = lengthLinkedhead1(head)
	fmt.Println("\nNumber of nodes:", n)
	fmt.Println("\nThe head is nil?", isEmpty(head))
	fmt.Println("\nOne is in the head?", findElement(head, 7))
	aheadInsert(&head, -1)
	fmt.Println("\nNow, insert negative one at the front of the head")
	treavlhead(head)
	rearInsert(&head, -1)
	fmt.Println("\nNow, insert negative one at the front of the rear")
	treavlhead(head)
	q := head.link.link
	fmt.Println("\nNow, insert negative one at the front of the rear")
	beforeQpointernodeInsert(&head, q, 11111)
	treavlhead(head)
	// Insert after the two node
	beforeInodeInsert(&head, 2, 2222)
	fmt.Println("\nNow, Insert after the two node")
	treavlhead(head)

	q = head.link
	q = head.link.link
	//q = head.link.link.link.link.link.link.link
	fmt.Println("\nNow, Deletes a node  q pointed")
	delNodeqPoint(&head, q)
	treavlhead(head)

	delAllDataIsTwoNode(&head, -1)
	fmt.Println("\nNow, Deletes a node whose data fiels is negative one")
	treavlhead(head)

	reverseLinkList(&head)
	fmt.Println("\nNow, Reverse a linked list ")
	treavlhead(head)

	insertList := []ElemType{1, 3, 5, 7}
	head1 := createLinkedhead1(insertList) //main.linkList,This is a order link list
	treavlhead(head1)
	orderListInsert(&head1, 2)
	treavlhead(head1)

	insertList = []ElemType{2, 4, 6, 8}
	head2 := createLinkedhead1(insertList) //main.linkList,This is a order link list
	treavlhead(head2)

	fmt.Println("\nNow, Merge two ordered linked lists ")
	head3 := mergeOrederedLinkList(head1, head2)
	treavlhead(head3)

	head4 := copyLinkList(&head3)
	fmt.Println("\nNow, Copy a linked list ")
	treavlhead(*head4)

}
