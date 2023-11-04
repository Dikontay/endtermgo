package internal

import "fmt"

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

type IAccessory interface {
	AnimalStaregy
	Print()
}

type Hat struct {
	animal AnimalStaregy
}

func (h *Hat) ChangeCleanness(b bool) {
	h.animal.ChangeCleanness(b)
}

func (h *Hat) ChangeEat(b bool) {
	h.animal.ChangeEat(b)
}

func (h *Hat) ChangeMood(b bool) {
	h.animal.ChangeMood(b)
}

func (h *Hat) GetTimeToBeDirty() int {
	return h.animal.GetTimeToBeDirty()
}

func (h *Hat) GetTimeToEat() int {
	return h.animal.GetTimeToEat()
}

func (h *Hat) GetTimeToPlay() int {
	return h.animal.GetTimeToPlay()
}

func (h *Hat) GetMood() int {
	return h.animal.GetMood()
}

func (h *Hat) GetFood() int {
	return h.animal.GetFood()
}

func (h *Hat) GetCleanness() int {
	return h.animal.GetCleanness()
}

func (h *Hat) notifyKiller() {
	h.animal.notifyKiller()
}

func (h *Hat) Print() {
	for i := 0; i < len(accessoryPictures["hat"]); i++ {
		fmt.Println(accessoryPictures["hat"][i])
	}
	h.animal.Print()
}
