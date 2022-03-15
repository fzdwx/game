package config

import (
	"game/util"
)

type UniqueTaskCfg struct {
	TaskId    int
	SortType  int
	OpenLevel int
	TaskType  int
	Condition int
}

var UniqueTaskMap map[int]*UniqueTaskCfg

func init() {
	UniqueTaskMap = make(map[int]*UniqueTaskCfg)
	var taskArray []*UniqueTaskCfg
	err := util.CsvParse("../resource/UniqueTask.csv", &taskArray)
	if err != nil {
		panic(err)
	}
	for i, cfg := range taskArray {
		UniqueTaskMap[i] = cfg
	}

}
