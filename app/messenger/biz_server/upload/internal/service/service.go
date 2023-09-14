package service

import (
	dfs_facade "open.chat/app/service/dfs/facade"
	_ "open.chat/app/service/dfs/facade/dfs"
	media_client "open.chat/app/service/media/client"
)

type Service struct {
	dfs_facade.DfsFacade
}

func New(dataPath string) *Service {
	if dataPath == "" {
		dataPath = "/opt/nbfs"
	}

	var err error

	s := &Service{}
	s.DfsFacade, err = dfs_facade.NewDfsFacade("dfs")
	if err != nil {
		panic(err)
	}

	media_client.New()
	return s
}

func (s *Service) Close() {

}
