Para hablar de channels tenemos que hablar de goroutines.

Una goroutine es muy parecido a lo que en otros lenguajes se conoce como thread o hilo. Es básicamente una tarea que se ejecuta de forma independiente de otras.

Existen más cosas a considerar pero para ponerlo simple, si queremos ejecutar una función como una goroutine sólo debemos anteponer la palabra go antes de la llamada de dicha función, así:

go function()
Ahora bien, para que las goruotines puedan comunicarse usamos los channels (esa es su verdadera función).

Un channel se puede declarar así:

count := make(chan int)
Ahora, cabe mencionar que hay dos tipos de channels, están los Buffered channels y los Unbuffered channels. Pongo un ejemplo de ambos:

// Unbuffered channel
count := make(chan int)
// Buffered channel
count:= make(chan int, 10)
Los Buffered channels tienen una capacidad máxima y si intentamos usarlos más de esas veces obtendremos un error.

La mayor diferencia sería que con los Buffered channels no se puede garantizar que los mensajes han sido entregados.

Para enviar un mensaje a un channel hacemos lo siguiente:

// Buffered channel
messages := make(chan string, 2)
// Enviar un mensaje a un channel.
messages <- "Golang"
Para recibir el mensaje del channel hacemos esto:

value := <-messages
Finalmente te dejo un ejemplo completo usando channels y goroutines, explico el código con comentarios dentro del mismo.

package main

import (
    "fmt"
    "sync"
)

// wg se usa para esperar a que el programa termine.
var wg sync.WaitGroup

func main() {

    count := make(chan int)
    // Esto es para saber cuánto esperar, se agrega 1 para cada goroutine.
    wg.Add(2)

    fmt.Println("Inician Goroutines")
    //Empieza una goroutine con la etiqueta "A"
    go printCounts("A", count)
    //Empieza una goroutine con la etiqueta "B"
    go printCounts("B", count)
    fmt.Println("Empieza el channel")
    count <- 1
    // Espera a que las goroutines terminen.
    fmt.Println("Esperando para finalizar")
    wg.Wait()
    fmt.Println("\nFinalizando el programa")
}

func printCounts(label string, count chan int) {
    // Vamos a ejecutar el Done cuando terminemos.
    defer wg.Done()
    for {
        //Recibe un mensaje del Channel
        val, ok := <-count
        if !ok {
            fmt.Println("Se ha cerrado el channel")
            return
        }
        fmt.Printf("Contador: %d recibió desde %s \n", val, label)
        if val == 10 {
            fmt.Printf("El channel fue cerrado desde %s \n", label)
            // Cerrar el channel
            close(count)
            return
        }
        val++
        // Enviar de regreso el contador a la otra goroutine.
        count <- val
    }
}
La salida sería algo como esto:

Inician Goroutines
Empieza el channel
Contador: 1 recibió desde A
Contador: 2 recibió desde B
Esperando para finalizar
Contador: 3 recibió desde A
Contador: 4 recibió desde B
Contador: 5 recibió desde A
Contador: 6 recibió desde B
Contador: 7 recibió desde A
Contador: 8 recibió desde B
Contador: 9 recibió desde A
Contador: 10 recibió desde B
El channel fue cerrado desde B
Se ha cerrado el channel
Finalizando el programa