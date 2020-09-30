package push

// PushWorkflowHistory
type PushWorkflowHistory struct {
	ProjectId   string   `json:"projectId" comment:"项目id" optional:"false"`
	ProcInstId  string   `json:"procInstId" comment:"流程id" optional:"false"`
	Title       string   `json:"title" comment:"流程名称" optional:"false"`
	TaskId      string   `json:"taskId" comment:"任务id" optional:"false"`
	Name        string   `json:"name" comment:"任务名称" optional:"false"`
	ExecutionId string   `json:"executionId" comment:"执行流id" optional:"false"`
	Value       string   `json:"value" comment:"推送内容" optional:"false"`
	PushCode    string   `json:"pushCode" comment:"推送code" optional:"false"`
	Kind        string   `json:"kind" comment:"推送类型,normal,start等" optional:"false"`
	UserIds     []string `json:"userIds" comment:"用户账号列表" optional:"false"`
	OperatorId  string   `json:"operatorId" comment:"operatorId" optional:"false"`
}
