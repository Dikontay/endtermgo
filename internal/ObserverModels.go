package internal

import "time"

type Itimer interface{
	notify()
	updateState()
}

type Timer struct{
	secondSinceStart time.Time
	animal AnimalContext
}

type StatsObserver interface{
	updateAll()
	updateRadost()
	updateCleanness()
	updateFood()
}

func (t *Timer) updateState(timenow int){
	
	t.animal.ChangeEat(timenow-t.secondSinceStart.Second() %==0)
	t.animal.ChangeMood(timenow-t.secondSinceStart.Second() % 4==0)
	t.animal.ChangeCleanness(t.secondSinceStart.Second() % 8==0)


	if (timenow-t.secondSinceStart.Second()) % 2==0{
		t.animal.ChangeEat( true)  
	} else {
		t.animal.ChangeEat(false)  
	}
	
	if (timenow-t.secondSinceStart.Second()) % 4==0 {
		t.animal.ChangeCleanness(true)  
	} else {
		t.animal.ChangeCleanness(false)  
	}
	
	
	if t.secondSinceStart.Second() % 8==0{
		t.animal.ChangeMood(true)
	} else {
		t.animal.ChangeMood(false)
	}

	// switch (t.secondSinceStart.Second()) {
	// 	case t.secondSinceStart.Second() % 2 :
	// 		t.animal.ChangeEat(true) 
		
	// 	break;
	// 	case t.secondSinceStart.Second() % 4: 
	// 	t.animal.ChangeEat(true) 
		
	// 	break;
	// 	case t.secondSinceStart.Second() % 8:
		
	// 	break;
	// }
}

func (t *Timer) notify(){
	
}

// интерфейс таймера(observable) <- анимал(subscriber)
// убиватор -> анимал
 