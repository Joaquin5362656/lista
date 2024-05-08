package diccionario

import (
	"fmt"
	TDALista "tdas/lista"
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

func convertirABytes[K comparable](clave K) []byte {
	return []byte(fmt.Sprintf("%v", clave))
}

func (hash hashAbierto[K, V]) buscarConClave(clave K) (iterLista TDALista.IteradorLista[parClaveValor[K, V]]) {

	listaABuscar := hash.tabla[HashBernstein(convertirABytes(clave))%uint32(hash.tam)]

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

// nombre de la funcion hash = HAsh Bernstein
func HashBernstein(cadena []byte) uint32 {
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
	tam := 17
	return &hashAbierto[K, V]{tabla: crearTablaHash[K, V](tam), tam: tam}
}

func (h *hashAbierto[K, V]) Pertenece(clave K) bool {

	elementoEnLista := h.buscarConClave(clave)

	if elementoEnLista.HaySiguiente() {
		return true
	} else {
		return false
	}
}

func (h *hashAbierto[K, V]) Guardar(clave K, dato V) {

	elementoEnLista := h.buscarConClave(clave)

	if elementoEnLista.HaySiguiente() {
		elementoEnLista.Borrar()
	} else {
		h.cantidad++
	}

	elementoEnLista.Insertar(parClaveValor[K, V]{clave: clave, dato: dato})

}

func (h *hashAbierto[K, V]) Obtener(clave K) V {

	elementoEnLista := h.buscarConClave(clave)

	if elementoEnLista.HaySiguiente() {
		return elementoEnLista.VerActual().dato
	} else {
		panic("")
	}

}

func (h *hashAbierto[K, V]) Borrar(clave K) V {

	elementoEnLista := h.buscarConClave(clave)

	if elementoEnLista.HaySiguiente() {
		datoEliminado := elementoEnLista.VerActual().dato
		elementoEnLista.Borrar()
		h.cantidad--
		return datoEliminado
	} else {
		panic("")
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

type iteradorHash[K comparable, V any] struct {
	tablaDeHash   []TDALista.Lista[parClaveValor[K, V]]
	iteradorLista TDALista.IteradorLista[parClaveValor[K, V]]
}

func crearIterador[K comparable, V any](tablaDeHash []TDALista.Lista[parClaveValor[K, V]]) *iteradorHash[K, V] {
	nuevoIterador := iteradorHash[K, V]{tablaDeHash: tablaDeHash, iteradorLista: nil}
	nuevoIterador.iteradorLista = nuevoIterador.encontrarSiguiente()
	return &nuevoIterador
}

func (iterHash *iteradorHash[K, V]) encontrarSiguiente() TDALista.IteradorLista[parClaveValor[K, V]] {

	listaActual := iterHash.tablaDeHash[0]
	var i int = 1

	for i < len(iterHash.tablaDeHash) && listaActual.EstaVacia() {
		listaActual = iterHash.tablaDeHash[i]
		i++
	}

	if i >= len(iterHash.tablaDeHash) {
		return nil
	} else {
		iterHash.tablaDeHash = iterHash.tablaDeHash[i:]
		return listaActual.Iterador()
	}

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

	if iterHash.iteradorLista == nil {
		panic("")
	}

	iterHash.iteradorLista.Siguiente()

	if !iterHash.iteradorLista.HaySiguiente() {
		iterHash.iteradorLista = iterHash.encontrarSiguiente()
	}

}

func (iterHash *iteradorHash[K, V]) VerActual() (K, V) {

	if iterHash.iteradorLista.HaySiguiente() {
		return iterHash.iteradorLista.VerActual().clave, iterHash.iteradorLista.VerActual().dato
	} else {
		panic("")
	}
}
