package main

import (
	"fmt"
	"sync"
	"time"
)

func dance(signalChan chan bool, doneChan chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-signalChan:
			fmt.Println("Людина почала танцювати")
			time.Sleep(time.Second)
		case <-doneChan:
			fmt.Println("Людина зупинила танець")
			return
		}
	}
}

func main() {
	signalChan := makeSignalChan()
	doneChan := makeDoneChan()
	var wg sync.WaitGroup

	// Запуск воркера, який буде обробляти сигнали про початок танцю
	wg.Add(1)
	go dance(signalChan, doneChan, &wg)

	// Моделюємо ситуацію: початок танцю
	signalChan <- true
	time.Sleep(3 * time.Second) // Затримка в танцю

	// Моделюємо ситуацію: зупинка танцю
	close(doneChan)
	doneChan = makeDoneChan()
	wg.Wait()

	// Після зупинки танцю можна знову запустити воркера
	wg.Add(1)
	go dance(signalChan, doneChan, &wg)

	// Моделюємо ситуацію: знову початок танцю
	signalChan <- true
	time.Sleep(2 * time.Second) // Затримка в танцю

	// Завершення роботи
	close(doneChan)
	wg.Wait()
	close(signalChan)
}

func makeDoneChan() chan struct{} {
	return make(chan struct{})
}

func makeSignalChan() chan bool {
	return make(chan bool)
}
