package lista_test

import (
	"strings"
	TDALista "tdas/lista"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestListaVacia(t *testing.T) {

	listaInt := TDALista.CrearListaEnlazada[int]()

	require.True(t, listaInt.EstaVacia(), "Una lista recien creada esta vacia")
}

func TestAgregarBorrarElemento(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(1)
	lista.InsertarPrimero(2)
	require.Equal(t, 2, lista.BorrarPrimero(), "Se espera que el primer elemento insertado sea el ultimo en ser borrado")
	require.Equal(t, 1, lista.BorrarPrimero(), "Se espera que el segundo elemento insertado sea el primero en ser borrado")
	require.True(t, lista.EstaVacia(), "La lista debera estar vacia despuees de desapilar todos los elementos")
}

func TestPruebaVolumen(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	cantidadElementos := 10000
	for i := 1; i <= cantidadElementos; i++ {
		lista.InsertarPrimero(i)
	}

	for i := cantidadElementos; i > 0; i-- {
		require.Equal(t, i, lista.VerPrimero(), "El tope de la lista debería ser el elemento correcto en cada iteracion")
		require.Equal(t, i, lista.BorrarPrimero(), "Se espera que el elemento borrado sea el correcto")
	}
	require.True(t, lista.EstaVacia(), "La lista deberia estar vacia despues de borrar todos los elementos")
}

func TestAccionesInvalidasEnListaVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	require.True(t, lista.EstaVacia(), "La lista recin creada deberia estar vacia")

	require.Panics(t, func() { lista.BorrarPrimero() }, "Debe producirse un panic al intentar borrar un elemento de una lista vacia")
	require.Panics(t, func() { lista.VerPrimero() }, "Debe producirse un panic al intentar ver el primer elemento de una lista vacia")
}

func TestEstaVaciaEnListaRecienCreada(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	require.True(t, lista.EstaVacia(), "La lista recien creada deberia estar vacia")

	lista.InsertarPrimero(1)
	lista.InsertarPrimero(2)
	require.False(t, lista.EstaVacia(), "La lista a la que se insertaron elementos no deberia estar vacia")
}

func TestAgregarDiferentesTiposDeDatos(t *testing.T) {
	listaInt := TDALista.CrearListaEnlazada[int]()
	listaInt.InsertarPrimero(1)
	listaInt.InsertarPrimero(2)
	require.Equal(t, 2, listaInt.BorrarPrimero())
	require.Equal(t, 1, listaInt.BorrarPrimero())
	require.True(t, listaInt.EstaVacia())

	listaString := TDALista.CrearListaEnlazada[string]()
	listaString.InsertarPrimero("Hola")
	listaString.InsertarPrimero("MUNDO")
	require.Equal(t, "MUNDO", listaString.BorrarPrimero())
	require.False(t, listaString.EstaVacia())
	require.Equal(t, "Hola", listaString.BorrarPrimero())
	require.True(t, listaString.EstaVacia())

	listaFloat := TDALista.CrearListaEnlazada[float64]()
	listaFloat.InsertarPrimero(3.14)
	listaFloat.InsertarPrimero(2.71)
	require.NotEqual(t, 3.14, listaFloat.BorrarPrimero())
	require.Equal(t, 3.14, listaFloat.BorrarPrimero())
}

