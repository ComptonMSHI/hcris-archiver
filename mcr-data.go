package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"

	mcr "github.com/dcinformatics/datatools"
)

func main() {
	err := mcr.LoadConfig("config.yaml")

	if err != nil {
		mcr.Error(mcr.AppMsg{"Fatal", "LoadConfig", "Could not load configuration.", 11, err}, true)
	}

	mcr.Debug(fmt.Sprintf("Starting with output to %s.", mcr.AppConfig.Settings.Output))

	if mcr.AppConfig.Input.Download {
		files := mcr.DownloadFiles(
			mcr.AppConfig.Input.Url,
			mcr.AppConfig.Input.Path,
			mcr.AppConfig.Input.TagAttr,
			mcr.AppConfig.Input.List,
			mcr.AppConfig.Input.TagMatch,
			mcr.AppConfig.Output.Prefix,
			mcr.AppConfig.Output.Ext)

		mcr.Debug(strings.Join(files, ", "))

		if mcr.AppConfig.Settings.Output == "database" || mcr.AppConfig.Settings.Output == "sql" {
			for _, file := range files {
				mcr.Debug(file)
				mcr.ExtractFile(file)
			}
		}
	} else {
		mcr.Debug("Skipping Download per Config")
	}

	if mcr.AppConfig.Output.Extract {
		fileList, err := ioutil.ReadDir(mcr.GetInputFolder())
		mcr.Check(err)

		rx := regexp.MustCompile("(?i)zip$")

		for _, file := range fileList {
			// matched, err := rx.MatchString(file.Name())
			// mcr.Check(err)

			if rx.MatchString(file.Name()) {
				mcr.Debug(fmt.Sprintf("Extracting and preparing CSV files from %s into the store directory. (%s)", file.Name(), mcr.GetOutputFolder()))
				mcr.ExtractToFolder(mcr.GetInputFolder()+"/"+file.Name(), mcr.GetOutputFolder())
			}
		}

		mcr.CheckAndMoveExtractFiles()
		mcr.Check(err)

		os.Remove(mcr.GetOutputFolder())
		mcr.Check(err)
	}

	mcr.Debug("Finished")
}
