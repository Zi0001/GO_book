package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Передайте аргумент пожалуйста!")
		return
	}
	
	file := arguments[1]
	// Добавляем расширения .exe, .bat, .cmd если их нет
	if !hasExecutableExtension(file) {
		file = file + ".exe" // пробуем сначала с .exe
	}
	
	path := os.Getenv("PATH")
	pathSplit := filepath.SplitList(path)
	
	for _, directory := range pathSplit {
		fullPath := filepath.Join(directory, file)
		
		// Проверяем существование файла
		if fileExists(fullPath) {
			fmt.Println(fullPath)
			return
		}
		
		// Если изначально не было расширения, пробуем другие варианты
		if !hasExecutableExtension(arguments[1]) {
			tryExtensions := []string{".exe", ".bat", ".cmd", ".com"}
			for _, ext := range tryExtensions {
				fullPathWithExt := filepath.Join(directory, arguments[1]+ext)
				if fileExists(fullPathWithExt) {
					fmt.Println(fullPathWithExt)
					return
				}
			}
		}
	}
	
	fmt.Println("Файла нет в PATH")
}

func hasExecutableExtension(filename string) bool {
	ext := strings.ToLower(filepath.Ext(filename))
	return ext == ".exe" || ext == ".bat" || ext == ".cmd" || ext == ".com"
}

func fileExists(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return !info.IsDir() 
}