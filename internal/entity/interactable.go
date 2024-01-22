package entity

import "fmt"

type Interactable interface {
	GetName() string
	GetApplicableItems() map[string]func(Interactable, string) string
}

type Door struct {
	Name            string
	Between         [2]*Location
	IsClosed        bool
	ApplicableItems []*Item
}

func NewDoor(name string, between [2]*Location, isClosed bool, applicableItems []*Item) *Door {
	return &Door{
		Name: name,
		Between: between,
		IsClosed: isClosed,
		ApplicableItems: applicableItems,
	}
}

func (d *Door) AddApplicableItems(items []*Item) {
	d.ApplicableItems = append(d.ApplicableItems, items...)
}

func (d *Door) IsApplicable(item *Item) bool {
	for _, itm := range d.ApplicableItems {
		if itm.Name == item.Name {
			return true
		}
	}
	return false
}

func DoorBetweenLocations(locA, locB *Location) (*Door, error) {
	for _, d := range locA.Doors {
		if (d.Between[0] == locA && d.Between[1] == locB) || (d.Between[0] == locB && d.Between[1] == locA) {
			return d, nil
		}
	}
	return nil, fmt.Errorf("Door not found")
}

type StaticObject struct {
	Name                  string
	ListItemsStartMessage string
	Items                 []*Item
	ApplicableItems       []*Item
}

func NewStaticObject(name, listItemStartMessage string) *StaticObject {
	return &StaticObject{
		Name: name,
		ListItemsStartMessage: listItemStartMessage,
	}
}

func (so *StaticObject) AddItems(items ...*Item) {
	so.Items = append(so.Items, items...)
}

func (so *StaticObject) AddApplicableItems(items []*Item) {
	so.ApplicableItems = append(so.ApplicableItems, items...)
}

func (so *StaticObject) ApplyItemToStatObject(item *Item) string {
	return ""
}

func (so *StaticObject) DeleteItem(item *Item) {
	for idx, itm := range so.Items {
		if itm.Name == item.Name {
			so.Items[idx] = so.Items[len(so.Items)-1]
			so.Items[len(so.Items)-1] = nil
			so.Items = so.Items[:len(so.Items)-1]
			break
		}
	}
}

func (so *StaticObject) IsApplicable(item *Item) bool {
	for _, itm := range so.ApplicableItems {
		if itm.Name == item.Name {
			return true
		}
	}
	return false
}

func (d *Door) ApplyItemToDoor(item *Item) string {
	if !d.IsApplicable(item) {
		return "нельзя применить"
	}
	switch item.Name {
	case "ключи":
		d.IsClosed = !d.IsClosed
		if d.IsClosed {
			return "дверь закрыта"
		}
		return "дверь открыта"
	default:
		return "нельзя применить"
	}
}

func (i *StaticObject) GetItems() string {
	if len(i.Items) == 0 {
		return ""
	}
	itemsString := i.ListItemsStartMessage
	for idx, item := range i.Items {
		itemsString = itemsString + item.Name
		if idx != len(i.Items)-1 {
			itemsString = itemsString + ", "
		}
	}
	// itemsString = itemsString + ". "
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
