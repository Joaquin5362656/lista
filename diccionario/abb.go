package diccionario

import (
	TDAPila "tdas/pila"
)

type nodoAbb[K comparable, V any] struct {
	izquierdo *nodoAbb[K, V]
	derecho   *nodoAbb[K, V]
	nodoRaiz  parClaveValor[K, V]
}

type abb[K comparable, V any] struct {
	raiz     *nodoAbb[K, V]
	cantidad int
	funcCmp  func(K, K) int
}

func CrearABB[K comparable, V any](funcion_cmp func(K, K) int) DiccionarioOrdenado[K, V] {
	return &abb[K, V]{raiz: nil, cantidad: 0, funcCmp: funcion_cmp}
}

func crearNodoAbb[K comparable, V any](nuevoElemento parClaveValor[K, V]) *nodoAbb[K, V] {
	return &nodoAbb[K, V]{izquierdo: nil, derecho: nil, nodoRaiz: nuevoElemento}
}

func (arbol *abb[K, V]) Pertenece(clave K) bool {

	if arbol.raiz == nil {
		return false
	} else {
		ramaEncontrada := arbol.raiz.buscarRama(clave, arbol.funcCmp)
		return *ramaEncontrada == nil
	}

}

func (arbol *abb[K, V]) Guardar(clave K, dato V) {

	if arbol.raiz == nil {
		arbol.raiz = crearNodoAbb(parClaveValor[K, V]{clave, dato})
		arbol.cantidad++
	} else {
		ramaEncontrada := arbol.raiz.buscarRama(clave, arbol.funcCmp)

		if *ramaEncontrada == nil {
			*ramaEncontrada = crearNodoAbb(parClaveValor[K, V]{clave, dato})
			arbol.cantidad++
		} else {
			nodoAModificar := *ramaEncontrada
			nodoAModificar.nodoRaiz.dato = dato
		}

	}
}

func (arbol *abb[K, V]) Obtener(clave K) V {

	ramaEncontrada := arbol.raiz.buscarRama(clave, arbol.funcCmp)

	if *ramaEncontrada == nil {
		panic("La clave no pertenece al diccionario")
	} else {
		nodoEncontrado := *ramaEncontrada
		return nodoEncontrado.nodoRaiz.dato
	}
}

func (arbol *abb[K, V]) Borrar(clave K) V {

	ramaABorrar := arbol.raiz.buscarRama(clave, arbol.funcCmp)

	if *ramaABorrar == nil {
		panic("La clave no pertenece al diccionario")
	}

	datoBorrado := (*ramaABorrar).nodoRaiz.dato

	if (*ramaABorrar).derecho != nil && (*ramaABorrar).izquierdo != nil {

		nodoAModificar := (*ramaABorrar)
		ramaDerecha := &(*ramaABorrar).derecho
		ramaABorrar = buscarRamaSucesorInmediato((*ramaABorrar).derecho)
		if *ramaABorrar == *ramaDerecha {
			ramaABorrar = ramaDerecha
		}
		nodoAModificar.nodoRaiz.clave = (*ramaABorrar).nodoRaiz.clave
		nodoAModificar.nodoRaiz.dato = (*ramaABorrar).nodoRaiz.dato

	}

	(*ramaABorrar) = hallarHijoNoNulo((*ramaABorrar).izquierdo, (*ramaABorrar).derecho)

	arbol.cantidad--

	return datoBorrado
}

func (arbol *abb[K, V]) Cantidad() int {
	return arbol.cantidad
}

func (arbol *abb[K, V]) Iterar(funcion func(clave K, dato V) bool) {

	arbol.raiz.iterarRec(funcion)
}

func (nodoArbol *nodoAbb[K, V]) iterarRec(funcion func(clave K, dato V) bool) bool {

	if nodoArbol == nil {
		return true
	}

	seguirIterando := nodoArbol.izquierdo.iterarRec(funcion)

	if seguirIterando {
		seguirIterando = funcion(nodoArbol.nodoRaiz.clave, nodoArbol.nodoRaiz.dato)
	}

	if seguirIterando {
		return nodoArbol.derecho.iterarRec(funcion)
	} else {
		return false
	}
}

func buscarRamaSucesorInmediato[K comparable, V any](raiz *nodoAbb[K, V]) **nodoAbb[K, V] {

	if raiz.izquierdo == nil {
		return &raiz
	}

	sucesorInmediato := buscarRamaSucesorInmediato(raiz.izquierdo)

	if *sucesorInmediato == raiz.izquierdo {
		return &raiz.izquierdo
	}

	return sucesorInmediato
}

func hallarHijoNoNulo[K comparable, V any](izquierdo *nodoAbb[K, V], derecho *nodoAbb[K, V]) *nodoAbb[K, V] {

	if izquierdo == nil {
		return derecho
	} else {
		return izquierdo
	}
}

func (rama *nodoAbb[K, V]) buscarRama(clave K, comparar func(K, K) int) **nodoAbb[K, V] {

	if rama == nil {
		return nil
	}

	if rama.nodoRaiz.clave == clave {
		return &rama
	}

	var ramaEncontrada **nodoAbb[K, V]

	if comparar(clave, rama.nodoRaiz.clave) < 0 {
		ramaEncontrada = rama.izquierdo.buscarRama(clave, comparar)
		if ramaEncontrada == nil || *ramaEncontrada == rama.izquierdo {
			ramaEncontrada = &rama.izquierdo
		}
	} else {
		ramaEncontrada = rama.derecho.buscarRama(clave, comparar)
		if ramaEncontrada == nil || *ramaEncontrada == rama.derecho {
			ramaEncontrada = &rama.derecho
		}
	}

	return ramaEncontrada
}

type iteradorAbb[K comparable, V any] struct {
	nodosEnOrden TDAPila.Pila[*nodoAbb[K, V]]
}

func (arbol *abb[K, V]) Iterador() IterDiccionario[K, V] {
	return crearIteradorAbb(arbol.raiz)
}

func crearIteradorAbb[K comparable, V any](raiz *nodoAbb[K, V]) *iteradorAbb[K, V] {

	nodosAIterar := TDAPila.CrearPilaDinamica[*nodoAbb[K, V]]()
	apilarNodosMenores(nodosAIterar, raiz)
	return &iteradorAbb[K, V]{nodosEnOrden: nodosAIterar}
}

func (iterAbb *iteradorAbb[K, V]) HaySiguiente() bool {
	return !iterAbb.nodosEnOrden.EstaVacia()
}

func (iterAbb *iteradorAbb[K, V]) Siguiente() {

	nodoIterado := iterAbb.nodosEnOrden.Desapilar()
	apilarNodosMenores(iterAbb.nodosEnOrden, nodoIterado.derecho)
}

func (iterAbb *iteradorAbb[K, V]) VerActual() (K, V) {

	actual := iterAbb.nodosEnOrden.VerTope()
	return actual.nodoRaiz.clave, actual.nodoRaiz.dato
}

func apilarNodosMenores[K comparable, V any](pila TDAPila.Pila[*nodoAbb[K, V]], raiz *nodoAbb[K, V]) {

	for raiz != nil {
		pila.Apilar(raiz)
		raiz = raiz.izquierdo
	}

}
