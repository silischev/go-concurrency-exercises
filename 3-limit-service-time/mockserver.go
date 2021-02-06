//////////////////////////////////////////////////////////////////////
//
// DO NOT EDIT THIS PART
// Your task is to edit `main.go`
//

package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

// RunMockServer pretends to be a video processing service. It
// simulates user interacting with the Server.
func RunMockServer() {
	u1 := User{ID: 0, IsPremium: false}
	u2 := User{ID: 1, IsPremium: true}

	wg.Add(5)

	go createMockRequest(1, shortProcess, &u1)
	time.Sleep(1 * time.Second)

	go createMockRequest(2, longProcess, &u2)
	time.Sleep(2 * time.Second)

	go createMockRequest(3, shortProcess, &u1)
	time.Sleep(1 * time.Second)

	go createMockRequest(4, longProcess, &u1)
	go createMockRequest(5, shortProcess, &u2)

	wg.Wait()
}

func createMockRequest(pid int, fn func(), u *User) {
	fmt.Println("UserID:", u.ID, "\tProcess", pid, "started.")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	status := "Basic"
	if u.ID == 1 {
		ctx = context.Background()
		status = "Premium"
	}

	res := HandleRequest(ctx, fn, u)

	if res {
		fmt.Println("UserID:", u.ID, " Status:", status, "\tProcess", pid, "done.")
	} else {
		fmt.Println("UserID:", u.ID, " Status:", status, "\tProcess", pid, "killed. (No quota left)")
	}

	wg.Done()
}

func shortProcess() {
	time.Sleep(6 * time.Second)
}

func longProcess() {
	time.Sleep(11 * time.Second)
}
