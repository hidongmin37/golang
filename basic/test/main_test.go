package test

import (
	"golang"
	"log"
	"testing"
)

func TestMain(t *testing.T) {
	t.Run("processTruck", func(t *testing.T) {
		t.Run("should load and unload a truck cargo", func(t *testing.T) {
			nt := &main.NormalTruck{id: "1"}
			et := &main.ElectricTruck{id: "2"}

			person := make(map[string]interface{}, 0)

			person["name"] = "Tiago"
			person["age"] = 18

			age, exists := person["age"].(int)
			if !exists {
				log.Fatal("Age is not an integer")
				return
			}
			log.Println("Age is", age)

			err := main.processTruck(nt)
			if err != nil {
				log.Fatalf("error processing truck: %v", err)
			}
			err = main.processTruck(et)
			if err != nil {
				log.Fatalf("error processing truck: %v", err)
			}

			// asserting
			if nt.cargo != 0 {
				t.Fatal("Normal truck cargo should be 0")
			}
			if et.battery != -2 {
				t.Fatal("Electric truck battery should be  -2")
			}
		})
	})
}
