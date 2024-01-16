package main

import (
	"concurrency2/pkg"
	"fmt"
	"time"
)

const totalPatients = 100
const maxGoroutines = 3

func main() {
	patients := make([]pkg.Patient, totalPatients)
	clinic := pkg.NewClinic()

	for i := 0; i < totalPatients; i++ {
		patients[i] = pkg.GenerateRandomPatient()
	}

	// Горутина для системи сповіщення лікарів
	go func() {
		for p := range clinic.Chan() {
			fmt.Printf("Doctor was notified about patient: %s\n", p)
		}
	}()

	wayChan := make(chan chan string)
	go func() {
		for where := range wayChan {
			where <- pkg.GenerateRandomString(10)
		}
	}()

	// Старт горутин для обробки даних пацієнтів
	closeClinic := make(chan struct{})
	gGuard := make(chan struct{}, maxGoroutines)
	go func() {
		for i, patient := range patients {
			select {
			case <-closeClinic:
				fmt.Println("Clinic closed")
				return
			default:
				gGuard <- struct{}{} // would block if gGuard channel is already filled
				go clinic.ProcessData(i, patient, gGuard, wayChan)
			}

		}
	}()

	time.Sleep(5 * time.Second)
	close(closeClinic)
	done := make(chan struct{}, 1)
	go clinic.Stop(done)
	<-done

	fmt.Println("All patients processed for today")
}
