package internal

func newDog() *Dog {
	return &Dog{
		Clean: 10,
		Mood: 10,
		Food: 10,
		TimeToBeDirty: 4,
		TimeToEat: 6,
		TimeToPlay: 3,
		killer: Killer{},
	}
}

func newOwl() *Owl {
	return &Owl{
		Clean: 10,
		Mood: 10,
		Food: 10,
		TimeToBeDirty: 10,
		TimeToEat: 12,
		TimeToPlay: 12,
		killer: Killer{},
	}
}

func newBear() *Bear {
	return &Bear{
		Clean: 10,
		Mood: 10,
		Food: 10,
		TimeToBeDirty: 20,
		TimeToEat: 2,
		TimeToPlay: 6,
		killer: Killer{},
	}
}

func newCat() *Cat {
	return &Cat{
		Clean: 10,
		Mood: 10,
		Food: 10,
		TimeToBeDirty: 6,
		TimeToEat: 4,
		TimeToPlay: 8,
		killer: Killer{},
	}
}