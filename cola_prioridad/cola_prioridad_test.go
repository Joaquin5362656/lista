package cola_prioridad_test

import (
	TDAHeap "tdas/cola_prioridad"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHeapVacio(t *testing.T) {
	heap := TDAHeap.CrearHeap[int](func(a, b int) int {
		return a - b
	})

	require.True(t, heap.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })
}

func TestEncolarYVerMax(t *testing.T) {
	heap := TDAHeap.CrearHeap[int](func(a, b int) int {
		return a - b
	})

	heap.Encolar(5)
	require.False(t, heap.EstaVacia())
	require.Equal(t, 5, heap.VerMax())

	heap.Encolar(10)
	require.Equal(t, 10, heap.VerMax())

	heap.Encolar(3)
	require.Equal(t, 10, heap.VerMax())
}

func TestEncolarYDesencolar(t *testing.T) {
	heap := TDAHeap.CrearHeap[int](func(a, b int) int {
		return a - b
	})

	heap.Encolar(5)
	heap.Encolar(10)
	heap.Encolar(3)

	require.Equal(t, 10, heap.Desencolar())
	require.Equal(t, 5, heap.Desencolar())
	require.Equal(t, 3, heap.Desencolar())
	require.True(t, heap.EstaVacia())
}

func TestHeapify(t *testing.T) {
	arreglo := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}
	heap := TDAHeap.CrearHeapArr(arreglo, func(a, b int) int {
		return a - b
	})

	require.False(t, heap.EstaVacia())
	require.Equal(t, 9, heap.VerMax())
}

func TestHeapSort(t *testing.T) {
	arreglo := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}
	expected := []int{1, 1, 2, 3, 3, 4, 5, 5, 5, 6, 9}

	TDAHeap.HeapSort(arreglo, func(a, b int) int {
		return a - b
	})

	require.Equal(t, expected, arreglo, "El arreglo deberia estar ordenado despues de HeapSort")
}

func TestCantidad(t *testing.T) {
	heap := TDAHeap.CrearHeap[int](func(a, b int) int {
		return a - b
	})

	require.Equal(t, 0, heap.Cantidad())

	heap.Encolar(1)
	require.Equal(t, 1, heap.Cantidad())

	heap.Encolar(2)
	require.Equal(t, 2, heap.Cantidad())

	heap.Desencolar()
	require.Equal(t, 1, heap.Cantidad())
}

func TestHeapVolumen(t *testing.T) {
	heap := TDAHeap.CrearHeap[int](func(a, b int) int {
		return a - b
	})

	for i := 0; i < 1000000; i++ {
		heap.Encolar(i)
	}
	require.Equal(t, 1000000, heap.Cantidad())
	require.Equal(t, 999999, heap.VerMax())

	for i := 999999; i >= 0; i-- {
		require.Equal(t, i, heap.Desencolar())
	}
	require.True(t, heap.EstaVacia())
}

func TestHeapUnElemento(t *testing.T) {
	heap := TDAHeap.CrearHeap[int](func(a, b int) int {
		return a - b
	})

	heap.Encolar(42)
	require.Equal(t, 1, heap.Cantidad())
	require.Equal(t, 42, heap.VerMax())
	require.Equal(t, 42, heap.Desencolar())
	require.True(t, heap.EstaVacia())
}

func TestHeapDuplicados(t *testing.T) {
	heap := TDAHeap.CrearHeap[int](func(a, b int) int {
		return a - b
	})

	valores := []int{5, 1, 5, 3, 5, 2}
	for _, v := range valores {
		heap.Encolar(v)
	}

	require.Equal(t, 6, heap.Cantidad())
	require.Equal(t, 5, heap.VerMax())

	for i := 0; i < 3; i++ {
		require.Equal(t, 5, heap.Desencolar())
	}
	require.Equal(t, 3, heap.Desencolar())
	require.Equal(t, 2, heap.Desencolar())
	require.Equal(t, 1, heap.Desencolar())
	require.True(t, heap.EstaVacia())
}

func TestHeapifyConArregloVacio(t *testing.T) {
	arreglo := []int{}
	heap := TDAHeap.CrearHeapArr(arreglo, func(a, b int) int {
		return a - b
	})

	require.True(t, heap.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })
}

func TestHeapSortArregloVacio(t *testing.T) {
	arreglo := []int{}
	expected := []int{}

	TDAHeap.HeapSort(arreglo, func(a, b int) int {
		return a - b
	})

	require.Equal(t, expected, arreglo)
}

func TestHeapSortUnElemento(t *testing.T) {
	arreglo := []int{42}
	expected := []int{42}

	TDAHeap.HeapSort(arreglo, func(a, b int) int {
		return a - b
	})

	require.Equal(t, expected, arreglo)
}

func TestHeapSortElementosIguales(t *testing.T) {
	arreglo := []int{1, 1, 1, 1, 1}
	expected := []int{1, 1, 1, 1, 1}

	TDAHeap.HeapSort(arreglo, func(a, b int) int {
		return a - b
	})

	require.Equal(t, expected, arreglo)
}
