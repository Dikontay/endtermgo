package internal

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/nsf/termbox-go"
)

var (
	animationTicker *time.Ticker
	animationMutex  sync.Mutex
	stopAnimation   chan bool
	isShop 			bool
)

func Keycontrol(animal *AnimalContext) {
	isShop = false
	
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	stopAnimation = make(chan bool)
	animationTicker = time.NewTicker(100 * time.Millisecond)

	go runAnimation(animal)

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-signalCh
		stopAnimation <- true
	}()


	eventQueue := make(chan termbox.Event)
	go func() {
		for {
			eventQueue <- termbox.PollEvent()
		}
	}()

mainLoop: 	// Main loop for user input
	for {
		switch ev := <-eventQueue;{
		
		case  ev.Type == termbox.EventKey:
			if ev.Key == termbox.KeyEsc{
				break mainLoop
			}
			if !isShop{ 
				handleKeyEvent(ev, animal)
			} else {
				ShophandleKeyEvent(ev, animal)
			}
		}
	}

	
	stopAnimation <- true	// Stop the animation and wait for it to finish
	animationMutex.Lock()
}

func runAnimation(animal *AnimalContext) {
	for {
		select {
		case <-stopAnimation:
			return
		case <-animationTicker.C:
			animationMutex.Lock()
			printAnimation(animal)
			animationMutex.Unlock()
		}
	}
}

func printAnimation(animal *AnimalContext) {
	fmt.Print("\033[H") // Move cursor to the top-left corner of the screen

	fmt.Print("Food: ")
	for i := 0; i < animal.Behavior.GetFood(); i++ {
		fmt.Print(string('⭓'))
	}
	fmt.Println()

	fmt.Print("Cleanness: ")
	for i := 0; i < animal.Behavior.GetCleanness(); i++ {
		fmt.Print(string('⭓'))
	}
	fmt.Println()

	fmt.Print("Mood: ")
	for i := 0; i < animal.Behavior.GetMood(); i++ {
		fmt.Print(string('⭓'))
	}
	fmt.Print("\n\n")

	animal.Behavior.Print()
	fmt.Print("\n\n")
	fmt.Printf("1) Feed 	2) Clean	3) Play		4) Shop \n")

	time.Sleep(time.Second)
	cmd := exec.Command("cmd", "/c", "cls")
    cmd.Stdout = os.Stdout
    cmd.Run()
}

func handleKeyEvent(ev termbox.Event, animal *AnimalContext) {
	switch ev.Ch {
	case '1' :
		animal.Behavior.ChangeEat(true)
		break
	case '2':
		animal.Behavior.ChangeCleanness(true)
		break
	case '3':
		animal.Behavior.ChangeMood(true)
		break
	case '4':
		isShop = true
		stopAnimation <- true
		go runShopAnimation(animal)

	default:
		fmt.Println("Wrong Button!")
	}

	fmt.Printf("you pressed %c\n", ev.Ch)
}

func runShopAnimation(animal *AnimalContext){
	for {
		select {
		case <-stopAnimation:
			return
		case <-animationTicker.C:
			animationMutex.Lock()
			printShopAnimation(animal)
			animationMutex.Unlock()
		}
	}
}

func printShopAnimation(animal *AnimalContext) {
	fmt.Println("Buy: 1) Hat	2) Boots	3) Umbrella    4) Return to pet")
	time.Sleep(time.Second)
	cmd := exec.Command("cmd", "/c", "cls")
    cmd.Stdout = os.Stdout
    cmd.Run()
}

func ShophandleKeyEvent(ev termbox.Event, animal *AnimalContext) {
	switch ev.Ch {
	case '1' :
		setUpdatedInstance(&AnimalContext{Behavior: &Hat{animal: animal.Behavior}, killer: Killer{}})
		break
	case '2':
		setUpdatedInstance(&AnimalContext{Behavior: &Hat{animal: animal.Behavior}, killer: Killer{}})
		break
	case '3':
		setUpdatedInstance(&AnimalContext{Behavior: &Hat{animal: animal.Behavior}, killer: Killer{}})
		break
	case '4':
		isShop = false
		stopAnimation <- true
		go runAnimation(animalContext)
		

	default:
		fmt.Println("Wrong Button!")
	}

	fmt.Printf("you pressed %c\n", ev.Ch)
}