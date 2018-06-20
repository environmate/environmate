package main

import "testing"

import "environmate/libs/envutils"

func TestReadWriteAddEnv(t *testing.T) {
    key := "ehIEOgie$4c~rVy{[;U_;&.&-K9gV&yp"
    envName := "local"
    newEnv := envutils.Environment{}
    newEnv.Vars = make([]envutils.Var, 0)
    envutils.WriteEnv("local", key, newEnv)
    v := envutils.Var{
      Name: "blah",
      Value: "blahblah",
    }
    v2 := envutils.Var{
      Name: "blah2",
      Value: "blahblah",
    }
    envutils.AddVar(envName, key, v)
    envutils.AddVar(envName, key, v2)
    if err := envutils.AddVar(envName, key, v); err == nil {
       t.Errorf("Duplicate detection when adding new environment variable is not working")
    }
    env := envutils.ReadEnv(envName, key)

    if env.Vars[0].Name != "blah" {
       t.Errorf("Could not create and read a new Environment")
    }
    if len(env.Vars) != 2 {
       t.Errorf("Could not add variable to existing environment")
    }
}
