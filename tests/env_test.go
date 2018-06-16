package main

import "testing"

import "environmate/libs/envutils"

func TestReadWriteEnv(t *testing.T) {
    key := "ehIEOgie$4c~rVy{[;U_;&.&-K9gV&yp"
    newEnv := envutils.Environment{}
    newEnv.Vars = make([]envutils.Var, 0)
    v := envutils.Var{
      Name: "blah",
      Value: "blahblah",
    }
    newEnv.Vars = append(newEnv.Vars, v)
    envutils.WriteEnv("local", key, newEnv)
    env := envutils.ReadEnv("local", key)

    if env.Vars[0].Name != "blah" {
       t.Errorf("Could not create and read a new Environment")
    }
}
