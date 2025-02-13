package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	name := "TEST"
	value, exists := os.LookupEnv(name)
	fmt.Printf("%s - %s - %v\n", name, value, exists)

	for i := 0; i < 10; i++ {
		do(i, value)
		time.Sleep(time.Second)
	}
}
func do(i int, value string) {
	response, err := http.Get(value)
	if err != nil {
		panic(err)
		os.Exit(0)
	}
	fmt.Println(i, response.StatusCode)
	io.Copy(os.Stdout, response.Body)
	fmt.Println("=--------------")

}
