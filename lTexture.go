package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

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

	if clip != nil {
		renderQuad.W = clip.W
		renderQuad.H = clip.H
	} else {
		renderQuad.W = screenWidth
		renderQuad.H = screenHeight
	}

	gRenderer.Copy(l.mTexture, clip, &renderQuad)
}

func (l *lTexture) setBlendMode(blending sdl.BlendMode) {
	l.mTexture.SetBlendMode(blending)
}

func (l *lTexture) setAlpha(alpha uint8) {
	l.mTexture.SetAlphaMod(alpha)
}
