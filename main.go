package main

import (
	"fmt"
	"net/http"

	"github.com/tomvictor/try_go/controllers"
)

func main() {
	fmt.Println("Server running in port 3000")
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
