

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter your full Name")
	name, _ := reader.Readstring('\n') //includes newline character
	fmt.Println("Your name is", name)
	// fmt.Println(reflect.Typeof(name))
	fmt.Print("Enter your favourite number")
	favNUm, _ := reader.Readstring('\n') // includes newline character
	fmt.Println("Your Name is", favNum)
	fmt.Println("Enter your fun fact")
	funfact, _ := reader.Readstring('\n') //includes newline character
	fmt.Println("your fun fact is", funFact)
}
