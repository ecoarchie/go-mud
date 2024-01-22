package entity

import "fmt"

type Item struct {
	Name            string
	CanBeAppliedTo  string
	CanBePutInto    string
	IsInventoryItem bool
}


func NewItem(name, canBeAppliedTo, canBePutInto string, isInventory bool) *Item {
	return &Item{
		Name: name,
		CanBeAppliedTo: canBeAppliedTo,
		CanBePutInto: canBePutInto,
		IsInventoryItem: isInventory,
	}
}

func ItemApplied(itemName string, itr Interactable) (string, error) {
	funcMap := itr.GetApplicableItems()
	_, ok := funcMap[itemName]
	if !ok {
		return "", fmt.Errorf("нельзя применить")
	}
	return funcMap[itemName](itr, itemName), nil
}
