package main

import (
	"fmt"
	"html/template"
	"os"

	"gopkg.in/yaml.v2"
)

func buildFromGlob(file, glob string, data map[string]interface{}) error {
	out, err := os.Create(file)
	defer out.Close()
	if err != nil {
		return err
	}
	t, err := template.ParseGlob(glob)
	if err != nil {
		return err
	}
	return t.Execute(out, data)
}

func buildFromFile(file, tmpl string, data map[string]interface{}) error {
	out, err := os.Create(file)
	defer out.Close()
	if err != nil {
		return err
	}
	t, err := template.ParseFiles(tmpl)
	if err != nil {
		return err
	}
	return t.Execute(out, data)
}

func main() {
	data := map[string]interface{}{}
	buf, err := os.ReadFile("data.yml")
	if err != nil {
		return
	}
	err = yaml.Unmarshal(buf, &data)
	if err != nil {
		return
	}
	entries, err := os.ReadDir("tmpl")
	if err != nil {
		return
	}
	for _, entry := range entries {
		fmt.Println(entry.Name())
	}

	// TODO detect tmpl and fail if not found
	// TODO iterate over tmpl directory and
	//    if directory build a file matching the name of the directory
	//    from the files in the directory
	//    or,
	//    if a file just build from that file
	//    make sure to detect hte template/html or template/text based on
	//    suffix
}
