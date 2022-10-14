package libs

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type Resources struct {
	Name string   `yaml:"name"`
	IP   []string `yaml:"IP"`
}

func ParseResources(path string) []Resources {
	var resources []Resources
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		log.Println(err)
	}

	err = yaml.Unmarshal(yamlFile, &resources)
	if err != nil {
		log.Println(err)
	}

	// fmt.Println("xxxxxx")
	// for _, sign := range signs {
	// 	fmt.Println(sign)
	// }
	return resources
}
