package main

import "fmt"

type Interactable interface {
	getName() string
	getApplicableItems() map[string]func(Interactable, string) string
}

type Door struct {
	name            string
	between         [2]*Location
	isClosed        bool
	applicableItems []*Item
}

func (d *Door) isApplicable(item *Item) bool {
	for _, itm := range d.applicableItems {
		if itm.name == item.name {
			return true
		}
	}
	return false
}

func doorBetweenLocations(locA, locB *Location) (*Door, error) {
	for _, d := range locA.doors {
		if (d.between[0] == locA && d.between[1] == locB) || (d.between[0] == locB && d.between[1] == locA) {
			return d, nil
		}
	}
	return nil, fmt.Errorf("Door not found")
}

type StaticObject struct {
	name                  string
	items                 []*Item
	listItemsStartMessage string
	applicableItems       []*Item
}

func (so *StaticObject) applyItemToStatObject(item *Item) string {
	return ""
}

func (so *StaticObject) isApplicable(item *Item) bool {
	for _, itm := range so.applicableItems {
		if itm.name == item.name {
			return true
		}
	}
	return false
}

func (d *Door) applyItemToDoor(item *Item) string {
	if !d.isApplicable(item) {
		return "нельзя применить"
	}
	switch item.name {
	case "ключи":
		d.isClosed = !d.isClosed
		if d.isClosed {
			return "Дверь закрыта"
		}
		return "Дверь открыта"
	default:
		return "нельзя применить"
	}
}

func (i *StaticObject) getItems() string {
	if len(i.items) == 0 {
		return ""
	}
	itemsString := i.listItemsStartMessage
	for idx, item := range i.items {
		itemsString = itemsString + item.name
		if idx != len(i.items)-1 {
			itemsString = itemsString + ", "
		}
	}
	itemsString = itemsString + ". "
	return itemsString
}

// func (s *StaticObject) getApplicableItems() map[string]func(*StaticObject, string) string {
// 	return s.applicableItems
// }

// func (d *Door) getApplicableItems() map[string]func(*Door, string) string {
// 	return d.applicableItems
// }

// func (s *StaticObject) getName() string {
// 	return s.name
// }

// func (d *Door) getName() string {
// 	return d.name
// }

// func (s *StaticObject) isApplicable(itemName string) bool {
// 	for _, item := range s.getApplicableItems() {
// 		if item.name == itemName {
// 			return true
// 		}
// 	}
// 	return false
// }

// func (d *Door) isApplicable(itemName string) bool {
// 	for _, item := range d.getApplicableItems() {
// 		if item.name == itemName {
// 			return true
// 		}
// 	}
// 	return false
// }
