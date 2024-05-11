package diccionario_test

import (
	"fmt"
	"testing"

	TDADiccionarioOrdenado "tdas/diccionario"

	"github.com/stretchr/testify/require"
)

func TestDiccionarioVacioo(t *testing.T) {
	t.Log("Comprueba que el diccionario vacío no tiene claves")
	dic := TDADiccionarioOrdenado.CrearABB[string, string](funcion_cmp)
	require.EqualValues(t, 0, dic.Cantidad())
	require.False(t, dic.Pertenece("A"))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener("A") })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar("A") })
}

func TestUnElemento(t *testing.T) {
	t.Log("Comprueba que el diccionario con un elemento tiene esa clave únicamente")
	dic := TDADiccionarioOrdenado.CrearABB[string, int](funcion_cmp)
	dic.Guardar("A", 10)
	require.EqualValues(t, 1, dic.Cantidad())
	require.True(t, dic.Pertenece("A"))
	require.False(t, dic.Pertenece("B"))
	require.EqualValues(t, 10, dic.Obtener("A"))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener("B") })
}

func TestDiccionarioGuardaar(t *testing.T) {
	t.Log("Guarda algunos pocos elementos en el diccionario y comprueba su correcto funcionamiento")
	dic := TDADiccionarioOrdenado.CrearABB[string, string](funcion_cmp)
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	valor1 := "miau"
	valor2 := "guau"
	valor3 := "moo"
	claves := []string{clave1, clave2, clave3}
	valores := []string{valor1, valor2, valor3}

	for i := 0; i < len(claves); i++ {
		require.False(t, dic.Pertenece(claves[i]))
		dic.Guardar(claves[i], valores[i])
		require.EqualValues(t, i+1, dic.Cantidad())
		require.True(t, dic.Pertenece(claves[i]))
		require.EqualValues(t, valores[i], dic.Obtener(claves[i]))
	}
}

func TestDiccionarioBoorrar(t *testing.T) {
	t.Log("Guarda algunos elementos en el diccionario y luego los borra")
	dic := TDADiccionarioOrdenado.CrearABB[string, string](funcion_cmp)
	clave1 := "Rojo"
	clave2 := "Verde"
	clave3 := "Azul"
	valor1 := "#FF0000"
	valor2 := "#00FF00"
	valor3 := "#0000FF"
	claves := []string{clave1, clave2, clave3}
	valores := []string{valor1, valor2, valor3}

	for i := 0; i < len(claves); i++ {
		dic.Guardar(claves[i], valores[i])
	}

	require.EqualValues(t, len(claves), dic.Cantidad())

	for i := 0; i < len(claves); i++ {
		require.True(t, dic.Pertenece(claves[i]))
		dic.Borrar(claves[i])
		require.False(t, dic.Pertenece(claves[i]))
		require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener(claves[i]) })
		require.EqualValues(t, len(claves)-(i+1), dic.Cantidad())
	}
}

func TestDiccionarioActualizar(t *testing.T) {
	t.Log("Guarda un elemento en el diccionario y luego lo actualiza")
	dic := TDADiccionarioOrdenado.CrearABB[string, int](funcion_cmp)
	clave := "Edad"
	valorOriginal := 25
	valorActualizado := 30

	dic.Guardar(clave, valorOriginal)
	require.EqualValues(t, valorOriginal, dic.Obtener(clave))

	dic.Guardar(clave, valorActualizado)
	require.EqualValues(t, valorActualizado, dic.Obtener(clave))
}

func TestReutlizacionDeBoorrados(t *testing.T) {
	t.Log("Prueba de caja blanca: revisa, para el caso que fuere un HashCerrado, que no haya problema " +
		"reinsertando un elemento borrado")
	dic := TDADiccionarioOrdenado.CrearABB[string, string](funcion_cmp)
	clave := "hola"
	dic.Guardar(clave, "mundo!")
	dic.Borrar(clave)
	require.EqualValues(t, 0, dic.Cantidad())
	require.False(t, dic.Pertenece(clave))
	dic.Guardar(clave, "mundooo!")
	require.True(t, dic.Pertenece(clave))
	require.EqualValues(t, 1, dic.Cantidad())
	require.EqualValues(t, "mundooo!", dic.Obtener(clave))
}

