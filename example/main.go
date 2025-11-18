package main

import (
	"cmp"
	"crypto/sha256"
	"fmt"
	"slices"
	"strings"
	"time"
	"unicode/utf8"

	b64 "encoding/base64"
	s "strings"
)

func main() {
	PriorityQueueFunc()
}

func printFunc() {
	fmt.Println("hello")
	fmt.Println("world")

	fmt.Print("hello")
	fmt.Print("world")
	fmt.Print("\n")

	n, err := fmt.Printf("%s, %s", "hello", "world\n")
	if err != nil {
		fmt.Println("error")
	}

	fmt.Println("the count is", n)
}

func cmpfunc() {
	fmt.Println(cmp.Compare(1, 2))
	fmt.Println(cmp.Compare(1, 1))
	fmt.Println(cmp.Compare(2, 1))

	fmt.Println(cmp.Less(1, 2))
	fmt.Println(cmp.Less(2, 1))

	fmt.Println(cmp.Or("", "default"))
	fmt.Println(cmp.Or("hhhh", "default"))

	type order struct {
		product  string
		customer string
		price    int
	}

	orders := []order{
		{"foo", "alice", 1.00},
		{"bar", "bob", 3.00},
		{"baz", "carol", 4.00},
		{"foo", "alice", 2.00},
		{"bar", "carol", 1.00},
		{"foo", "bob", 4.00},
	}

	slices.SortFunc(orders, func(a, b order) int {
		return cmp.Or(
			strings.Compare(a.customer, b.customer),
			strings.Compare(a.product, b.product),
			cmp.Compare(b.price, a.price),
		)
	})

	for _, order := range orders {
		fmt.Println(order)
	}
}

func encoder() {
	s := "Hello, World!"
	h := sha256.New()
	h.Write([]byte(s))
	bs := h.Sum(nil)

	fmt.Printf("sha256 hash of '%s': %x\n", s, bs)

	ss := "dffd6021bb2bd5b0af676290809ec3a53191dd81c7f70a4b28688a362182986f"
	fmt.Printf("the length of ss is %d", len(ss))

	data := "abc123!?$*&()'-=@~"
	fmt.Println("the length of data is:", len(data))
	sEnc := b64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println("Encoded:", sEnc)
	fmt.Println("the length of sEnc is:", len(sEnc))

	sDec, _ := b64.StdEncoding.DecodeString(sEnc)
	fmt.Println("Decoded:", string(sDec))
}

func stringfunc() {
	var p = fmt.Println

	p("contains:", s.Contains("test", "es"))
	p("count:", s.Count("test", "t"))
	p("hasPrefix:", s.HasPrefix("test", "te"))
	p("hasSuffix:", s.HasSuffix("test", "st"))
	p("index:", s.Index("test", "e"))
	p("join:", s.Join([]string{"a", "b", "c"}, "-"))
	p("repeat:", s.Repeat("a", 5))
	p("replace:", s.Replace("foo bar foo", "foo", "baz", -1))
	p("split:", s.Split("a-b-c", "-"))
	p("toLower:", s.ToLower("TEST"))
	p("toUpper:", s.ToUpper("test"))

	type point struct {
		x int
		y int
	}

	po := point{x: 1, y: 2}
	fmt.Printf("struct1: %v\n", po)
	fmt.Printf("struct2: %+v\n", po)
	fmt.Printf("struct3: %#v\n", po)

	fmt.Printf("type: %T\n", po)
}

func mypanic() {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main:", r)
		}
	}()

	fmt.Println("After panic")

	panic("a problem occurred")

}

func sortfunc() {
	s := []string{"cat", "bat", "rat", "elephant"}
	slices.Sort(s)
	fmt.Println(s)

	type Person struct {
		name string
		age  int
	}

	people := []Person{
		{"Bob", 31},
		{"John", 42},
		{"Michael", 17},
		{"Jenny", 26},
	}

	slices.SortFunc(people, func(a, b Person) int {
		return cmp.Compare(a.age, b.age)
	})

	fmt.Println("Sorted by age:", people)
}

