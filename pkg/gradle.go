package pkg

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

const (
	lineSplitSeparator = "\n"
	lineSearchString   = "versionCode"

	fileWalkStartPoint = "."

	gradleExtensionString = ".gradle"
	kotlinExtensionString = ".kts"
)

func getGradleFiles() []string {
	var paths []string
	err := filepath.Walk(fileWalkStartPoint, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		var extension = filepath.Ext(info.Name())
		if extension == gradleExtensionString || extension == kotlinExtensionString {
			paths = append(paths, path)
		}
		return nil
	})
	if err != nil {
		log.Println(err)
	}
	return paths
}

func UpdateVersionCode(versionCode int64, step int) {
	for _, path := range getGradleFiles() {
		input, err := ioutil.ReadFile(path)
		if err != nil {
			log.Fatalln(err)
		}
		lines := strings.Split(string(input), lineSplitSeparator)
		for i, line := range lines {
			if strings.Contains(line, lineSearchString) {
				var re = regexp.MustCompile(`[\d]+`)
				temp := re.ReplaceAllString(line, strconv.Itoa(int(versionCode)+step))
				lines[i] = temp
			}
		}
		output := strings.Join(lines, lineSplitSeparator)
		err = ioutil.WriteFile(path, []byte(output), 0644)
		if err != nil {
			log.Fatalln(err)
		}
	}
}
