package entity

import "fmt"

type Location struct {
	Name                string
	LookAroundMessage   string
	ComeToMessage       string
	StaticObjs          []*StaticObject
	AccessableLocations []*Location
	Doors               []*Door
}

func NewEmptyLocation(name, lookAroundMessage, ComeToMessage string) *Location {
	return &Location{
		Name: name,
		LookAroundMessage: lookAroundMessage,
		ComeToMessage: ComeToMessage,
	}
}

func (l *Location) AddStaticObjects(so ...*StaticObject) {
	l.StaticObjs = append(l.StaticObjs, so...)
}

func (l *Location) AddAccessableLocations(al ...*Location) {
	l.AccessableLocations = append(l.AccessableLocations, al...)
}

func (l *Location) AddDoors(d []*Door) {
	l.Doors = append(l.Doors, d...)
}

func (l *Location) CanGoToMessage() string {
	if l.Name == "улица" {
		return "можно пройти - домой"
	}
	mes := "можно пройти - "
	for i, loc := range l.AccessableLocations {
		mes = mes + loc.Name
		if i != len(l.AccessableLocations)-1 {
			mes = mes + ", "
		}
	}
	return mes
}

func (l *Location) ListItems() []*Item {
	items := []*Item{}
	for _, so := range l.StaticObjs {
		items = append(items, so.Items...)
	}
	return items
}

func (l *Location) FindItemByName(name string) (*Item, *StaticObject, error) {
	for _, so := range l.StaticObjs {
		for _, item := range so.Items {
			if item.Name == name {
				// foundItem := item
				return item, so, nil
			}
		}
	}
	return nil, nil, fmt.Errorf("нет такого")
}

func (l *Location) FindStaticObjectByName(name string) (*StaticObject, error) {
	for _, so := range l.StaticObjs {
		if so.Name == name {
			return so, nil
		}
	}
	return nil, fmt.Errorf("не к чему применить")
}

// func (l *Location) findInteractableByName(name string) (Interactable, error) {
// 	for _, so := range l.staticObjs {
// 		if so.name == name {
// 			return so, nil
// 		}
// 	}
// 	for _, d := range l.doors {
// 		if d.name == name {
// 			return d, nil
// 		}
// 	}
// 	return nil, fmt.Errorf("не к чему применить")
// }

func (l *Location) IsDoor(name string) (*Door, bool) {
	for _, d := range l.Doors {
		if d.Name == name {
			return d, true
		}
	}
	return nil, false
}

func (l *Location) IsStatObj(name string) (*StaticObject, bool) {
	for _, so := range l.StaticObjs {
		if so.Name == name {
			return so, true
		}
	}
	return nil, false
}

func (l *Location) DoorClosedFrom(loc *Location) bool {
	return false
}
