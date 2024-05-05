package diccionario

import (
	TDALista "tdas/lista"
)

type hashCampo[K comparable, V any] struct {
	clave K
	dato  V
}

type diccionario[K comparable, V any] struct {
	tablaDeHash []TDALista.Lista[hashCampo[K, V]]
	ocupados    int
}

type iteradorHash[K comparable, V any] struct {
	tablaDeHash   []TDALista.Lista[hashCampo[K, V]]
	iteradorLista TDALista.IteradorLista[hashCampo[K, V]]
}

func crearIterador[K comparable, V any](tablaDeHash []TDALista.Lista[hashCampo[K, V]]) *iteradorHash[K, V] {
	primeraListaNoVacia := encontrarSiguiente(tablaDeHash)
	return &iteradorHash[K, V]{tablaDeHash: tablaDeHash, iteradorLista: primeraListaNoVacia.Iterador()}
}

func encontrarSiguiente[K comparable, V any](tablaDeHash []TDALista.Lista[hashCampo[K, V]]) TDALista.Lista[hashCampo[K, V]] {

	listaActual := tablaDeHash[0]
	var i int = 1

	for i <= len(tablaDeHash) && listaActual.EstaVacia() {
		listaActual = tablaDeHash[i]
		i++
	}

	if i <= len(tablaDeHash) {
		return nil
	} else {
		tablaDeHash = tablaDeHash[i:]
		return listaActual
	}

}

func (diccionario *diccionario[K, V]) Iterador() IterDiccionario[K, V] {
	return crearIterador[K, V](diccionario.tablaDeHash)
}

func (iterHash iteradorHash[K, V]) HaySiguiente() bool {

	if !iterHash.iteradorLista.HaySiguiente() && encontrarSiguiente(iterHash.tablaDeHash) == nil {
		return false
	} else {
		return true
	}

}

func (iterHash iteradorHash[K, V]) Siguiente() {

	if iterHash.HaySiguiente() {
		iterHash.iteradorLista.Siguiente()
	} else {
		listaActual := encontrarSiguiente[K, V](iterHash.tablaDeHash)
		if listaActual == nil {
			panic("")
		}
		iterHash.iteradorLista = listaActual.Iterador()
	}
}

func (iterHash iteradorHash[K, V]) VerActual() (K, V) {

	if iterHash.iteradorLista.HaySiguiente() {
		return iterHash.iteradorLista.VerActual().clave, iterHash.iteradorLista.VerActual().dato
	} else {
		panic("")
	}
}
