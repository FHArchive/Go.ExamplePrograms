//go 1.10.4

package main
import "fmt"

func sqrt(x float64) float64 {
	
	z := 1.0
	for iteration := 0; iteration < 10; iteration ++{
		z -= (z*z - x) / (2*z)
		
		fmt.Printf("%f\n", z)
		
	}
	return z
}


func main() {
	val := 2.0
	fmt.Printf("The root of %f is %f", val, sqrt(val))
	
}