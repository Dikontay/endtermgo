package internal

import (
	"time"
)

type Itimer interface {
	notify()
	updateState()
}

type Timer struct {
	secondSinceStart time.Time
	animal           AnimalStaregy
}

type StatsObserver interface {
	updateAll()
	updateRadost()
	updateCleanness()
	updateFood()
}

func (t *Timer) updateState() {
	
	for {
		timenow := time.Now().Second()

		if (timenow-t.secondSinceStart.Second())%t.animal.GetTimeToEat() == 0 {
			t.animal.ChangeEat(false)
		}
		if (timenow-t.secondSinceStart.Second())%t.animal.GetTimeToPlay() == 0 {
			t.animal.ChangeMood(false)
		}
		if (timenow-t.secondSinceStart.Second())%t.animal.GetTimeToBeDirty() == 0{
			t.animal.ChangeCleanness(false)
		}
		time.Sleep(time.Second)
	}
	
}