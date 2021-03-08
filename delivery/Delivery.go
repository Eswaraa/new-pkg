package delivery

import (
	"errors"
	"fmt"
)

type Representative struct {
	Name   string
	Mobile string
	Place  string
}
type DeliveryOrg struct {
	Reps map[string]Representative
}

func (d *DeliveryOrg) AddDeliveryRep(rep Representative) error {

	if d.Reps == nil {
		d.Reps = make(map[string]Representative)

	}
	for _, v := range d.Reps {
		if v.Mobile == rep.Mobile {
			return errors.New("Mobile already exists")
		}
	}
	// if d.Reps[rep.Mobile]  {
	// 	return errors.New("Mobile already exists")
	// }
	// d.Reps = append(d.Reps, rep)
	d.Reps[rep.Mobile] = rep
	fmt.Printf("Delivery Rep Added ", rep.Name)
	return nil
}
