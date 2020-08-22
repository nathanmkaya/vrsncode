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

func getGradleFiles() []string {
	var paths []string
	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		var extension = filepath.Ext(info.Name())
		if extension == ".gradle" || extension == ".kts" {
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
		lines := strings.Split(string(input), "\n")
		for i, line := range lines {
			if strings.Contains(line, "versionCode") {
				var re = regexp.MustCompile(`[\d]+`)
				temp := re.ReplaceAllString(line, strconv.Itoa(int(versionCode)+step))
				lines[i] = temp
			}
		}
		output := strings.Join(lines, "\n")
		err = ioutil.WriteFile(path, []byte(output), 0644)
		if err != nil {
			log.Fatalln(err)
		}
	}
}
