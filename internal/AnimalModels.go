package internal

//STRATEGY

type AnimalContext struct {
	Behavior AnimalStaregy
}

type AnimalStaregy interface {
	ChangeCleanness(bool)
	ChangeEat(bool)
	ChangeMood(bool)

	GetTimeToBeDirty() int
	GetTimeToEat() int
	GetTimeToPlay() int

	GetMood() int
	GetFood() int
	GetCleanness() int

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

func (d *Dog) GetCleanness() int {
	return d.Clean
}

func (d *Cat) GetCleanness() int {
	return d.Clean
}

func (d *Owl) GetCleanness() int {
	return d.Clean
}

func (d *Bear) GetCleanness() int {
	return d.Clean
}

func (d *Dog) GetMood() int {
	return d.Mood
}

func (d *Cat) GetMood() int {
	return d.Mood
}

func (d *Owl) GetMood() int {
	return d.Mood
}

func (d *Bear) GetMood() int {
	return d.Mood
}

func (d *Dog) GetFood() int {
	return d.Food
}

func (d *Cat) GetFood() int {
	return d.Food
}

func (d *Owl) GetFood() int {
	return d.Food
}

func (d *Bear) GetFood() int {
	return d.Food
}

func (d *Dog) GetTimeToBeDirty() int {
	return d.TimeToBeDirty
}

func (d *Cat) GetTimeToBeDirty() int {
	return d.TimeToBeDirty
}

func (d *Owl) GetTimeToBeDirty() int {
	return d.TimeToBeDirty
}

func (d *Bear) GetTimeToBeDirty() int {
	return d.TimeToBeDirty
}

func (d *Dog) GetTimeToEat() int {
	return d.TimeToEat
}

func (d *Cat) GetTimeToEat() int {
	return d.TimeToEat
}

func (d *Owl) GetTimeToEat() int {
	return d.TimeToEat
}

func (d *Bear) GetTimeToEat() int {
	return d.TimeToEat
}

func (d *Dog) GetTimeToPlay() int {
	return d.TimeToPlay
}

func (d *Cat) GetTimeToPlay() int {
	return d.TimeToPlay
}

func (d *Owl) GetTimeToPlay() int {
	return d.TimeToPlay
}

func (d *Bear) GetTimeToPlay() int {
	return d.TimeToPlay
}
//////////////////////////////////////////////////////////////////
