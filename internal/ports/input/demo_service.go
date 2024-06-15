package input

import "RuoYi-Go/internal/domain/model"

type DemoService interface {
	GetDemoByID(id uint) (*model.Demo, error)
	GenerateRandomCode() (string, error)
}
