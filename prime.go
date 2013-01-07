package main

import "fmt"

func main() {
	fmt.Println("生成素数算法")
	prime(10)
}

func gen(ch chan<- int) {
	for i := 2; ; i++ {
		ch <- i
	}
}

func filter(cin <-chan int, cout chan<- int, prime int) {
	for {
		i := <-cin
		if i%prime != 0 {
			cout <- i
		}
	}
}

func prime(n int) {
	cin := make(chan int)
	go gen(cin)
	for i := 0; i < n; i++ {
		prime := <-cin
		fmt.Print(prime, ", ")
		cout := make(chan int)
		go filter(cin, cout, prime)
		cin = cout
	}
	fmt.Println()
}
