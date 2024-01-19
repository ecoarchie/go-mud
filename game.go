package main

import (
	"strings"
)

type Command func(*Player, []string) string

type Game struct {
	player             *Player
	locations          []*Location
	applicableCommands map[string]Command
}

func newGame(player *Player, locations []*Location) *Game {
	return &Game{
		player:    player,
		locations: locations,
		applicableCommands: map[string]Command{
			"осмотреться": lookAround,
			"идти":        goTo,
			"надеть":      initInventory,
			"взять":       pickItem,
			"применить":   applyItem,
			"посмотреть":  checkInventory,
		},
	}
}

func initGame() {

	//locations
	// kitchen
	kitchen := newLocation("кухня")
	kitchen.lookAroundMessage = "ты находишься на кухне, "
	kitchen.comeToMessage = "кухня, ничего интересного. "
	tea := &Item{name: "чай", isInventoryItem: false, canBePutInto: "рюкзак"}
	kitchen.staticObjs = []*StaticObject{
		{
			name:                  "стол",
			items:                 []*Item{tea},
			listItemsStartMessage: "на столе: ",
		},
	}

	//room
	room := newLocation("комната")
	room.comeToMessage = "ты в своей комнате. "
	keys := &Item{name: "ключи", isInventoryItem: false, canBeAppliedTo: "дверь", canBePutInto: "рюкзак"}
	papers := &Item{name: "конспекты", isInventoryItem: false, canBePutInto: "рюкзак"}
	sack := &Item{name: "рюкзак", isInventoryItem: true, canBePutInto: "рюкзак"}
	room.staticObjs = []*StaticObject{
		{
			name:                  "стол",
			items:                 []*Item{keys, papers},
			listItemsStartMessage: "на столе: ",
		},
		{
			name:                  "стул",
			items:                 []*Item{sack},
			listItemsStartMessage: "на стуле: ",
		},
	}

	//hall
	hall := newLocation("коридор")
	hall.comeToMessage = "ничего интересного. "

	//street
	street := newLocation("улица")
	street.comeToMessage = "на улице весна. "

	//accessible locations
	kitchen.accessableLocations = []*Location{
		hall,
	}
	room.accessableLocations = []*Location{
		hall,
	}
	hall.accessableLocations = []*Location{
		kitchen,
		room,
		street,
	}
	street.accessableLocations = []*Location{
		hall,
	}

	doorHallStreet := &Door{
		name:            "дверь",
		between:         [2]*Location{hall, street},
		isClosed:        true,
		applicableItems: []*Item{keys},
	}

	hall.doors = append(hall.doors, doorHallStreet)
	street.doors = append(street.doors, doorHallStreet)

	//Player
	P = newPlayer(kitchen)
	P.tasks = append(P.tasks, "собрать рюкзак", "идти в универ")

	//Game
	G = newGame(P, []*Location{kitchen, room, hall, street})

}

func handleCommand(command string) string {
	instructions := strings.Split(command, " ")
	mainCommand := strings.TrimSpace(instructions[0])
	options := instructions[1:]
	c, ok := G.applicableCommands[mainCommand]
	if !ok {
		return "неизвестная команда"
	}
	return c(P, options)
}
