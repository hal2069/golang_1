package main
import (
"fmt"
"net/http"
// "os"
)

func zeroval(ival int){
  ival = 0
}

func zeroptr(iptr *int){
  *iptr = 0
}

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {

    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)

 i := 1
 fmt.Println("initial:", i)
 zeroval(i)
 fmt.Println("zeroval:", i)
 zeroptr(&i)
 fmt.Println("zeroptr:", i)
 fmt.Println("pointer:", &i)
}
