package internal

import "time"

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

func (t *Timer) updateState(timenow int) {

	t.animal.ChangeEat((timenow-t.secondSinceStart.Second())%t.animal.GetTimeToEat()== 0)
	t.animal.ChangeMood(timenow-t.secondSinceStart.Second()%t.animal.GetTimeToPlay() == 0)
	t.animal.ChangeCleanness(t.secondSinceStart.Second()%t.animal.GetTimeToBeDirty() == 0)

}