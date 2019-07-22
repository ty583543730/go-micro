package main

import "./logic"

func main()  {
	model, err := logic.CreateModel()
	logic.CheckErr(err)


}
