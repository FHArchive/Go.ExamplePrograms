//go 1.10.4

// Copyright 2019 Kieran W.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

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
