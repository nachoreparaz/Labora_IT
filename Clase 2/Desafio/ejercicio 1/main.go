package main

import "fmt"

// ================ EJERCICIO 1 ===================
// Una pequeña tienda necesita un programa para gestionar su inventario de productos.
// El programa debe permitir la adición de nuevos productos, la actualización de la cantidad disponible de un producto,
// la eliminación de productos y la visualización del inventario completo.

// Requisitos:

// Implementar una estructura de datos para almacenar la información de cada producto :(nombre, precio y cantidad disponible).
// Permitir la adición de nuevos productos con sus respectivas cantidades.
// Permitir la actualización de la cantidad disponible de un producto :existente.
// Permitir la eliminación de productos del inventario.
// Mostrar el inventario completo, incluyendo el nombre, precio y cantidad disponible de cada producto.

type Inventario struct {
	Nombre   string
	Precio   int
	Cantidad int
}

func agregarProducto(inventario map[string]Inventario, nombre string, precio int, cantidad int) {
	inventario[nombre] = Inventario{
		Nombre:   nombre,
		Precio:   precio,
		Cantidad: cantidad,
	}
	fmt.Printf("Producto :'%s' agregado al inventario\n", nombre)
}

func actualizarCantidad(inventario map[string]Inventario, nombre string, cantidad int) {
	if prod, encontrado := inventario[nombre]; encontrado {
		prod.Cantidad = cantidad
		inventario[nombre] = prod
		fmt.Print("Producto :actualizado:\n", nombre, prod)
	} else {
		fmt.Println("Producto :no encontrado")
	}
}

func eliminarProducto(inventario map[string]Inventario, nombre string) {
	if _, encontrado := inventario[nombre]; encontrado {
		delete(inventario, nombre)
		fmt.Printf("Producto :'%s' eliminado", nombre)
	} else {
		fmt.Println("Producto :no encontrado")
	}
}

func mostrarInventario(inventario map[string]Inventario) {
	for _, producto := range inventario {
		fmt.Println("\nProducto: ", producto.Nombre, producto.Precio, producto.Cantidad)
	}
}

func buscarProducto(inventario map[string]Inventario, nombre string) {
	if _, encontrado := inventario[nombre]; encontrado {
		fmt.Printf("Producto :'%s' encontrado", nombre)
	} else {
		fmt.Println("Producto :no encontrado")
	}
}

/*
	Simular compra de prod
	Carrito => array
	Calcular suma de precio de los prods del carrito
	Dinero en billetera y restarle la suma de los prods del carrito
*/

func agregarProductoLista(productos *[]string, saldo_a_pagar *[]int) {
	for {
		var choice int
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			producto := "Producto1"
			*productos = append(*productos, producto)
			agregarPrecios(saldo_a_pagar, choice)
		case 2:
			producto := "Producto2"
			*productos = append(*productos, producto)
			agregarPrecios(saldo_a_pagar, choice)
		case 3:
			producto := "Producto3"
			*productos = append(*productos, producto)
			agregarPrecios(saldo_a_pagar, choice)
		case 0:
			return
		default:
			break
		}
	}
}

func agregarPrecios(saldo_a_pagar *[]int, choice int) {
	switch choice {
	case 1:
		precio := 100
		*saldo_a_pagar = append(*saldo_a_pagar, precio)
		fmt.Print("Saldo en agregar Precio Case 1: ", saldo_a_pagar)
	case 2:
		precio := 200
		*saldo_a_pagar = append(*saldo_a_pagar, precio)
		fmt.Print("Saldo en agregar Precio Case 2: ", saldo_a_pagar)
	case 3:
		precio := 300
		*saldo_a_pagar = append(*saldo_a_pagar, precio)
		fmt.Print("Saldo en agregar Precio Case 3: ", saldo_a_pagar)
	default:
		break
	}
}

func comprar(saldo_a_pagar []int) {
	billetera := 2000
	total_a_pagar := 0
	valor := 0
	for i := 0; i < len(saldo_a_pagar); i++ {
		valor = saldo_a_pagar[i]
		total_a_pagar += valor
	}
	fmt.Print("Total a pagar: ", total_a_pagar)
	if billetera >= total_a_pagar {
		billetera -= total_a_pagar
		fmt.Print(saldo_a_pagar)
		fmt.Printf("\nCompra Exitosa, tu saldo restante es: '%d'", billetera)
	} else {
		fmt.Println("Saldo insuficiente")
	}

}

func main() {
	inventario := make(map[string]Inventario)
	// fmt.Print("\nInventario Al Crearse\n")
	agregarProducto(inventario, "Producto1", 100, 10)
	agregarProducto(inventario, "Producto2", 200, 20)
	agregarProducto(inventario, "Producto3", 300, 30)

	var saldo_a_pagar []int
	var productos []string

	agregarProductoLista(&productos, &saldo_a_pagar)
	comprar(saldo_a_pagar)
}
