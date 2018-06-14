package main

import (
  "io/ioutil"
	"fmt"
	//"encoding/json"
)

var env string
type Var struct {
	Name	string `json:"name"`
	Value string `json:"value"`
}
type Environment struct {
	Vars [] Var `json:"vars"`
}

func loadEnv(env string) {
  data, err := ioutil.ReadFile(fmt.Sprintf("%s.json.encrypted", env))
  if err != nil {
    fmt.Println("Error")
    fmt.Println(err)
  } else {
    fmt.Println(string(data))
  }
}

func main() {
  loadEnv("testing")
}
