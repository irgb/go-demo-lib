package common

import (
    "io/ioutil"
    "gopkg.in/yaml.v2"
)

// generated by https://mengzhuo.github.io/yaml-to-go/
type T struct {
    Log  struct {
        Level            string   `yaml:"level"`
        Encoding         string   `yaml:"encoding"`
        OutputPaths      []string `yaml:"outputPaths"`
        ErrorOutputPaths []string `yaml:"errorOutputPaths"`
    } `yaml:"log"`
}

var Config T

func init() {
    content, err := ioutil.ReadFile("config.yaml")
    if err != nil {
        panic("read config file failed")
    }
    err = yaml.Unmarshal(content, &Config)
    if err != nil {
        panic("unmarshal config file failed")
    }
}
