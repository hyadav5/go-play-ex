package main

import "time"

func main() {
	//wp := MakeWorkerPool(3, 1)
	//start := time.Now()
	//c := make(chan bool, 1)
	//foo := func() {
	//	time.Sleep(10 * time.Second)
	//	print("I ran 10 secs")
	//	c <- true
	//}
	//wp.DoWork(foo)
	//wp.DoWork(foo)
	//wp.DoWork(foo)
	//for i := 0; i < 3; i++ {
	//	<-c
	//}
	//ti := time.Now().Sub(start)
	//if ti < 18*time.Millisecond || ti > 22*time.Millisecond {
	//	print("hey their")
	//}

	wp := MakeKeyedWorkerPool(3, 1)
	c := make(chan string, 100)
	wp.DoWork("service1", func() {
		print("service1 is doing some work. ID: ", wp.GetWorkerID("service1"), "\n")
		time.Sleep(5 * time.Second)
		c <- "service1"
	})
	wp.DoWork("service3", func() {
		print("service3 is doing some work. ID: ", wp.GetWorkerID("service3"), "\n")
		time.Sleep(15 * time.Second)
		c <- "service3"
	})
	wp.DoWork("service2", func() {
		print("service2 is doing some work. ID: ", wp.GetWorkerID("service2"), "\n")
		time.Sleep(10 * time.Second)
		c <- "service2"
	})

	for i := 0; i < 3; i++ {
		x := <-c
		print("Got: ", x, "\n")
	}
}
