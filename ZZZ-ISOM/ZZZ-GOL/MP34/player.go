package main

import (
	"github.com/gen2brain/raylib-go/raylib"
	"math"
)

func main() {
	// Configurações da janela
	rl.InitWindow(800, 600, "Cubo 3D Wireframe Animado")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	// Câmera 3D
	camera := rl.Camera3D{}
	camera.Position = rl.NewVector3(10.0, 10.0, 10.0) // Posição da câmera
	camera.Target = rl.NewVector3(0.0, 0.0, 0.0)     // Ponto para onde a câmera olha
	camera.Up = rl.NewVector3(0.0, 1.0, 0.0)         // Vetor "para cima" da câmera
	camera.Fovy = 45.0                                // Campo de visão
	camera.Projection = rl.CameraPerspective          // Projeção perspectiva

	// Ângulo de rotação
	rotationAngle := float32(0.0)

	// Loop principal
	for !rl.WindowShouldClose() {
		// Atualiza o ângulo de rotação
		rotationAngle += 1.0

		// Inicia o modo 3D
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		rl.BeginMode3D(camera)

		// Desenha o cubo em wireframe
		rl.DrawCubeWires(rl.NewVector3(0.0, 0.0, 0.0), 2.0, 2.0, 2.0, rl.Black)

		// Aplica rotação ao cubo
		rl.Rotatef(rotationAngle, 1.0, 1.0, 1.0)

		rl.EndMode3D()

		// Desenha texto na tela
		rl.DrawText("Cubo 3D Wireframe Animado", 10, 10, 20, rl.DarkGray)

		rl.EndDrawing()
	}
}