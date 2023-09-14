package dataobject

type UserContactsDO struct {
	Id               int32  `db:"id"`
	OwnerUserId      int32  `db:"owner_user_id"`
	ContactUserId    int32  `db:"contact_user_id"`
	ContactPhone     string `db:"contact_phone"`
	ContactFirstName string `db:"contact_first_name"`
	ContactLastName  string `db:"contact_last_name"`
	Mutual           int8   `db:"mutual"`
	IsDeleted        int8   `db:"is_deleted"`
	Date2            int32  `db:"date2"`
}
