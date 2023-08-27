package systems

import (
	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo/common"
)

type ButtonComponent struct {
	OnClick func()
}

func (c *ButtonComponent) GetButtonComponent() *ButtonComponent { return c }

type NotButtonComponent struct{}

func (n *NotButtonComponent) GetNotButtonComponent() *NotButtonComponent { return n }

type ButtonFace interface {
	GetButtonComponent() *ButtonComponent
}

type NotButtonFace interface {
	GetNotButtonComponent() *NotButtonComponent
}

type Buttonable interface {
	ecs.BasicFace

	common.MouseFace
	ButtonFace
}

type NotButtonable interface {
	NotButtonFace
}

type buttonEntity struct {
	*ecs.BasicEntity

	*common.MouseComponent
	*ButtonComponent
}

type ButtonSystem struct {
	entities []buttonEntity
}

func (s *ButtonSystem) Add(basic *ecs.BasicEntity, mouse *common.MouseComponent, button *ButtonComponent) {
	s.entities = append(s.entities, buttonEntity{basic, mouse, button})
}

func (s *ButtonSystem) AddByInterface(i ecs.Identifier) {
	o, _ := i.(Buttonable)
	s.Add(o.GetBasicEntity(), o.GetMouseComponent(), o.GetButtonComponent())
}

func (s *ButtonSystem) Remove(basic ecs.BasicEntity) {
	var delete = -1
	for index, entity := range s.entities {
		if entity.ID() == basic.ID() {
			delete = index
			break
		}
	}
	if delete >= 0 {
		s.entities = append(s.entities[:delete], s.entities[delete+1:]...)
	}
}

func (s *ButtonSystem) Update(dt float32) {
	for _, ent := range s.entities {
		if ent.Clicked {
			ent.OnClick()
			return
		}
	}
}
