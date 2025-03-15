package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	//delayTime := time.Second * 1
	//delay, exists := os.LookupEnv("DELAY")
	//if exists {
	//	strconv.Atoi(delay)
	//	var err error
	//	delayTime, err = time.ParseDuration(delay)
	//	if err != nil {
	//		delayTime = time.Second
	//	}
	//}
	//
	//fmt.Printf("starting in .... %v\n", delayTime)
	//time.Sleep(delayTime)
	fmt.Printf("started\n")

	http.HandleFunc("/", show)
	http.HandleFunc("/healthcheck", healthcheck)
	http.HandleFunc("/readinesscheck", readinesscheck)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func show(w http.ResponseWriter, r *http.Request) {
	//	time.Sleep(time.Second * 10)
	fmt.Fprintf(w, "Hello, from %s ", r.Host)
}

func readinesscheck(w http.ResponseWriter, r *http.Request) {
	if os.Getenv("READINESS") == "true" {
		fmt.Fprintln(w, "ready")
		return
	}
	if _, err := os.Stat("r"); err != nil && os.IsExist(err) {
		w.WriteHeader(500)
		fmt.Println("FAIL - readinesscheck")
		fmt.Fprintf(w, "FAIL - readinesscheck")
		return
	}
	fmt.Println("OK - readinesscheck")
	fmt.Fprintf(w, "OK - readinesscheck")
}

func healthcheck(w http.ResponseWriter, r *http.Request) {
	if _, err := os.Stat("h"); err != nil && os.IsNotExist(err) {
		fmt.Println("OK - healthcheck")
		fmt.Fprintf(w, "OK - healthcheck")
		return
	}
	w.WriteHeader(500)
	fmt.Println("FAIL - healthcheck")
	fmt.Fprintf(w, "FAIL - healthcheck")
}
