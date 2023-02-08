package main
import (
	"fmt"
)

func main () {
	// x := type{values} | Composite Literal
	x := []int{4,5,7,8,42}
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(x[4])
	fmt.Println(x[2:])
	fmt.Println(x[1:3])

	fmt.Println("===<::>===")
	
	for i, v := range x {
		fmt.Println("| ", i, v)
	}

	fmt.Println("===<::>===")

	for i := 0; i <= (len(x) - 1); i++ {
		fmt.Printf("x[%d]: %d\n", i, x[i]);
	}
}