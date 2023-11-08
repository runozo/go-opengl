package main

import (
	"log"

	"github.com/go-gl/glfw/v3.3/glfw"
)

func createWindow() *glfw.Window {
	// Initialize GLFW
	if err := glfw.Init(); err != nil {
		log.Fatalln("failed to initialize glfw:", err)
	}
	// defer glfw.Terminate()
	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 3)
	glfw.WindowHint(glfw.ContextVersionMinor, 3)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
	window, err := glfw.CreateWindow(windowWidth, windowHeight, "Hello Triangle", nil, nil)
	if err != nil {
		log.Fatalln(err)
	}
	window.MakeContextCurrent()
	return window
}

func destroyWindow(window *glfw.Window) {
	window.Destroy()
	glfw.Terminate()
}
