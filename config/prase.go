package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

func main() {

	if bytes, err := ioutil.ReadFile("./config/config.yaml"); err == nil{
		var f interface{}
		if e := yaml.Unmarshal(bytes, &f); e ==nil{
			fmt.Printf("%+v",f)
		}else {
			fmt.Println(e)
		}
	}else {
		fmt.Println(err)
	}

}
