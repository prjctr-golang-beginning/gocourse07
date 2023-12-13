package pkg

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"time"
)

// Clinic - структура, що містить мапу пацієнтів і м'ютекс для синхронізації доступу
type Clinic struct {
	patients map[string]Patient
	mu       sync.RWMutex
}

// NewClinic створює нову клініку з порожньою мапою пацієнтів
func NewClinic() *Clinic {
	return &Clinic{
		patients: make(map[string]Patient),
	}
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

// GetPatient повертає пацієнта за ID
func (c *Clinic) GetPatient(id string) (Patient, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	p, exists := c.patients[id]
	return p, exists
}

// UpdatePatient оновлює дані пацієнта
func (c *Clinic) UpdatePatient(id string, newPatient Patient) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if _, exists := c.patients[id]; exists {
		c.patients[id] = newPatient
	}
}

// DeletePatient видаляє пацієнта за ID
func (c *Clinic) DeletePatient(id string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.patients, id)
}

// FindPatientsByBloodType знаходить пацієнтів за групою крові
func (c *Clinic) FindPatientsByBloodType(bloodType string) []Patient {
	c.mu.RLock()
	defer c.mu.RUnlock()
	var found []Patient
	for _, p := range c.patients {
		if p.BloodType == bloodType {
			found = append(found, p)
		}
	}
	return found
}

// SerializePatients серіалізує мапу пацієнтів в JSON
func (c *Clinic) SerializePatients() (string, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	data, err := json.Marshal(c.patients)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// DeserializePatients десеріалізує JSON в мапу пацієнтів
func (c *Clinic) DeserializePatients(data string) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	var patients map[string]Patient
	if err := json.Unmarshal([]byte(data), &patients); err != nil {
		return err
	}
	c.patients = patients
	return nil
}
