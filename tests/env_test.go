package main

import "testing"

import "environmate/libs/envutils"

func TestReadWriteEnv(t *testing.T) {
    newEnv := envutils.Environment{}
    newEnv.Vars = make([]envutils.Var, 0)
    v := envutils.Var{
      Name: "blah",
      Value: "blahblah",
    }
    newEnv.Vars = append(newEnv.Vars, v)
    envutils.WriteEnv("local", newEnv)
    env := envutils.ReadEnv("local")

    if env.Vars[0].Name != "blah" {
       t.Errorf("Could not create and read a new Environment")
    }
}
