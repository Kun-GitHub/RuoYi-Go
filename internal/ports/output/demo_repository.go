package output

import (
	"RuoYi-Go/internal/domain/model"
)

type DemoRepository interface {
	GetDemoByID(id uint) (*model.Demo, error)
}
