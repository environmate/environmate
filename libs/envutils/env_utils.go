package envutils

import (
	"encoding/json"
  "io/ioutil"
	"fmt"

	"environmate/libs/encryption"
)

type Var struct {
	Name	string `json:"name"`
	Value string `json:"value"`
}
type Environment struct {
	Vars [] Var `json:"vars"`
}

func ReadEnv(env string) Environment {
	var e Environment
	key := []byte("ehIEOgie$4c~rVy{[;U_;&.&-K9gV&yp")
  data, _ := ioutil.ReadFile(fmt.Sprintf("%s.encrypted", env))
	envJson, _ := aes.Decrypt(key, string(data))
	err := json.Unmarshal([]byte(envJson), &e)
	if err != nil {
		panic(err)
	}
	return e
}

func WriteEnv(env string, envData Environment) {
	key := []byte("ehIEOgie$4c~rVy{[;U_;&.&-K9gV&yp")
	envJson, _ := json.Marshal(envData)
	encryptedJson, _ := aes.Encrypt(key, string(envJson))
	if err := ioutil.WriteFile(fmt.Sprintf("%v.encrypted", env), []byte(encryptedJson), 0644); err != nil {
		fmt.Println(fmt.Sprintf("There was a problem writing the env (%v)", env))
	}
}

func main() {
	createEnv := Environment{}
	createEnv.Vars = make([]Var, 0)
	v := Var{
		Name: "blah",
		Value: "blahblah",
	}
	createEnv.Vars = append(createEnv.Vars, v)
	WriteEnv("local", createEnv)
	env := ReadEnv("local")
	fmt.Println(env)
	fmt.Println(env.Vars[0].Name)
}
