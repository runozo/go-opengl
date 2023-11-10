package main

import (
	"log"
	"runtime"
	"unsafe"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

const windowWidth = 1280
const windowHeight = 720

func init() {
	// GLFW event handling must run on the main OS thread
	runtime.LockOSThread()
}
func main() {
	window := createWindow()
	defer destroyWindow(window)

	// Initialize Glow
	if err := gl.Init(); err != nil {
		log.Fatalln(err)
	}

	// Print opengl version
	version := gl.GoStr(gl.GetString(gl.VERSION))
	// max_vertex_attributes := gl.GoStr(gl.GetString(gl.MAX_VERTEX_ATTRIBS))

	log.Println("OpenGL version:", version)
	// log.Println("Maximum nr of vertex attributes supported:", max_vertex_attributes)

	vertices := []float32{
		0.0, 0.5, 0.0,
		0.5, -0.5, 0.0,
		-0.5, -0.5, 0.0,
	}

	// Vertex Array Object
	var VAO uint32
	gl.GenVertexArrays(1, &VAO)
	// bind the Vertex Array Object first, then bind and set vertex buffer(s), and then configure vertex attributes(s).
	gl.BindVertexArray(VAO)

	// Vertex Buffer Object
	var VBO uint32
	gl.GenBuffers(1, &VBO)
	gl.BindBuffer(gl.ARRAY_BUFFER, VBO)
	gl.BufferData(gl.ARRAY_BUFFER, len(vertices)*int(unsafe.Sizeof(vertices[0])), gl.Ptr(vertices), gl.STATIC_DRAW)

	vertexShader := createVertexShader()
	fragmentShader := createFragmentShader()

	shaderProgram := createShaderProgram(vertexShader, fragmentShader)

	gl.VertexAttribPointerWithOffset(0, 3, gl.FLOAT, false, 3*4, 0)
	gl.EnableVertexAttribArray(0)

	// You can unbind the VAO afterwards so other VAO calls won't accidentally modify this VAO, but this rarely happens. Modifying other
	// VAOs requires a call to glBindVertexArray anyways so we generally don't unbind VAOs (nor VBOs) when it's not directly necessary.
	// gl.BindVertexArray(0)

	for !window.ShouldClose() {
		gl.ClearColor(0.2, 0.3, 0.3, 1.0)
		gl.Clear(gl.COLOR_BUFFER_BIT)
		// draw our first triangle
		gl.UseProgram(shaderProgram)
		gl.BindVertexArray(VAO)
		gl.DrawArrays(gl.TRIANGLES, 0, 3)

		// Maintenance
		window.SwapBuffers()
		glfw.PollEvents()
	}

	gl.DeleteVertexArrays(1, &VAO)
	gl.DeleteBuffers(1, &VBO)
	gl.DeleteProgram(shaderProgram)
}
