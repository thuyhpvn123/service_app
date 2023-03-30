package models

// Header struct
type Header struct {
	Success bool      `json:"success"`
	// Count  int         `json:"count"`
	Data   interface{} `json:"data"`
}


