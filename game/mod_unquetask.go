package game

const (
	TASK_STATE_INIT = iota
	TASK_STATE_DOINT
	TASK_STATE_FINISH
)

type ModUniqueTask struct {
	myTaskInfo map[int]*TaskInfo
	//mux        sync.RWMutex
}

type TaskInfo struct {
	TaskId int
	State  int
}

// IsFinish  判断某个任务是否完成
func (this *ModUniqueTask) IsFinish(taskId int) bool {
	task := this.GetTaskInfo(taskId)
	if task != nil {
		return task.State == TASK_STATE_FINISH
	}
	return false
}

func (this *ModUniqueTask) GetTaskInfo(taskId int) *TaskInfo {
	//this.mux.RLock()
	info := this.myTaskInfo[taskId]
	//this.mux.RUnlock()
	return info
}
