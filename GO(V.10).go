//GO (V.10)
package main
import (
    "fmt"
    "math/rand"
    "time"
)

type Celda struct {
    data int
    next *Celda
}

type Fila struct {
    firstC *Celda
    lastC *Celda
    nextF *Fila
}

type Hoja struct{
    nfilas int
    ncolum int
    firstF *Fila
}

// Operación 1: Crear Hoja
func (hoja*Hoja) crearHoja(){
    // FUNCIÓN QUE CREA UNA HOJA 1X1 (1 FILA POR 1 COLUMNA)
    //1. CREAMOS LA NUEVA COLUMNA
    newCelda := Celda{
        data: 0,
        next : nil}
    //2. CREAMOS LA NUEVA FILA
    newFil := Fila{
        nextF : nil}
    //3. CONFIGURAMOS LA HOJA
    newFil.firstC = &newCelda
    newFil.lastC = &newCelda
    hoja.firstF = &newFil
    hoja.nfilas = 1 // Tiene una fila
    hoja.ncolum = 1 //Tiene una columna
}

// Operación 2: Imprimir Hoja
func (hoja*Hoja) imprimirHoja(){
  auxf := hoja.firstF  //Se apunta a la primera fila
  _ = auxf
  for index := 0; index < hoja.nfilas; index++{ //Se recorre fila por fila
      auxc := auxf.firstC //Se apunta a la primera columna de la fila actual
      _ = auxc
      for index := 0; index < hoja.ncolum; index++ { //Se recorre columna por columna de la fila actual
          fmt.Print(auxc.data, " 🠪 ") //Se printea el valor de la celda actual
          auxc = auxc.next //Se apunta a la siguiente celda
      }
      fmt.Println("nil")
      auxf = auxf.nextF //Se apunta a la siguiente fila
      fmt.Println("↓")
  }
  fmt.Println("nil")
}

//Operación 3: Agregar Columna
func (hoja*Hoja) agregarColumna(n int){
    for index := 0; index < n; index++{
        auxf := hoja.firstF //Se apunta a la primera fila
        _ = auxf
        for index := 0; index < hoja.nfilas; index++{ //Se recorre fila por fila
            newCelda := Celda{ //Se añade una nueva celda a la fila actual
                data: 0,
                next : nil}
            auxf.lastC.next = &newCelda
            auxf.lastC = &newCelda
            auxf = auxf.nextF //Se apunta a la siguiente fila
        }
    }
    hoja.ncolum += n //La hoja tiene más columnas
}

//Operación 4: Agregar FIla
func (hoja*Hoja) agregarFila(n int){
    auxf := hoja.firstF //Se apunta a la primera fila
    _ = auxf
    for index := 1; index < hoja.nfilas; index++{ //Se recorren todas las filas hasta apuntar a la última existente
        auxf = auxf.nextF
    }
    for index := 0; index < n; index++{
        newFil := Fila{ //Se añade una nueva fila al final
            nextF : nil}
        auxf.nextF = &newFil
        newCelda := Celda{ //Se añade una nueva celda a la nueva fila
            data: 0,
            next: nil}
        newFil.firstC = &newCelda
        newFil.lastC = &newCelda
        for index := 1; index < hoja.ncolum; index++{ //Se añaden las celdas restantes a la nueva fila de ser necesario
            newCelda := Celda{
                data: 0,
                next: nil}
            newFil.lastC.next = &newCelda
            newFil.lastC = &newCelda
        }
        auxf = auxf.nextF
    }
    hoja.nfilas += n //La hoja tiene una fila más
}

// Operación 5: Rellana filas con número máximo de Celdas
func (hoja*Hoja) rellenar(nFil int, nCol int){
    if nFil - hoja.nfilas > 0 {
        hoja.agregarFila(nFil - hoja.nfilas)
        
    }
    if nCol - hoja.ncolum > 0 {
        hoja.agregarColumna(nCol - hoja.ncolum)
        
    }
}

//Operación 6: Agrega Data en la Hoja
func (hoja*Hoja) agregar(nFil int, nCol int, valor int){
  hoja.rellenar(nFil, nCol)
  auxf := hoja.firstF
  _ = auxf
  for index := 1; index < nFil; index++{
        auxf = auxf.nextF
  }
  auxc := auxf.firstC
  _ = auxc
  for index := 1; index < nCol; index++{
        auxc = auxc.next
  }
  auxc.data = valor
}

