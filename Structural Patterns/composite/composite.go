package main

import "fmt"

type Component interface {
	search(string)
}

type Folder struct {
	components []Component
	name       string
}

func (f *Folder) search(keyword string) {
	fmt.Printf("SEARCHING FOR %s IN FOLDER %s\n", keyword, f.name)

	for _, composite := range f.components {
		composite.search(keyword)
	}
}

func (f *Folder) add(c Component) {
	f.components = append(f.components, c)
}

type File struct {
	name string
}

func (f *File) search(keyword string) {
	fmt.Printf("SEARCHING FOR %s IN FILE %s\n", keyword, f.name)
}

func (f *File) getName() string {
	return f.name
}

func main() {

	file1 := &File{name: "file1"}
	file2 := &File{name: "file2"}
	file3 := &File{name: "file3"}

	folder1 := &Folder{name: "folder1"}
	folder1.add(file1)

	folder2 := &Folder{name: "folder2"}

	folder2.add(file2)
	folder2.add(file3)
	folder2.add(folder1)

	folder2.search("Nikhil")
}
