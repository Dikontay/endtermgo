package internal

//STRATEGY

type AnimalContext struct {
	Behavior AnimalStaregy
}

type AnimalStaregy interface {
	Eat() string
	BeDirty() string
	Play() string
	
	Print()
}

type Dog struct {
	Name string
}

type Bear struct {
	Name string
}

type Owl struct {
	Name string
}

func (d *Bear) Eat() string {
	return "++++++++++"
}

func (d *Bear) BeDirty() string {
	return "++++++++++"
}

func (d *Bear) Play() string {
	return "++++++++++"
}

func (d *Bear) Print() {
	PrintAnimalPicture("bear")
}

func (d *Owl) Eat() string {
	return "++++++++++"
}

func (d *Owl) BeDirty() string {
	return "++++++++++"
}

func (d *Owl) Play() string {
	return "++++++++++"
}

func (d *Owl) Print() {
	PrintAnimalPicture("owl")
}

func (d *Dog) Eat() string {
	return "++++++++++"
}

func (d *Dog) BeDirty() string {
	return "++++++++++"
}

func (d *Dog) Play() string {
	return "++++++++++"
}

func (d *Dog) Print() {
	PrintAnimalPicture("dog")
}

type Cat struct {
	Name string
}

func (c *Cat) Eat() string {
	return "++++++++++"
}

func (c *Cat) BeDirty() string {
	return "++++++++++"
}

func (c *Cat) Play() string {
	return "++++++++++"
}

//////////////////////////////////////////////////////////////////
