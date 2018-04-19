package main
import "fmt"
func main()  {
	/*
	if !(5 < 10){
		fmt.Println("verdadero")
	}

	fmt.Println("fin . ")
	*/
/*
	//la varible isValid solo va a exitir dentro del if
	if isValid := true ; isValid{
		fmt.Println("verdadero")
	}else{
		fmt.Println("falso")
	}
	fmt.Println("fin . ")
	*/

	if edad ,nombre := 15 ,"nano";edad < 14{
		fmt.Println(nombre,"es un niño")
	}else if edad < 18{
		fmt.Println(nombre,"es un adolescente")
	}else if edad < 30{
		fmt.Println(nombre,"es un adulto")
	}else{
		fmt.Println(nombre,"es un adulto grandecito")
	}
}