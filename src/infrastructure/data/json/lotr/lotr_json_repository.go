package lotr

import jsoniter "github.com/json-iterator/go"

type lotrJSONRepository struct {
}

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func NewLOTRJSONRepository() *lotrJSONRepository {
	return &lotrJSONRepository{}
}
