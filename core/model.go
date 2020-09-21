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

// CommandOut ...
type CommandOut struct {
	StdOutput []byte
	StdError  []byte
}

// Commands ...
type Commands struct {
	Commands JsonFile `json:"commands"`
}

// JsonFile ...
type JsonFile struct {
	Concurrent bool        `json:"concurrent"`
	List       []JsonInner `json:"list"`
}

// HandlerObjects ...
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

// ListItems ...
type ListItems struct {
	CMD string `json:"cmd"`
}
