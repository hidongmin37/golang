package main

import (
	"errors"
	"fmt"
	"log"
)

var (
	ErrorNotImplemented = errors.New("Not Implemented yet")
	ErrorTruckNotFound  = errors.New("Truck not found")
)

type Truck interface {
	LoadCargo() error
	UnloadCargo() error
}

type NormalTruck struct {
	id    string
	cargo int
}

type ElectricTruck struct {
	id      string
	cargo   int
	battery int
}

func (e *ElectricTruck) LoadCargo() error {
	e.cargo += 2
	e.battery -= 1
	return nil
}

func (e *ElectricTruck) UnloadCargo() error {
	e.cargo = 0
	e.battery -= 1
	return nil
}
func (t *NormalTruck) LoadCargo() error {
	t.cargo += 2
	return nil
}

func (t *NormalTruck) UnloadCargo() error {
	t.cargo = 0
	return nil
}

func processTruck(truck Truck) error {
	fmt.Printf("Processing truck: %+v \n", truck)
	err := truck.LoadCargo()
	if err != nil {
		return fmt.Errorf("Error loading cargo: %v", err)
	}

	err = truck.UnloadCargo()
	if err != nil {
		return fmt.Errorf("Error unloading cargo: %w", err)
	}

	return nil
}

func main() {
	nt := &NormalTruck{id: "1"}
	et := &ElectricTruck{id: "2"}

	person := make(map[string]interface{}, 0)

	person["name"] = "Tiago"
	person["age"] = 18

	age, exists := person["width"].(int)
	if !exists {
		log.Fatal("Age is not an integer")
		return
	}
	log.Println("Age is", age)

	err := processTruck(nt)
	if err != nil {
		log.Fatalf("error processing truck: %v", err)
	}
	err = processTruck(et)
	if err != nil {
		log.Fatalf("error processing truck: %v", err)
	}
	log.Println(nt.cargo)
	log.Println(et.battery)
}
