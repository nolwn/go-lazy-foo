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

// Walking animation
const walkingAnimationFrames = 4

var gSpriteClips [walkingAnimationFrames]sdl.Rect
var gSpriteSheetTexture lTexture

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
	}
}

func loadMedia() (err error) {
	//Load textures
	if err = gSpriteSheetTexture.loadFromFile("media/foo.png"); err != nil {
		fmt.Printf("Could not load media/foo.png")
		return
	}

	gSpriteClips[0] = sdl.Rect{
		X: 0,
		Y: 0,
		W: 64,
		H: 205,
	}

	gSpriteClips[1] = sdl.Rect{
		X: 62,
		Y: 0,
		W: 64,
		H: 205,
	}

	gSpriteClips[2] = sdl.Rect{
		X: 128,
		Y: 0,
		W: 64,
		H: 205,
	}

	gSpriteClips[3] = sdl.Rect{
		X: 196,
		Y: 0,
		W: 64,
		H: 205,
	}

	return
}

func close() {
	gSpriteSheetTexture.free()

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

		//current animation frame
		frame := 0

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

			frame++
			if frame/4 >= walkingAnimationFrames {
				frame = 0
			}

			//Clear screen
			gRenderer.SetDrawColor(0xFF, 0xFF, 0xFF, 0xFF)
			gRenderer.Clear()

			currentClip := &gSpriteClips[frame/4]
			gSpriteSheetTexture.render((screenWidth-currentClip.W)/2, (screenHeight-currentClip.H)/2, currentClip)

			//Update screen
			gRenderer.Present()
		}
	}
}
