package utils

import (
	"encoding/json"
	"log"
	"strconv"
)

// ResultTransformer structure
type ResultTransformer struct {
	// success bool`db:"success" json:"success"`
	Data interface{} `db:"data" json:"data"`
}

// NewResultTransformer constructor
func NewResultTransformer(data interface{}) *ResultTransformer {
	return &ResultTransformer{data}
}

// Set value
func (rt *ResultTransformer) Set(data interface{}) {
	rt.Data = data
	// rt.success = true
}

// Get value
func (rt *ResultTransformer) Get() interface{} {
	return rt.Data
}

// ToJSON return json
func (rt *ResultTransformer) ToJSON() (string, error) {

	json, err := json.Marshal(rt.Data)
	if err != nil {
		return "", err
	}

	return string(json), nil
}
func ConvertInt(str string)int{
	kq,err:=strconv.Atoi(str)
	if err!=nil{
		log.Fatal()
	}
	return kq
}