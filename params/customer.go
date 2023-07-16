package params

type Customer struct {
	CustomerID          int          `json:"id,omitempty"`
	CustomerName        string       `json:"nama"`
	CustomerDoB         string       `json:"tanggal_lahir"`
	CustomerPhoneNumber string       `json:"telepon"`
	CustomerEmail       string       `json:"email"`
	Nationality         string       `json:"kewarganegaraan"`
	FamilyList          []FamilyList `json:"keluarga"`
}
