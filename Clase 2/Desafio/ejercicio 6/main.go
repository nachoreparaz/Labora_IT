package main

import (
	"fmt"
	"time"
)

type Jugador struct {
	Nombre  string
	Nivel   int
	Salud   int
	Ataque  int
	Defensa int
	Lugar   Lugar
}

type Monstruo struct {
	Nombre  string
	Salud   int
	Ataque  int
	Defensa int
	Lugar   Lugar
}
type Lugar string

const (
	PANTANO  Lugar = "Pantano"
	BOSQUE   Lugar = "Bosque"
	CASTILLO Lugar = "Castillo"
	CASA     Lugar = "Casa"
)

type Objeto struct {
	Nombre string
	Tipo   string
}

func crearJugador(nombre string) *Jugador {
	jugador := &Jugador{
		Nombre:  nombre,
		Nivel:   1,
		Salud:   100,
		Ataque:  35,
		Defensa: 20,
		Lugar:   CASA,
	}
	return jugador
}

func actualizarLugarJugador(jugador *Jugador, lugar Lugar) {
	jugador.Lugar = lugar
}

func crearMonstruo(monstruos_map map[Lugar]Monstruo, lugares *[]Lugar) {
	ogro := Monstruo{
		Nombre:  "Ogro",
		Salud:   100,
		Ataque:  70,
		Defensa: 20,
		Lugar:   PANTANO,
	}
	monstruos_map[PANTANO] = ogro
	*lugares = append(*lugares, PANTANO)

	goblin := Monstruo{
		Nombre:  "Ogro",
		Salud:   40,
		Ataque:  20,
		Defensa: 70,
		Lugar:   BOSQUE,
	}
	monstruos_map[BOSQUE] = goblin
	*lugares = append(*lugares, BOSQUE)

	fantasma := Monstruo{
		Nombre:  "Fantasma",
		Salud:   100,
		Ataque:  50,
		Defensa: 5,
		Lugar:   CASTILLO,
	}
	monstruos_map[CASTILLO] = fantasma
	*lugares = append(*lugares, CASTILLO)
}

func batalla(jugador *Jugador, lugar Lugar, monstruos_map map[Lugar]Monstruo) {
	monstruo, existe := monstruos_map[lugar]
	if !existe {
		fmt.Println("No hay ningún monstruo en este lugar.")
		return
	}

	fmt.Printf("\nHas encontrado un %s en %s\n", monstruo.Nombre, lugar)

	turno := 0
	for jugador.Salud > 0 && monstruo.Salud > 0 {
		if turno%2 == 0 {
			fmt.Println("\nTurno del Jugador:")
			if atacar(jugador.Ataque, monstruo.Defensa, &monstruo.Salud) {
				fmt.Println("\nEl Jugador ha derrotado al", monstruo.Nombre)
				statusJugador(jugador)
				monstruos_map[lugar] = monstruo
				delete(monstruos_map, lugar)
				fmt.Printf("\nDerrotaste a los monstruos del %s! Si quieres seguir luchando ve a los lugares disponibles!\n", lugar)
				time.Sleep(3 * time.Second)
				break
			}
		} else {
			fmt.Println("\nTurno del Monstruo:")
			if atacar(monstruo.Ataque, jugador.Defensa, &jugador.Salud) {
				fmt.Println("\nEl", monstruo.Nombre, "ha derrotado al Jugador!")
				monstruos_map = make(map[Lugar]Monstruo)
				break
			}
		}
		turno++
		time.Sleep(3 * time.Second)
	}
}

func atacar(ataque, defensa int, salud *int) bool {
	daño := max(0, ataque-defensa)
	*salud -= daño
	fmt.Printf("Inflige %d de daño. Salud restante: %d\n", daño, *salud)
	return *salud <= 0
}

func statusJugador(jugador *Jugador) {
	fmt.Printf("Luego del Monstruo tu Personaje quedo con %d de vida\n", jugador.Salud)
}

func mostrarLugaresDisponibles(lugares []Lugar, monstruos_map map[Lugar]Monstruo) {
	for i, el := range lugares {
		if _, ok := monstruos_map[el]; ok {
			fmt.Printf("\n%d: %s\n", i+1, el)
		}
	}
}

func main() {
	// var objeto Objeto
	monstruos_map := make(map[Lugar]Monstruo)
	var lugares []Lugar
	crearMonstruo(monstruos_map, &lugares)
	fmt.Println("LUGARES: ", lugares)
	fmt.Println("Bienvenido al juego de Rol de Labora IT\n Ahora debes empezar nombrando a tu Heroe!\n")
	var nombrePersonaje string
	fmt.Scanln(&nombrePersonaje)

	jugador := crearJugador(nombrePersonaje)
	fmt.Println("\n Comenzando La aventura, donde quieres ir?")
	for {
		mostrarLugaresDisponibles(lugares, monstruos_map)
		if len(monstruos_map) == 0 {
			fmt.Println("Espero que te hayas divertido! Si quieres seguir jugando, corre nuevamente el juego de LABORA IT :)")
			break
		}
		var choice int
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			actualizarLugarJugador(jugador, PANTANO)
			batalla(jugador, PANTANO, monstruos_map)
			continue
		case 2:
			actualizarLugarJugador(jugador, BOSQUE)
			batalla(jugador, BOSQUE, monstruos_map)
			continue
		case 3:
			actualizarLugarJugador(jugador, CASTILLO)
			batalla(jugador, CASTILLO, monstruos_map)
			continue
		case 4:
			actualizarLugarJugador(jugador, CASA)
			fmt.Println("\n Has llegado a la casa, el juego ha terminado!")
			break
		default:
			break
		}
	}

	// TODO: Crear Monstruo en base al lugar en donde se encuentra el jugador. Crear objetos que dependiendo de una probabilidad, el jugador puede encontrar
	// monstruo.Nombre = "Goblin"
	// monstruo.Salud = 100
	// monstruo.Ataque = 10
	// monstruo.Defensa = 10
	// monstruo.Lugar = "Pantano"
	// monstruo.Tipo = "Monstruo"

}
