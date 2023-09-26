package main

import ("bufio"
        "fmt"
	     "os"
		 "strings")

const (
	CmdHello = "hello"
	CmdGoodbye = "bye"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	numLines :=0
	numcommands :=0
	for scanner.Scan(){
		if strings.ToUpper(scanner.Text()) == "Q"{
			break
		}else{
			text := strings.TrimSpace(scanner.Text())

			switch text{
			case CmdHello:
				numcommands +=1
				fmt.Println("commmand Response : hi")

			case CmdGoodbye:
				numcommands +=1
				fmt.Println("Command Response  : bye")
			}
			if text !="" {
				numLines +=1
			}
		}
	}

	fmt.Printf("You entered %v lines",numLines)
	fmt.Printf("You entered %v commands",numcommands)
}