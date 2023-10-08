package main

import "github.com/o-ga09/GO_TEMPLATE_RESTAPI/api/controller"

func main(){
	s, err := controller.NewServer()
	if err != nil {
		panic(err)
	}
	s.Run()
}