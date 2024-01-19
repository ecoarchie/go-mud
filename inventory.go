package main

import "fmt"

type Inventory struct {
	items map[string][]*Item
}

func (i *Inventory) findItemByName(name string) (*Item, error) {
	for _, inv := range i.items {
		for _, itm := range inv {
			if itm.name == name {
				return itm, nil
			}
		}
	}
	return nil, fmt.Errorf("нет предмета в инвентаре - %s", name)
}
