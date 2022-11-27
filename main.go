package main

import (
	"fmt"

	"github.com/abhishek2966/builder-demo/pkg/resource"
)

func main() {
	eb := resource.NewEmployeeBuilder()
	e1 := eb.PersonalInfo().
		AtHouseNo(212).
		InSociety("Wildstone").
		HavingPin(560001).
		WorkInfo().
		WithID(101).
		WithEmail("abc@email.com").
		Build()

	fmt.Printf("%+v\n", e1)
	//&{Id:101 Email:abc@email.com Address:{HouseNo:212 Society:Wildstone Pin:560001}}
}
