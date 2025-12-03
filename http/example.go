package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"sync"
)

type countHandler struct {
	mu  sync.Mutex
	cnt int
}

func (h *countHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.mu.Lock()
	defer h.mu.Unlock()

	h.cnt++
	fmt.Println("the count is", h.cnt)
}

func newPeopleHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "this is the people handler.")
	})
}

func main() {
	/*http.Handle("/count", new(countHandler))
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatal(err)
	}*/

	/*mux := http.NewServeMux()

	mux.Handle("/resource", http.NotFoundHandler())
	mux.Handle("/resource/people", newPeopleHandler())

	err := http.ListenAndServe(":8082", mux)
	if err != nil {
		log.Fatal(err)
	}*/

	resp, err := http.Get("https://httpbin.org/get")
	if err != nil {
		log.Panic(err)
	}

	respData, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		log.Panic(err)
	}

	fmt.Printf("%s\n", respData)

	fmt.Println("get method", resp.Status)

	respPost, err := http.Post("https://httpbin.org/post", "application/json", nil)
	if err != nil {
		log.Panic(err)
	}

	fmt.Println("post method", respPost.Status)

	req, err := http.NewRequest(http.MethodPut, "https://httpbin.org/put", nil)
	if err != nil {
		log.Panic(err)
	}

	req.Header.Set("Content-Type", "application/json")

	respPut, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Panic(err)
	}

	fmt.Println("put method", respPut.Status)

	req2, err := http.NewRequest(http.MethodDelete, "https://httpbin.org/delete", nil)
	if err != nil {
		log.Panic(err)
	}

	respDelete, err := http.DefaultClient.Do(req2)
	if err != nil {
		log.Panic(err)
	}

	fmt.Println("delete method", respDelete.Status)

	req3, err := http.NewRequest(http.MethodGet, "https://httpbin.org/get", nil)
	if err != nil {
		log.Panic(err)
	}

	params := make(url.Values)
	params.Set("page_num", "1")
	params.Add("page_size", "10")
	req3.URL.RawQuery = params.Encode()

	respGetParams, err := http.DefaultClient.Do(req3)
	if err != nil {
		log.Panic(err)
	}

	fmt.Println("get method with url params", respGetParams.Status)
}
