package diccionario

import (
	"reflect"
	"strings"
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

	ramaABuscar := arbol.raiz.buscarRama(clave, arbol.funcCmp)

	if *ramaABuscar == nil {
		return false
	} else {
		return true
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

	if ramaEncontrada == nil || *ramaEncontrada == nil {
		panic("La clave no pertenece al diccionario")
	}
	nodoEncontrado := *ramaEncontrada
	return nodoEncontrado.nodoRaiz.dato
}

func (arbol *abb[K, V]) Borrar(clave K) V {
	ramaABorrar := arbol.raiz.buscarRama(clave, arbol.funcCmp)

	if ramaABorrar == nil || *ramaABorrar == nil {
		panic("La clave no pertenece al diccionario")
	}

	if *ramaABorrar == arbol.raiz {
		ramaABorrar = &arbol.raiz
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

func (arbol *abb[K, V]) IterarRango(desde *K, hasta *K, visitar func(clave K, dato V) bool) {

	arbol.raiz.iterarRangoRec(desde, hasta, visitar, arbol.funcCmp)
}

func (nodoArbol *nodoAbb[K, V]) iterarRangoRec(desde *K, hasta *K, visitar func(clave K, dato V) bool, comparar func(K, K) int) (seguirIterando bool) {

	if nodoArbol == nil {
		return true
	}

	mayorRangoInferior := false
	seguirIterando = true

	if desde == nil || comparar(*desde, nodoArbol.nodoRaiz.clave) <= 0 {
		seguirIterando = nodoArbol.izquierdo.iterarRangoRec(desde, hasta, visitar, comparar)
		mayorRangoInferior = true
	}
	if hasta == nil || comparar(*hasta, nodoArbol.nodoRaiz.clave) >= 0 {
		if mayorRangoInferior {
			seguirIterando = visitar(nodoArbol.nodoRaiz.clave, nodoArbol.nodoRaiz.dato)
		}

		if seguirIterando {
			seguirIterando = nodoArbol.derecho.iterarRangoRec(desde, hasta, visitar, comparar)
		}
	}

	return seguirIterando

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
	}
	return izquierdo
}

func (rama *nodoAbb[K, V]) buscarRama(clave K, comparar func(K, K) int) **nodoAbb[K, V] {

	if rama == nil {
		return &rama
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
	desde        *K
	hasta        *K
	comparar     func(K, K) int
}

func (arbol *abb[K, V]) Iterador() IterDiccionario[K, V] {
	return crearIteradorAbb(arbol.raiz, nil, nil, arbol.funcCmp)
}

func (arbol *abb[K, V]) IteradorRango(desde *K, hasta *K) IterDiccionario[K, V] {
	return crearIteradorAbb(arbol.raiz, desde, hasta, arbol.funcCmp)
}

func crearIteradorAbb[K comparable, V any](raiz *nodoAbb[K, V], desde *K, hasta *K, comparar func(K, K) int) *iteradorAbb[K, V] {

	nodosAIterar := TDAPila.CrearPilaDinamica[*nodoAbb[K, V]]()
	iterAbb := iteradorAbb[K, V]{nodosAIterar, desde, hasta, comparar}
	iterAbb.apilarNodosMenores(raiz)
	return &iterAbb
}

func (iterAbb *iteradorAbb[K, V]) HaySiguiente() bool {
	return !iterAbb.nodosEnOrden.EstaVacia()
}

func (iterAbb *iteradorAbb[K, V]) Siguiente() {

	if !iterAbb.HaySiguiente() {
		panic("El iterador termino de iterar")
	}

	nodoIterado := iterAbb.nodosEnOrden.Desapilar()
	iterAbb.apilarNodosMenores(nodoIterado.derecho)

}

func (iterAbb *iteradorAbb[K, V]) VerActual() (K, V) {

	if !iterAbb.HaySiguiente() {
		panic("El iterador termino de iterar")
	}

	actual := iterAbb.nodosEnOrden.VerTope()
	return actual.nodoRaiz.clave, actual.nodoRaiz.dato
}

func (iterAbb *iteradorAbb[K, V]) apilarNodosMenores(raiz *nodoAbb[K, V]) {
	if raiz == nil {
		return
	}

	if (iterAbb.desde == nil || iterAbb.comparar(*iterAbb.desde, raiz.nodoRaiz.clave) <= 0) &&
		(iterAbb.hasta == nil || iterAbb.comparar(*iterAbb.hasta, raiz.nodoRaiz.clave) >= 0) {
		iterAbb.nodosEnOrden.Apilar(raiz)
	}

	iterAbb.apilarNodosMenores(raiz.izquierdo)
}

func Funcion_cmp[K comparable](clave1, clave2 K) int {
	tipoClave := reflect.TypeOf(clave1)

	switch tipoClave.Kind() {
	case reflect.String:
		c1 := reflect.ValueOf(clave1).String()
		c2 := reflect.ValueOf(clave2).String()
		return strings.Compare(c1, c2)
	case reflect.Int:
		c1 := reflect.ValueOf(clave1).Int()
		c2 := reflect.ValueOf(clave2).Int()
		if c1 < c2 {
			return -1
		} else if c1 > c2 {
			return 1
		}
		return 0
	}
	return 0
}
