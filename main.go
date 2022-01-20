package main

import (
	"fmt"
	"github.com/jinzhu/copier"
	"google.golang.org/api/androidpublisher/v3"
	"google.golang.org/api/googleapi"
	"net/http"
)

func main() {
	fmt.Println("-- Start --")

	product := androidpublisher.InAppProduct{
		ServerResponse: googleapi.ServerResponse{
			Header: http.Header{
				"Content-Type": []string{"application/json"}, // Removing this line fixes it. Still happens if slice is empty
			},
		},
	}

	productTarget := &androidpublisher.InAppProduct{}
	cloneErr := copier.CopyWithOption(productTarget, product, copier.Option{DeepCopy: true})
	if cloneErr != nil {
		fmt.Println(fmt.Sprintf("Clone Error: %s", cloneErr))
	}

	fmt.Println("-- Fin --")
}
