package main

type node struct {
	coef int
	exp  int
	link *node
}

func generateCNode(co int, ex int, r *node) (w *node) {
	w = &node{co, ex, nil}
	r.link = w
	return w
}

func addNode(a node, b node) {
	var c, r *node
	var x int
	p, q := &a, &b
	r = c
	for p != nil && q != nil {
		if p.exp == q.exp {
			x = p.coef + q.coef
			if x != 0 {
				r = generateCNode(x, p.exp, r)
				p = p.link
				q = q.link
			} else if p.exp < q.exp {
				r = generateCNode(q.coef, q.exp, r)
				q = q.link
			} else {
				r = generateCNode(p.coef, p.exp, r)
				p = p.link
			}
		}
	}
	for p != nil {
		r = generateCNode(p.coef, p.exp, r)
		p = p.link
	}
	for q != nil {
		r = generateCNode(q.coef, p.exp, r)
		p = p.link
	}
	r.link = nil
	p = c
	c = c.link

}
