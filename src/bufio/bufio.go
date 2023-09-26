package main

import ("fmt"
        "bufio"
	)


func main (){
	r := bufio.NewReader(os.Stdin)

	sum := 0

	for {

		input, inputErr := r.ReadString(' ')
		n := strings.TrimSpace(input)

		if n =="" {
			 continue

		}

		num,convErr := strconv.Atoi(n)
		if convErr != nil{

			fmt.Println(convErr)

		}else{
			sum += sum
		}

		if inputErr == io.EOF{
			break
		}

		if inputErr != nil{
			fmt.Println("Error Reading Stdin: ",inputErr)
		}


	}
	fmt.Println("Sum : %v \n",sum)
}