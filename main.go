package main

import (
	"fmt"
	"github.com/jinzhu/copier"
)

func main() {
	fmt.Println("-- Start --")

	// map[string][]string is the type of the http.Header in googleapi.ServerResponse on the thing we're tyring to clone googleapi.InAppProduct
	// Its anonymous on the InAppProduct, but thats not the issue as seen in non-anonymous example
	// https://pkg.go.dev/google.golang.org/api/androidpublisher/v3#InAppProduct
	httpHeader := map[string][]string{
		"Content-Type": []string{}, // Removing this line fixes it. Still happens if slice is empty
	}
	httpHeaderTarget := &map[string][]string{}

	cloneErr := copier.CopyWithOption(httpHeaderTarget, httpHeader, copier.Option{DeepCopy: true})
	if cloneErr != nil {
		fmt.Println(fmt.Sprintf("httpHeader Clone Error: %s", cloneErr))
	}

	fmt.Println("-- Fin --")
}
