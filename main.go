package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

// Screen dimension constants
const screenWidth = 640
const screenHeight = 480

//The window we'll be rendering to
var gWindow *sdl.Window

//The surface contained by the window
var gScreenSurface *sdl.Surface

//The image we will load and show on the screen
var gHelloWorld *sdl.Surface

func init() {
	var err error

	//Initialize SDL
	if err = sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		fmt.Printf("SDL could not initialize! SDL_Error: %s\n", err)
		panic(err)
	}

	if gWindow, err = sdl.CreateWindow(
		"SDL Tutorial",
		sdl.WINDOWPOS_UNDEFINED,
		sdl.WINDOWPOS_UNDEFINED,
		screenWidth,
		screenHeight,
		sdl.WINDOW_SHOWN); err != nil {
		fmt.Printf("SDL could not initialize! SDL_Error: %s\n", err)

		panic(err)
	}

	if gScreenSurface, err = gWindow.GetSurface(); err != nil {
		fmt.Printf("SDL could not initialize! SDL_Error: %s\n", err)

		panic(err)
	}
}

func loadMedia() {
	var err error

	//Load splash image
	gHelloWorld, err = sdl.LoadBMP("media/hello_world.bmp")
	if err != nil {
		fmt.Printf("Unable to load image %s! SDL Error: %s\n", "02_getting_an_image_on_the_screen/hello_world.bmp", err)
	}
}

func close() {
	//Deallocate surface
	gHelloWorld.Free()

	//Destroy window
	gWindow.Destroy()

	//Quit SDL subsystems
	sdl.Quit()
}

func main() {
	defer close()

	//Load media
	loadMedia()

	running := true

	for running {
		e := sdl.PollEvent()

		for e != nil {
			fmt.Printf("%v", e)
			if e.GetType() == sdl.QUIT {
				running = false
			}

			e = sdl.PollEvent()
		}

		//Apply the image
		gHelloWorld.Blit(nil, gScreenSurface, nil)

		//Update the surface
		gWindow.UpdateSurface()
	}
}
