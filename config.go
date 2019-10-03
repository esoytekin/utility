package utility

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"os/user"
	"path"
)

var open = os.Open

// ReadConfiguration opens file in path $HOME/.<appname>/config.json
// and loads properties into config var
func ReadConfiguration(appname string, config interface{}) error {
	usr, err := user.Current()

	if err != nil {
		return err
	}

	finalPath := path.Join(usr.HomeDir, "."+appname, "config.json")
	jsonFile, err := open(finalPath)

	if err != nil {
		return err
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &config)
	return nil
}
