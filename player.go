package main

import (
	"fmt"
	"strings"
)

type Player struct {
	inventory       *Inventory
	currentLocation *Location
	tasks           []string
}

func newPlayer(currentLocation *Location) *Player {
	return &Player{
		inventory:       nil,
		currentLocation: currentLocation,
	}
}

func (p *Player) tasksMessage() string {
	if len(p.tasks) == 0 {
		return ""
	}
	mes := "надо "
	for i, t := range p.tasks {
		mes = mes + t
		if i != len(p.tasks)-1 {
			mes = mes + " и "
		}
	}
	mes = mes + ". "
	return mes
}

var lookAround = func(p *Player, options []string) string {
	fmt.Printf("двери - %+v\n", p.currentLocation.doors)
	var message string
	var staticObjsMessage string
	tasksMessage := ""
	for _, so := range p.currentLocation.staticObjs {
		itemsMes := so.getItems()
		if itemsMes == "" {
			continue
		}
		staticObjsMessage = staticObjsMessage + itemsMes
	}
	if staticObjsMessage == "" {
		staticObjsMessage = "пустая комната. "
	}
	if p.currentLocation.name == "кухня" {
		tasksMessage = p.tasksMessage()
	}
	message = p.currentLocation.lookAroundMessage +
		staticObjsMessage +
		tasksMessage +
		p.currentLocation.canGoToMessage()
	return message
}

var goTo = func(p *Player, options []string) string {
	var message string
	location := strings.TrimSpace(options[0])
	for _, l := range p.currentLocation.accessableLocations {
		if l.name == location {
			door, err := doorBetweenLocations(l, p.currentLocation)
			if err == nil {
				if isClosed := door.isClosed; isClosed {
					return "дверь закрыта"
				}
			}
			p.currentLocation = l
			message = p.currentLocation.comeToMessage + l.canGoToMessage()
			return message
		}
	}
	message = "нет пути в " + location
	return message
}

var initInventory = func(p *Player, options []string) string {
	inventory := strings.TrimSpace(options[0])
	// items := p.currentLocation.listItems()
	sack, err := p.currentLocation.findItemByName(inventory)
	if err != nil {
		return err.Error()
	}
	if sack.isInventoryItem {
		p.inventory.items[inventory] = []*Item{}
		return "вы надели: " + inventory
	}
	return "нет инвентарного предмета"
}

var pickItem = func(p *Player, options []string) string {
	itemName := strings.TrimSpace(options[0])
	if p.inventory == nil {
		return fmt.Sprintf("нет предмета в инвентаре - %s", itemName)
	}
	item, err := p.currentLocation.findItemByName(itemName)
	if err != nil {
		return err.Error()
	}
	_, ok := p.inventory.items[item.canBePutInto]
	if !ok {
		return "некуда класть"
	}
	p.inventory.items[item.canBePutInto] = append(p.inventory.items[item.canBePutInto], item) 

	return "предмет добавлен в инвентарь: " + itemName
}

var applyItem = func(p *Player, options []string) string {
	itemName := strings.TrimSpace(options[0])
	applyToName := strings.TrimSpace(options[1])
	item, err := p.inventory.findItemByName(itemName)
	if err != nil {
		return err.Error()
	}
	door, isDoor := p.currentLocation.isDoor(applyToName)
	if isDoor {
		return door.applyItemToDoor(item)
	}

	statObj, isStatObj := p.currentLocation.isStatObj(applyToName)
	if isStatObj {
		return statObj.applyItemToStatObject(item)
	}
	// interactable, err := p.currentLocation.findInteractableByName(applyToName)
	// if err != nil {
	// 	return err.Error()
	// }
	// if item.canBeAppliedTo != interactable.getName() {
	// 	return ""
	// }
	// mes, err :=  itemApplied(item.name, interactable)
	// if err != nil {
	// 	return err.Error()
	// }
	// return mes
	return "не к чему применить"
}
