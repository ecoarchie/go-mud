package app

import (
	"strings"

	"github.com/ecoarchie/go-mud/internal/entity"
)

var G = new(entity.Game)
var P = new(entity.Player)


func HandleCommand(command string) string {
	instructions := strings.Split(command, " ")
	mainCommand := strings.TrimSpace(instructions[0])
	options := instructions[1:]
	c, ok := G.ApplicableCommands[mainCommand]
	if !ok {
		return "неизвестная команда"
	}
	return c(P, options)
}

func InitGame() {
	//items
	tea := entity.NewItem("чай", "", "рюкзак", false)
	keys := entity.NewItem("ключи", "дверь", "рюкзак", false)
	papers := entity.NewItem("конспекты", "", "рюкзак", false)
	sack := entity.NewItem("рюкзак", "", "рюкзак", true)

	//static objects
	kitchenTable := entity.NewStaticObject("стол", "на столе: ")
	kitchenTable.AddItems(tea)

	roomTable := entity.NewStaticObject("стол", "на столе: ")
	roomTable.AddItems(keys, papers)
	roomChair := entity.NewStaticObject("стул", "на стуле: ")
	roomChair.AddItems(sack)

	//locations
	// kitchen
	kitchen := entity.NewEmptyLocation("кухня", "ты находишься на кухне, ", "кухня, ничего интересного. ")
	kitchen.AddStaticObjects(kitchenTable)
	//room
	room := entity.NewEmptyLocation("комната", "", "ты в своей комнате. ")
	room.AddStaticObjects(roomTable, roomChair)

	//hall
	hall := entity.NewEmptyLocation("коридор", "", "ничего интересного. ")

	//street
	street := entity.NewEmptyLocation("улица", "", "на улице весна. ")

	//accessible locations
	kitchen.AddAccessableLocations(hall)
	room.AddAccessableLocations(hall)
	hall.AddAccessableLocations(
		kitchen,
		room,
		street,
	)
	street.AddAccessableLocations(
		hall,
	)

	//doors
	doorHallStreet := entity.NewDoor("дверь", [2]*entity.Location{hall, street}, true, []*entity.Item{keys})
	hall.AddDoors([]*entity.Door{doorHallStreet})
	street.AddDoors([]*entity.Door{doorHallStreet})

	//Player
	P = entity.NewPlayer(kitchen)
	P.Tasks = append(P.Tasks, "собрать рюкзак", "идти в универ")

	//Game
	G = entity.NewGame(P, []*entity.Location{kitchen, room, hall, street})

}

