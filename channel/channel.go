package channel

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string, 2)
	people := [5]string{"nico", "flym", "dal", "woo", "dwd"}
	for _, person := range people {
		go isSexy(person, c)
	}
	for i := 0; i < len(people); i++ {
		fmt.Println("waiting for " + people[i])
		fmt.Println(<-c)
	}
}

func sexyCount(person string) {
	for i := 0; i < 10; i++ {
		fmt.Println(person, "is sexy", i)
		time.Sleep(time.Second)
	}
}

func isSexy(person string, c chan string) {
	time.Sleep(time.Second * 5)
	c <- person + " is sexy"
}
