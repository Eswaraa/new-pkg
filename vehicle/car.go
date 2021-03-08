package vehicle

import "fmt"

type Car struct {
	CarNo string
	Model string
	Make  string
}

func print(c Car) {
	fmt.Printf("CarNo: %s", c.CarNo)
}