//Operación 7: Operación
func (hoja*Hoja) operacion(niFil int, niCol int, nfFil int, nfCol int, ndFil int, ndCol int, op int){
    if niFil > nfFil {
        aux := nfFil
        nfFil = niFil
        niFil = aux
    }
    if niCol > nfCol {
        aux := nfCol
        nfCol = niCol
        niCol = aux
    }
    hoja.rellenar(nfFil, nfCol)
    suma := 0;
    _ = suma;
    cant := 0;
    _ = cant;
    auxf := hoja.firstF
    _ = auxf
    for index := 1; index < niFil; index++{
        auxf = auxf.nextF
    }
    for index := 0; index < nfFil - niFil + 1; index++{
        auxc := auxf.firstC
        for index := 1; index < niCol; index++{
            auxc = auxc.next
        }
        for index := 0; index < nfCol - niCol + 1; index++{
            suma += auxc.data
            cant++
            auxc = auxc.next
        }
        auxf = auxf.nextF
    }
    if op == 1 {
        suma /= cant
    }
    hoja.agregar(ndFil,ndCol,suma)
}

//operacion 8: Sumar
func (hoja*Hoja) sumar(niFil int, niCol int, nfFil int, nfCol int, ndFil int, ndCol int){
    hoja.operacion(niFil, niCol, nfFil, nfCol, ndFil, ndCol, 0)
}

//Operación 9: Sumar
func (hoja*Hoja) promediar(niFil int, niCol int, nfFil int, nfCol int, ndFil int, ndCol int){
    hoja.operacion(niFil, niCol, nfFil, nfCol, ndFil, ndCol, 1)
}

//Extra 1: Máximo valor Hoja
func (hoja *Hoja) getMaxVal() int{
    auxf := hoja.firstF
    mayor := 0;
    
    for index := 0; index < hoja.nfilas; index++{
        auxC := auxf.firstC
        
        for index := 0; index < hoja.ncolum; index++{
            if(mayor < auxC.data){
                mayor = auxC.data
            }
            auxC = auxC.next
        }
        
        auxf = auxf.nextF
    }
    return mayor
}

//Extra 2: Mínimo valor Hoja
func (hoja *Hoja) getMinVal() int{
    auxf := hoja.firstF
    menor := hoja.getMaxVal()
    
    for index := 0; index < hoja.nfilas; index++{
        auxC := auxf.firstC
        
        for index := 0; index < hoja.ncolum; index++{
            if(menor > auxC.data){
                menor = auxC.data
            }
            auxC = auxC.next
        }
        
        auxf = auxf.nextF
    }
    return menor
}

// Extra 3: Matriz Cuadrada
func (hoja*Hoja) cuadrado(){
    if hoja.nfilas > hoja.ncolum {
        hoja.rellenar(hoja.nfilas,hoja.nfilas)
    }else if hoja.ncolum > hoja.nfilas {
        hoja.rellenar(hoja.ncolum,hoja.ncolum)
    }
}

//Extra 4: Vacía la hoja
func (hoja*Hoja) system32(){
     auxf := hoja.firstF
    _ = auxf
  for index := 0; index < hoja.nfilas; index++{
      auxc := auxf.firstC
      _ = auxc
      for index := 0; index < hoja.ncolum; index++ {
          auxc.data = 0
          auxc = auxc.next
      }
      auxf = auxf.nextF
  }
}

//Extra 5: Llena aleatoriamente la hoja
func (hoja*Hoja) llenadoAleatorio(){
     auxf := hoja.firstF
    _ = auxf
  for index := 0; index < hoja.nfilas; index++{
      auxc := auxf.firstC
      _ = auxc
      for index := 0; index < hoja.ncolum; index++ {
          if auxc.data == 0{
              semilla := rand.New(rand.NewSource(time.Now().UnixNano()))
              valor := semilla.Intn(1000)
              auxc.data = valor
          }
          auxc = auxc.next
      }
      auxf = auxf.nextF
  }
}

