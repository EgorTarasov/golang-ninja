package main

import "fmt"

type AutomobileSpecification struct {
	Name               string
	NDoors             int
	RimCenterBore      float32
	RimBoltPatter      string
	Petrol             string
	EngineVolumeLiters float32
	EngineName         string
	HorsePower         float32
	WeightKilogram     int
	HasSunroof         bool
	HasRoof            bool
}


func isSportCar(spec *AutomobileSpecification) bool {
	var NSportFeatures int
	if spec.NDoors < 4 {
		NSportFeatures += 1
	}

	if spec.HorsePower > 150 {
		NSportFeatures += 1
	}

	if spec.WeightKilogram < 1050 {
		NSportFeatures += 1
	}

	return NSportFeatures > 2
}

func main() {
	car1 := &AutomobileSpecification{
		Name:               "ВАЗ2106",
		NDoors:             4,
		RimCenterBore:      58.5,
		RimBoltPatter:      "4x98",
		Petrol:             "A92",
		EngineVolumeLiters: 1.57,
		EngineName:         "2106",
		HorsePower:         71.5,
		WeightKilogram:     1035,
		HasSunroof:         false,
		HasRoof:            true,
	}

	fmt.Printf("is %s a sport car? %v\n", car1.Name, isSportCar(car1))
	// did something change?
	fmt.Printf("is %s a sport car? %v\n", car1.Name, isSportCar(car1))
	// did something change?
	fmt.Printf("is %s a sport catr %v\n", car1.Name, isSportCar(car1))
}
