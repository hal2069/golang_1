package main
import ("fmt"
    "time")

func main() {
    ch := make(chan string)
    ch1 := make(chan time.Time)
    g := "hello"
    go fillChWithSmth(ch1)
    select {
        case e := <-ch :
            fmt.Println("Z kanału ch otrzymaliśmy wartość", e)
        case f := <-ch1:
            fmt.Println("Z kanału ch1 otrzymaliśmy wartość", f)
        case ch <- g:
            fmt.Println("Do kanału ch wysłaliśmy wartość", g)
    }
}

func fillChWithSmth(ch chan time.Time) {
    time.Sleep(10 * time.Second)
    ch <- time.Now()
}
