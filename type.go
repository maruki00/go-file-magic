package github.com/maruki00/go-magic


type FileType struct {
	Description string
	Ext string
	Offset uint8
}

var Magics = map[[]byte]FileType {
	[]byte{} : FileType{}
}
