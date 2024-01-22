package entity

import (
	"fmt"
	"strings"
)

type Player struct {
	Inventory       Inventory
	CurrentLocation *Location
	Tasks           []string
}

func NewPlayer(currentLocation *Location) *Player {
	inventories := make(map[string][]*Item)
	return &Player{
		Inventory:       inventories,
		CurrentLocation: currentLocation,
	}
}

func (p *Player) TasksMessage() string {
	if len(p.Tasks) == 0 {
		return ""
	}
	mes := "надо "
	for i, t := range p.Tasks {
		mes = mes + t
		if i != len(p.Tasks)-1 {
			mes = mes + " и "
		}
	}
	// mes = mes + ". "
	return mes
}

func LookAround(p *Player, options []string) string {
	staticObjsMessage := ""
	itemsMes := []string{}
	for _, so := range p.CurrentLocation.StaticObjs {
		if len(so.Items) == 0 {
			continue
		}
		itemsMes = append(itemsMes, so.GetItems())
	}
	staticObjsMessage = staticObjsMessage + strings.Join(itemsMes, ", ")

	if staticObjsMessage == "" {
		staticObjsMessage = "пустая комната"
	}
	if p.CurrentLocation.Name == "кухня" {
		staticObjsMessage = staticObjsMessage + ", " + p.TasksMessage() 
	}
	return p.CurrentLocation.LookAroundMessage +
		staticObjsMessage +
		". " +
		p.CurrentLocation.CanGoToMessage()
}

func GoTo(p *Player, options []string) string {
	var message string
	location := strings.TrimSpace(options[0])
	for _, l := range p.CurrentLocation.AccessableLocations {
		if l.Name == location {
			door, err := DoorBetweenLocations(l, p.CurrentLocation)
			if err == nil {
				if isClosed := door.IsClosed; isClosed {
					return "дверь закрыта"
				}
			}
			p.CurrentLocation = l
			message = p.CurrentLocation.ComeToMessage + l.CanGoToMessage()
			return message
		}
	}
	message = "нет пути в " + location
	return message
}

func InitInventory(p *Player, options []string) string {
	inventoryName := strings.TrimSpace(options[0])
	// items := p.currentLocation.listItems()
	sack, so, err := p.CurrentLocation.FindItemByName(inventoryName)
	if err != nil {
		return err.Error()
	}
	if sack.IsInventoryItem {
		so.DeleteItem(sack)
		p.Inventory[inventoryName] = []*Item{}
		return "вы надели: " + inventoryName
	}
	return "нет инвентарного предмета"
}

func PickItem(p *Player, options []string) string {
	itemName := strings.TrimSpace(options[0])
	if p.Inventory == nil {
		return fmt.Sprintf("нет предмета в инвентаре - %s", itemName)
	}
	item, staticObj, err := p.CurrentLocation.FindItemByName(itemName)
	if err != nil {
		return err.Error()
	}
	_, ok := p.Inventory[item.CanBePutInto]
	if !ok {
		return "некуда класть"
	}
	staticObj.DeleteItem(item)
	p.Inventory[item.CanBePutInto] = append(p.Inventory[item.CanBePutInto], item)
	if len(p.Inventory["рюкзак"]) == 2 {
		task, idx, err := p.GetTaskByName("собрать рюкзак")
		if err == nil {
			p.DeleteTask(task, idx)
		}
	}

	return "предмет добавлен в инвентарь: " + itemName
}

func (p *Player) DeleteTask(name string, idx int) {
	p.Tasks[idx] = p.Tasks[len(p.Tasks)-1]
	p.Tasks = p.Tasks[:len(p.Tasks)-1]
}

func (p *Player) GetTaskByName(name string) (string, int, error) {
	for idx, t := range p.Tasks {
		if t == name {
			return t, idx, nil
		}
	}
	return "", 0, fmt.Errorf("task not found")
}

func ApplyItem(p *Player, options []string) string {
	itemName := strings.TrimSpace(options[0])
	applyToName := strings.TrimSpace(options[1])
	item, err := p.Inventory.FindItemByName(itemName)
	if err != nil {
		return err.Error()
	}
	door, isDoor := p.CurrentLocation.IsDoor(applyToName)
	if isDoor {
		return door.ApplyItemToDoor(item)
	}

	statObj, isStatObj := p.CurrentLocation.IsStatObj(applyToName)
	if isStatObj {
		return statObj.ApplyItemToStatObject(item)
	}
	return "не к чему применить"
}

func CheckInventory(p *Player, options []string) string {
	inventoryName := strings.TrimSpace(options[0])
	items := p.Inventory.GetItems(inventoryName)
	return items
}
