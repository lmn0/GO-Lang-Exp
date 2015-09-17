package main
import "fmt"

func Fibonacci(n int) int{
    if n==1 {
        return 0
    }else if n==2 {
        return 1
    }else{
        return(Fibonacci(n - 1) + Fibonacci(n - 2))}
}

func main(){
    var n int
    var value int
    fmt.Println("Enter n: ")
    fmt.Scanf("%d", &n)
    if n<0 {
        fmt.Println("N/A")
    }else{
        value = Fibonacci(n)
        fmt.Println(value)
    }
}

