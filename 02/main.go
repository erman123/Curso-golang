package main
import "fmt"
func main(){
//todas estas formas son validas para declarar variables

	//var nombre,apellido string
	//nombre = "erman"
	//apellido ="meneses"
	
	//nombre :="erman"
	//apellido :="meneses"

	nombre ,apellido := "erman","meneses"

	fmt.Println(nombre)
	fmt.Println(apellido)
}