func TestConClavesNumericas(t *testing.T) {
	t.Log("Valida que no solo funcione con strings")
	dic := TDADiccionarioOrdenado.CrearABB[int, string](funcion_cmp)
	clave := 10
	valor := "Gatito"

	dic.Guardar(clave, valor)
	require.EqualValues(t, 1, dic.Cantidad())
	require.True(t, dic.Pertenece(clave))
	require.EqualValues(t, valor, dic.Obtener(clave))
	require.EqualValues(t, valor, dic.Borrar(clave))
	require.False(t, dic.Pertenece(clave))
}

func TestConClavesStructs(t *testing.T) {
	t.Log("Valida que tambien funcione con estructuras mas complejas")
	type basico struct {
		a string
		b int
	}
	type avanzado struct {
		w int
		x basico
		y basico
		z string
	}

	dic := TDADiccionarioOrdenado.CrearABB[avanzado, int](funcion_cmp)

	a1 := avanzado{w: 10, z: "hola", x: basico{a: "mundo", b: 8}, y: basico{a: "!", b: 10}}
	a2 := avanzado{w: 10, z: "aloh", x: basico{a: "odnum", b: 14}, y: basico{a: "!", b: 5}}
	a3 := avanzado{w: 10, z: "hello", x: basico{a: "world", b: 8}, y: basico{a: "!", b: 4}}

	dic.Guardar(a1, 0)
	dic.Guardar(a2, 1)
	dic.Guardar(a3, 2)

	require.True(t, dic.Pertenece(a1))
	require.True(t, dic.Pertenece(a2))
	require.True(t, dic.Pertenece(a3))
	require.EqualValues(t, 0, dic.Obtener(a1))
	require.EqualValues(t, 1, dic.Obtener(a2))
	require.EqualValues(t, 2, dic.Obtener(a3))
	dic.Guardar(a1, 5)
	require.EqualValues(t, 5, dic.Obtener(a1))
	require.EqualValues(t, 2, dic.Obtener(a3))
	require.EqualValues(t, 5, dic.Borrar(a1))
	require.False(t, dic.Pertenece(a1))
	require.EqualValues(t, 2, dic.Obtener(a3))

}

func TestClaveVacia(t *testing.T) {
	t.Log("Guardamos una clave vacía (i.e. \"\") y deberia funcionar sin problemas")
	dic := TDADiccionarioOrdenado.CrearABB[string, string](funcion_cmp)
	clave := ""
	dic.Guardar(clave, clave)
	require.True(t, dic.Pertenece(clave))
	require.EqualValues(t, 1, dic.Cantidad())
	require.EqualValues(t, clave, dic.Obtener(clave))
}

func TestValorNulo(t *testing.T) {
	t.Log("Probamos que el valor puede ser nil sin problemas")
	dic := TDADiccionarioOrdenado.CrearABB[string, *int](funcion_cmp)
	clave := "Pez"
	dic.Guardar(clave, nil)
	require.True(t, dic.Pertenece(clave))
	require.EqualValues(t, 1, dic.Cantidad())
	require.EqualValues(t, (*int)(nil), dic.Obtener(clave))
	require.EqualValues(t, (*int)(nil), dic.Borrar(clave))
	require.False(t, dic.Pertenece(clave))
}

func TestCadenaLargaParticular(t *testing.T) {
	t.Log("Se han visto casos problematicos al utilizar la funcion de hashing de K&R, por lo que " +
		"se agrega una prueba con dicha funcion de hashing y una cadena muy larga")
	// El caracter '~' es el de mayor valor en ASCII (126).
	claves := make([]string, 10)
	cadena := "%d~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~" +
		"~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~"
	dic := TDADiccionarioOrdenado.CrearABB[string, string](funcion_cmp)
	valores := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J"}
	for i := 0; i < 10; i++ {
		claves[i] = fmt.Sprintf(cadena, i)
		dic.Guardar(claves[i], valores[i])
	}
	require.EqualValues(t, 10, dic.Cantidad())

	ok := true
	for i := 0; i < 10 && ok; i++ {
		ok = dic.Obtener(claves[i]) == valores[i]
	}

	require.True(t, ok, "Obtener clave larga funciona")
}

func TestGuardarYBorrarRepetidasVeces(t *testing.T) {
	t.Log("Esta prueba guarda y borra repetidas veces. Esto lo hacemos porque un error comun es no considerar " +
		"los borrados para agrandar en un Hash Cerrado. Si no se agranda, muy probablemente se quede en un ciclo " +
		"infinito")

	dic := TDADiccionarioOrdenado.CrearABB[int, int]()
	for i := 0; i < 1000; i++ {
		dic.Guardar(i, i)
		require.True(t, dic.Pertenece(i))
		dic.Borrar(i)
		require.False(t, dic.Pertenece(i))
	}
}
