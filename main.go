// package main


// import "fmt"

// var nemo = []string {"nemo","jarky","mainn"}

// func main (){

// 	for i := 0; i < len(nemo); i++{
		
// 		if nemo[i] == "nemo"{
// 			fmt.Println("Found nemo!!!")
//         nemo = append(nemo,"NEMOO")
// 			continue
// 		}

// 		fmt.Println(nemo[i])
// 	}

// 	fmt.Println(nemo)
// }

package main

import "fmt"


var nemo = []string {"nemo","jarchy","nonty"}

func main(){

	for index, element :=range nemo{

		if element == "nemo"{
			fmt.Println("Found Nemo")
		}

		fmt.Println(index," : ",element)
	}
}