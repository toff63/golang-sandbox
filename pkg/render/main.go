package main

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"
)

func main() {
	templates := parseTemplates()
	fmt.Println("Render Hello:")
	templates["hello.gohtml"].Execute(os.Stdout, nil)
	fmt.Println("Render Greeting:")
	templates["greeting.gohtml"].Execute(os.Stdout, map[string]string{"Name": "John"})
}

func parseTemplates() map[string]*template.Template {
	// First create base layout
	base := baseLayout()

	// List common files to all templates
	layoutFiles := files("layout/*.gohtml")

	// List views depending on the layout
	viewFiles := files("view/*.gohtml")

	templates := make(map[string]*template.Template, 0)
	for _, file := range viewFiles {
		// Here the file name will be used as key to lookup for the template later
		fileName := filepath.Base(file)
		// Add the view to layout files so we have all the files on which depend the view
		files := append(layoutFiles, file)
		// Clone base layout template struct so it can be used for every view
		baseLayout, err := base.Clone()
		if err != nil {
			panic(fmt.Sprintf("Failed to clone base layout: %v", err))
		}
		// Parse all the files in one template
		templates[fileName] = template.Must(baseLayout.ParseFiles(files...))

	}
	return templates
}

func files(pattern string) []string {
	files, err := filepath.Glob(pattern)
	if err != nil {
		panic(fmt.Sprintf("Failed to find files following pattern %s: %v", pattern, err))
	}
	return files

}

func baseLayout() *template.Template {
	baseLayout := template.New("base")
	baseLayout, err := baseLayout.ParseFiles("layout/base.gohtml")
	if err != nil {
		panic(fmt.Sprintf("Failed to parse base layout: %v", err))

	}
	return baseLayout
}
