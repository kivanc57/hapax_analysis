package main

import (
	"path/filepath"
	"strings"
	utils"hapax_analysis/utils"
)

func main() {
	absDirPath := "/home/kivanc57/Desktop/hapax_analysis/data"
	var sb strings.Builder

  dir  := utils.ReadDir(absDirPath)
	for _, f := range dir {
		file_name := f.Name()

		if strings.HasSuffix(file_name, ".txt") {
			absFilePath := filepath.Join(absDirPath, file_name)
			sb.WriteString(utils.ReadWriteText(absFilePath))
		}
	}
	corpus := sb.String()
	freqMap := utils.GetFreqMap(corpus, true)

	utils.WriteExcel(freqMap)
}

