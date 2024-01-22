package entity


type Command func(*Player, []string) string

type Game struct {
	Player             *Player
	Locations          []*Location
	ApplicableCommands map[string]Command
}

func NewGame(player *Player, locations []*Location) *Game {
	return &Game{
		Player:    player,
		Locations: locations,
		ApplicableCommands: map[string]Command{
			"осмотреться": LookAround,
			"идти":        GoTo,
			"надеть":      InitInventory,
			"взять":       PickItem,
			"применить":   ApplyItem,
			"посмотреть":  CheckInventory,
		},
	}
}