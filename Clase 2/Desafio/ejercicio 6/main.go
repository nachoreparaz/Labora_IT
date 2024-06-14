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
		Nombre:  "Goblin",
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
	fmt.Printf("\nEl %s tiene Salud: %d, Ataque %d y Defensa %d\n", monstruo.Nombre, monstruo.Salud, monstruo.Ataque, monstruo.Defensa)
	fmt.Printf("\nTu heroe %s tiene Salud: %d, Ataque %d y Defensa %d\n", jugador.Nombre, jugador.Salud, jugador.Ataque, jugador.Defensa)
	fmt.Println("\nQuieres combatir al monstruo?\n 1: Si\n 2: No")
	var choice int
	fmt.Scanln(&choice)
	switch choice {
	case 1:
		comenzarBatalla(jugador, lugar, monstruos_map, monstruo)
	case 2:
		return
	default:
		break
	}
}

func comenzarBatalla(jugador *Jugador, lugar Lugar, monstruos_map map[Lugar]Monstruo, monstruo Monstruo) {
	turno := 0
	for jugador.Salud > 0 && monstruo.Salud > 0 {
		if turno%2 == 0 {
			fmt.Println("\nTurno del Jugador:")
			if atacar(jugador.Ataque, &monstruo.Defensa, &monstruo.Salud) {
				fmt.Println("\nEl Jugador ha derrotado al", monstruo.Nombre)
				statusJugador(jugador)
				delete(monstruos_map, lugar)
				fmt.Printf("\nDerrotaste a los monstruos del %s!\n", lugar)
				if jugador.Salud < 100 && jugador.Defensa < 50 {
					elegirRecompensa(jugador)
				}
				if jugador.Defensa < 20 {
					elegirObjeto(jugador, lugar)
				}
				time.Sleep(3 * time.Second)
				break
			}
		} else {
			fmt.Println("\nTurno del Monstruo:")
			if atacar(monstruo.Ataque, &jugador.Defensa, &jugador.Salud) {
				fmt.Println("\nEl", monstruo.Nombre, "ha derrotado al Jugador!")
				monstruos_map = make(map[Lugar]Monstruo)
				break
			}
		}
		turno++
		time.Sleep(3 * time.Second)
	}
}

func atacar(ataque int, defensa *int, salud *int) bool {
	hit := *defensa - ataque
	if hit > 0 {
		*defensa = hit
	} else {
		*defensa = 0
		*salud += hit
	}
	fmt.Printf("Inflige %d de daño de ataque. Salud restante: %d\n", ataque, *salud)
	return *salud <= 0
}

func statusJugador(jugador *Jugador) {
	fmt.Printf("\nTu heroe tiene %d de vida, %d de ataque y %d de defensa\n", jugador.Salud, jugador.Ataque, jugador.Defensa)
	time.Sleep(3 * time.Second)
}

func mostrarLugaresDisponibles(lugares []Lugar, monstruos_map map[Lugar]Monstruo) {
	for i, el := range lugares {
		if _, ok := monstruos_map[el]; ok {
			fmt.Printf("\n%d: %s\n", i+1, el)
		}
	}
}

func elegirRecompensa(jugador *Jugador) {
	println("\n Has dado una gran batalla y como agradecimiento le ofrecemos las siguientes recompensas. Elija sabiamente:\n")
	fmt.Println(" 1. Pocion Curativa: +100 salud y +50 defensa\n 2. Armadura de Metal: +150 defensa y +50 ataque")
	var choice int
	fmt.Scanln(&choice)
	switch choice {
	case 1:
		jugador.Salud += 100
		jugador.Defensa += 50
		fmt.Println("Has recibido una pocion curativa")
	case 2:
		jugador.Defensa += 150
		jugador.Ataque += 50
		fmt.Println("Has recibido una armadura de metal")
	default:
		fmt.Println("No has elegido ninguna recompensa")
	}
	time.Sleep(3 * time.Second)
	statusJugador(jugador)
}

func elegirObjeto(jugador *Jugador, lugar Lugar) {
	fmt.Printf("\nHas encontrado un tesoro en %s y contiene algunos objetos dentro, cual te quieres llevar?\n", lugar)
	fmt.Println("\n1. Bolsa de sangre: +40 salud\n")
	fmt.Println("\n2. Escudo: +90 defensa\n")
	fmt.Println("\n3. Espada: +65 ataque\n")
	var choice int
	fmt.Scanln(&choice)
	switch choice {
	case 1:
		jugador.Salud += 40
	case 2:
		jugador.Defensa += 90
	case 3:
		jugador.Ataque += 65
	default:
		break
	}
	statusJugador(jugador)
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
		if len(monstruos_map) == 0 {
			fmt.Println("\nFelicitaciones, has ganado el juego! Espero que te hayas divertido! Si quieres seguir jugando, corre nuevamente el juego de LABORA IT :)")
			break
		}
		if jugador.Salud < 1 {
			fmt.Println("Casi lo consigues! Sigue intentando")
			break
		}
		mostrarLugaresDisponibles(lugares, monstruos_map)
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
}

// Al terminar subir nivel de mostruo x2 y al heroe x0.5
// Hacer las estadisticas de los monstuos de manera random