func wker(id int) {
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func multigo() {
	/*gofunc("main")

	go gofunc("goroutine")

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	time.Sleep(time.Second)

	done := make(chan bool, 1)
	go worker(done)

	if <-done {
		fmt.Println("Worker completed")
	}*/

	/*c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(time.Second)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	for range 3 {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		case <-time.After(3 * time.Second):
			fmt.Println("timeout")
		}
	}*/

	/*jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	for j := range 5 {
		jobs <- j
		fmt.Println("sent job", j)
	}

	close(jobs)
	fmt.Println("sent all jobs")

	<-done
	_, ok := <-jobs
	fmt.Println("jobs channel closed?", !ok)

	q := make(chan int, 2)

	q <- 1
	q <- 2
	close(q)

	for i := range q {
		fmt.Println(i)
	}*/

	/*timer1 := time.NewTimer(3 * time.Second)
	timer2 := time.NewTimer(time.Second)

	<-timer1.C
	fmt.Println("Timer 1 expired")

	go func() {
		<-timer2.C
		fmt.Println("Timer 2 expired")
	}()

	if timer1.Stop() {
		fmt.Println("Timer 1 stopped")
	} else {
		fmt.Println("Timer 1 already expired")
	}

	if timer2.Stop() {
		fmt.Println("Timer 2 stopped")
	}*/

	/*ticker := time.NewTicker(500 * time.Millisecond)
		done := make(chan bool)

		go func() {
			for {
				select {
				case <-done:
					return
				case t := <-ticker.C:
					fmt.Println("Tick at", t)
				}
			}
		}()

		time.Sleep(2200 * time.Millisecond)
		ticker.Stop()
		done <- true
		fmt.Println("Ticker stopped")

	}

	func worker(done chan bool) {
		fmt.Println("working...")
		time.Sleep(time.Second)
		fmt.Println("done")

		done <- true*/

	const n = 6
	jobs := make(chan int, n)
	results := make(chan int, n)

	for i := range 3 {
		go workerforjobs(i, jobs, results)
	}

	for i := range n {
		jobs <- i
	}
	close(jobs)

	for range n {
		<-results
	}
}

func workerforjobs(id int, jobs <-chan int, result chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		result <- j
	}
}

func gofunc(name string) {
	for i := range 3 {
		fmt.Println(name, ":", i)
	}
}

func other() {
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend!")
	default:
		fmt.Println("It's a weekday.")
	}

	var a []int = []int{1, 2, 3, 4, 5}
	fmt.Println(a[0])

	s := make([]string, 3)
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println(s)

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println(s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println(c)

	l := s[2:5]
	fmt.Println(l)

	l = s[:5]
	fmt.Println(l)

	l = s[2:]
	fmt.Println(l)

	m := make(map[string]int)
	m["foo"] = 1
	m["bar"] = 2
	fmt.Println(m)

	delete(m, "foo")
	fmt.Println(m)

	clear(m)
	fmt.Println(m)

	result := plus(3, 4)
	fmt.Println(result)

	va, vb := vals()
	fmt.Println(va)
	fmt.Println(vb)

	sum(1, 2)
	sum(1, 2, 3, 4, 5)

	nums := []int{6, 7, 8}
	sum(nums...)

	for i, c := range "letusgo" {
		fmt.Println(i, c)
	}

	const ss = "我是中国人"
	fmt.Println(len(ss))
	fmt.Println("rune count:", utf8.RuneCountInString(ss))
	for i, runeval := range ss {
		fmt.Printf("%d %c\n", i, runeval)
	}
}

func plus(a int, b int) int {
	return a + b
}

func vals() (int, int) {
	return 3, 7
}

func sum(nums ...int) {
	fmt.Print(nums, ",")

	total := 0
	for _, num := range nums {
		total += num
	}

	fmt.Println(total)
}

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
