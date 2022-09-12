package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)


func addr() string {
	return fmt.Sprintf("%s:%s", app.Host, app.Port)
} // addr


func readConf() {

	_, err := os.Stat(*conf)

	if err != nil || os.IsNotExist(err) {
		log.Fatal(err)
	} else {

		buf, err := ioutil.ReadFile(*conf)

		if err != nil {
			log.Fatal(err)
		} else {

			err := json.Unmarshal(buf, &app)

			if err != nil {
				log.Println(err)
			}

		}
	
	}

} // readConf
