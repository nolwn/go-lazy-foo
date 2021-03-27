package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

//Screen dimension constants
const screenWidth = 640
const screenHeight = 480

//The window we'll be rendering to
var gWindow *sdl.Window

//The window renderer
var gRenderer *sdl.Renderer

//Current displayed texture
// var gTexture *sdl.Texture

// graphic
var gBackgroundTexture lTexture
var gModulatedTexture lTexture

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
	if gRenderer, err = sdl.CreateRenderer(gWindow, -1, sdl.RENDERER_ACCELERATED); err != nil {
		fmt.Printf("Renderer could not be created! SDL Error: %s\n", err)
		panic(err)
	}

	//Initialize renderer color
	gRenderer.SetDrawColor(0xFF, 0xFF, 0xFF, 0xFF)

	//Initialize PNG loading
	if err := img.Init(img.INIT_PNG); err != nil {
		fmt.Printf("SDL_image could not initialize! SDL_image Error: %s\n", err)
	}
}

func loadMedia() (err error) {
	//Load textures
	if err = gBackgroundTexture.loadFromFile("media/fadein.png"); err != nil {
		fmt.Printf("Could not load media/fadein.png")
		return
	}

	if err = gModulatedTexture.loadFromFile("media/fadeout.png"); err != nil {
		fmt.Printf("Could not load media/fadeout.png")
		return
	}

	gModulatedTexture.setBlendMode(sdl.BLENDMODE_BLEND)

	return
}

func close() {
	gBackgroundTexture.free()

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

		// Modulation component
		var a uint8 = 255

		//While application is running
		for !quit {
			//Event handler
			e := sdl.PollEvent()

			//Handle events on queue
			for e != nil {
				//User requests quit
				if e.GetType() == sdl.QUIT {
					quit = true
				} else if e.GetType() == sdl.KEYDOWN {
					switch e.(*sdl.KeyboardEvent).Keysym.Sym {
					case sdl.K_w:
						if a > 255-32 {
							a = 255
						} else {
							a += 32
						}
					case sdl.K_s:
						if a < 32 {
							a = 0
						} else {
							a -= 32
						}
					}
				}

				e = sdl.PollEvent()
			}

			//Clear screen
			gRenderer.SetDrawColor(0xFF, 0xFF, 0xFF, 0xFF)
			gRenderer.Clear()

			gBackgroundTexture.render(0, 0, nil)

			gModulatedTexture.setAlpha(a)
			gModulatedTexture.render(0, 0, nil)

			//Update screen
			gRenderer.Present()
		}
	}
}
