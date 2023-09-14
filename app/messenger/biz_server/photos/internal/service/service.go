package service

import (
	sync_client "open.chat/app/messenger/sync/client"
	user_client "open.chat/app/service/biz_service/user/client"
	media_client "open.chat/app/service/media/client"
)

type Service struct {
	user_client.UserFacade
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func New() *Service {
	var (
		err error
		s   = new(Service)
	)
	s.UserFacade, err = user_client.NewUserFacade("local")
	checkErr(err)

	sync_client.New()
	media_client.New()

	return s
}

func (s *Service) Close() {

}
