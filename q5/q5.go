package main
import "fmt"

func runtime(){
	defer func(){
		str:=recover()
		fmt.Println(str)
	}()
	a:=[]int{1,2,3}
	fmt.Println(a[4])
}
func main(){
	fmt.Println("Calling function runtimeerror")
	runtime()
	fmt.Println("Done!")
}
