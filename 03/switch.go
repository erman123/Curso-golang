package main
 
import "fmt"
func main()  {
	/*
	var a int 
	a = 1
	
	switch a {
	case 1:
		fmt.Println("el valor es 1")
	case 2:
		fmt.Print("el valor es 2")
	default:
		fmt.Println("no es ninguno de los casos")
	}
	*/
	/*
	switch a {
	case 1:
		fallthrough
	case 2:
		fallthrough
	case 3:
		fallthrough
	case 4:
		fallthrough
	case 5:
		fmt.Println("estas entre semana")
	case 6:
		fallthrough
	case 7:
		fmt.Println("estas en fin de semana")
	default:
		fmt.Println("no es un dia valido")
	}
	*/
	
	switch a :=4;{
	case a >= 0 && a <= 5:
		fmt.Println("estas entre semana")
	case a >= 6 && a <= 7:
		fmt.Println("estas en fin de semana")
	default:
		fmt.Println("no es un dia valido")
	}
}