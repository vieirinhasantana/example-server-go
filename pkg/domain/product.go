package domain

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type Product struct {
	Base
	Name     string `json:"Name"`
	ClientId string `json:"clientId"`
}

func NewProduct() *Product {
	return &Product{}
}

func (p *Product) process() {
	p.ID = uuid.New().String()
	p.CreatedAt = time.Now()
	p.UpdateAt = time.Now()
}

func (p *Product) Create(name string, clientId string) (string, error) {
	p.Name = name
	p.ClientId = clientId

	p.process()

	b, err := json.Marshal(p)
	if err != nil {
		panic("Error marshal object")
	}

	return string(b), nil
}
