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
var gTexture *sdl.Texture

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
	//Load texture

	if gTexture = loadTexture("media/viewport.png"); err != nil {
		fmt.Printf("Failed to load texture image!\n")
	}

	return
}

func close() {
	//Free loaded image
	gTexture.Destroy()

	//Destroy window
	gRenderer.Destroy()
	gWindow.Destroy()

	//Quit SDL subsystems
	img.Quit()
	sdl.Quit()
}

func loadTexture(path string) (newTexture *sdl.Texture) {
	var err error

	//Load image at specified path
	loadedSurface, err := img.Load(path)

	if err != nil {
		fmt.Printf("Unable to load image %s! SDL_image Error: %s\n", path, err)
	}
	//Create texture from surface pixels
	if newTexture, err = gRenderer.CreateTextureFromSurface(loadedSurface); err != nil {
		fmt.Printf("Unable to create texture from %s! SDL Error: %s\n", path, err)
	}

	//Get rid of old loaded surface
	loadedSurface.Free()

	return
}

func main() {
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
			gRenderer.SetDrawColor(0x00, 0xFF, 0xFF, 0xFF)
			gRenderer.Clear()

			//Top left corner viewport
			topLeftViewport := sdl.Rect{
				X: 0,
				Y: 0,
				W: screenWidth,
				H: screenHeight,
			}

			gRenderer.Copy(gTexture, nil, nil)
			gRenderer.SetViewport(&topLeftViewport)

			//Render texture to screen
			gRenderer.Copy(gTexture, nil, nil)

			//Top right viewport
			topRightViewport := sdl.Rect{
				X: screenWidth / 2,
				Y: 0,
				W: screenWidth / 2,
				H: screenHeight / 2,
			}
			gRenderer.SetViewport(&topRightViewport)

			//Render texture to screen
			gRenderer.Copy(gTexture, nil, nil)

			// //Bottom viewport
			// bottomViewport := sdl.Rect{
			// 	X: 0,
			// 	Y: screenHeight / 2,
			// 	W: screenWidth,
			// 	H: screenHeight / 2,
			// }
			// gRenderer.SetViewport(&bottomViewport)

			// //Render texture to screen
			// gRenderer.Copy(gTexture, nil, nil)

			//Update screen
			gRenderer.Present()
		}
	}

	//Free resources and close SDL
	close()
}
