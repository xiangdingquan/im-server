package dataobject

type BlogTopicMappingsDo struct {
	TopicId  int32 `db:"topic_id"`
	MomentId int32 `db:"moment_id"`
}
