package main

import "testing"
import "os"
import "fmt"

import "environmate/libs/envutils"

func TestReadWriteAddEnv(t *testing.T) {
    key := "ehIEOgie$4c~rVy{[;U_;&.&-K9gV&yp"
    envName := "local"
    newEnv := envutils.Environment{}
    newEnv.Vars = make([]envutils.Var, 0)
    envutils.CreateEnv("local", key)
    v := envutils.Var{
      Name: "blah",
      Value: "blahblah",
    }
    v2 := envutils.Var{
      Name: "blah2",
      Value: "blahblah",
    }
    envutils.AddVar(envName, key, v.Name, v.Value)
    envutils.AddVar(envName, key, v2.Name, v2.Value)
    if err := envutils.AddVar(envName, key, v.Name, v.Value); err == nil {
       t.Errorf("Duplicate detection when adding new environment variable is not working")
    }
    os.Remove(fmt.Sprintf("%s.encrypted", envName))
}
