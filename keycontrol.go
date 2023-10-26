package main

import (
	"fmt"
	"os"
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
)

func keycontrol() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	stopAnimation = make(chan bool)
	animationTicker = time.NewTicker(100 * time.Millisecond)

	go runAnimation()

	// Handle signals to gracefully stop the animation
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-signalCh
		stopAnimation <- true
	}()

	// Main loop for user input
mainLoop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			if ev.Key == termbox.KeyEsc {
				break mainLoop
			}
			// Handle other key events as needed
			handleKeyEvent(ev)
		}
	}

	// Stop the animation and wait for it to finish
	stopAnimation <- true
	animationMutex.Lock()
}

func runAnimation() {
	for {
		select {
		case <-stopAnimation:
			return
		case <-animationTicker.C:
			animationMutex.Lock()
			printAnimation()
			animationMutex.Unlock()
		}
	}
}

func printAnimation() {
	// Replace this with your animation logic
	fmt.Print("\033[H") // Move cursor to the top-left corner of the screen
	fmt.Println("Animating...")
}

func handleKeyEvent(ev termbox.Event) {
	// Handle user input events as needed
	// You can add logic here to respond to specific key presses
	fmt.Printf("Key pressed: %c\n", ev.Ch)
}
