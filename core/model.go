package core

import "go.uber.org/zap"

type Handler interface {
	concurrent()
	sync()
}

// enum for error types ...
const (
	Success int = iota
	Failed
	Error
)

type CommandOut struct {
	StdOutput []byte
	StdError  []byte
}

type Commands struct {
	Commands JsonFile `json:"commands"`
}

// JsonFile represents the structure of the json file...
type JsonFile struct {
	Concurrent bool        `json:"concurrent"`
	List       []JsonInner `json:"list"`
}

type HandlerObjects struct {
	Logger *zap.SugaredLogger
	Config JsonInner
}

// JsonInner represents the structure of the innser json ...
type JsonInner struct {
	Directory  string      `json:"directory"`
	Concurrent bool        `json:"concurrent"`
	List       []ListItems `json:"list"`
}

type ListItems struct {
	CMD string `json:"cmd"`
}
