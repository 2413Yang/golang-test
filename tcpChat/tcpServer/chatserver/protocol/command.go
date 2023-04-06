package protocol

import "errors"

var (
	UnKonwCommand = errors.New("Unkonw command")
)
type SendCmd struct{
	Message string
}
type NameCmd struct{
	Name string
}
type MessCmd struct{
	Name string
	Message string
}