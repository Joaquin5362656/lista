package diccionario

//import ("strings")

type FuncCmp[K any] func(K, K) int

type nodoAbb[K comparable, V any] struct {
	izquierdo *nodoAbb[K, V]
	derecho   *nodoAbb[K, V]
	clave     K
	dato      V
}

type abb[K comparable, V any] struct {
	raiz     *nodoAbb[K, V]
	cantidad int
	cmp      FuncCmp[K]
}

func CrearABB[K comparable, V any](funcion_cmp func(K, K) int) *abb[K, V] {
	return &abb[K, V]{raiz: nil, cantidad: 0, cmp: funcion_cmp}
}

// Guardar inserta un nuevo par clave-valor en el ABB
func (a *abb[K, V]) Guardar(clave K, dato V) {
	a.raiz = a.insertar(a.raiz, clave, dato)
}

// insertar inserta recursivamente un nuevo nodo en el Abb
func (a *abb[K, V]) insertar(nodo *nodoAbb[K, V], clave K, dato V) *nodoAbb[K, V] {
	if nodo == nil {
		return &nodoAbb[K, V]{clave: clave, dato: dato}
	}

	cmpResultado := a.cmp(clave, nodo.clave)
	switch {
	case cmpResultado < 0:
		nodo.izquierdo = a.insertar(nodo.izquierdo, clave, dato)
	case cmpResultado > 0:
		nodo.derecho = a.insertar(nodo.derecho, clave, dato)
	default:
		nodo.dato = dato
	}
	return nodo
}

// Pertenece verifica si una clave está presente en el Abb
func (a *abb[K, V]) Pertenece(clave K) bool {
	_, encontrado := a.buscar(a.raiz, clave)
	return encontrado
}

// busca recursivamente un nodo con la clave dada en el Abb
func (a *abb[K, V]) buscar(nodo *nodoAbb[K, V], clave K) (*nodoAbb[K, V], bool) {
	if nodo == nil {
		return nil, false
	}

	cmpResultado := a.cmp(clave, nodo.clave)
	switch {
	case cmpResultado < 0:
		return a.buscar(nodo.izquierdo, clave)
	case cmpResultado > 0:
		return a.buscar(nodo.derecho, clave)
	default:
		return nodo, true
	}
}

// Obtiene el valor asociado con la clave dada en el Abb
func (a *abb[K, V]) Obtener(clave K) V {
	nodo, encontrado := a.buscar(a.raiz, clave)
	if !encontrado {
		panic("La clave no pertenece al diccionario")
	}
	return nodo.dato
}

// Borrar elimina un par clave-valor del Abb
func (a *abb[K, V]) Borrar(clave K) V {
	var borrado *nodoAbb[K, V]
	a.raiz, borrado = a.borrar(a.raiz, clave)
	if borrado == nil {
		panic("La clave no pertenece al diccionario")
	}
	return borrado.dato
}

// borrar elimina recursivamente un nodo con la clave dada del Abb
func (a *abb[K, V]) borrar(nodo *nodoAbb[K, V], clave K) (*nodoAbb[K, V], *nodoAbb[K, V]) {
	if nodo == nil {
		return nil, nil
	}

	cmpResultado := a.cmp(clave, nodo.clave)
	var borrado *nodoAbb[K, V]

	switch {
	case cmpResultado < 0:
		nodo.izquierdo, borrado = a.borrar(nodo.izquierdo, clave)
		return nodo, borrado
	case cmpResultado > 0:
		nodo.derecho, borrado = a.borrar(nodo.derecho, clave)
		return nodo, borrado
	default:
		if nodo.izquierdo == nil {
			return nodo.derecho, nodo
		}
		if nodo.derecho == nil {
			return nodo.izquierdo, nodo
		}

		reemplazo := a.minimo(nodo.derecho)
		nodo.clave, nodo.dato = reemplazo.clave, reemplazo.dato
		nodo.derecho, borrado = a.borrar(nodo.derecho, reemplazo.clave)
		return nodo, borrado
	}
}

// minimo encuentra y devuelve el nodo con la clave ms pequeña en el Abb
func (a *abb[K, V]) minimo(nodo *nodoAbb[K, V]) *nodoAbb[K, V] {
	for nodo.izquierdo != nil {
		nodo = nodo.izquierdo
	}
	return nodo
}

// Cantidad devuelve la cantidad de nodos en el ABB
func (a *abb[K, V]) Cantidad() int {
	return a.contar(a.raiz)
}

// cuenta recursivamente la cantidad de nodos en el Abb
func (a *abb[K, V]) contar(nodo *nodoAbb[K, V]) int {
	if nodo == nil {
		return 0
	}
	return 1 + a.contar(nodo.izquierdo) + a.contar(nodo.derecho)
}
