package diccionario

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
	} else {
		ramaEncontrada := arbol.raiz.buscarRama(clave, arbol.funcCmp)

		nodoAModificar := *ramaEncontrada

		if nodoAModificar == nil {
			nodoAModificar = crearNodoAbb(parClaveValor[K, V]{clave, dato})
		} else {
			nodoAModificar.nodoRaiz.dato = dato
		}

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

	if comparar(rama.nodoRaiz.clave, clave) < 0 {
		ramaEncontrada = rama.izquierdo.buscarRama(clave, comparar)
	} else {
		ramaEncontrada = rama.derecho.buscarRama(clave, comparar)
	}

	if ramaEncontrada == nil {
		return &rama
	} else {
		return ramaEncontrada
	}
}
