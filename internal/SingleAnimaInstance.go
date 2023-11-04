package internal

var animalContext *AnimalContext 

func getInstance() *AnimalContext{
	return animalContext
}

func setUpdatedInstance(a *AnimalContext){
	animalContext = a
}