package main

import "fmt"
import "io/ioutil"
import "net/http"
//import "os"

func main() {
	var a [2]string
	a[0] = "http://google.pl"

	for _, url:= range a {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Println( "test", err)
		}
		b, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			fmt.Println("Blad2")

		}
    defer resp.Body.Close()
		fmt.Printf("%s100",b)
	}
}
