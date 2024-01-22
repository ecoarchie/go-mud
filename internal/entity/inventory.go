package entity

import (
	"fmt"
	"strings"
)

type Inventory map[string][]*Item

func (i Inventory) FindItemByName(name string) (*Item, error) {
	for _, items := range i {
		for _, item := range items {
			if item.Name == name {
				return item, nil
			}
		}
	}
	return nil, fmt.Errorf("нет предмета в инвентаре - %s", name)
}

func (i Inventory) GetItems(name string) string {
	items := i[name]
	var mes []string
	for _, item := range items {
		mes = append(mes, item.Name)
	}
	if len(mes) == 0 {
		return fmt.Sprintf("%s пуст", name)
	}
	return strings.Join(mes, " ")
}
