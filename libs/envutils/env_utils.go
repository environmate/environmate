package envutils

import (
  "encoding/json"
  "errors"
  "fmt"
  "io/ioutil"
  "log"
  "strings"

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

func readEnv(env string, key string) Environment {
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

func writeEnv(env string, key string, envData Environment) error {
  keyBytes := []byte(key)
  envJson, _ := json.Marshal(envData)
  encryptedJson, _ := aes.Encrypt(keyBytes, string(envJson))
  if err := ioutil.WriteFile(fmt.Sprintf("%v.encrypted", env), []byte(encryptedJson), 0644); err != nil {
    return err
  }
  return nil
}

func AddVar(env string, key string, varName string, varValue string) error {
  e := readEnv(env, key)
  v := Var{varName, varValue}
  if (!isUniqueVar(e, v)) {
    return errors.New(fmt.Sprintf("The variable %v already exists", v.Name))
  }
  e.Vars = append(e.Vars, v)
  writeEnv(env, key, e)
  return nil
}

func RemoveVar(env string, key string, varName string) error {
  e := readEnv(env, key)
  tmpVars := e.Vars[:0]
  for _, ev := range e.Vars {
    if ev.Name != varName {
      tmpVars = append(tmpVars, ev)
    }
  }
  e.Vars = tmpVars
  writeEnv(env, key, e)
  return nil
}

func UpdateVar(env string, key string, varName string, varValue string) error {
  e := readEnv(env, key)
  v := Var{varName, varValue}
  for i := 0; i < len(e.Vars); i++ {
    envVar := &e.Vars[i]
    if envVar.Name == v.Name {
      envVar.Value = v.Value
    }
  }
  writeEnv(env, key, e)
  return nil
}

func CreateEnv(env string, key string) {
  fmt.Println(fmt.Sprintf("Creating env (%v)...", env))
  newEnv := Environment{}
  newEnv.Vars = make([]Var, 0)

  if err := writeEnv(env, key, newEnv); err == nil {
    fmt.Println(fmt.Sprintf("Successfully created env (%v)", env))
  } else {
    fmt.Println(fmt.Sprintf("There was a problem creating the env (%v)", env))
  }
}

func ListEnvs(env string, key string) {
  extension := ".json.encrypted"
  files, err := ioutil.ReadDir("./*" + extension)
  if err != nil {
    log.Fatal(err)
  }

  for _, f := range files {
    if (strings.Contains(f.Name(), extension)) {
      envName := strings.Replace(f.Name(), extension, "", -1)
      fmt.Println(envName)
    }
  }
}

func GenEnv(env string, key string) error {
  envData := readEnv(env, key)
  envJson, _ := json.Marshal(envData)

  if err := ioutil.WriteFile(fmt.Sprintf("%s.env", env), []byte(envJson), 0644); err != nil {
    return err
  }
  return nil
}
