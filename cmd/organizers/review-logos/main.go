package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"sort"
	"strings"

	"github.com/readytotouch/readytotouch/internal/domain"
	"github.com/readytotouch/readytotouch/internal/generated/organizers"
	"github.com/readytotouch/readytotouch/internal/organizer/db"
)

func main() {
	var (
		companies = db.CloneCompanies()
	)

	_ = companies

	/*
		review("./public/logos-v0/original/")
		review("./public/logos-v0/adapted/")
	*/
	/*
		prepareLogosV1Mapping(companies)
	*/

	review("./public/logos-v1/adapted/")
}

func review(dir string) {
	files, err := os.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	imageExistsMap := make(map[string]string)

	for _, file := range files {
		if file.IsDir() {
			panic("dir unexpected")
		}

		fileName := file.Name()
		if fileName == ".gitkeep" {
			continue
		}

		ext := filepath.Ext(file.Name())

		switch ext {
		case ".png", ".jpg", ".jpeg":
		default:
			panic(fmt.Sprintf("unexpected ext: %s", ext))
		}

		alias := strings.TrimSuffix(fileName, ext)

		imageExistsMap[fileName] = alias
	}

	aliasImageMap, err := fetchAliasImageMap("./public/logos-v1/mapping.txt")
	if err != nil {
		panic(err)
	}

	assertCorrectnessAliasImageMap(aliasImageMap, organizers.CompanyAliasToCodeMap, imageExistsMap)

	// sortAndStoreAliasImageMap(aliasImageMap, "./public/logos/mapping_sorted.txt")
}

func assertCorrectnessAliasImageMap(aliasImageMap map[string]string, aliasMap map[string]int64, imageExistsMap map[string]string) {
	for alias, image := range aliasImageMap {
		if _, ok := aliasMap[alias]; !ok {
			panic(fmt.Sprintf("alias not found: %s", alias))
		}

		if image == "" {
			// can be empty

			continue
		}

		if _, ok := imageExistsMap[image]; !ok {
			panic(fmt.Sprintf("image not found: %s", image))
		}
	}

	allowDuplicateMap := map[[2]string]struct{}{
		{"applied-systems", "applied.png"}:        {},
		{"applied-systems-canada", "applied.png"}: {},
	}

	duplicateImageMap := make(map[string][]string)
	for alias, image := range aliasImageMap {
		if _, ok := allowDuplicateMap[[2]string{alias, image}]; ok {
			continue
		}

		duplicateImageMap[image] = append(duplicateImageMap[image], alias)
	}

	for image, aliases := range duplicateImageMap {
		if len(aliases) > 1 {
			panic(fmt.Sprintf("duplicate image: %s, aliases: %v", image, aliases))
		}
	}

	assertAllImageMapped(imageExistsMap, aliasImageMap)
}

func assertAllImageMapped(imageExistsMap map[string]string, aliasImageMap map[string]string) {
	mappedImageExistsMap := make(map[string]struct{})
	for _, image := range aliasImageMap {
		mappedImageExistsMap[image] = struct{}{}
	}

	for image := range imageExistsMap {
		if _, ok := mappedImageExistsMap[image]; !ok {
			panic(fmt.Sprintf("image not mapped: %s", image))
		}
	}
}

func fetchAliasImageMap(filename string) (map[string]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	dataMap := make(map[string]string)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")

		var (
			key   string
			value string
		)

		if len(parts) == 1 {
			key = parts[0]
			value = ""
		} else if len(parts) == 2 {
			key = parts[0]
			value = parts[1]
		} else {
			return nil, fmt.Errorf("invalid line: %s", line)
		}

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

func sortAndStoreAliasImageMap(imageMap map[string]string, dst string) {
	rows := make([][2]string, 0, len(imageMap))
	for alias, image := range imageMap {
		rows = append(rows, [2]string{alias, image})
	}

	sort.Slice(rows, func(i, j int) bool {
		return rows[i][0] < rows[j][0]
	})

	file, err := os.Create(dst)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	for _, row := range rows {
		_, err := file.WriteString(fmt.Sprintf("%s %s\n", row[0], row[1]))
		if err != nil {
			panic(err)
		}
	}
}

func prepareLogosV1Mapping(companies []domain.CompanyProfile) {
	aliases := make([]string, 0, len(companies))
	for _, company := range companies {
		aliases = append(aliases, company.LinkedInProfile.Alias)
	}
	slices.Sort(aliases)

	err := os.WriteFile("./public/logos-v1/mapping.txt", []byte(strings.Join(aliases, " \n")), 0644)
	if err != nil {
		panic(err)
	}
}
