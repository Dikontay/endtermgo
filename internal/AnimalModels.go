package internal

//STRATEGY

type AnimalContext struct {
	Behavior AnimalStaregy
}

type AnimalStaregy interface {
	ChangeCleanness(bool)
	ChangeEat(bool)
	ChangeMood(bool)
	Print()
}

type Dog struct {
	Clean int
	Mood  int
	Food  int

	TimeToBeDirty int
	TimeToEat     int
	TimeToPlay    int
}

type Bear struct {
	Clean int
	Mood  int
	Food  int

	TimeToBeDirty int
	TimeToEat     int
	TimeToPlay    int
}

type Owl struct {
	Clean int
	Mood  int
	Food  int

	TimeToBeDirty int
	TimeToEat     int
	TimeToPlay    int
}

type Cat struct {
	Clean int
	Mood  int
	Food  int

	TimeToBeDirty int
	TimeToEat     int
	TimeToPlay    int
}

func (d *Bear) ChangeEat(b bool) {
	if b {
		d.Food++
	} else {
		d.Food--
	}
}

func (d *Bear) ChangeCleanness(b bool) {
	if b {
		d.Clean++
	} else {
		d.Clean--
	}
}

func (d *Bear) ChangeMood(b bool) {
	if b {
		d.Mood++
	} else {
		d.Mood--
	}
}

func (d *Bear) Print() {
	PrintAnimalPicture("bear")
}

func (d *Owl) ChangeEat(b bool) {
	if b {
		d.Food++
	} else {
		d.Food--
	}
}

func (d *Owl) ChangeCleanness(b bool) {
	if b {
		d.Clean++
	} else {
		d.Clean--
	}
}

func (d *Owl) ChangeMood(b bool) {
	if b {
		d.Mood++
	} else {
		d.Mood--
	}
}

func (d *Owl) Print() {
	PrintAnimalPicture("owl")
}

func (d *Dog) ChangeEat(b bool) {
	if b {
		d.Food++
	} else {
		d.Food--
	}
}

func (d *Dog) ChangeCleanness(b bool) {
	if b {
		d.Clean++
	} else {
		d.Clean--
	}
}

func (d *Dog) ChangeMood(b bool) {
	if b {
		d.Mood++
	} else {
		d.Mood--
	}
}

func (d *Dog) Print() {
	PrintAnimalPicture("dog")
}

func (c *Cat) ChangeEat(b bool) {
	if b {
		c.Food++
	} else {
		c.Food--
	}
}

func (c *Cat) ChangeCleanness(b bool) {
	if b {
		c.Clean++
	} else {
		c.Clean--
	}
}

func (c *Cat) ChangeMood(b bool) {
	if b {
		c.Mood++
	} else {
		c.Mood--
	}
}

//////////////////////////////////////////////////////////////////
