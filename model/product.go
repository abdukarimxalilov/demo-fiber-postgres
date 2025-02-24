package model

type Product struct {
    Name        string  `json:"name"`        
    Description string  `json:"description"` 
    Category    string  `json:"category"`    
    Amount      float64 `json:"amount"`      
}

type Products struct {
    Products []Product `json:"products"` 
}
