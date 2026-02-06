package something

import (
	"fmt"
	"strings"
)

//	func main() {
//		fmt.Println("Hello World")
//		something.SayHello()
//		something.SayBye()
//
// }
//func main() {
//	name := true
//	name = "min"
//	//const name string = "dongmin"
//	fmt.Println(name)
//}

func multiply(a, b int) int {
	return a * b
}

func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

//func main() {
//	result := multiply(3, 4)
//	println("Result: ", result)
//	length, upperName := lenAndUpper("dongmin")
//	println("Length:", length, "Uppercase Name:", upperName)
//}

//func repeatMe(name ...string) {
//	fmt.Println(strings.Join(name, ", "))
//}
//
//func main() {
//	repeatMe("dongmin", "min", "dm")
//}

// naked return
//func lenAndUpper2(name string) (length int, uppercase string) {
//	defer fmt.Println("finished") // defer는 함수가 끝나기 직전에 실행됨
//	defer func() {
//		for i := 0; i < 10; i++ {
//			fmt.Println(i)
//		}
//	}()
//	length = len(name)
//	uppercase = strings.ToUpper(name)
//	fmt.Println("not finished")
//	return
//}
//
//func main() {
//	length, upperName := lenAndUpper2("dongmin")
//	println("Length:", length, "Uppercase Name:", upperName)
//}

//func superAdd(numbers ...int) int {
//	total := 0
//	for _, number := range numbers {
//		total += number
//	}
//	return total
//}
//
//func main() {
//	result := superAdd(1, 2, 3, 4, 5)
//	fmt.Println("Result:", result)
//}

//func canIDrink(age int) bool {
//	if korean := age + 2; korean >= 20 {
//		return true
//	}
//	return false
//}
//
//func main() {
//	result := canIDrink(18)
//	println("Can I drink?", result)
//}

//func main() {
//	a := 10
//	b := &a
//	*b = 20
//	fmt.Println(*b)
//	fmt.Println(a)
//}
//
//func main() {
//	names := [5]string{"dongmin", "min", "dm", "hello", "world"}
//	slice := names[0:3]
//	fmt.Println(names)
//	slice[0] = "changed"
//	fmt.Println(slice)
//
//	longName := []string{"dongmin", "min", "dm", "hello", "world"}
//	i := append(longName, "new1", "new2")
//	fmt.Println(longName)
//	fmt.Println(i)
//}

//func main() {
//	dongmin := map[string]int{
//		"name":    12,
//		"age":     23,
//		"address": 23,
//	}
//	fmt.Println(dongmin)
//	dongmin["age"] = 40
//	fmt.Println(dongmin)
//
//	for _, v := range dongmin {
//		fmt.Println(v)
//	}
//}

type person struct {
	name         string
	age          int
	favoriteFood []string
}

func main() {
	dongmin := person{
		name:         "dongmin",
		age:          23,
		favoriteFood: []string{"Chicken", "Pizza"},
	}
	fmt.Println(dongmin)
	fmt.Println(dongmin.name)
}
