package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

//Screen dimension constants
const screenWidth = 640
const screenHeight = 480

//The window we'll be rendering to
var gWindow *sdl.Window

//The window renderer
var gRenderer *sdl.Renderer

//Globallty used font
var gFont *ttf.Font

//Rendered texture
var gCurrentTexture *lTexture

var gUpTexture lTexture
var gRightTexture lTexture
var gDownTexture lTexture
var gLeftTexture lTexture
var gPressTexture lTexture

func init() {
	var err error

	//Initialize SDL
	if err = sdl.Init(sdl.INIT_VIDEO); err != nil {
		fmt.Printf("SDL could not initialize! SDL Error: %s\n", err)
		panic(err)
	}

	if enabled := sdl.SetHint(
		sdl.HINT_RENDER_SCALE_QUALITY,
		"1"); !enabled { //Set texture filtering to linear
		fmt.Printf("Warning: Linear texture filtering not enabled!\n")
	}

	//Create window
	if gWindow, err = sdl.CreateWindow(
		"SDL Tutorial",
		sdl.WINDOWPOS_UNDEFINED,
		sdl.WINDOWPOS_UNDEFINED,
		screenWidth,
		screenHeight,
		sdl.WINDOW_SHOWN); err != nil {
		fmt.Printf("Window could not be created! SDL Error: %s\n", err)
		panic(err)
	}

	//Create renderer for window
	if gRenderer, err = sdl.CreateRenderer(gWindow, -1, sdl.RENDERER_ACCELERATED|sdl.RENDERER_PRESENTVSYNC); err != nil {
		fmt.Printf("Renderer could not be created! SDL Error: %s\n", err)
		panic(err)
	}

	//Initialize renderer color
	gRenderer.SetDrawColor(0xFF, 0xFF, 0xFF, 0xFF)

	//Initialize PNG loading
	if err := img.Init(img.INIT_PNG); err != nil {
		fmt.Printf("SDL_image could not initialize! SDL_image Error: %s\n", err)
		panic(err)
	}

	//Initialize SDL_ttf
	if err := ttf.Init(); err != nil {
		fmt.Printf("SDL_ttf could not initialize! SDL_ttf Error: %s\n", err)
		panic(err)
	}
}

func loadMedia() (err error) {
	if err = gUpTexture.loadFromFile("media/up.png"); err != nil {
		fmt.Printf("Failed to render text texture!\n")
	}

	if err = gRightTexture.loadFromFile("media/right.png"); err != nil {
		fmt.Printf("Failed to render text texture!\n")
	}

	if err = gDownTexture.loadFromFile("media/down.png"); err != nil {
		fmt.Printf("Failed to render text texture!\n")
	}

	if err = gLeftTexture.loadFromFile("media/left.png"); err != nil {
		fmt.Printf("Failed to render text texture!\n")
	}

	if err = gPressTexture.loadFromFile("media/press.png"); err != nil {
		fmt.Printf("Failed to render text texture!\n")
	}

	return
}

func close() {
	gCurrentTexture.free()

	gFont.Close()
	gFont = nil

	gRenderer.Destroy()
	gWindow.Destroy()
	gRenderer = nil
	gWindow = nil

	//Quit SDL subsystems
	img.Quit()
	sdl.Quit()
}

func main() {
	//Free resources and close SDL
	defer close()

	//Load media
	if err := loadMedia(); err != nil {
		fmt.Printf("Failed to load media!\n")
	} else {
		//Main loop flag
		quit := false

		//While application is running
		for !quit {
			//Event handler
			e := sdl.PollEvent()

			//Handle events on queue
			for e != nil {
				//User requests quit
				if e.GetType() == sdl.QUIT {
					quit = true
				}

				e = sdl.PollEvent()
			}

			//Clear screen
			gRenderer.SetDrawColor(0xFF, 0xFF, 0xFF, 0xFF)
			gRenderer.Clear()

			//set texture based on current keystate
			currentKeyStates := sdl.GetKeyboardState()
			if currentKeyStates[sdl.SCANCODE_UP] == 1 {
				gCurrentTexture = &gUpTexture
			} else if currentKeyStates[sdl.SCANCODE_RIGHT] == 1 {
				gCurrentTexture = &gRightTexture
			} else if currentKeyStates[sdl.SCANCODE_DOWN] == 1 {
				gCurrentTexture = &gDownTexture
			} else if currentKeyStates[sdl.SCANCODE_LEFT] == 1 {
				gCurrentTexture = &gLeftTexture
			} else {
				gCurrentTexture = &gPressTexture
			}

			gCurrentTexture.render(0, 0, nil, 0, nil, 0)

			//Update screen
			gRenderer.Present()
		}
	}
}
