package main

import (
	"fmt"
	"testing"
	"time"
)

func TestAfterFunc(t *testing.T) {
	var c chan int

	select {
	case m := <-c:
		fmt.Println(m)
	case <-time.After(2 * time.Second):
		fmt.Println("timeout")
	}
}

func TestTickFunc(t *testing.T) {
	timer := time.Tick(time.Second)

	for next := range timer {
		fmt.Println(next)
	}
}

func TestDurationFunc(t *testing.T) {
	t1 := time.Now()
	time.Sleep(200 * time.Millisecond)
	t2 := time.Now()
	d1 := t2.Sub(t1)
	fmt.Println("duration is", d1)

	d2, _ := time.ParseDuration("1h20m53s")
	fmt.Println("parse duration is", d2)
	fmt.Println("parse duration hours is", d2.Hours())

	d3 := time.Since(t1)
	fmt.Println("duration since t1 is", d3)
}

func TestLocationFunc(t *testing.T) {
	location, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		panic(err)
	}

	timeInBJ := time.Date(2025, 11, 15, 11, 45, 0, 0, time.Local)
	fmt.Println(timeInBJ.In(location))
}

func TestMonthFunc(t *testing.T) {
	y, m, d := time.Now().Date()

	if m == time.November {
		fmt.Println(y, m, d)
	}
}

func TestTickerFunc(t *testing.T) {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	done := make(chan bool, 1)

	go func() {
		time.Sleep(5 * time.Second)
		done <- true
	}()

	for {
		select {
		case <-done:
			fmt.Println("done done done")
			return
		case t := <-ticker.C:
			fmt.Println(t)
		}
	}
}
