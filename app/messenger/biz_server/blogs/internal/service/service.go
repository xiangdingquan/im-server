package service

import (
	sync_client "open.chat/app/messenger/sync/client"
	blog_facade "open.chat/app/service/biz_service/blog/facade"
	user_client "open.chat/app/service/biz_service/user/client"
	wallet_facade "open.chat/app/service/biz_service/wallet/facade"
)

type Service struct {
	user_client.UserFacade
	blog_facade.BlogFacade
	wallet_facade.WalletFacade
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
	s.BlogFacade, err = blog_facade.NewBlogFacade("local")
	checkErr(err)
	s.WalletFacade, err = wallet_facade.NewWalletFacade("local")
	checkErr(err)
	sync_client.New()
	return s
}

func (s *Service) Close() {
}