//Permitir al usuario modificar la Hoja
func main() {
  hoja := Hoja{}
  _ = hoja
  hoja.crearHoja()
  var n1,n2,n3,n4,n5,n6 int
  for index := 0; index < 1; index += 0{
      fmt.Print("\033[H\033[2J")
      fmt.Println("-----Hoja-----")
      hoja.imprimirHoja()
      fmt.Println("--------------")
      fmt.Println("N° minimo: ", hoja.getMinVal())
      fmt.Println("N° maximo: ", hoja.getMaxVal())
      fmt.Println()
      fmt.Println("0: Salir")
      fmt.Println("1: Agregar Fila(s)")
      fmt.Println("2: Agregar Columna(s)")
      fmt.Println("3: Asignar Valor a Celda")
      fmt.Println("4: Hacer la Hoja cuadrada")
      fmt.Println("5: Sumar Celdas")
      fmt.Println("6: Promediar Celdas")
      fmt.Println("7: Colocar valores aleatorios")
      fmt.Println("8: Vaciar la hoja")
      fmt.Println()
      fmt.Print("Ingrese el n° del comando deseado: ")
      fmt.Scanln(&n1)
      if n1 == 0 {
          fmt.Println()
          fmt.Println("Gracias por usar GO")
          index++
      }else if n1 == 1 {
          fmt.Println()
          fmt.Print("Inserte la cantidad de filas a añadir: ")
          fmt.Scanln(&n1)
          if n1 > 0 {
              hoja.agregarFila(n1)
          }
      }else if n1 == 2 {
          fmt.Println()
          fmt.Print("Inserte la cantidad de columnas a añadir: ")
          fmt.Scanln(&n1)
          if n1 > 0 {
              hoja.agregarColumna(n1)
          }
      }else if n1 == 3 {
          fmt.Println()
          fmt.Print("Inserte la fila de la celda: ")
          fmt.Scanln(&n1)
          if n1 > 0 {
              fmt.Println()
              fmt.Print("Inserte la columna de de la celda: ")
              fmt.Scanln(&n2)
              if n2 > 0 {
                  fmt.Println()
                  fmt.Print("Inserte el valor de la celda: ")
                  fmt.Scanln(&n3)
                  if n3 > 0 {
                      hoja.agregar(n1,n2,n3)
                  }
              }
          }
      }else if n1 == 4 {
          hoja.cuadrado()
      }else if n1 == 5 {
          fmt.Println()
          fmt.Print("Inserte la fila de la celda inicial: ")
          fmt.Scanln(&n1)
          if n1 > 0 {
              fmt.Println()
              fmt.Print("Inserte la columna de la celda inicial: ")
              fmt.Scanln(&n2)
              if n2 > 0 {
                  fmt.Println()
                  fmt.Print("Inserte la fila de la celda final: ")
                  fmt.Scanln(&n3)
                  if n3 > 0 {
                      fmt.Println()
                      fmt.Print("Inserte la columna de la celda final: ")
                      fmt.Scanln(&n4)
                      if n4 > 0 {
                          fmt.Println()
                          fmt.Print("Inserte la fila de la celda donde se guardara la suma: ")
                          fmt.Scanln(&n5)
                          if n5 > 0 {
                              fmt.Println()
                              fmt.Print("Inserte la columna de la celda donde se guardara la suma: ")
                              fmt.Scanln(&n6)
                              if n6 > 0 {
                                  hoja.sumar(n1,n2,n3,n4,n5,n6)
                              }
                          }
                      }
                  }
              }
          }
      }else if n1 == 6 {
          fmt.Println()
          fmt.Print("Inserte la fila de la celda inicial: ")
          fmt.Scanln(&n1)
          if n1 > 0 {
              fmt.Println()
              fmt.Print("Inserte la columna de la celda inicial: ")
              fmt.Scanln(&n2)
              if n2 > 0 {
                  fmt.Println()
                  fmt.Print("Inserte la fila de la celda final: ")
                  fmt.Scanln(&n3)
                  if n3 > 0 {
                      fmt.Println()
                      fmt.Print("Inserte la columna de la celda final: ")
                      fmt.Scanln(&n4)
                      if n4 > 0 {
                          fmt.Println()
                          fmt.Print("Inserte la fila de la celda donde se guardara el promedio: ")
                          fmt.Scanln(&n5)
                          if n5 > 0 {
                              fmt.Println()
                              fmt.Print("Inserte la columna de la celda donde se guardara el promedio: ")
                              fmt.Scanln(&n6)
                              if n6 > 0 {
                                  hoja.promediar(n1,n2,n3,n4,n5,n6)
                              }
                          }
                      }
                  }
              }
          }
      }else if n1 == 7 {
        hoja.llenadoAleatorio()
      }else if n1 == 8{
          hoja.system32()
      }
}
}