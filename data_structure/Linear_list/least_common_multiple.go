package main

import "fmt"

//最小公倍数
//func processing(m int,n int,product_value int) (result int){
//	if m == n{
//		result = product_value/m
//		fmt.Println("least common multiple",product_value/m)
//	}else{
//		difference := int(math.Abs(float64(m - n)))
//		if m >n{
//			m = difference
//		}else{
//			n = difference
//		}
//		processing(m, n, product_value)
//	}
//	return
//}

func processing(m int, n int) (result int) {
	r := m * n
	for m != n {
		if m > n {
			m -= n
		} else {
			n -= m
		}
	}
	return r / n
}
func main() {
	m, n := 4, 7
	result := processing(m, n)
	fmt.Println("-=-=", result)

}
