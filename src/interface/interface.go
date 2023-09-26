// --Summary:
//  Create a program that directs vehicles at a mechanic shop
//  to the correct vehicle lift, based on vehicle size.
//
//--Requirements:
//* The shop has lifts for multiple vehicle sizes/types:
//  - Motorcycles: small lifts
//  - Cars: standard lifts
//  - Trucks: large lifts
//* Write a single function to handle all of the vehicles
//  that the shop works on.
//* Vehicles have a model name in addition to the vehicle type:
//  - Example: "Truck" is the vehicle type, "Road Devourer" is a model name
//* Direct at least 1 of each vehicle type to the correct
//  lift, and print out the vehicle information.
//
//--Notes:
//* Use any names for vehicle models

// package main

// import "fmt"

// const (Smalllift =iota
//        StandardLift
// 	   LargeLift
// 	)

// type Lift int 

// type LiftPicker interface{
// 	PickLift() Lift
// }

// type Motorcycle string
// type Car string
// type Truck string


// func (m Motorcycle) String() string{
// 	return fmt.Sprintf("MotorCycle : %v",string(m))
// }

// func (c Car) string()string{
// 	return fmt.Sprintf("Car : %v",string(c))
// }

// func (t Truck) String()string{
// 	return fmt.Sprintf("Truck : %v",string(t))
// }

// func (m Motorcycle) PickLift() Lift{
// 	return Smalllift

// }

// func (c Car)PickLift()Lift{
// 	return StandardLift
// }

// func (t Truck)PickLift()Lift{
// 	return LargeLift
// }

// func sendToLift(p LiftPicker){
// 	switch p.PickLift(){
// 	case Smalllift:
// 		fmt.Printf("send %v to small lift",p)

// 	case StandardLift:
// 		fmt.Printf("send %v to standard lift",p)

// 	case LargeLift:
// 		fmt.Printf("send %v to large lift ",p)
// 	}
// }

// func main() {
// car :=Car("Sporty")
// truck := Truck("MountainCrusher")
// motorcycle := Motorcycle("Croozer")

// sendToLift(car)
// sendToLift(truck)
// sendToLift(motorcycle)
// }


package main 

import ("fmt")

const (
	SmallLift =iota
	LargeLift
	StandardLift
)

type Lift int

type LiftPicker interface{
	PickLift() Lift
}


type Motocycle string
type Car string
type Truck string 


func (m Motocycle) String()string{
	return fmt.Sprintf("Motorcycle : %v",string(m))
}

func (c Car) String()string{
	return fmt.Sprintf("Car : %v",string(c))
}

func (t Truck)String()string{
	return fmt.Sprintf("Truck :%v",string(t))
}

func (m Motocycle)PickLift()Lift{
	return SmallLift
}

func (c Car)PickLift()Lift{
	return StandardLift
}

func (t Truck)PickLift()Lift{
	return LargeLift
}


func sendToLift(p LiftPicker){
	switch p.PickLift(){
	case SmallLift :
		fmt.Printf("send %v to small lidt",p)

	case LargeLift:
		fmt.Printf("send %v to large lift",p)

	case StandardLift:
		fmt.Printf("send %v to standard lift",p)
	}
}

func main(){
	car := Car("Sporty")
	motorcycle := Motocycle("Powerbike")
	truck :=Truck("Croozer")

	sendToLift(car)
	sendToLift(motorcycle)
	sendToLift(truck)
}