func TestIteradorInternoTodosLosElementos(t *testing.T) {

	var (
		unNumero      = []*int{proElemento(5)}
		variosNumeros = []*int{proElemento(8), proElemento(2), proElemento(3), proElemento(4), proElemento(12)}
		mismosNumeros = []*int{proElemento(3), proElemento(3), proElemento(3), proElemento(3), proElemento(3), proElemento(3)}
		variosString  = []*string{proElemento("andres"), proElemento("lorenzo"), proElemento("carla"), proElemento("manuel")}
	)
	var (
		listaUnNumero      = TDALista.CrearListaEnlazada[*int]()
		listaVariosNumeros = TDALista.CrearListaEnlazada[*int]()
		listaMismosNumeros = TDALista.CrearListaEnlazada[*int]()
		listaVariosString  = TDALista.CrearListaEnlazada[*string]()
	)
	var (
		aumentarUno = func(numero *int) bool {
			*numero++
			return true
		}
		eliminarVocales = func(nombre *string) bool {
			*nombre = strings.ReplaceAll(*nombre, "a", "")
			*nombre = strings.ReplaceAll(*nombre, "e", "")
			*nombre = strings.ReplaceAll(*nombre, "i", "")
			*nombre = strings.ReplaceAll(*nombre, "o", "")
			*nombre = strings.ReplaceAll(*nombre, "u", "")
			return true
		}
	)

	insertarArrayALista(unNumero, listaUnNumero, false)
	listaUnNumero.Iterar(aumentarUno)
	visitarArray(unNumero, aumentarUno)
	require.Equalf(t, unNumero, borrarPrimerosNElementos(listaUnNumero, listaUnNumero.Largo()), "Se puede iterar un unico elemento de la lista aplicando la funcion indicada")

	insertarArrayALista(variosNumeros, listaVariosNumeros, false)
	listaVariosNumeros.Iterar(aumentarUno)
	visitarArray(variosNumeros, aumentarUno)
	require.Equalf(t, variosNumeros, borrarPrimerosNElementos(listaVariosNumeros, listaVariosNumeros.Largo()), "Iterar con iterador interno sobre todos los elementos aplica la funcion correctamente a cada uno de ellos")

	insertarArrayALista(mismosNumeros, listaMismosNumeros, false)
	listaMismosNumeros.Iterar(aumentarUno)
	visitarArray(mismosNumeros, aumentarUno)
	require.Equalf(t, mismosNumeros, borrarPrimerosNElementos(listaMismosNumeros, listaMismosNumeros.Largo()), "Se puede iterar una lista con mismos elementos, comportandose de la misma forma para cada elemento")

	insertarArrayALista(variosString, listaVariosString, false)
	listaVariosString.Iterar(eliminarVocales)
	visitarArray(variosString, eliminarVocales)
	require.Equalf(t, variosString, borrarPrimerosNElementos(listaVariosString, listaVariosString.Largo()), "Se puede iterar una lista de strings, aplicando la funcion a cada elemento en el orden y forma correcta")
}

