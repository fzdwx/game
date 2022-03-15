package game

const (
	TASK_STATE_INIT = iota
	TASK_STATE_DOINT
	TASK_STATE_FINISH
)

type ModUniqueTask struct {
	MyTaskInfo map[int]*TaskInfo
}

type TaskInfo struct {
	TaskId int
	State  int
}

// IsFinish  判断某个任务是否完成
func (this *ModUniqueTask) IsFinish(taskId int) bool {
	if task, ok := this.MyTaskInfo[taskId]; ok {
		return task.State == TASK_STATE_FINISH
	} else {
		return false
	}
}
