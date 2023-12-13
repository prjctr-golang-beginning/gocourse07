package main

import (
	"concurrency2/pkg"
	"context"
	"fmt"
	"strconv"
	"time"
)

func main() {
	clinic := pkg.NewClinic()

	// Створення контексту з відміною
	cancelCtx, cancelFunc := context.WithCancel(context.Background())

	// Додавання пацієнта з можливістю відміни
	pChan := make(chan pkg.Patient)
	go clinic.AddPatientWhileCtx(cancelCtx, pChan)

	id := 1
	go func() {
		for { //горутина зависає до кінця виконання програми
			pChan <- pkg.Patient{strconv.Itoa(id), "John Doe", 30, "A+"}
			time.Sleep(1 * time.Second)
			id++
		}
	}()

	// Скасування додавання
	time.Sleep(5 * time.Second)
	cancelFunc()

	// Створення контексту з таймаутом
	timeoutCtx, _ := context.WithTimeout(context.Background(), 3*time.Second)

	// Обробка пацієнта з таймаутом
	go processPatient(timeoutCtx, "2")

	// Передача значень через контекст
	valueCtx := context.WithValue(context.Background(), "key", "importantValue")
	go func(ctx context.Context) {
		if val := ctx.Value("key"); val != nil {
			fmt.Println("Received value from context:", val)
		}
	}(valueCtx)

	// Даємо час на завершення горутин
	time.Sleep(5 * time.Second)
}

func processPatient(ctx context.Context, patientID string) {
	// Використання властивості 'Deadline' з контексту
	if deadline, ok := ctx.Deadline(); ok {
		fmt.Printf("Processing patient %s must be completed by %v\n", patientID, deadline)
	}

	// Імітуємо тривалу обробку
	time.Sleep(2 * time.Second)

	// Перевірка на відміну
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Processing cancelled for patient:", patientID)
			return
		default:
			fmt.Println("Processed patient:", patientID)
			time.Sleep(1 * time.Second)
		}
	}
}
