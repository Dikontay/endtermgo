// internal/accessories.go

package internal

var accessoryPictures = map[string][]string{
	"hat": {
		// Hat picture
		"    .~~~~`\\~~\\",
		"     ;       ~~ \\",
		"     |           ;",
		" ,--------,______|---.",
		"/          \\-----`    \\",
		"`.__________`-_______-'",
	},
	"cap": {
		// Cap picture
		"  _____",
		" .-     -.",
		"/           \\",
		"|-.           |",
		"|  \\          |",
		"[__|__________|_______",
	},
	"official": {
		// official picture
		"    __________",
		"    |        |",
		"    |________|",
		"    |        |",
		"    |--------|",
		"    |        |",
		"----|--------|----",
		"|                |",
		"------------------",
	},
}

type Accessory interface {
	Apply(AnimalStaregy)
}

type HatAccessory struct{}
type CapAccessory struct{}
type OfficialAccessory struct{}

func (h *HatAccessory) Apply(animal AnimalStaregy) {
	animal.PrintAccessory("hat")
}

func (c *CapAccessory) Apply(animal AnimalStaregy) {
	animal.PrintAccessory("cap")
}

func (g *OfficialAccessory) Apply(animal AnimalStaregy) {
	animal.PrintAccessory("official")
}
