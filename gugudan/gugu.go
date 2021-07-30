package gugudan

import "fmt"

var a int
var b int

func Gugu(x, y int) int{
	return x*y
}

func main(){
	n, err := fmt.Scan(&a, &b)
	if err != nil{
		fmt.Println(err, n)
	} else {
		result := Gugu(a, b)
		fmt.Println(a, "X", b, "의 값은 =", result)
	}
}