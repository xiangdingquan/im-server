package dataobject

type PollAnswerVotersDO struct {
	Id         int64  `db:"id"`
	PollId     int64  `db:"poll_id"`
	VoteUserId int32  `db:"vote_user_id"`
	Options    string `db:"options"`
	Option0    int8   `db:"option0"`
	Option1    int8   `db:"option1"`
	Option2    int8   `db:"option2"`
	Option3    int8   `db:"option3"`
	Option4    int8   `db:"option4"`
	Option5    int8   `db:"option5"`
	Option6    int8   `db:"option6"`
	Option7    int8   `db:"option7"`
	Option8    int8   `db:"option8"`
	Option9    int8   `db:"option9"`
	Date2      int64  `db:"date2"`
	Deleted    int8   `db:"deleted"`
}
