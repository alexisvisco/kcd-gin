package main

import (
	"fmt"
	"github.com/alexisvisco/kcd-gin/pkg/kcdgin"
	"github.com/alexisvisco/kcd/pkg/errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.New()

	kcdgin.Setup() // Do not forget this part otherwise you will not be able to recover the path parameters

	r.GET("/:name", kcdgin.Handler(YourHttpHandler, http.StatusOK))
	//                          ^ Here the magic happen this is the only thing you need
	//                            to do. Adding kcdgin.Handler(your handler)

	_ = r.Run(":3000")
}

// CreateCustomerInput is an example of input for an http request.
type CreateCustomerInput struct {
	Name   string   `path:"name"`
	Emails []string `query:"emails" exploder:","`
}

// CustomerOutput is the output type of your handler it contain the input for simplicity.
type CustomerOutput struct {
	Name string `json:"name"`
}

// YourHttpHandler is your http handler but in a shiny version.
// You can add *http.ResponseWriter or http.Request in params if you want.
func YourHttpHandler(in *CreateCustomerInput) (CustomerOutput, error) {
	// do some stuff here

	fmt.Printf("%+v", in)

	return CustomerOutput{}, errors.NewWithKind(errors.KindInternal, "c'est fini !")
}
