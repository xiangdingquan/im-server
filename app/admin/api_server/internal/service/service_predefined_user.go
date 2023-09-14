package service

import (
	"context"
	"github.com/gogo/protobuf/types"
	"math/rand"
	"open.chat/app/admin/api_server/api"
	"open.chat/mtproto"
	"open.chat/pkg/log"
	"open.chat/pkg/random2"
)

func (s *Service) CreatePredefinedUser(ctx context.Context, i *api.CreatePredefinedUser) (r *mtproto.PredefinedUser, err error) {
	log.Debugf("createPredefinedUser - request: %s", i.DebugString())

	var (
		result mtproto.TLObject
		ok     bool
	)

	code := i.Code
	if code == "" {
	} else if len(code) != 5 {
		log.Errorf("createPredefinedUser - error: code's length != 5")
		err = mtproto.ErrPhoneCodeInvalid
		return
	}

	req := &mtproto.TLAccountCreatePredefinedUser{
		Phone:     i.Phone,
		FirstName: &types.StringValue{Value: i.FirstName},
		LastName:  &types.StringValue{Value: i.LastName},
		Username:  nil,
		Code:      code,
		Verified:  i.Verified,
	}
	if i.Username != "" {
		req.Username = &types.StringValue{Value: i.Username}
	}

	result, err = s.Invoke(ctx, 777000, req)
	if err != nil {
		log.Errorf("createPredefinedUser - error: %v", err)
		return
	} else if r, ok = result.(*mtproto.PredefinedUser); !ok {
		log.Errorf("createPredefinedUser - error: invalid PredefinedUser type")
		err = mtproto.ErrInternelServerError
		return
	}

	log.Debugf("createPredefinedUser - reply: %s", r.DebugString())
	return
}

func (s *Service) UpdatePredefinedUsername(ctx context.Context, i *api.UpdatePredefinedUsername) (r *mtproto.PredefinedUser, err error) {
	log.Debugf("updatePredefinedUsername - request: %s", i.DebugString())

	var (
		result mtproto.TLObject
		ok     bool
	)

	req := &mtproto.TLAccountUpdatePredefinedUsername{
		Phone:    i.Phone,
		Username: i.Username,
	}

	result, err = s.Invoke(ctx, 777000, req)
	if err != nil {
		log.Errorf("updatePredefinedUsername - error: %v", err)
		return
	} else if r, ok = result.(*mtproto.PredefinedUser); !ok {
		log.Errorf("updatePredefinedUsername - error: invalid PredefinedUser type")
		err = mtproto.ErrInternelServerError
		return
	}

	log.Debugf("updatePredefinedUsername - reply: %s", r.DebugString())
	return
}

func (s *Service) UpdatePredefinedProfile(ctx context.Context, i *api.UpdatePredefinedProfile) (r *mtproto.PredefinedUser, err error) {
	log.Debugf("updatePredefinedProfile - request: %s", i.DebugString())

	var (
		result mtproto.TLObject
		ok     bool
	)

	req := &mtproto.TLAccountUpdatePredefinedProfile{
		Phone:     i.Phone,
		FirstName: nil,
		LastName:  nil,
		About:     nil,
	}
	if i.FirstName != "" {
		req.FirstName = &types.StringValue{Value: i.FirstName}
	}
	if i.LastName != "" {
		req.LastName = &types.StringValue{Value: i.LastName}
	}
	if i.About != "" {
		req.About = &types.StringValue{Value: i.About}
	}

	result, err = s.Invoke(ctx, 777000, req)
	if err != nil {
		log.Errorf("updatePredefinedProfile - error: %v", err)
		return
	} else if r, ok = result.(*mtproto.PredefinedUser); !ok {
		log.Errorf("updatePredefinedProfile - error: invalid PredefinedUser type")
		err = mtproto.ErrInternelServerError
		return
	}

	log.Debugf("updatePredefinedProfile - reply: %s", r.DebugString())
	return
}

