package main

import (
	"fmt"
	"io/ioutil"

	"github.com/ghodss/yaml"
)

func main() {
	yamlFile, err := ioutil.ReadFile("example1.yaml")

	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}

	y, err := yaml.JSONToYAML(yamlFile)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	fmt.Println("Yaml File")
	fmt.Println(string(y))

	j2, err := yaml.YAMLToJSON(y)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	fmt.Println("Coverted to JSON")
	fmt.Println(string(j2))
}
