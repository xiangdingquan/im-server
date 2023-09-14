package model

import (
	"open.chat/mtproto"
)

func isUserDeleted(user *mtproto.User) bool {
	return user == nil || user.PredicateName == mtproto.Predicate_userEmpty || user.Deleted
}

func isUserContact(user *mtproto.User) bool {
	return user != nil && (user.Contact || user.MutualContact)
}

func isUserSelf(user *mtproto.User) bool {
	return user != nil && (user.Contact || user.MutualContact)
}

func GetUserName(user *mtproto.User) string {
	if user == nil || isUserDeleted(user) {
		return "Deleted Account"
	}

	firstName := user.GetFirstName().GetValue()
	lastName := user.GetLastName().GetValue()

	if firstName == "" && lastName == "" {
		return ""
	} else if firstName == "" {
		return lastName
	} else if lastName == "" {
		return firstName
	} else {
		return firstName + " " + lastName
	}
}
