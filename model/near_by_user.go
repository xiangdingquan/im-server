package model

import "strconv"

type NearByUser struct {
	Id         int32
	Distance   float64
	Expiration int32
}

func NearByUserFromString(id string, distance string) (*NearByUser, error) {
	out := NearByUser{}

	if err := out.setId(id); err != nil {
		return nil, err
	}
	if err := out.setDistance(distance); err != nil {
		return nil, err
	}

	return &out, nil
}

func (n *NearByUser) setId(s string) error {
	id, err := strconv.Atoi(s)
	if err == nil {
		n.Id = int32(id)
	}
	return err
}

func (n *NearByUser) setDistance(s string) error {
	d, err := strconv.ParseFloat(s, 64)
	if err == nil {
		n.Distance = d
	}
	return err
}
