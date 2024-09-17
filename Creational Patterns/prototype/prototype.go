package main

import "fmt"

type Inode interface {
	print(string)
	clone() Inode
}

type File struct {
	name string
}

func (f *File) print(indentation string) {
	fmt.Println(indentation + f.name)
}

func (f *File) clone() Inode {
	return &File{name: f.name + "_clone"}
}

type Folder struct {
	children []Inode
	name     string
}

func (f *Folder) print(indentation string) {
	fmt.Println(indentation + f.name)

	for _, i := range f.children {
		i.print(indentation + indentation)
	}
}

func (f *Folder) clone() Inode {

	cloneFolder := &Folder{name: f.name + "_clone"}

	var tempChildren []Inode

	for _, i := range f.children {
		copy := i.clone()
		tempChildren = append(tempChildren, copy)
	}

	cloneFolder.children = tempChildren

	return cloneFolder

}

func main() {
	file1 := &File{name: "File1"}
	file2 := &File{name: "File2"}
	file3 := &File{name: "File3"}
	file4 := &File{name: "File4"}

	folder1 := &Folder{
		children: []Inode{file1},
		name:     "Folder1",
	}

	folder2 := &Folder{
		children: []Inode{file2, file3, file4},
		name:     "Folder2",
	}

	fmt.Println("\nFolder1 Structure")
	folder1.print(" ")

	fmt.Println("\nFolder2 Structure")
	folder2.print(" ")

	fmt.Println("\nFolder2 Clone")
	cloneFolder2 := folder2.clone()
	cloneFolder2.print(" ")

}
