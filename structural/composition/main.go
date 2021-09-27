package main

// Composite is a structural design pattern
//that allows composing objects into a tree-like structure and work with it
//as if it was a singular object.
// 树形结构，递归使用
// 例子：文件夹


func main() {
	file1 := &file{name: "File1"}
	file2 := &file{name: "File2"}
	file3 := &file{name: "File3"}

	folder1 := &folder{
		name: "Folder1",
	}

	folder1.add(file1)

	folder2 := &folder{
		name: "Folder2",
	}
	folder2.add(file2)
	folder2.add(file3)
	folder2.add(folder1)

	folder2.search("rose")
}