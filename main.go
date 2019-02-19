package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

func main() {
	f, err := os.Open("yamls/router.yml")
	if err != nil {
		fmt.Println("open error:", err)
	}
	defer f.Close()
	yamlInput, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println("read error:", err)
	}
	var objInput map[interface{}]interface{}
	yamlError := yaml.UnmarshalStrict(yamlInput, &objInput)
	if yamlError != nil {
		fmt.Println("unmarshal error:", err)
	}
	container := map[string]interface{}{}
	evaluated, err := NewArgument(objInput["main"]).Evaluate(container)
	fmt.Println("container", container)
	fmt.Println("evaluated", evaluated)
	fmt.Println("error", err)
	select {}
}
