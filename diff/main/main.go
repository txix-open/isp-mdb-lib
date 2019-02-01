package main

import (
	"encoding/json"
	"fmt"
	"github.com/integration-system/isp-mdb-lib/diff"
	"io/ioutil"
)

func main() {
	a := map[string]interface{}{}
	b := map[string]interface{}{}
	x, err := ioutil.ReadFile("diff/main/a.json")
	if err != nil {
		panic(err)
	}
	y, err := ioutil.ReadFile("diff/main/b.json")
	if err != nil {
		panic(err)
	}
	json.Unmarshal(x, &a)
	json.Unmarshal(y, &b)
	_, d := diff.EvalDiff(a, b)
	by, _ := json.Marshal(d)
	fmt.Println(string(by))

	fmt.Println()

	m := diff.FlattenDelta(d)
	by, _ = json.Marshal(m)
	fmt.Println(string(by))
}
