package dataobject

type PredefinedUsersDO struct {
	Id               int32  `db:"id"`
	Phone            string `db:"phone"`
	FirstName        string `db:"first_name"`
	LastName         string `db:"last_name"`
	Username         string `db:"username"`
	Code             string `db:"code"`
	Verified         int8   `db:"verified"`
	RegisteredUserId int32  `db:"registered_user_id"`
	Deleted          int8   `db:"deleted"`
}
