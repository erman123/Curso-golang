package main
import "fmt"
func main(){
//todas estas formas son validas para declarar variables

	//var nombre,apellido string
	//nombre = "erman"
	//apellido ="meneses"
	
	//nombre :="erman"
	//apellido :="meneses"
	
	nu := 6
	nombre ,apellido := "erman","meneses"

	fmt.Println(nombre)
	fmt.Println(apellido)
	fmt.Printf("hola %s como estas?\n",nombre)
	fmt.Printf("hola %s %s como estas?\n",nombre,apellido)
	fmt.Printf("hola %s %s como estas?\n",apellido,nombre)
	//par aimprimir numeros
	fmt.Printf("hola %s %d como estas?",nombre,nu)
}