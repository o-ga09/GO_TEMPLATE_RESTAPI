package main

import "main/api/controller"

func main(){
	s, err := controller.NewServer()
	if err != nil {
		panic(err)
	}
	s.Run()
}