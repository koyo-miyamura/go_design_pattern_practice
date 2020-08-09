package main

import "fmt"

// DirectoryEntity is XXX
type DirectoryEntity interface {
	Remove()
}

// File is XXX
type File struct {
	name string
}

// NewFile is XXX
func NewFile(name string) *File {
	return &File{
		name: name,
	}
}

// Remove is XXX
func (f *File) Remove() {
	fmt.Printf("%s を削除しました\n", f.name)
}

// Directory is XXX
type Directory struct {
	list []DirectoryEntity
	name string
}

// NewDirectory is XXX
func NewDirectory(name string) *Directory {
	return &Directory{
		list: []DirectoryEntity{},
		name: name,
	}
}

// Add is XXX
func (d *Directory) Add(entity DirectoryEntity) {
	d.list = append(d.list, entity)
}

// Remove is XXX
func (d *Directory) Remove() {
	for _, entity := range d.list {
		entity.Remove()
	}
	fmt.Printf("%sを削除しました\n", d.name)
}

func main() {
	file1 := NewFile("file1")
	file2 := NewFile("file2")
	file3 := NewFile("file3")
	file4 := NewFile("file4")
	dir1 := NewDirectory("dir1")
	dir1.Add(file1)
	dir2 := NewDirectory("dir2")
	dir2.Add(file2)
	dir2.Add(file3)
	dir1.Add(dir2)
	dir1.Add(file4)

	dir1.Remove()
}
