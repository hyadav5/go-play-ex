package main

import "time"

func doKeyedWork(wp *KeyedWorkerPool, c chan string, key string, waitInSec int) {
	wp.DoWork(key, func() {
		print(key, " is doing some work. ID: ", wp.GetWorkerID(key), "\n")
		time.Sleep(time.Duration(waitInSec) * time.Second)
		c <- key
	})
}

func doKeyedWorkCustom(wp *KeyedWorkerPool, c chan string, key string, myWork func()) {
	wp.DoWork(key, func() {
		print(key, " is doing some work. ID: ", wp.GetWorkerID(key), "\n")
		myWork()
		c <- key
	})
}

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

	wp := MakeKeyedWorkerPool(3, 100)
	c := make(chan string, 100)

	//doKeyedWork(wp, c, "service1", 5)
	//doKeyedWork(wp, c, "service2", 10)
	//doKeyedWork(wp, c, "service3", 15)

	doKeyedWorkCustom(wp, c, "service1", func() {
		print("service1 doing its custom work \n")
		time.Sleep(time.Duration(5) * time.Second)
	})
	doKeyedWorkCustom(wp, c, "service2", func() {
		print("service2 doing its custom work \n")
		time.Sleep(time.Duration(10) * time.Second)
	})
	doKeyedWorkCustom(wp, c, "service3", func() {
		print("service3 doing its custom work \n")
		time.Sleep(time.Duration(15) * time.Second)
	})

	//workerChannelLength := wp.WorkerChannelLength(wp.GetWorkerID("service3"))
	//print("service3 workerChannelLength = ", workerChannelLength, "\n")
	//
	//length := wp.Length()
	//print("length = ", length, "\n")

	for i := 0; i < 3; i++ {
		x := <-c
		print("Got: ", x, "\n")
	}
}
