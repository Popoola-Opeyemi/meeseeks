package core

// JsonFile represents the structure of the json file...
type JsonFile struct {
	Commands   string      `json:"commands"`
	Concurrent bool        `json:"concurrent"`
	List       []JsonInner `json:"list"`
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
