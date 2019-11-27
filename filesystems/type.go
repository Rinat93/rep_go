package filesystems

type DataBody struct {
	body string
	serialize interface{}
	len int
}

type Files struct {
	Name string
	Size uint64
	Extension string // Расширения файла
	Hash uint64 // hash текст
	Data []*DataBody
}

type Directory struct {
	Name string
	files []Files
}
