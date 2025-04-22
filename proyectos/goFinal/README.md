Enunciado:

Crear un programa que calcule el cuadrado de un número y lo imprima en la consola. Utilizaremos una goroutine y un canal para comunicar el resultado.

Requisitos:

1) Interfaz: Crear una interfaz NumberCalculator que defina un método Square(int) int.
2) Goroutine: Crear una goroutine que calcule el cuadrado de un número.
3) Canal: Utilizar un canal para comunicar el resultado del cálculo de la goroutine al programa principal.
4) Close: Utilizar close para cerrar el canal cuando se haya enviado el resultado.
5) Módulo echo: Utilizar el módulo echo de Go para imprimir el resultado en la consola.
Pasos a seguir:

Crear la interfaz NumberCalculator con el método Square(int) int.
Crear una estructura Calculator que implemente la interfaz NumberCalculator.
Crear una función CalculateSquare(n int, c chan int) que calcule el cuadrado de un número y envíe el resultado a través del canal.
Crear el programa principal que:
Cree un canal de enteros.
Inicie una goroutine que llame a CalculateSquare(n, c).
Utilice select para leer del canal y, cuando reciba el resultado, imprímelo en la consola utilizando echo.Println().
Cierra el canal cuando se haya recibido el resultado.