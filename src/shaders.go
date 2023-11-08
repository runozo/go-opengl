package main

import (
	"log"
	"strings"

	"github.com/go-gl/gl/v4.1-core/gl"
)

func createVertexShader() uint32 {
	// Vertex Shader
	vertexShaderSource := `
		#version 330 core
		layout (location = 0) in vec3 aPos;
		void main() {
			gl_Position = vec4(aPos.x, aPos.y, aPos.z, 1.0);
		}
	`

	vertexShader := gl.CreateShader(gl.VERTEX_SHADER)
	csource, free := gl.Strs(vertexShaderSource)
	defer free()
	gl.ShaderSource(vertexShader, 1, csource, nil)
	gl.CompileShader(vertexShader)

	var success int32
	gl.GetShaderiv(vertexShader, gl.COMPILE_STATUS, &success)

	if success == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(vertexShader, gl.INFO_LOG_LENGTH, &logLength)
		text := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(vertexShader, logLength, nil, gl.Str(text))
		log.Fatalln("failed to compile vertex shader:", text)
	}
	return vertexShader
}

func createFragmentShader() uint32 {
	fragmentShaderSource := `
		#version 330 core
		out vec4 FragColor;
		
		void main()
		{
			FragColor = vec4(1.0f, 0.5f, 0.2f, 1.0f);
		} 
	`
	fragmentShader := gl.CreateShader(gl.FRAGMENT_SHADER)
	csource2, free2 := gl.Strs(fragmentShaderSource)
	defer free2()
	gl.ShaderSource(fragmentShader, 1, csource2, nil)
	gl.CompileShader(fragmentShader)

	return fragmentShader
}
