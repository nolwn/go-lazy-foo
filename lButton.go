package main

import "github.com/veandco/go-sdl2/sdl"

const (
	buttonSpriteMouseOut = iota
	buttonSpriteMouseOverMotion
	buttonSpriteMouseDown
	buttonSpriteMouseUp
	buttonSpriteTotal
)

type lButton struct {
	mPosition      sdl.Point
	mCurrentSprite int
}

func (b *lButton) setPosition(x int32, y int32) {
	b.mPosition.X = x
	b.mPosition.Y = y
}

func (b *lButton) handleEvent(e *sdl.Event) {
	var x, y int32

	//if mouse event happened
	if (*e).GetType() == sdl.MOUSEMOTION ||
		(*e).GetType() == sdl.MOUSEBUTTONUP ||
		(*e).GetType() == sdl.MOUSEBUTTONDOWN {
		//get mouse position
		x, y, _ = sdl.GetMouseState()
	}

	isInside := true

	if x < b.mPosition.X { //mouse is left of the button
		isInside = false
	} else if x > b.mPosition.X+buttonWidth { //mouse is right of the button
		isInside = false
	} else if y < b.mPosition.Y { //mouse is above the button
		isInside = false
	} else if y > b.mPosition.Y+buttonHeight { //mouse is belove the button
		isInside = false
	}

	if !isInside { //mouse is outside the button
		b.mCurrentSprite = buttonSpriteMouseOut
	} else {
		switch (*e).GetType() {
		case sdl.MOUSEMOTION:
			b.mCurrentSprite = buttonSpriteMouseOverMotion
		case sdl.MOUSEBUTTONUP:
			b.mCurrentSprite = buttonSpriteMouseUp
		case sdl.MOUSEBUTTONDOWN:
			b.mCurrentSprite = buttonSpriteMouseDown
		}
	}
}

func (b *lButton) render() {
	//show current button sprite
	gButtonSpriteSheetTexture.render(b.mPosition.X, b.mPosition.Y, &gSpriteClips[b.mCurrentSprite], 0, nil, 0)
}
