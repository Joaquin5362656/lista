package cola_prioridad

import (
	"reflect"
	"strings"
)

type colaConPrioridad[T any] struct {
	datos []T
	cant  int
	cmp   func(T, T) int
}

func CrearHeap[T any](funcion_cmp func(T, T) int) ColaPrioridad[T] {
	return &colaConPrioridad[T]{
		datos: []T{},
		cant:  0,
		cmp:   funcion_cmp,
	}
}

func CrearHeapArr[T any](arreglo []T, funcion_cmp func(T, T) int) ColaPrioridad[T] {
	heap := &colaConPrioridad[T]{
		datos: arreglo,
		cant:  len(arreglo),
		cmp:   funcion_cmp,
	}
	heap.heapify()
	return heap
}

func (cola *colaConPrioridad[T]) EstaVacia() bool {
	return cola.cant == 0
}

func (cola *colaConPrioridad[T]) Cantidad() int {
	return cola.cant
}

func (cola *colaConPrioridad[T]) Encolar(elem T) {
	cola.datos = append(cola.datos, elem)
	cola.cant++
	cola.upHeap(cola.cant - 1)
}

func (cola *colaConPrioridad[T]) VerMax() T {
	if cola.EstaVacia() {
		panic("La cola esta vacia")
	}
	return cola.datos[0]
}

func (cola *colaConPrioridad[T]) Desencolar() T {
	if cola.EstaVacia() {
		panic("La cola esta vacia")
	}
	maxElem := cola.datos[0]
	cola.datos[0] = cola.datos[cola.cant-1]
	cola.datos = cola.datos[:cola.cant-1]
	cola.cant--
	cola.downHeap(0)
	return maxElem
}

func (cola *colaConPrioridad[T]) upHeap(pos int) {
	for pos > 0 {
		padre := (pos - 1) / 2
		if cola.cmp(cola.datos[pos], cola.datos[padre]) <= 0 {
			return
		}
		cola.datos[pos], cola.datos[padre] = cola.datos[padre], cola.datos[pos]
		pos = padre
	}
}

func (cola *colaConPrioridad[T]) downHeap(pos int) {
	for {
		hijoIzq := 2*pos + 1
		hijoDer := 2*pos + 2
		mayor := pos

		if hijoIzq < cola.cant && cola.cmp(cola.datos[hijoIzq], cola.datos[mayor]) > 0 {
			mayor = hijoIzq
		}
		if hijoDer < cola.cant && cola.cmp(cola.datos[hijoDer], cola.datos[mayor]) > 0 {
			mayor = hijoDer
		}
		if mayor == pos {
			return
		}
		cola.datos[pos], cola.datos[mayor] = cola.datos[mayor], cola.datos[pos]
		pos = mayor
	}
}

func (cola *colaConPrioridad[T]) heapify() {
	for i := cola.cant/2 - 1; i >= 0; i-- {
		cola.downHeap(i)
	}
}

func HeapSort[T any](elementos []T, funcion_cmp func(T, T) int) {
	heap := CrearHeapArr(elementos, funcion_cmp)
	for i := len(elementos) - 1; i >= 0; i-- {
		elementos[i] = heap.Desencolar()
	}
}

func funcion_cmp[K comparable](clave1, clave2 K) int {
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
