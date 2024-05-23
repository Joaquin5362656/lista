package cola_prioridad

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
	copia := make([]T, len(arreglo))
	copy(copia, arreglo)

	heap := &colaConPrioridad[T]{
		datos: copia,
		cant:  len(copia),
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
	cola.cant--
	cola.downHeap(0)
	cola.datos = cola.datos[:cola.cant]

	if 4*cola.cant <= len(cola.datos) && cola.cant > 0 {
		cola.redimensionar()
	}

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
	heapify(cola.datos, cola.cant, cola.cmp)
}

func HeapSort[T any](elementos []T, funcion_cmp func(T, T) int) {
	heapify(elementos, len(elementos), funcion_cmp)

	for i := len(elementos) - 1; i > 0; i-- {
		elementos[0], elementos[i] = elementos[i], elementos[0]
		downHeap(elementos, i, 0, funcion_cmp)
	}
}

func heapify[T any](datos []T, n int, cmp func(T, T) int) {
	for i := n/2 - 1; i >= 0; i-- {
		downHeap(datos, n, i, cmp)
	}
}

func downHeap[T any](datos []T, n, pos int, cmp func(T, T) int) {
	for {
		hijoIzq := 2*pos + 1
		hijoDer := 2*pos + 2
		mayor := pos

		if hijoIzq < n && cmp(datos[hijoIzq], datos[mayor]) > 0 {
			mayor = hijoIzq
		}
		if hijoDer < n && cmp(datos[hijoDer], datos[mayor]) > 0 {
			mayor = hijoDer
		}
		if mayor == pos {
			return
		}
		datos[pos], datos[mayor] = datos[mayor], datos[pos]
		pos = mayor
	}
}

func (cola *colaConPrioridad[T]) redimensionar() {
	nuevaCapacidad := len(cola.datos) / 2
	nuevaCopia := make([]T, cola.cant, nuevaCapacidad)
	copy(nuevaCopia, cola.datos[:cola.cant])
	cola.datos = nuevaCopia
}
