package pkg

// Patient - структура, що представляє пацієнта
type Patient struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Age       int    `json:"age"`
	BloodType string `json:"blood_type"`
}
