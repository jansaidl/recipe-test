package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {

	for i := 0; i < 10; i++ {
		fmt.Fprintf(os.Stderr, "hele to je chyba %d\n", i)
		fmt.Fprintf(os.Stdout, "hele to je chyba %d\n", i)
	}

	os.Exit(1)
	//name := "TEST"
	//value, exists := os.LookupEnv(name)
	//fmt.Printf("%s - %s - %v\n", name, value, exists)
	//
	//for i := 0; i < 10; i++ {
	//	do(i, value)
	//	time.Sleep(time.Second)
	//}
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
