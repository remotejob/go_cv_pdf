package toml_parser

import (
	"github.com/remotejob/go_cv_pdf/domains"
	//	"fmt"
	"io/ioutil"
	"os"

	"github.com/naoina/toml"
)

func Parse(file string) domains.Config {

	var cv domains.Config

	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	buf, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}

	if err := toml.Unmarshal(buf, &cv); err != nil {
		panic(err)
	}

	return cv

}

func ParseWorkPlaces(file string) domains.Job {

	var job domains.Job

	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	buf, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}
	//    var config tomlConfig
	if err := toml.Unmarshal(buf, &job); err != nil {
		panic(err)
	}

	return job

}
