package main

import (
	"fmt"
	"github.com/Bigous/go-test/utils"
)

func main() {
	fmt.Println("Returned: ", utils.If(15 > 14, complex(5, -3)*complex(8, 2), 3.1415927))
}
