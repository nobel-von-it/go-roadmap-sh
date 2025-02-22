package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

type tmpUser struct {
	id   int
	name string
}

var (
	id    int
	users = make(map[int]tmpUser)
)

func event(u tmpUser, ctx context.Context) {
	users[u.id] = u
	select {
	case <-ctx.Done():
		delete(users, u.id)
		log.Printf("Deleted user %s", u.name)
	}
}

func addUser(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), time.Duration(rand.Intn(3))*time.Second)
	defer cancel()

	query := r.URL.Query()
	name := query.Get("n")
	if name == "" {
		http.Error(w, "Missing name parameter", http.StatusBadRequest)
		return
	}

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for range 10 {
			select {
			case <-ctx.Done():
				break
			default:
				time.Sleep(time.Millisecond * 100)
			}
		}
	}()

	done := make(chan struct{})
	go func() {
		defer close(done)
		wg.Wait()
	}()

	select {
	case <-ctx.Done():
		http.Error(w, ctx.Err().Error(), http.StatusInternalServerError)
		return
	case <-done:
		fmt.Fprintf(w, "Added user %s\n", name)
	}
}

func showAll(w http.ResponseWriter, r *http.Request) {
	for i, u := range users {
		fmt.Fprintf(w, "Id: %d Name is %s\n", i, u.name)
	}
	// Bruh
}

func main() {
	http.HandleFunc("/new", addUser)
	http.HandleFunc("/", showAll)

	log.Fatalln(http.ListenAndServe(":8080", nil))
}
