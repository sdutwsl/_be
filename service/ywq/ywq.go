package ywq

import (
	"_be/model/ywq"
)

func GetAllRecords() []ywq.YWQ {
	return ywq.GetAllRecords()
}

func AddRecord() int {
	return ywq.AddRecord()
}
