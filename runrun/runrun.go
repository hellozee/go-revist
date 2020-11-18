package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"os/exec"
)

type Config struct {
	Jobs []string `json:jobs`
}

func main() {
	file, err := os.Open("runfile.json")

	if err != nil {
		panic(err)
	}

	bytes, _ := ioutil.ReadAll(file)

	file.Close()

	var config Config
	err = json.Unmarshal(bytes, &config)

	if err != nil {
		panic(err)
	}

	for _, val := range config.Jobs {
		cmd := exec.Command(os.Getenv("SHELL"), "-c", val)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Stdin = os.Stdin
		err = cmd.Run()
		if err != nil {
			panic(err)
		}
	}

}
