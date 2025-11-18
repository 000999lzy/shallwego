package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func syscFunc() {
	//onceFunc()
	//onceValFunc()
	//waitGroupFunc()
	//mutexFunc()
	//atomicFunc()
	condFunc()
}

func condFunc() {
	var mu sync.Mutex
	cond := sync.NewCond(&mu)

	for i := range 10 {
		go func(x int) {
			cond.L.Lock()
			defer cond.L.Unlock()

			cond.Wait()
			fmt.Println(x)
		}(i)
	}

	time.Sleep(time.Second)
	fmt.Println("signal...")
	cond.Signal()

	time.Sleep(time.Second)
	fmt.Println("signal...")
	cond.Signal()

	time.Sleep(time.Second)
	fmt.Println("broadcast...")
	cond.Broadcast()

	select {}
}

func atomicFunc() {
	//var counter atomic.Int32

	counter := int32(0)

	var wg sync.WaitGroup
	for range 1000 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomic.AddInt32(&counter, 1)
		}()
	}

	wg.Wait()
	fmt.Println("final counter value:", counter)

}

func mutexFunc() {
	var mu sync.Mutex
	counter := 0

	var wg sync.WaitGroup
	for range 1000 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			counter++
			mu.Unlock()
		}()
	}

	wg.Wait()
	fmt.Println("final counter value:", counter)
}

func waitGroupFunc() {
	var wg sync.WaitGroup

	for i := range 5 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Printf("goroutine %d done\n", i)
		}()
	}

	wg.Wait()
	fmt.Println("all goroutines done")
}

func onceValFunc() {
	once := sync.OnceValue(func() int {
		sum := 0
		for i := 0; i < 1000; i++ {
			sum += i
		}

		fmt.Println("compute once:", sum)
		return sum
	})

	done := make(chan bool)

	for range 10 {
		go func() {
			const wanted = 499500
			v := once()
			if v != wanted {
				fmt.Printf("unexpected value: got %d, want %d\n", v, wanted)
			}
			done <- true
		}()
	}

	for i := range 10 {
		<-done
		fmt.Printf("main done--%d\n", i)
	}
}

func onceFunc() {
	var once sync.Once

	onceBody := func() {
		fmt.Println("only once")
	}

	done := make(chan bool)

	for i := range 10 {
		go func() {
			once.Do(onceBody)
			fmt.Printf("goroutine done--%d\n", i)
			done <- true
		}()
	}

	for i := range 10 {
		<-done
		fmt.Printf("main done--%d\n", i)
	}
}
