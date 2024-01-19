package main

import (
	"fmt"
	"strings"
)

type Inventory map[string][]*Item

func (i Inventory) findItemByName(name string) (*Item, error) {
	for _, items := range i {
		for _, item := range items {
			if item.name == name {
				return item, nil
			}
		}
	}
	return nil, fmt.Errorf("нет предмета в инвентаре - %s", name)
}

func (i Inventory) getItems(name string) string {
	items := i[name]
	var mes []string
	for _, item := range items {
		mes = append(mes, item.name)
	}
	if len(mes) == 0 {
		return fmt.Sprintf("%s пуст", name)
	}
	return strings.Join(mes, " ")
}
