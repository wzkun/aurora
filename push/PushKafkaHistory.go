package push

// PushKafkaHistory
type PushKafkaHistory struct {
	ProjectId  string `json:"projectId" comment:"项目id" optional:"false"`
	MappingKey string `json:"mappingKey" comment:"MappingKey" optional:"false"`
	DataClass  string `json:"dataClass" comment:"DataClass" optional:"false"`
	OperatorId string `json:"operatorId" comment:"操作人id" optional:"false"`
	Item       string `json:"item" comment:"Item" optional:"false"`
	Operation  int    `json:"operation" comment:"Operation" optional:"false"`
}
