package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"

	mcr "github.com/ComptonMSHI/hcris-tools"
)

func main() {
	err := mcr.LoadConfig("config.yaml")

	if err != nil {
		mcr.Error(mcr.AppMsg{"Fatal", "LoadConfig", "Could not load configuration.", 11, err}, true)
	}

	mcr.Debug(fmt.Sprintf("Starting with output to %s.", mcr.AppConfig.Settings.Output))

	if mcr.AppConfig.Source.Download {
		files := mcr.DownloadFiles(
			mcr.AppConfig.Source.Url,
			mcr.AppConfig.Source.Path,
			mcr.AppConfig.Source.TagAttr,
			mcr.AppConfig.Source.List,
			mcr.AppConfig.Source.TagMatch,
			mcr.AppConfig.Store.Prefix,
			mcr.AppConfig.Store.Ext)

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

	if mcr.AppConfig.Store.Extract {
		fileList, err := ioutil.ReadDir(mcr.GetDataFolder())
		mcr.Check(err)

		rx := regexp.MustCompile("(?i)zip$")

		for _, file := range fileList {
			// matched, err := rx.MatchString(file.Name())
			// mcr.Check(err)

			if rx.MatchString(file.Name()) {
				mcr.Debug(fmt.Sprintf("Extracting and preparing CSV files from %s into the store directory. (%s)", file.Name(), mcr.GetOutputFolder()))
				mcr.ExtractToFolder(mcr.GetDataFolder()+"/"+file.Name(), mcr.GetOutputFolder())
			}
		}

		mcr.CheckExtractFiles()
		mcr.Check(err)
	}

	mcr.Debug("Finished")
}
