package main

import (
	"fmt"
)

const port = 3000

func main() {
	r := setupRoutes()

	r.Run(fmt.Sprintf(":%d", port))
}