func (s *Service) UpdatePredefinedProfilePhoto(ctx context.Context, i *api.UpdatePredefinedProfilePhoto) (r *api.VoidRsp, err error) {
	log.Debugf("updatePredefinedProfilePhoto - request: %s", i.DebugString())

	var (
		result mtproto.TLObject
		ok     bool
		user   *mtproto.PredefinedUser
		r2     *mtproto.Photos_Photo
	)

	if user, err = s.GetPredefinedUser(ctx, &api.GetPredefinedUser{Phone: i.Phone}); err != nil {
		log.Errorf("updatePredefinedProfilePhoto - error: %v", err)
		return nil, err
	}

	req := &mtproto.TLPhotosUploadProfilePhoto{
		Constructor: mtproto.CRC32_photos_uploadProfilePhoto_4f32c098,
		File:        nil,
	}

	creatorId := rand.Int63()
	req.File, err = urlToInputFile(i.Photo, func(fileId int64, filePart int32, bytes []byte) error {
		s.DfsFacade.WriteFilePartData(ctx, creatorId, fileId, filePart, bytes)
		return nil
	})

	if err != nil {
		log.Errorf("updatePredefinedProfilePhoto - error: %v", err)
		return nil, err
	}

	result, err = s.Invoke2(ctx, creatorId, user.GetRegisteredUserId().GetValue(), req)
	if err != nil {
		log.Errorf("updatePredefinedProfilePhoto - error: %v", err)
		return
	} else if r2, ok = result.(*mtproto.Photos_Photo); !ok {
		log.Errorf("updatePredefinedProfile - error: invalid PredefinedUser type")
		err = mtproto.ErrInternelServerError
		return
	}

	_ = r2
	r = new(api.VoidRsp)
	log.Debugf("updatePredefinedProfilePhoto - reply: %s", r.DebugString())
	return
}

func (s *Service) UpdatePredefinedVerified(ctx context.Context, i *api.UpdatePredefinedVerified) (r *mtproto.PredefinedUser, err error) {
	log.Debugf("updatePredefinedVerified - request: %s", i.DebugString())

	var (
		result mtproto.TLObject
		ok     bool
	)

	req := &mtproto.TLAccountUpdatePredefinedVerified{
		Phone:    i.Phone,
		Verified: i.Verified,
	}

	result, err = s.Invoke(ctx, 777000, req)
	if err != nil {
		log.Errorf("updatePredefinedVerified - error: %v", err)
		return
	} else if r, ok = result.(*mtproto.PredefinedUser); !ok {
		log.Errorf("updatePredefinedVerified - error: invalid PredefinedUser type")
		err = mtproto.ErrInternelServerError
		return
	}

	log.Debugf("updatePredefinedVerified - reply: %s", r.DebugString())
	return
}

func (s *Service) UpdatePredefinedCode(ctx context.Context, i *api.UpdatePredefinedCode) (r *mtproto.PredefinedUser, err error) {
	log.Debugf("updatePredefinedCode - request: %s", i.DebugString())

	var (
		result mtproto.TLObject
		ok     bool
	)

	if i.Code == "" {
		i.Code = random2.RandomNumeric(5)
	} else if len(i.Code) != 5 {
		err = mtproto.ErrPhoneCodeInvalid
		log.Debugf("updatePredefinedCode - error: %v", err)
		return
	}

	req := &mtproto.TLAccountUpdatePredefinedCode{
		Phone: i.Phone,
		Code:  i.Code,
	}

	result, err = s.Invoke(ctx, 777000, req)
	if err != nil {
		log.Errorf("updatePredefinedCode - error: %v", err)
		return
	} else if r, ok = result.(*mtproto.PredefinedUser); !ok {
		log.Errorf("updatePredefinedCode - error: invalid PredefinedUser type")
		err = mtproto.ErrInternelServerError
		return
	}

	log.Debugf("updatePredefinedCode - reply: %s", r.DebugString())
	return
}

func (s *Service) GetPredefinedUsers(ctx context.Context, i *api.GetPredefinedUsers) (r *mtproto.Vector_PredefinedUser, err error) {
	log.Debugf("getPredefinedUsers - request: %s", i.DebugString())

	var (
		result mtproto.TLObject
		ok     bool
	)

	req := &mtproto.TLUsersGetPredefinedUsers{}

	result, err = s.Invoke(ctx, 777000, req)
	if err != nil {
		log.Errorf("getPredefinedUsers - error: %v", err)
		return
	} else if r, ok = result.(*mtproto.Vector_PredefinedUser); !ok {
		log.Errorf("getPredefinedUsers - error: invalid Vector_PredefinedUser type")
		err = mtproto.ErrInternelServerError
		return
	}

	log.Debugf("getPredefinedUsers - reply: %s", r.DebugString())
	return
}

func (s *Service) GetPredefinedUser(ctx context.Context, i *api.GetPredefinedUser) (r *mtproto.PredefinedUser, err error) {
	log.Debugf("getPredefinedUser - request: %s", i.DebugString())

	var (
		result mtproto.TLObject
		ok     bool
	)

	req := &mtproto.TLUsersGetPredefinedUser{
		Phone: i.Phone,
	}

	result, err = s.Invoke(ctx, 777000, req)
	if err != nil {
		log.Errorf("getPredefinedUser - error: %v", err)
		return
	} else if r, ok = result.(*mtproto.PredefinedUser); !ok {
		log.Errorf("getPredefinedUser - error: invalid PredefinedUser type")
		err = mtproto.ErrInternelServerError
		return
	}

	log.Debugf("getPredefinedUser - reply: %s", r.DebugString())
	return
}
