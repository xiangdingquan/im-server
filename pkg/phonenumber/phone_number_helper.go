package phonenumber

import (
	"errors"
	"fmt"
	"strings"

	"github.com/nyaruka/phonenumbers"
)

type PhoneNumberHelper struct {
	*phonenumbers.PhoneNumber
}

func IsNotPhoneUser(number string) bool {
	if len(number) == 13 {
		return strings.HasPrefix(number, "86100")
	} else if len(number) == 14 {
		return strings.HasPrefix(number, "+86100")
	}
	return false
}

func IsVirtualUser(number string) bool {
	if len(number) == 13 {
		return strings.HasPrefix(number, "86101")
	} else if len(number) == 14 {
		return strings.HasPrefix(number, "+86101")
	}
	return false
}

func MakePhoneNumberHelper(number, region string) (*PhoneNumberHelper, error) {
	var (
		pNumber *phonenumbers.PhoneNumber
		err     error
	)

	if number == "" {
		return nil, errors.New("empty phone number")
	}

	if region == "" && number[:1] != "+" {
		number = "+" + number
	}

	// check phone invalid
	pNumber, err = phonenumbers.Parse(number, region)
	if err != nil {
		err = fmt.Errorf("parse phone number %s err: %v", number, err)
	} else {
		if !phonenumbers.IsValidNumber(pNumber) {
			if !IsNotPhoneUser(number) && !IsVirtualUser(number) {
				err = fmt.Errorf("invalid phone number: %s - %v", number, pNumber)
			}
		}
	}

	if err != nil {
		return nil, err
	} else {
		return &PhoneNumberHelper{pNumber}, nil
	}
}

func (p *PhoneNumberHelper) GetNormalizeDigits() string {
	// DB store normalize phone number
	return phonenumbers.NormalizeDigitsOnly(phonenumbers.Format(p.PhoneNumber, phonenumbers.E164))
}

func (p *PhoneNumberHelper) GetRegionCode() string {
	return phonenumbers.GetRegionCodeForNumber(p.PhoneNumber)
}

// Check number
func CheckAndGetPhoneNumber(number string) (phoneNumber string, err error) {
	var (
		pNumber *PhoneNumberHelper
	)

	pNumber, err = MakePhoneNumberHelper(number, "")
	if err != nil {
		return
	}

	return pNumber.GetNormalizeDigits(), nil
}