func TestIteradorInternoHastaDevolverFalse(t *testing.T) {

	var (
		unNumeroCumpleCond            = []*int{proElemento(5)}
		unNumeroNoCumple              = []*int{proElemento(2)}
		numerosImparesMenosUno        = []*int{proElemento(3), proElemento(7), proElemento(13), proElemento(4), proElemento(11), proElemento(1)}
		positivosInicioNegativosFinal = []*int{proElemento(3), proElemento(6), proElemento(7), proElemento(-3), proElemento(-8), proElemento(-31)}
		primerNumeroNoCumple          = []*int{proElemento(6), proElemento(5), proElemento(2), proElemento(7), proElemento(8)}
	)
	var (
		listaUnNumeroCumpleCond            = TDALista.CrearListaEnlazada[*int]()
		listaUnNumeroNoCumple              = TDALista.CrearListaEnlazada[*int]()
		listaNumerosImparesMenosUno        = TDALista.CrearListaEnlazada[*int]()
		listaPositivosInicioNegativosFinal = TDALista.CrearListaEnlazada[*int]()
		listaPrimerNumeroNoCumple          = TDALista.CrearListaEnlazada[*int]()
	)
	var (
		multiplicarx2TodosLosElementos = func(numero *int) bool {
			*numero = *numero * 2
			return true
		}
		multiplicarx2HastaEncontrarPar = func(numero *int) bool {
			if *numero%2 == 1 {
				*numero = *numero * 2
				return true
			} else {
				return false
			}
		}
		invertirSignoATodos = func(numero *int) bool {
			*numero = *numero * -1
			return true
		}
		pasarAElementosNegativosHastaEncontraUno = func(numero *int) bool {

			if *numero <= 0 {
				return false
			}

			*numero = *numero * -1
			return true
		}
	)

	insertarArrayALista(unNumeroCumpleCond, listaUnNumeroCumpleCond, false)
	listaUnNumeroCumpleCond.Iterar(multiplicarx2HastaEncontrarPar)
	visitarArray(unNumeroCumpleCond, multiplicarx2TodosLosElementos)
	require.Equalf(t, unNumeroCumpleCond, borrarPrimerosNElementos(listaUnNumeroCumpleCond, listaUnNumeroCumpleCond.Largo()), "Se itera una lista con un solo elemento que cumple la condicion aplicandole la funcion correspondiente")

	insertarArrayALista(unNumeroNoCumple, listaUnNumeroNoCumple, false)
	listaUnNumeroNoCumple.Iterar(multiplicarx2HastaEncontrarPar)
	visitarArray(unNumeroNoCumple, multiplicarx2HastaEncontrarPar)
	require.Equalf(t, unNumeroNoCumple, borrarPrimerosNElementos(listaUnNumeroNoCumple, listaUnNumeroNoCumple.Largo()), "Se itera una lista con un unico elemento que devuelve false y se cumple lo indicado en la funcion correctamente")

	insertarArrayALista(numerosImparesMenosUno, listaNumerosImparesMenosUno, false)
	listaNumerosImparesMenosUno.Iterar(multiplicarx2HastaEncontrarPar)
	visitarArray(numerosImparesMenosUno, multiplicarx2HastaEncontrarPar)
	require.Equalf(t, numerosImparesMenosUno, borrarPrimerosNElementos(listaNumerosImparesMenosUno, listaNumerosImparesMenosUno.Largo()), "Se iteran correctamente aplicando la funcion a cada elemento hasta encontrar uno que devuelva false, dejando de iterar elementos")

	insertarArrayALista(positivosInicioNegativosFinal, listaPositivosInicioNegativosFinal, false)
	listaPositivosInicioNegativosFinal.Iterar(pasarAElementosNegativosHastaEncontraUno)
	visitarArray(positivosInicioNegativosFinal, invertirSignoATodos)
	require.NotEqualf(t, positivosInicioNegativosFinal, borrarPrimerosNElementos(listaPositivosInicioNegativosFinal, listaPositivosInicioNegativosFinal.Largo()), "Al iterar una lista donde un elemento devuelve false, no se iteran todos los elementos y no se le aplica la funcion a todos estos")

	visitarArray(positivosInicioNegativosFinal, invertirSignoATodos)

	insertarArrayALista(positivosInicioNegativosFinal, listaPositivosInicioNegativosFinal, false)
	listaPositivosInicioNegativosFinal.Iterar(pasarAElementosNegativosHastaEncontraUno)
	visitarArray(positivosInicioNegativosFinal, pasarAElementosNegativosHastaEncontraUno)
	require.Equalf(t, positivosInicioNegativosFinal, borrarPrimerosNElementos(listaPositivosInicioNegativosFinal, listaPositivosInicioNegativosFinal.Largo()), "Al iterar una lista donde un elementos devuelve false, se iteran correctamente los primeros elementos hasta encontrar el que devuelve false")

	insertarArrayALista(primerNumeroNoCumple, listaPrimerNumeroNoCumple, false)
	listaPrimerNumeroNoCumple.Iterar(multiplicarx2HastaEncontrarPar)
	require.Equal(t, primerNumeroNoCumple, borrarPrimerosNElementos(listaPrimerNumeroNoCumple, listaPrimerNumeroNoCumple.Largo()), "En una lista que se itera, si el primer elemento devuelve false, deja de iterar no aplicando la funcion a ningun elemento y dejando la lista igual que al inicio")
}

func proElemento[T any](dato T) *T {

	nuevoProT := new(T)
	*nuevoProT = dato
	return nuevoProT

}

func visitarArray[T any](datos []T, visitar func(T) bool) {

	var seguirRecorriendo bool = true
	var i int = 0

	for seguirRecorriendo && len(datos) > i {
		seguirRecorriendo = visitar(datos[i])
		i++
	}

}

// insertarArrayALista agrega todos los elementos del array pasado por parametro en el mismo
// orden dispuesto por el array o en el orden inverso en caso de que el parametro ordenInverso
// sea true
func insertarArrayALista[T any](datos []*T, lista TDALista.Lista[*T], ordenInverso bool) {

	if ordenInverso {
		for _, valor := range datos {
			lista.InsertarPrimero(proElemento(*valor))
		}
	} else {
		for _, valor := range datos {
			lista.InsertarUltimo(proElemento(*valor))
		}
	}

}

// borrarPrimerosNElementos saca los primeros N elementos de la lista indicados por parametro,
// devolviendo un array con el mismo orden en que se encontraban en la lista
func borrarPrimerosNElementos[T any](lista TDALista.Lista[T], numeroABorrar int) []T {

	arrayBorrados := make([]T, 0, numeroABorrar)

	for numeroABorrar > 0 {
		arrayBorrados = append(arrayBorrados, lista.BorrarPrimero())
		numeroABorrar--
	}

	return arrayBorrados
}
