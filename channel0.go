// This sample program demonstrates the basic channel mechanics
// for goroutine signaling.
package main

import (
	"context"
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {

	waitForResult()
	// fanOut()

	// waitForTask()
	// pooling()

	// Advanced patterns
	// fanOutSem()
	// boundedWorkPooling()
	// drop()

	// Cancellation Pattern
	// cancellation()

	// Retry Pattern
	// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// defer cancel()
	// retryTimeout(ctx, time.Second, func(ctx context.Context) error { return errors.New("always fail") })
}

// waitForResult: You are a manager and you hire a new employee. Your new
// employee knows immediately what they are expected to do and starts their
// work. You sit waiting for the result of the employee's work. The amount
// of time you wait on the employee is unknown because you need a
// guarantee that the result sent by the employee is received by you.

func waitForResult() {
	ch := make(chan string)

	go func() {
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
		ch <- "paper"
		fmt.Println("employee : sent signal")
	}()


	p := <-ch
	fmt.Println("manager : recv'd signal :", p)

	time.Sleep(time.Second)
	fmt.Println("-------------------------------------------------------------")
}

// fanOut: You are a manager and you hire one new employee for the exact amount
// of work you have to get done. Each new employee knows immediately what they
// are expected to do and starts their work. You sit waiting for all the results
// of the employees work. The amount of time you wait on the employees is
// unknown because you need a guarantee that all the results sent by employees
// are received by you. No given employee needs an immediate guarantee that you
// received their result.
func fanOut() {
	emps := 2000
	ch := make(chan string, emps)

	for e := 0; e < emps; e++ {
		go func(emp int) {
			time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
			ch <- "paper"
			fmt.Println("employee : sent signal :", emp)
		}(e)
	}


	for emps > 0 {
		p := <-ch
		emps--
		fmt.Println(p)
		fmt.Println("manager : recv'd signal :", emps)
	}

	time.Sleep(time.Second)
	fmt.Println("-------------------------------------------------------------")
}

// waitForTask: You are a manager and you hire a new employee. Your new
// employee doesn't know immediately what they are expected to do and waits for
// you to tell them what to do. You prepare the work and send it to them. The
// amount of time they wait is unknown because you need a guarantee that the
// work your sending is received by the employee.
func waitForTask() {
	ch := make(chan string)

	go func() {
		p := <-ch
		fmt.Println("employee : recv'd signal :", p)
	}()


	time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
	ch <- "paper"
	fmt.Println("manager : sent signal")

	time.Sleep(time.Second)
	fmt.Println("-------------------------------------------------------------")
}
// pooling: You are a manager and you hire a team of employees. None of the new
// employees know what they are expected to do and wait for you to provide work.
// When work is provided to the group, any given employee can take it and you
// don't care who it is. The amount of time you wait for any given employee to
// take your work is unknown because you need a guarantee that the work your
// sending is received by an employee.
func pooling() {
	ch := make(chan string)

	g := runtime.GOMAXPROCS(0)
	for e := 0; e < g; e++ {
		go func(emp int) {
			for p := range ch {
				fmt.Printf("employee %d : recv'd signal : %s\n", emp, p)
			}
			fmt.Printf("employee %d : recv'd shutdown signal\n", emp)
		}(e)
	}
