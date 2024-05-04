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

// nombre de la funcion hash = HAsh Bernstein
func HashBernstein(cadena string) uint32 {
	var hash uint32 = 5381
	for _, c := range cadena {
		hash = (hash << 5) + hash + uint32(c)
	}
	return hash
}

func CrearHash[K comparable, V any]() Diccionario[K, V] {
	tam := 17
	tabla := make([]TDALista.Lista[parClaveValor[K, V]], tam)
	return &hashAbierto[K, V]{tabla: tabla, tam: tam}
}

func (h *hashAbierto[K comparable, V any]) Guardar(clave K, dato V) {
	indice := int(HashBernstein(convertirABytes(clave))) % h.tam
	lista := h.tabla[indice]

}


func (h *hashAbierto[K comparable, V any]) Pertenece(clave K) bool {
	return true
}

func (h *hashAbierto[K comparable, V any]) Obtener(clave K) V {
	var dato V
	return dato
}

func (h *hashAbierto[K comparable, V any]) Borrar(clave K) V {
	var dato V
	return dato
}

func (h *hashAbierto[K comparable, V any]) Cantidad() int {
	return 0
}
