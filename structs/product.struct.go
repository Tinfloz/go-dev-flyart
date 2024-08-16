package structs

type Product struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Dimension   string `json:"dimension"`
	Medium      string `json:"medium"`
	Price       string `json:"price"`
	Base64Str   string `json:"imageStr"`
}
