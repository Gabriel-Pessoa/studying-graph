package main

import (
	"fmt"

	"github.com/Gabriel-Pessoa/studying-graph/data-structures/dictionary"
)

func main() {
	dic := dictionary.NewDictionary()
	dic.Set("Gandalf", "gandalf@email.com")
	dic.Set("John", "johnsnow@email.com")
	dic.Set("Tyrion", "tyrion@email.com")

	fmt.Println(dic.Get("Tyrion"))

	err := dic.Set("Gabriel", "")
	if err != nil {
		fmt.Println(err)
	}

	value, err := dic.Get("Gabriel")
	fmt.Printf("Value: %v, Error: %v", value, err)
}
