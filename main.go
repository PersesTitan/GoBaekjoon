package main

//import "fmt"
//
//func main() {
//	fmt.Println("강한친구 대한육군")
//	fmt.Println("강한친구 대한육군")
//}

func sum(a []int) int {
	s := 0
	for value := range a {
		s += a[value]
	}
	return s
}

func main() {
	nums := []int{1, 5, 16, 2}
	print(sum(nums))
}
