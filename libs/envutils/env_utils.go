package envutils

import (
	"encoding/json"
  "errors"
	"fmt"
  "io/ioutil"

	"environmate/libs/encryption"
)

type Var struct {
	Name	string `json:"name"`
	Value string `json:"value"`
}
type Environment struct {
	Vars [] Var `json:"vars"`
}

func isUniqueVar(e Environment, v Var) bool {
  for _, variable := range e.Vars {
    if (variable.Name == v.Name) {
      return false
    }
  }
  return true
}

func ReadEnv(env string, key string) Environment {
	var e Environment
	keyBytes := []byte(key)
  data, _ := ioutil.ReadFile(fmt.Sprintf("%s.encrypted", env))
	envJson, _ := aes.Decrypt(keyBytes, string(data))
	err := json.Unmarshal([]byte(envJson), &e)
	if err != nil {
		panic(err)
	}
	return e
}

func WriteEnv(env string, key string, envData Environment) {
	keyBytes := []byte(key)
	envJson, _ := json.Marshal(envData)
	encryptedJson, _ := aes.Encrypt(keyBytes, string(envJson))
	if err := ioutil.WriteFile(fmt.Sprintf("%v.encrypted", env), []byte(encryptedJson), 0644); err != nil {
		fmt.Println(fmt.Sprintf("There was a problem writing the env (%v)", env))
	}
}

func AddVar(env string, key string, v Var) error {
  e := ReadEnv(env, key)
  if (!isUniqueVar(e, v)) {
    return errors.New(fmt.Sprintf("The variable %v already exists", v.Name))
  }
  e.Vars = append(e.Vars, v)
  WriteEnv(env, key, e)
  return nil
}
