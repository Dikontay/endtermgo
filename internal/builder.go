package internal

type AccessoryBuilder struct {
	animal    AnimalStaregy
	accessory Accessory
}

func NewAccessoryBuilder(animal AnimalStaregy) *AccessoryBuilder {
	return &AccessoryBuilder{
		animal:    animal,
		accessory: nil,
	}
}

func (b *AccessoryBuilder) AddHat() *AccessoryBuilder {
	b.accessory = &HatAccessory{}
	return b
}

func (b *AccessoryBuilder) AddCap() *AccessoryBuilder {
	b.accessory = &CapAccessory{}
	return b
}

func (b *AccessoryBuilder) AddOfficial() *AccessoryBuilder {
	b.accessory = &OfficialAccessory{}
	return b
}

func (b *AccessoryBuilder) Build() AnimalStaregy {
	if b.accessory != nil {
		b.accessory.Apply(b.animal)
	}
	return b.animal
}
