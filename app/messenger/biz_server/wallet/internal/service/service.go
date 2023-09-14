package service

import (
	wallet_facade "open.chat/app/service/biz_service/wallet/facade"
)

type Service struct {
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
	s.WalletFacade, err = wallet_facade.NewWalletFacade("local")
	checkErr(err)
	return s
}

func (s *Service) Close() {
}
