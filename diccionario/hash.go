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

func CrearHash[K comparable, V any]() Diccionario[K, V] {
	tam := 17
	tabla := make([]TDALista.Lista[parClaveValor[K, V]], tam)
	return &hashAbierto[K, V]{tabla: tabla, tam: tam}
}

func convertirABytes[K comparable](clave K) []byte {
	return []byte(fmt.Sprintf("%v", clave))
}

// nombre de la funcion hash = HAsh Bernstein
func HashBernstein(cadena string) uint32 {
	var hash uint32 = 5381
	for _, c := range cadena {
		hash = (hash << 5) + hash + uint32(c)
	}
	return hash
}

func (h *hashAbierto[K comparable, V any]) Guardar(clave K, dato V) {
	indice := int(HashBernstein(convertirABytes(clave))) % h.tam
	lista := h.tabla[indice]
	var encontrado bool

	lista.Iterar(func(par parClaveValor[K, V]) bool {
		if par.clave == clave {
			par.dato = dato
			encontrado = true
			return false
		}
		return true
	})

	if !encontrado {
		lista.InsertarUltimo(parClaveValor[K, V]{clave: clave, dato: dato})
		h.cantidad++
	}
}

func (h *hashAbierto[K comparable, V any]) Pertenece(clave K) bool {
	indice := int(HashBernstein(string(convertirABytes(&clave)))) % h.tam
	lista := h.tabla[indice]

	var encontrado bool
	lista.Iterar(func(par parClaveValor[K, V]) bool {
		if par.clave == clave{
			encontrado = true
			return false
		}
		return true
	})
	return encontrado
}

func (h *hashAbierto[K comparable, V any]) Obtener(clave K) V {
	indice := int(HashBernstein(string(convertirABytes(&clave)))) % h.tam
	lista := h.tabla[indice]
	var dato V
	var encontrado bool

	lista.Iterar(func(par parClaveValor[K, V]) bool {
		if par.clave == clave {
			dato = par.dato
			encontrado = true
			return false
		}
		return true
	})

	if !encontrado{
		panic("La clave no pertenece al diccionario")
	}

	return dato
}

func (h *hashAbierto[K comparable, V any]) Borrar(clave K) V {
	indice := int(HashBernstein(string(convertirABytes(&clave)))) % h.tam
	lista := h.tabla[indice]
	var dato V
	var encontrado bool

	lista.Iterar(func(par parClaveValor[K, V]) bool {
		if par.clave == clave {
			dato = par.dato
			lista.Borrar()
			encontrado = true
			return false
		}
		return true
	})

	if !encontrado{
		panic("La clave no pertenece al diccionario")
	}

	h.cantidad--
	
	return dato
}

func (h *hashAbierto[K comparable, V any]) Cantidad() int {
	return h.cantidad
}
