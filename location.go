package main

import "fmt"

type Location struct {
	name                string
	staticObjs          []*StaticObject
	accessableLocations []*Location
	lookAroundMessage   string
	comeToMessage       string
	doors               []*Door
}

func newLocation(name string) *Location {
	return &Location{
		name: name,
	}
}

func (l *Location) canGoToMessage() string {
	if l.name == "улица" {
		return "можно пройти - домой"
	}
	mes := "можно пройти - "
	for i, loc := range l.accessableLocations {
		mes = mes + loc.name
		if i != len(l.accessableLocations)-1 {
			mes = mes + ", "
		}
	}
	return mes
}

func (l *Location) listItems() []*Item {
	items := []*Item{}
	for _, so := range l.staticObjs {
		items = append(items, so.items...)
	}
	return items
}

func (l *Location) findItemByName(name string) (*Item, *StaticObject, error) {
	for _, so := range l.staticObjs {
		for _, item := range so.items {
			if item.name == name {
				// foundItem := item
				return item, so, nil
			}
		}
	}
	return nil, nil, fmt.Errorf("нет такого")
}

func (l *Location) findStaticObjectByName(name string) (*StaticObject, error) {
	for _, so := range l.staticObjs {
		if so.name == name {
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

func (l *Location) isDoor(name string) (*Door, bool) {
	for _, d := range l.doors {
		if d.name == name {
			return d, true
		}
	}
	return nil, false
}

func (l *Location) isStatObj(name string) (*StaticObject, bool) {
	for _, so := range l.staticObjs {
		if so.name == name {
			return so, true
		}
	}
	return nil, false
}

func (l *Location) doorClosedFrom(loc *Location) bool {
	return false
}
