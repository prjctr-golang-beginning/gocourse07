package pkg

import (
	"fmt"
	"math/rand"
	"time"
)

// Patient - структура, що представляє пацієнта
type Patient struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Age       int    `json:"age"`
	BloodType string `json:"blood_type"`
	Data      string `json:"data"`
}

var (
	names      = []string{"John", "Jane", "Bob", "Alice", "Mike", "Emily", "David", "Sarah"}
	bloodTypes = []string{"A", "B", "AB", "O"}
)

func GenerateRandomPatient() Patient {
	rand.Seed(time.Now().UnixNano())
	id := fmt.Sprintf("P%03d", rand.Intn(1000))
	name := names[rand.Intn(len(names))]
	age := rand.Intn(80) + 1
	bloodType := bloodTypes[rand.Intn(len(bloodTypes))]

	return Patient{
		ID:        id,
		Name:      name,
		Age:       age,
		BloodType: bloodType,
	}
}
