package main

import (
	"fmt"
	"strings"
)

type Player struct {
	inventory       Inventory
	currentLocation *Location
	tasks           []string
}

func newPlayer(currentLocation *Location) *Player {
	inventories := make(map[string][]*Item)
	return &Player{
		inventory:       inventories,
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
	// mes = mes + ". "
	return mes
}

var lookAround = func(p *Player, options []string) string {
	var message string
	var staticObjsMessage string
	tasksMessage := ""
	itemsMes := []string{}
	for _, so := range p.currentLocation.staticObjs {
		if len(so.items) == 0 {
			continue
		}
		itemsMes = append(itemsMes, so.getItems())
	}
	staticObjsMessage = staticObjsMessage + strings.Join(itemsMes, ", ")

	if staticObjsMessage == "" {
		staticObjsMessage = "пустая комната"
	}
	if p.currentLocation.name == "кухня" {
		tasksMessage = ", " + p.tasksMessage()
	}
	message = p.currentLocation.lookAroundMessage +
		staticObjsMessage +
		tasksMessage + ". " +
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
	inventoryName := strings.TrimSpace(options[0])
	// items := p.currentLocation.listItems()
	sack, so, err := p.currentLocation.findItemByName(inventoryName)
	if err != nil {
		return err.Error()
	}
	if sack.isInventoryItem {
		so.deleteItem(sack)
		p.inventory[inventoryName] = []*Item{}
		return "вы надели: " + inventoryName
	}
	return "нет инвентарного предмета"
}

var pickItem = func(p *Player, options []string) string {
	itemName := strings.TrimSpace(options[0])
	if p.inventory == nil {
		return fmt.Sprintf("нет предмета в инвентаре - %s", itemName)
	}
	item, staticObj, err := p.currentLocation.findItemByName(itemName)
	if err != nil {
		return err.Error()
	}
	_, ok := p.inventory[item.canBePutInto]
	if !ok {
		return "некуда класть"
	}
	staticObj.deleteItem(item)
	p.inventory[item.canBePutInto] = append(p.inventory[item.canBePutInto], item)
	if len(p.inventory["рюкзак"]) == 2 {
		task, idx, err := p.getTaskByName("собрать рюкзак")
		if err == nil {
			p.deleteTask(task, idx)
		}
	}

	return "предмет добавлен в инвентарь: " + itemName
}

func (p *Player) deleteTask(name string, idx int) {
	p.tasks[idx] = p.tasks[len(p.tasks)-1]
	p.tasks = p.tasks[:len(p.tasks)-1]
}

func (p *Player) getTaskByName(name string) (string, int, error) {
	for idx, t := range p.tasks {
		if t == name {
			return t, idx, nil
		}
	}
	return "", 0, fmt.Errorf("task not found")
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

func checkInventory(p *Player, options []string) string {
	inventoryName := strings.TrimSpace(options[0])
	items := p.inventory.getItems(inventoryName)
	return items
}
