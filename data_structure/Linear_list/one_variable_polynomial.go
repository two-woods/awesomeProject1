package main

import (
	"fmt"
	"strconv"
)

type node struct {
	coef int
	exp  int
	link *node
}

func createHead() (head node) {
	return node{0, 0, nil}
}
func generateCNode(co int, ex int, r *node) (w *node) {
	w = &node{co, ex, nil}
	r.link = w
	return w
}

func createList(head *node, coefList []int, expList []int) {
	p := head
	for k, coefItem := range coefList {
		q := &node{coefItem, expList[k], nil}
		p.link = q
		p = q
	}
}

func treavlOneVarPolynomial(head node) {
	p := head.link
	fmt.Print("\n**********One variable polynomial:" + "\n")
	for p != nil {
		fmt.Print(strconv.Itoa(p.coef) + "*X" + "**" + strconv.Itoa(p.exp) + "+")
		p = p.link
	}
}

func addNode(a node, b node) (c node) {
	var r *node
	var x int
	p, q := a.link, b.link
	r = &c
	fmt.Println("ewqeq")

	for p != nil && q != nil {
		if p.exp == q.exp {
			x = p.coef + q.coef
			if x != 0 {
				r = generateCNode(x, p.exp, r)
				p = p.link
				q = q.link
			}
		} else if p.exp < q.exp {
			r = generateCNode(q.coef, q.exp, r)
			q = q.link
		} else {
			fmt.Println("ewqeq")
			r = generateCNode(p.coef, p.exp, r)
			p = p.link
		}
	}
	for p != nil {
		r = generateCNode(p.coef, p.exp, r)
		p = p.link
	}
	for q != nil {
		r = generateCNode(q.coef, q.exp, r)
		q = q.link
	}
	fmt.Println("qweqweqw", r.coef)
	r.link = nil
	p = &c
	//c = *c.link
	return
}

func main() {
	head1, head2 := createHead(), createHead()
	coef_list1 := []int{3, 2, 1}
	coef_list2 := []int{3, 0, 1}
	exp_list1 := []int{4, 3, 1}
	exp_list2 := []int{3, 2, 1}
	createList(&head1, coef_list1, exp_list1)
	createList(&head2, coef_list2, exp_list2)

	treavlOneVarPolynomial(head1)
	treavlOneVarPolynomial(head2)
	fmt.Print("\n**********" + "\n")

	c := addNode(head1, head2)
	treavlOneVarPolynomial(c)
}
