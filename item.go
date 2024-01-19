package main

import "fmt"

type Item struct {
	name            string
	isInventoryItem bool
	canBeAppliedTo  string
	canBePutInto    string
}

func itemApplied(itemName string, itr Interactable) (string, error) {
	funcMap := itr.getApplicableItems()
	_, ok := funcMap[itemName]
	if !ok {
		return "", fmt.Errorf("нельзя применить")
	}
	return funcMap[itemName](itr, itemName), nil
}
