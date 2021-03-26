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

// Scene sprites
var gSpriteClips [4]sdl.Rect
var gSpriteSheetTexture *lTexture = &lTexture{}

// Texture wrapper
type lTexture struct {
	// The hardware texture
	mTexture *sdl.Texture

	// Image dimensions
	mWidth  int32
	mHeight int32
}

func (l *lTexture) loadFromFile(path string) (err error) {
	if l != nil {
		l.free()
	}

	var loadedSurface *sdl.Surface
	defer loadedSurface.Free()

	var newTexture *sdl.Texture

	if loadedSurface, err = img.Load(path); err != nil {
		fmt.Printf("Unable to load image %s! IMG error: %s", path, err)
		return
	}

	loadedSurface.SetColorKey(true, sdl.MapRGB(loadedSurface.Format, 0, 0xff, 0xff))

	if newTexture, err = gRenderer.CreateTextureFromSurface(loadedSurface); err != nil {
		fmt.Printf("Unable to load texture. SDL Error: %s", err)
		return
	}

	l.mWidth = loadedSurface.W
	l.mHeight = loadedSurface.H
	l.mTexture = newTexture

	return
}

func (l *lTexture) free() {
	fmt.Printf("Here's the deal %v\n", l)
	if l == nil {
		fmt.Print("not needed\n")

		return
	}

	if l.mTexture != nil {
		l.mTexture.Destroy()
		l.mTexture = nil
		l.mWidth = 0
		l.mHeight = 0
	}
}

func (l *lTexture) render(x int32, y int32, clip *sdl.Rect) {
	renderQuad := sdl.Rect{}
	renderQuad.X = x
	renderQuad.Y = y
	renderQuad.W = clip.W
	renderQuad.H = clip.H
	gRenderer.Copy(l.mTexture, clip, &renderQuad)
}

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
	if err = gSpriteSheetTexture.loadFromFile("media/dots.png"); err != nil {
		fmt.Printf("I broke.")
	}

	// Load rects
	gSpriteClips[0] = sdl.Rect{
		W: 100,
		H: 100,
		X: 0,
		Y: 0,
	}

	gSpriteClips[1] = sdl.Rect{
		W: 100,
		H: 100,
		X: 100,
		Y: 0,
	}

	gSpriteClips[2] = sdl.Rect{
		W: 100,
		H: 100,
		X: 0,
		Y: 100,
	}

	gSpriteClips[3] = sdl.Rect{
		W: 100,
		H: 100,
		X: 100,
		Y: 100,
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

			gSpriteSheetTexture.render(0, 0, &gSpriteClips[0])
			gSpriteSheetTexture.render(0, screenHeight-gSpriteClips[1].H, &gSpriteClips[1])
			gSpriteSheetTexture.render(screenWidth-gSpriteClips[2].W, 0, &gSpriteClips[2])
			gSpriteSheetTexture.render(screenWidth-gSpriteClips[3].W, screenHeight-gSpriteClips[3].H, &gSpriteClips[3])

			//Update screen
			gRenderer.Present()
		}
	}
}
