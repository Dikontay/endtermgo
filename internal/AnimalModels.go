package internal

//STRATEGY

type AnimalContext struct {
	Behavior AnimalStaregy
	killer Killer
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

	notifyKiller()
	Print()
}

type Dog struct {
	Clean int
	Mood  int
	Food  int

	TimeToBeDirty int
	TimeToEat     int
	TimeToPlay    int

	killer Killer
}

type Bear struct {
	Clean int
	Mood  int
	Food  int

	TimeToBeDirty int
	TimeToEat     int
	TimeToPlay    int

	killer Killer
}

type Owl struct {
	Clean int
	Mood  int
	Food  int

	TimeToBeDirty int
	TimeToEat     int
	TimeToPlay    int

	killer Killer
}

type Cat struct {
	Clean int
	Mood  int
	Food  int

	TimeToBeDirty int
	TimeToEat     int
	TimeToPlay    int

	killer Killer
}

func (d *Bear) ChangeEat(b bool) {
	if b && d.Food != 10 {
		d.Food++
	} else {
		d.Food--
		if d.Food == 0{
			d.notifyKiller()
		}
	}
}

func (d *Bear) ChangeCleanness(b bool) {
	if b && d.Clean != 10 {
		d.Clean++
	} else {
		d.Clean--
		if d.Clean == 0{
			d.notifyKiller()
		}
	}
}

func (d *Bear) ChangeMood(b bool) {
	if b && d.Mood != 10 {
		d.Mood++
	} else {
		d.Mood--
		if d.Mood == 0{
			d.notifyKiller()
		}
	}
}

func (d *Bear) Print() {
	PrintAnimalPicture("bear")
}

func (d *Owl) ChangeEat(b bool) {
	if b && d.Food != 10 {
		d.Food++
	} else {
		d.Food--
		if d.Food == 0{
			d.notifyKiller()
		}
	}
}

func (d *Owl) ChangeCleanness(b bool) {
	if b && d.Clean != 10 {
		d.Clean++
	} else {
		d.Clean--
		if d.Clean == 0{
			d.notifyKiller()
		}
	}
}

func (d *Owl) ChangeMood(b bool) {
	if b && d.Mood != 10 {
		d.Mood++
	} else {
		d.Mood--
		if d.Mood == 0{
			d.notifyKiller()
		}
	}
}

func (d *Owl) Print() {
	PrintAnimalPicture("owl")
}

func (d *Dog) ChangeEat(b bool) {
	if b && d.Food != 10 {
		d.Food++
	} else {
		d.Food--
		if d.Food == 0{
			d.notifyKiller()
		}
	}
}

func (d *Dog) ChangeCleanness(b bool) {
	if b && d.Clean != 10 {
		d.Clean++
	} else {
		d.Clean--
		if d.Clean == 0{
			d.notifyKiller()
		}
	}
}

func (d *Dog) ChangeMood(b bool) {
	if b && d.Mood != 10 {
		d.Mood++
	} else {
		d.Mood--
		if d.Mood == 0{
			d.notifyKiller()
		}
	}
}

func (d *Dog) Print() {
	PrintAnimalPicture("dog")
}

func (d *Cat) ChangeEat(b bool) {
	if b && d.Food != 10 {
		d.Food++
	} else {
		d.Food--
		if d.Food == 0{
			d.notifyKiller()
		}
	}
}

func (d *Cat) ChangeCleanness(b bool) {
	if b && d.Clean != 10 {
		d.Clean++
	} else {
		d.Clean--
		if d.Clean == 0{
			d.notifyKiller()
		}
	}
}

func (d *Cat) ChangeMood(b bool) {
	if b && d.Mood != 10 {
		d.Mood++
	} else {
		d.Mood--
		if d.Mood == 0{
			d.notifyKiller()
		}
	}
}

func (d *Cat) Print() {
	PrintAnimalPicture("cat")
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

func (d *Dog) notifyKiller(){
	d.killer.kill()
}

func (d *Cat) notifyKiller(){
	d.killer.kill()
}

func (d *Owl) notifyKiller(){
	d.killer.kill()
}

func (d *Bear) notifyKiller(){
	d.killer.kill()
}
//////////////////////////////////////////////////////////////////
