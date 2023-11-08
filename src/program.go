package main

import (
	"log"
	"strings"

	"github.com/go-gl/gl/v4.1-core/gl"
)

func createShaderProgram(vertexShader, fragmentShader uint32) uint32 {
	// Shader Program
	shaderProgram := gl.CreateProgram()
	gl.AttachShader(shaderProgram, vertexShader)
	gl.AttachShader(shaderProgram, fragmentShader)
	gl.LinkProgram(shaderProgram)

	var success int32
	gl.GetProgramiv(shaderProgram, gl.LINK_STATUS, &success)
	if success == gl.FALSE {
		var logLength int32
		gl.GetProgramiv(shaderProgram, gl.INFO_LOG_LENGTH, &logLength)
		text := strings.Repeat("\x00", int(logLength+1))
		gl.GetProgramInfoLog(shaderProgram, logLength, nil, gl.Str(text))
		log.Fatalln("failed to link shader program:", text)
	}

	gl.DeleteShader(vertexShader)
	gl.DeleteShader(fragmentShader)
	return shaderProgram
}
