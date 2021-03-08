package main

import (
	"eswork/go/src/github.com/Eswaraa/new-pkg/delivery"
	"eswork/go/src/github.com/Eswaraa/new-pkg/vehicle"
	"fmt"
	"log"
	"time"
)

func main() {
	fmt.Println("Go Packages / Module")
	t := time.Now()
	fmt.Printf("\n Time: %v , Year: %d , Location: %v", t, t.Year(), t.Local())
	log.Println("\nCompleted Main")
	const shortForm = "2006-Jan-02"
	t1, err := time.Parse(shortForm, "2013-Feb-03")

	if err != nil {
		log.Println("\n Error Parsing")
	}

	fmt.Printf("\n Time: %v", t1)

	c := vehicle.Car{CarNo: "KA01 GH 3245", Model: "XUV 500", Make: "Mahindra"}

	fmt.Printf("\n Time: %v", c)

	delOrg := delivery.DeliveryOrg{}
	delRep := delivery.Representative{
		Name:   "Logesh",
		Mobile: "3546437777",
		Place:  "Indiranagar",
	}
	err = delOrg.AddDeliveryRep(delRep)
	if err != nil {
		fmt.Println("Error: %v", err)
	} else {
		fmt.Println("\n Successfully added")
	}
	err = delOrg.AddDeliveryRep(delRep)
	if err != nil {
		fmt.Println("Error: %v", err)
	} else {
		fmt.Println("\n Successfully added")
	}

}
