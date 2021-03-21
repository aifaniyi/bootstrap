package main

import "io/ioutil"

func writeFile(content, filename string) error {
	err := ioutil.WriteFile(filename, []byte(content), 0755)
	if err != nil {
		return err
	}
	return nil
}
