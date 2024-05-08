package diccionario

import (
	"fmt"
	TDALista "tdas/lista"
)

type redimensionado int

const (
	factorDeCarga  float32        = 2.3
	tamanioInicial int            = 17
	ampliarTabla   redimensionado = iota
	reducirTabla
)

type parClaveValor[K comparable, V any] struct {
	clave K
	dato  V
}

type hashAbierto[K comparable, V any] struct {
	tabla    []TDALista.Lista[parClaveValor[K, V]]
	tam      int
	cantidad int
}

// nombre de la funcion hash = HAsh Bernstein
func hashBernstein(cadena []byte) uint32 {
	var hash uint32 = 5381
	for _, c := range cadena {
		hash = (hash << 5) + hash + uint32(c)
	}
	return hash
}

func crearTablaHash[K comparable, V any](tam int) []TDALista.Lista[parClaveValor[K, V]] {

	tabla := make([]TDALista.Lista[parClaveValor[K, V]], tam)

	for indice := range tabla {
		tabla[indice] = TDALista.CrearListaEnlazada[parClaveValor[K, V]]()
	}
	return tabla
}

func CrearHash[K comparable, V any]() Diccionario[K, V] {
	tam := tamanioInicial
	return &hashAbierto[K, V]{tabla: crearTablaHash[K, V](tam), tam: tam}
}

func (h *hashAbierto[K, V]) Pertenece(clave K) bool {

	elementoEnLista := buscarConClave(h.tabla, clave)

	if elementoEnLista.HaySiguiente() {
		return true
	} else {
		return false
	}
}

func (h *hashAbierto[K, V]) Guardar(clave K, dato V) {

	elementoEnLista := buscarConClave(h.tabla, clave)

	if elementoEnLista.HaySiguiente() {
		elementoEnLista.Borrar()
	} else {
		h.cantidad++

		if float32(h.cantidad)/float32(h.tam) >= factorDeCarga {
			h.redimensionar(ampliarTabla)
			elementoEnLista = buscarConClave(h.tabla, clave)
		}

	}

	elementoEnLista.Insertar(parClaveValor[K, V]{clave: clave, dato: dato})

}

func (h *hashAbierto[K, V]) Obtener(clave K) V {

	elementoEnLista := buscarConClave(h.tabla, clave)

	if elementoEnLista.HaySiguiente() {
		return elementoEnLista.VerActual().dato
	} else {
		panic("La clave no pertenece al diccionario")
	}

}

func (h *hashAbierto[K, V]) Borrar(clave K) V {

	elementoEnLista := buscarConClave(h.tabla, clave)

	if elementoEnLista.HaySiguiente() {
		datoEliminado := elementoEnLista.VerActual().dato
		elementoEnLista.Borrar()
		h.cantidad--

		if float32(h.cantidad)/float32(h.tam) <= factorDeCarga/4 {
			h.redimensionar(reducirTabla)
		}

		return datoEliminado
	} else {
		panic("La clave no pertenece al diccionario")
	}

}

func (h *hashAbierto[K, V]) Cantidad() int {
	return h.cantidad
}

func (h *hashAbierto[K, V]) Iterar(funcion func(clave K, dato V) bool) {

	var seguirIterando bool = true
	var posHash int = 0

	for seguirIterando && posHash < len(h.tabla) {

		listaActual := h.tabla[posHash]

		for iter := listaActual.Iterador(); seguirIterando && iter.HaySiguiente(); iter.Siguiente() {
			parClaveValor := iter.VerActual()
			seguirIterando = funcion(parClaveValor.clave, parClaveValor.dato)
		}

		posHash++
	}

}

func buscarConClave[K comparable, V any](tablaHash []TDALista.Lista[parClaveValor[K, V]], clave K) (iterLista TDALista.IteradorLista[parClaveValor[K, V]]) {

	listaABuscar := tablaHash[hashBernstein(convertirABytes(clave))%uint32(len(tablaHash))]

	var elementoEncontrado bool = false

	iterLista = listaABuscar.Iterador()
	for !elementoEncontrado && iterLista.HaySiguiente() {
		if clave == iterLista.VerActual().clave {
			elementoEncontrado = true
		} else {
			iterLista.Siguiente()
		}
	}

	return iterLista
}

func convertirABytes[K comparable](clave K) []byte {
	return []byte(fmt.Sprintf("%v", clave))
}

