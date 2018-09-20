package nec

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func loadSettings(path string) (*Settings, error) {
	settings := &Settings{}

	f, err := os.Open(path)
	if err != nil {
		return settings, fmt.Errorf("cannot open file: %s", err)
	}
	defer f.Close()

	b, err := ioutil.ReadAll(f)
	if err != nil {
		return settings, fmt.Errorf("cannot read file: %s", err)
	}

	if err := json.Unmarshal(b, &settings); err != nil {
		return settings, fmt.Errorf("cannot unmarshal: %s", err)
	}

	for k, v := range settings.Commands {
		if v.OnErr != "exit" && v.OnErr != "continue" {
			return settings, fmt.Errorf("%s \"onerror\" value must be \"exit\" or \"continue\"", k)
		}
	}

	return settings, nil
}
