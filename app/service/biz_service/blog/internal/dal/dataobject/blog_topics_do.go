package dataobject

type BlogTopicsDO struct {
	Id      int32  `db:"id"`
	Name    string `db:"name"`
	Ranking int32  `db:"ranking"`
}