func (h *hashAbierto[K, V]) redimensionar(tipoDeRedimension redimensionado) {

	nuevoTamanio := hallarPrimoMasOptimo(h.tam, tipoDeRedimension)

	if nuevoTamanio < tamanioInicial {
		nuevoTamanio = tamanioInicial
	}

	nuevaTabla := crearTablaHash[K, V](nuevoTamanio)

	h.Iterar(func(clave K, dato V) bool {
		elementoEnLista := buscarConClave(nuevaTabla, clave)
		elementoEnLista.Insertar(parClaveValor[K, V]{clave, dato})
		return true
	})

	h.tam = nuevoTamanio
	h.tabla = nuevaTabla

}

//		Esta funcion busca el numero mas optimo para la redimension buscando el numero primo
//		mas cercano al doble del numero anterior, en caso de que se busque ampliar tabla
//	 o el mas cercano a la cuarta parte del numero anterior en caso de que se busque
//		achicar la tabla.
//		La funcion se basa en el hecho de que si un numero es primo, entonces no debe ser par
//		ni terminar en 5 (ya que seria divisible por este) por lo tanto tendria que terminar
//		en 1, 3, 7 o 9, verificamos si alguno de los numeros mas cercanos al doble terminados
//		en esos numeros es primo y si encuentra uno, lo retorna, en caso contrario, que ninguno
//		sea primo, entonces devuelve el mayor de estos (se puede dar el caso que ninguno de
//		estos sea un numero primo pero en la mayoria de los casos si se cumple que alguno de
//		estos lo sea)
func hallarPrimoMasOptimo(primoAnterior int, tipoDeRedimension redimensionado) int {

	floatAux := float32(primoAnterior) / 10
	enteroDivisibleX10 := 0

	if tipoDeRedimension == ampliarTabla {
		enteroDivisibleX10 = int(floatAux * 2)
	} else if tipoDeRedimension == reducirTabla {
		enteroDivisibleX10 = int(floatAux / 4)
	}

	enteroDivisibleX10 *= 10

	candidatoPrimo1 := enteroDivisibleX10 + 1
	candidatoPrimo2 := enteroDivisibleX10 + 3
	candidatoPrimo3 := enteroDivisibleX10 + 7
	candidatoPrimo4 := enteroDivisibleX10 + 9

	if esPrimo(candidatoPrimo1) {
		return candidatoPrimo1
	} else if esPrimo(candidatoPrimo2) {
		return candidatoPrimo2
	} else if esPrimo(candidatoPrimo3) {
		return candidatoPrimo3
	} else {
		return candidatoPrimo4
	}
}

func esPrimo(numero int) bool {

	for x := 3; x*x < numero; x += 2 {
		if numero%x == 0 {
			return false
		}
	}
	return true
}

type iteradorHash[K comparable, V any] struct {
	tablaDeHash   []TDALista.Lista[parClaveValor[K, V]]
	iteradorLista TDALista.IteradorLista[parClaveValor[K, V]]
}

func crearIterador[K comparable, V any](tablaDeHash []TDALista.Lista[parClaveValor[K, V]]) *iteradorHash[K, V] {
	nuevoIterador := iteradorHash[K, V]{tablaDeHash: tablaDeHash, iteradorLista: nil}
	nuevoIterador.iteradorLista = nuevoIterador.encontrarSiguienteLista()
	return &nuevoIterador
}

func (hash *hashAbierto[K, V]) Iterador() IterDiccionario[K, V] {
	return crearIterador[K, V](hash.tabla)
}

func (iterHash *iteradorHash[K, V]) HaySiguiente() bool {

	if iterHash.iteradorLista == nil {
		return false
	} else {
		return true
	}

}

func (iterHash *iteradorHash[K, V]) Siguiente() {

	if iterHash.HaySiguiente() {

		iterHash.iteradorLista.Siguiente()

		if !iterHash.iteradorLista.HaySiguiente() {
			iterHash.iteradorLista = iterHash.encontrarSiguienteLista()
		}

	} else {
		panic("El iterador termino de iterar")
	}

}

func (iterHash *iteradorHash[K, V]) VerActual() (K, V) {

	if iterHash.HaySiguiente() {
		return iterHash.iteradorLista.VerActual().clave, iterHash.iteradorLista.VerActual().dato
	} else {
		panic("El iterador termino de iterar")
	}
}

func (iterHash *iteradorHash[K, V]) encontrarSiguienteLista() TDALista.IteradorLista[parClaveValor[K, V]] {

	if len(iterHash.tablaDeHash) == 0 {
		return nil
	}

	listaActual := iterHash.tablaDeHash[0]
	var i int = 1

	for i < len(iterHash.tablaDeHash) && listaActual.EstaVacia() {
		listaActual = iterHash.tablaDeHash[i]
		i++
	}

	if listaActual.EstaVacia() {
		return nil
	} else {
		iterHash.tablaDeHash = iterHash.tablaDeHash[i:]
		return listaActual.Iterador()
	}

}
