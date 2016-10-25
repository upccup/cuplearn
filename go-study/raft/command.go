package main

import (
	"github.com/goraft/raft"
)

type WriteCommand struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func NewWriteCommand(key, value string) *WriteCommand {
	return &WriteCommand{
		Key:   key,
		Value: value,
	}
}

func (c *WriteCommand) CommandName() string {
	return "write"
}

func (c *WriteCommand) Apply(server raft.Server) (interface{}, error) {
	db := server.Context().(*DB)
	db.Put(c.Key, c.Value)
	return nil, nil
}
