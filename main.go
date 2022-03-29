package main

import (
	"fmt"
	"time"

	u "example.com/utils"
)

func main() {
	user := u.CreateUser()
	fmt.Println("Initial user:", user)
	user.ChangeUserName("Tobi")
	fmt.Println("User changed passing struct's method", user)
	u.UpdateName(&user.Name)
	fmt.Println("User changed passing reference", user)

	c := make(chan int)
	go Wait(c)

	for i := 1; i <= 5; i++ {
		c <- i
		time.Sleep(time.Second * 2)
	}
}

func Wait(ch chan int) {
	for {
		n := <-ch
		fmt.Println("Received", n)
	}
}
