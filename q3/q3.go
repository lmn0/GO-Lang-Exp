package main
import ("fmt";"time")

func Sleep(x int) {
    <-time.After(time.Second * time.Duration(x))
}

func main(){
fmt.Println("Sleeping!")
Sleep(10)
fmt.Println("Awake!")
}
