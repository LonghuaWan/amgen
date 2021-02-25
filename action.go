package amgen

import (
	"gopkg.in/urfave/cli.v2"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"

	"gopkg.in/yaml.v2"
)

func MgoAction(context *cli.Context) error {
	configPath := context.String("config-file")
	configFilePath := context.String("config-file-path")
	if configFilePath != "" {
		files, _ := ioutil.ReadDir(configFilePath)
		for _, onefile := range files {
			 fileParts := strings.Split(onefile.Name(), ".")
			 if fileParts[len(fileParts) - 1] == "yaml" {
				 bytes, err := ioutil.ReadFile(path.Join(configFilePath, onefile.Name()))
				 if err != nil {
					 log.Fatalf("[ERROR] read config file error: %s\n", err)
					 return err
				 }

				 mg := new(ModelGenerator)
				 mg.ConfigName = path.Base(configPath)
				 if err := yaml.Unmarshal(bytes, mg); err != nil {
					 log.Fatalf("[ERROR] unmarshal yaml error: %s\n", err)
					 return err
				 }

				 response, err := TemplateBoxes.FindString("cmgo.tmpl")
				 if err != nil {
					 log.Fatal(err)
					 return err
				 }

				 t, err := template.New("cmgo").Funcs(template.FuncMap{
					 "ToLower":     strings.ToLower,
					 "SnakeString": SnakeString,
				 }).Parse(response)
				 if err != nil {
					 log.Fatalf("[ERROR] parse template files error: %s\n", err)
					 return err
				 }

				 filename := strings.Replace(onefile.Name(), path.Ext(onefile.Name()), ".mg.go", 1)
				 mg.FileName = filename
				 fp, err := os.Create(filename)
				 if err != nil {
					 log.Fatalf("[ERROR] create %s error: %s\n", filename, err)
					 return err
				 }
				 defer fp.Close()

				 if err := t.Execute(fp, mg); err != nil {
					 log.Fatalf("[ERROR] execute template error: %s\n", err)
					 return err
				 }
			 }
		}
	} else if configPath != "" {
		bytes, err := ioutil.ReadFile(configPath)
		if err != nil {
			log.Fatalf("[ERROR] read config file error: %s\n", err)
			return err
		}

		mg := new(ModelGenerator)
		mg.ConfigName = path.Base(configPath)
		if err := yaml.Unmarshal(bytes, mg); err != nil {
			log.Fatalf("[ERROR] unmarshal yaml error: %s\n", err)
			return err
		}

		response, err := TemplateBoxes.FindString("cmgo.tmpl")
		if err != nil {
			log.Fatal(err)
			return err
		}

		t, err := template.New("cmgo").Funcs(template.FuncMap{
			"ToLower":     strings.ToLower,
			"SnakeString": SnakeString,
		}).Parse(response)
		if err != nil {
			log.Fatalf("[ERROR] parse template files error: %s\n", err)
			return err
		}

		filename := strings.Replace(configPath, path.Ext(configPath), ".mg.go", 1)
		mg.FileName = filename
		fp, err := os.Create(filename)
		if err != nil {
			log.Fatalf("[ERROR] create %s error: %s\n", filename, err)
			return err
		}
		defer fp.Close()

		if err := t.Execute(fp, mg); err != nil {
			log.Fatalf("[ERROR] execute template error: %s\n", err)
			return err
		}
	}
	return nil
}
