package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

// Las tres opciones posibles del juego
var opciones = []string{"piedra", "papel", "tijera"}

// elegirComputador elige una opción al azar para la máquina
func elegirComputador() string {
	rand.Seed(time.Now().UnixNano())
	return opciones[rand.Intn(3)]
}

// decidirGanador compara las dos elecciones y devuelve el resultado
func decidirGanador(jugador, computador string) string {
	if jugador == computador {
		return "¡Empate! 🤝"
	}

	// Mapa de reglas: la clave le gana al valor
	gana := map[string]string{
		"piedra": "tijera",
		"tijera": "papel",
		"papel":  "piedra",
	}

	if gana[jugador] == computador {
		return "¡Ganaste! 🏆"
	}
	return "¡Perdiste! 💀"
}

// manejarInicio sirve el archivo HTML al usuario
func manejarInicio(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/index.html")
}

// manejarJugar procesa la elección del jugador y devuelve el resultado
func manejarJugar(w http.ResponseWriter, r *http.Request) {
	jugador := r.FormValue("jugador")

	// Validar que la opción sea válida
	if jugador != "piedra" && jugador != "papel" && jugador != "tijera" {
		fmt.Fprintf(w, "Opción inválida")
		return
	}

	computador := elegirComputador()
	resultado := decidirGanador(jugador, computador)

	emojis := map[string]string{
		"piedra": "🪨",
		"papel":  "📄",
		"tijera": "✂️",
	}

	// Genera la página HTML de resultado
	fmt.Fprintf(w, `<!DOCTYPE html>
<html lang="es">
<head>
  <meta charset="UTF-8">
  <title>Resultado</title>
  <style>
    * { box-sizing: border-box; margin: 0; padding: 0; }
    body {
      font-family: sans-serif;
      background: #1a1a2e;
      color: #eee;
      display: flex;
      flex-direction: column;
      align-items: center;
      justify-content: center;
      min-height: 100vh;
    }
    .tarjeta {
      background: #16213e;
      border: 2px solid #0f3460;
      border-radius: 20px;
      padding: 2rem 3rem;
      text-align: center;
      display: flex;
      flex-direction: column;
      gap: 1rem;
    }
    .versus { display: flex; align-items: center; justify-content: center; gap: 2rem; font-size: 4rem; }
    .versus span { font-size: 1rem; color: #aaa; display: block; margin-top: 4px; }
    .vs { font-size: 1.5rem; color: #eee; font-weight: bold; }
    .resultado { font-size: 1.5rem; font-weight: bold; color: #e94560; }
    .boton {
      margin-top: 0.5rem;
      padding: 0.6rem 2rem;
      font-size: 1rem;
      background: #e94560;
      border: none;
      border-radius: 10px;
      color: white;
      cursor: pointer;
      text-decoration: none;
      display: inline-block;
    }
    .boton:hover { background: #c73652; }
  </style>
</head>
<body>
  <div class="tarjeta">
    <div class="versus">
      <div>%s<span>Tú</span></div>
      <div class="vs">vs</div>
      <div>%s<span>Computador</span></div>
    </div>
    <div class="resultado">%s</div>
    <a class="boton" href="/">Jugar de nuevo</a>
  </div>
</body>
</html>`, emojis[jugador], emojis[computador], resultado)
}

func main() {
	http.HandleFunc("/", manejarInicio)
	http.HandleFunc("/jugar", manejarJugar)

	fmt.Println("Servidor corriendo en http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}