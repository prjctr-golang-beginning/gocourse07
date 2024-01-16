package pkg

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// Clinic - структура, що містить мапу пацієнтів і м'ютекс для синхронізації доступу
type Clinic struct {
	patients map[string]Patient
	mu       sync.RWMutex

	dataChan chan string
	done     atomic.Bool
}

// NewClinic створює нову клініку з порожньою мапою пацієнтів
func NewClinic() *Clinic {
	return &Clinic{
		patients: make(map[string]Patient),
		dataChan: make(chan string),
	}
}

// AddPatient додає нового пацієнта у мапу
func (c *Clinic) Chan() <-chan string {
	return c.dataChan
}

// AddPatient додає нового пацієнта у мапу
func (c *Clinic) AddPatient(p Patient) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.patients[p.ID] = p
}

func (c *Clinic) AddPatientWhileCtx(ctx context.Context, p <-chan Patient) {
	t := time.NewTicker(1 * time.Second)
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Adding patient cancelled\n")
			return
		case patient := <-p:
			c.mu.Lock()
			c.patients[patient.ID] = patient
			c.mu.Unlock()
			fmt.Println("Patient added:", patient.ID)
		case <-t.C:
			fmt.Println("Just tick")
		}
	}
}

// DeletePatient видаляє пацієнта за ID
func (c *Clinic) DeletePatient(id string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.patients, id)
}

// DeserializePatients десеріалізує JSON в мапу пацієнтів
func (c *Clinic) Stop(done chan<- struct{}) {
	c.mu.Lock()
	defer c.mu.Unlock()

	for id := range c.patients {
		delete(c.patients, id)
		time.Sleep(time.Second / 10)
	}

	c.done.Store(true)
	close(c.dataChan)

	done <- struct{}{}
}

func (c *Clinic) ProcessData(patientId int, p Patient, gg <-chan struct{}, wayChan chan<- chan string) {
	defer func() { <-gg }()

	if c.done.Load() {
		fmt.Println(`Clinic processing done`)
		return
	}

	var controlChan chan string
	if patientId%10 == 0 {
		controlChan = make(chan string)
		wayChan <- controlChan
	}

	var patientData string
	select {
	case way := <-controlChan:
		patientData = fmt.Sprintf("Patient %d processed in Special way: %s", patientId, way)
	case <-time.After(time.Second):
		patientData = fmt.Sprintf("Patient %d processed", patientId)
	}

	c.dataChan <- patientData
	p.Data = patientData

	c.AddPatient(p)
}
