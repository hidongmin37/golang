package cargo

//package main
//
//import (
//	"context"
//	"errors"
//	"fmt"
//	"log"
//	"sync"
//	"time"
//)
//
//type contextKey string
//
//var UserIDKey contextKey = "userID"
//
//var (
//	ErrorNotImplemented = errors.New("Not Implemented yet")
//	ErrorTruckNotFound  = errors.New("Truck not found")
//)
//
//type Truck interface {
//	LoadCargo() error
//	UnloadCargo() error
//}
//
//type NormalTruck struct {
//	id    string
//	cargo int
//}
//
//type ElectricTruck struct {
//	id      string
//	cargo   int
//	battery int
//}
//
//func (e *ElectricTruck) LoadCargo() error {
//	e.cargo += 2
//	e.battery -= 1
//	return nil
//}
//
//func (e *ElectricTruck) UnloadCargo() error {
//	e.cargo = 0
//	e.battery -= 1
//	return nil
//}
//func (t *NormalTruck) LoadCargo() error {
//	t.cargo += 2
//	return nil
//}
//
//func (t *NormalTruck) UnloadCargo() error {
//	t.cargo = 0
//	return nil
//}
//
//func processTruck(ctx context.Context, truck Truck) error {
//	fmt.Printf("Processing truck: %+v \n", truck)
//
//	// access the userId
//	//userID := ctx.Value(UserIDKey)
//	ctx, cancel := context.WithTimeout(ctx, time.Second*2)
//	defer cancel()
//
//	delay := time.Second * 1
//	select {
//	case <-ctx.Done():
//		return ctx.Err()
//	case <-time.After(delay):
//		break
//	}
//
//	err := truck.LoadCargo()
//	if err != nil {
//		return fmt.Errorf("Error loading cargo: %v", err)
//	}
//
//	err = truck.UnloadCargo()
//	if err != nil {
//		return fmt.Errorf("Error unloading cargo: %w", err)
//	}
//
//	fmt.Printf("Successfully unloaded cargo: %+v \n", truck)
//	return nil
//}
//
//func processFleet(ctx context.Context, truck []Truck) error {
//	var wg sync.WaitGroup
//	errorChan := make(chan error, len(truck))
//
//	defer close(errorChan)
//
//	for _, t := range truck {
//		// 카운터를 n만큼 증가. 고루틴 시작 전에 호출해야 race condition을 방지할 수 있음
//		wg.Add(1)
//		go func(t Truck) {
//			// 카운터를 1 감소. 내부적으로 Add(-1)과 동일
//			//defer wg.Done()
//
//			if err := processTruck(ctx, t); err != nil {
//				log.Printf("Error processing truck: %v", err)
//				errorChan <- err
//			}
//			wg.Done()
//		}(t)
//
//	}
//	// 모든 고루틴이 종료될 때까지 대기
//	wg.Wait()
//
//	var errs []error
//	for len(errorChan) > 0 {
//		errs = append(errs, <-errorChan)
//	}
//	if len(errs) > 0 {
//		return fmt.Errorf("Errors occurred during fleet processing: %v", errs)
//	}
//	return nil
//}
//
//func main() {
//	ctx := context.Background()
//	//ctx = context.WithValue(ctx, UserIDKey, 42)
//
//	fleet := []Truck{
//		&NormalTruck{id: "NT-001", cargo: 0},
//		&ElectricTruck{id: "ET-001", cargo: 0, battery: 100},
//		&NormalTruck{id: "NT-002", cargo: 0},
//		&ElectricTruck{id: "ET-002", cargo: 0, battery: 100},
//	}
//
//	// process all trucks concurrently
//	if err := processFleet(ctx, fleet); err != nil {
//		log.Fatalf("Error processing fleet: %v", err)
//		return
//	}
//
//	fmt.Println("All trucks processed successfully")
//}
