package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/readytotouch/readytotouch/internal/protos/organizers"
)

func main() {
	// review("./public/logos/original/")
	review("./public/logos/adapted/")
}

func review(dir string) {
	files, err := os.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		if file.IsDir() {
			panic("dir unexpected")
		}

		fileName := file.Name()
		if fileName == ".gitkeep" {
			continue
		}

		ext := filepath.Ext(file.Name())

		alias := strings.TrimSuffix(fileName, ext)

		switch ext {
		case ".png", ".jpg", ".jpeg":
		default:
			panic(fmt.Sprintf("unexpected ext: %s", ext))
		}

		_ = organizers.CompanyAliasMap[alias]
	}

	aliasImageMap, err := fetchAliasImageMap("./aliases_right.txt")
	if err != nil {
		panic(err)
	}
	_ = aliasImageMap
}

func fetchAliasImageMap(filename string) (map[string]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Ініціалізація мапи
	dataMap := make(map[string]string)

	// Читання рядків з файлу
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")

		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid line: %s", line)
		}

		key := parts[0]
		value := parts[1]

		if _, ok := dataMap[key]; ok {
			return nil, fmt.Errorf("duplicate key: %s", key)
		}

		dataMap[key] = value
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return dataMap, nil
}
