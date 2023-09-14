package core

import (
	"context"
	"math"
	"strconv"
	"time"

	"github.com/gogo/protobuf/types"
	"open.chat/app/service/biz_service/poll/internal/dal/dataobject"
	"open.chat/model"
	"open.chat/mtproto"
	"open.chat/pkg/database/sqlx"
	"open.chat/pkg/hack"
	"open.chat/pkg/log"
	"open.chat/pkg/util"
)

func (m *PollCore) makeMediaPoll(ctx context.Context, userId int32, pollDO *dataobject.PollsDO, pollAnswerDO *dataobject.PollAnswerVotersDO) *model.MediaPoll {
	var (
		poll = mtproto.MakeTLPoll(&mtproto.Poll{
			Id:             pollDO.Id,
			Closed:         util.Int8ToBool(pollDO.Closed),
			PublicVoters:   util.Int8ToBool(pollDO.PublicVoters),
			MultipleChoice: util.Int8ToBool(pollDO.MultipleChoice),
			Quiz:           util.Int8ToBool(pollDO.Quiz),
			Question:       pollDO.Question,
			Answers:        nil,
		}).To_Poll()

		pollResults = mtproto.MakeTLPollResults(&mtproto.PollResults{
			Min:          false,
			Results:      make([]*mtproto.PollAnswerVoters, 0, 10),
			TotalVoters:  &types.Int32Value{Value: 0},
			RecentVoters: nil,
		}).To_PollResults()

		chosen = false
	)

	// var optionIds
	if pollDO.Text0 != "" {
		poll.Answers = append(poll.Answers, mtproto.MakeTLPollAnswer(&mtproto.PollAnswer{
			Text:   pollDO.Text0,
			Option: hack.Bytes(pollDO.Option0),
		}).To_PollAnswer())

		if pollAnswerDO != nil && pollAnswerDO.Option0 == 1 {
			chosen = true
		}
		pollResults.Results = append(pollResults.Results, mtproto.MakeTLPollAnswerVoters(&mtproto.PollAnswerVoters{
			Chosen:  chosen,
			Correct: util.Int8ToBool(pollDO.CorrectAnswer0),
			Option:  hack.Bytes(pollDO.Option0),
			Voters:  pollDO.Voters0,
		}).To_PollAnswerVoters())

		pollResults.TotalVoters.Value += pollDO.Voters0
	} else {
		goto RET100
	}

	if pollDO.Text1 != "" {
		poll.Answers = append(poll.Answers, mtproto.MakeTLPollAnswer(&mtproto.PollAnswer{
			Text:   pollDO.Text1,
			Option: hack.Bytes(pollDO.Option1),
		}).To_PollAnswer())

		if pollAnswerDO != nil && pollAnswerDO.Option1 == 1 {
			chosen = true
		}
		pollResults.Results = append(pollResults.Results, mtproto.MakeTLPollAnswerVoters(&mtproto.PollAnswerVoters{
			Chosen:  chosen,
			Correct: util.Int8ToBool(pollDO.CorrectAnswer1),
			Option:  hack.Bytes(pollDO.Option1),
			Voters:  pollDO.Voters1,
		}).To_PollAnswerVoters())

		pollResults.TotalVoters.Value += pollDO.Voters1
	} else {
		goto RET100
	}

	if pollDO.Text2 != "" {
		poll.Answers = append(poll.Answers, mtproto.MakeTLPollAnswer(&mtproto.PollAnswer{
			Text:   pollDO.Text2,
			Option: hack.Bytes(pollDO.Option2),
		}).To_PollAnswer())

		if pollAnswerDO != nil && pollAnswerDO.Option2 == 1 {
			chosen = true
		}
		pollResults.Results = append(pollResults.Results, mtproto.MakeTLPollAnswerVoters(&mtproto.PollAnswerVoters{
			Chosen:  chosen,
			Correct: util.Int8ToBool(pollDO.CorrectAnswer2),
			Option:  hack.Bytes(pollDO.Option2),
			Voters:  pollDO.Voters2,
		}).To_PollAnswerVoters())

		pollResults.TotalVoters.Value += pollDO.Voters2
	} else {
		goto RET100
	}

	if pollDO.Text3 != "" {
		poll.Answers = append(poll.Answers, mtproto.MakeTLPollAnswer(&mtproto.PollAnswer{
			Text:   pollDO.Text3,
			Option: hack.Bytes(pollDO.Option3),
		}).To_PollAnswer())

		if pollAnswerDO != nil && pollAnswerDO.Option3 == 1 {
			chosen = true
		}
		pollResults.Results = append(pollResults.Results, mtproto.MakeTLPollAnswerVoters(&mtproto.PollAnswerVoters{
			Chosen:  chosen,
			Correct: util.Int8ToBool(pollDO.CorrectAnswer3),
			Option:  hack.Bytes(pollDO.Option3),
			Voters:  pollDO.Voters3,
		}).To_PollAnswerVoters())

		pollResults.TotalVoters.Value += pollDO.Voters3
	} else {
		goto RET100
	}

	if pollDO.Text4 != "" {
		poll.Answers = append(poll.Answers, mtproto.MakeTLPollAnswer(&mtproto.PollAnswer{
			Text:   pollDO.Text4,
			Option: hack.Bytes(pollDO.Option4),
		}).To_PollAnswer())

		if pollAnswerDO != nil && pollAnswerDO.Option4 == 1 {
			chosen = true
		}
		pollResults.Results = append(pollResults.Results, mtproto.MakeTLPollAnswerVoters(&mtproto.PollAnswerVoters{
			Chosen:  chosen,
			Correct: util.Int8ToBool(pollDO.CorrectAnswer4),
			Option:  hack.Bytes(pollDO.Option4),
			Voters:  pollDO.Voters4,
		}).To_PollAnswerVoters())

		pollResults.TotalVoters.Value += pollDO.Voters4
	} else {
		goto RET100
	}

	if pollDO.Text5 != "" {
		poll.Answers = append(poll.Answers, mtproto.MakeTLPollAnswer(&mtproto.PollAnswer{
			Text:   pollDO.Text5,
			Option: hack.Bytes(pollDO.Option5),
		}).To_PollAnswer())

		if pollAnswerDO != nil && pollAnswerDO.Option5 == 1 {
			chosen = true
		}
		pollResults.Results = append(pollResults.Results, mtproto.MakeTLPollAnswerVoters(&mtproto.PollAnswerVoters{
			Chosen:  chosen,
			Correct: util.Int8ToBool(pollDO.CorrectAnswer5),
			Option:  hack.Bytes(pollDO.Option5),
			Voters:  pollDO.Voters5,
		}).To_PollAnswerVoters())

		pollResults.TotalVoters.Value += pollDO.Voters5
	} else {
		goto RET100
	}

	if pollDO.Text6 != "" {
		poll.Answers = append(poll.Answers, mtproto.MakeTLPollAnswer(&mtproto.PollAnswer{
			Text:   pollDO.Text6,
			Option: hack.Bytes(pollDO.Option6),
		}).To_PollAnswer())

		if pollAnswerDO != nil && pollAnswerDO.Option6 == 1 {
			chosen = true
		}
		pollResults.Results = append(pollResults.Results, mtproto.MakeTLPollAnswerVoters(&mtproto.PollAnswerVoters{
			Chosen:  chosen,
			Correct: util.Int8ToBool(pollDO.CorrectAnswer6),
			Option:  hack.Bytes(pollDO.Option6),
			Voters:  pollDO.Voters6,
		}).To_PollAnswerVoters())

		pollResults.TotalVoters.Value += pollDO.Voters6
	} else {
		goto RET100
	}

	if pollDO.Text7 != "" {
		poll.Answers = append(poll.Answers, mtproto.MakeTLPollAnswer(&mtproto.PollAnswer{
			Text:   pollDO.Text7,
			Option: hack.Bytes(pollDO.Option7),
		}).To_PollAnswer())

		if pollAnswerDO != nil && pollAnswerDO.Option7 == 1 {
			chosen = true
		}
		pollResults.Results = append(pollResults.Results, mtproto.MakeTLPollAnswerVoters(&mtproto.PollAnswerVoters{
			Chosen:  chosen,
			Correct: util.Int8ToBool(pollDO.CorrectAnswer7),
			Option:  hack.Bytes(pollDO.Option7),
			Voters:  pollDO.Voters7,
		}).To_PollAnswerVoters())

		pollResults.TotalVoters.Value += pollDO.Voters7
	} else {
		goto RET100
	}

	if pollDO.Text8 != "" {
		poll.Answers = append(poll.Answers, mtproto.MakeTLPollAnswer(&mtproto.PollAnswer{
			Text:   pollDO.Text8,
			Option: hack.Bytes(pollDO.Option8),
		}).To_PollAnswer())

		if pollAnswerDO != nil && pollAnswerDO.Option8 == 1 {
			chosen = true
		}
		pollResults.Results = append(pollResults.Results, mtproto.MakeTLPollAnswerVoters(&mtproto.PollAnswerVoters{
			Chosen:  chosen,
			Correct: util.Int8ToBool(pollDO.CorrectAnswer8),
			Option:  hack.Bytes(pollDO.Option8),
			Voters:  pollDO.Voters8,
		}).To_PollAnswerVoters())

		pollResults.TotalVoters.Value += pollDO.Voters8
	} else {
		goto RET100
	}

	if pollDO.Text9 != "" {
		poll.Answers = append(poll.Answers, mtproto.MakeTLPollAnswer(&mtproto.PollAnswer{
			Text:   pollDO.Text9,
			Option: hack.Bytes(pollDO.Option9),
		}).To_PollAnswer())

		if pollAnswerDO != nil && pollAnswerDO.Option9 == 1 {
			chosen = true
		}
		pollResults.Results = append(pollResults.Results, mtproto.MakeTLPollAnswerVoters(&mtproto.PollAnswerVoters{
			Chosen:  chosen,
			Correct: util.Int8ToBool(pollDO.CorrectAnswer9),
			Option:  hack.Bytes(pollDO.Option9),
			Voters:  pollDO.Voters9,
		}).To_PollAnswerVoters())

		pollResults.TotalVoters.Value += pollDO.Voters9
	}

RET100:
	if pollDO.MultipleChoice == 1 {
		pollResults.TotalVoters = &types.Int32Value{Value: m.calcTotalVoters(ctx, pollDO.Id)}
	}
	if pollDO.PublicVoters == 1 {
		recentVoters, _ := m.Dao.PollAnswerVotersDAO.SelectRecentVoters(ctx, pollDO.Id, 15)
		for _, v := range recentVoters {
			pollResults.RecentVoters = append(pollResults.RecentVoters, v)
		}
	}

	return &model.MediaPoll{
		UserId:  userId,
		Poll:    poll,
		Results: pollResults,
	}
}

func (m *PollCore) CreateMediaPoll(ctx context.Context, userId int32, correctAnswers []int, poll *mtproto.Poll) (*model.MediaPoll, error) {
	var (
		pollDO = &dataobject.PollsDO{
			Id:             0,
			PollId:         poll.Id,
			Creator:        userId,
			Question:       poll.Question,
			Closed:         0,
			MultipleChoice: util.BoolToInt8(poll.MultipleChoice),
			PublicVoters:   util.BoolToInt8(poll.PublicVoters),
			Quiz:           util.BoolToInt8(poll.Quiz),
			Date2:          int32(time.Now().Unix()),
		}
		err error
	)

	hasCorrectAnswer := func(i int) int8 {
		for _, v := range correctAnswers {
			if i == v {
				return 1
			}
		}
		return 0
	}

	// for _, c := range poll.
	for i, answer := range poll.Answers {
		hasCorrectAnswer := hasCorrectAnswer(i)

		switch i {
		case 0:
			pollDO.Text0 = answer.Text
			pollDO.Option0 = hack.String(answer.Option)
			pollDO.CorrectAnswer0 = hasCorrectAnswer
		case 1:
			pollDO.Text1 = answer.Text
			pollDO.Option1 = hack.String(answer.Option)
			pollDO.CorrectAnswer1 = hasCorrectAnswer
		case 2:
			pollDO.Text2 = answer.Text
			pollDO.Option2 = hack.String(answer.Option)
			pollDO.CorrectAnswer2 = hasCorrectAnswer
		case 3:
			pollDO.Text3 = answer.Text
			pollDO.Option3 = hack.String(answer.Option)
			pollDO.CorrectAnswer3 = hasCorrectAnswer
		case 4:
			pollDO.Text4 = answer.Text
			pollDO.Option4 = hack.String(answer.Option)
			pollDO.CorrectAnswer4 = hasCorrectAnswer
		case 5:
			pollDO.Text5 = answer.Text
			pollDO.Option5 = hack.String(answer.Option)
			pollDO.CorrectAnswer5 = hasCorrectAnswer
		case 6:
			pollDO.Text6 = answer.Text
			pollDO.Option6 = hack.String(answer.Option)
			pollDO.CorrectAnswer6 = hasCorrectAnswer
		case 7:
			pollDO.Text7 = answer.Text
			pollDO.Option7 = hack.String(answer.Option)
			pollDO.CorrectAnswer7 = hasCorrectAnswer
		case 8:
			pollDO.Text8 = answer.Text
			pollDO.Option8 = hack.String(answer.Option)
			pollDO.CorrectAnswer8 = hasCorrectAnswer
		case 9:
			pollDO.Text9 = answer.Text
			pollDO.Option9 = hack.String(answer.Option)
			pollDO.CorrectAnswer9 = hasCorrectAnswer
		}
	}

	pollDO.Id, _, err = m.PollsDAO.Insert(ctx, pollDO)
	if err != nil {
		return nil, err
	}

	return m.makeMediaPoll(ctx, userId, pollDO, nil), nil
}

func (m *PollCore) CloseMediaPoll(ctx context.Context, userId int32, pollId int64) (*model.MediaPoll, error) {
	poll, err := m.PollsDAO.Select(ctx, pollId)
	if err != nil {
		return nil, err
	} else if poll == nil {
		return nil, mtproto.ErrMediaInvalid
	}

	if poll.Creator != userId {
		err := mtproto.ErrMediaInvalid
		return nil, err
	}

	_, err = m.PollsDAO.Update(ctx, map[string]interface{}{
		"Closed": 1,
	}, pollId)
	if err != nil {
		return nil, err
	}

	poll.Closed = 1
	pollAnswer, err := m.PollAnswerVotersDAO.Select(ctx, pollId, userId)
	if err != nil {
		err := mtproto.ErrMediaInvalid
		return nil, err
	}

	return m.makeMediaPoll(ctx, userId, poll, pollAnswer), nil
}

func (m *PollCore) GetMediaPoll(ctx context.Context, userId int32, pollId int64) (*model.MediaPoll, error) {
	poll, err := m.PollsDAO.Select(ctx, pollId)
	if err != nil {
		return nil, err
	} else if poll == nil {
		return nil, mtproto.ErrMediaInvalid
	}

	pollAnswer, err := m.PollAnswerVotersDAO.Select(ctx, pollId, userId)
	if err != nil {
		err := mtproto.ErrMediaInvalid
		return nil, err
	}

	return m.makeMediaPoll(ctx, userId, poll, pollAnswer), nil
}

func (m *PollCore) SendVote(ctx context.Context, userId int32, pollId int64, options []string) (*model.MediaPoll, error) {
	poll, err := m.PollsDAO.Select(ctx, pollId)
	if err != nil {
		return nil, err
	} else if poll == nil {
		return nil, mtproto.ErrMediaInvalid
	}

	pollAnswer, err := m.PollAnswerVotersDAO.Select(ctx, pollId, userId)
	if err != nil {
		err := mtproto.ErrMediaInvalid
		return nil, err
	}

	var hasPollAnswer = func(answer *dataobject.PollAnswerVotersDO) bool {
		if answer == nil {
			return false
		}
		return answer.Option0 == 1 ||
			answer.Option1 == 1 ||
			answer.Option2 == 1 ||
			answer.Option3 == 1 ||
			answer.Option4 == 1 ||
			answer.Option5 == 1 ||
			answer.Option6 == 1 ||
			answer.Option7 == 1 ||
			answer.Option8 == 1 ||
			answer.Option9 == 1
	}(pollAnswer)

	if pollAnswer == nil {
		if len(options) == 0 {
			goto RET100
		}
	} else {
		if !hasPollAnswer {
			if len(options) == 0 {
				goto RET100
			}
		} else {
			if len(options) > 0 {
				goto RET100
			}
		}
	}

	if len(options) == 0 {
		// 1. get voted options
		var (
			cMap = make(map[string]interface{}, 1)
		)
		if pollAnswer.Option0 == 1 {
			poll.Voters0 -= 1
			cMap["voters0"] = poll.Voters0
			pollAnswer.Option0 = 0
		}
		if pollAnswer.Option1 == 1 {
			poll.Voters1 -= 1
			cMap["voters1"] = poll.Voters1
			pollAnswer.Option1 = 0
		}
		if pollAnswer.Option2 == 1 {
			poll.Voters2 -= 1
			cMap["voters2"] = poll.Voters2
			pollAnswer.Option2 = 0
		}
		if pollAnswer.Option3 == 1 {
			poll.Voters3 -= 1
			cMap["voters3"] = poll.Voters3
			pollAnswer.Option3 = 0
		}
		if pollAnswer.Option4 == 1 {
			poll.Voters4 -= 1
			cMap["voters4"] = poll.Voters4
			pollAnswer.Option4 = 0
		}
		if pollAnswer.Option5 == 1 {
			poll.Voters5 -= 1
			cMap["voters5"] = poll.Voters5
			pollAnswer.Option5 = 0
		}
		if pollAnswer.Option6 == 1 {
			poll.Voters6 -= 1
			cMap["voters6"] = poll.Voters6
			pollAnswer.Option6 = 0
		}
		if pollAnswer.Option7 == 1 {
			poll.Voters7 -= 1
			cMap["voters7"] = poll.Voters7
			pollAnswer.Option7 = 0
		}
		if pollAnswer.Option8 == 1 {
			poll.Voters8 -= 1
			cMap["voters8"] = poll.Voters8
			pollAnswer.Option8 = 0
		}
		if pollAnswer.Option9 == 1 {
			poll.Voters9 -= 1
			cMap["voters9"] = poll.Voters9
			pollAnswer.Option9 = 0
		}
		pollAnswer.Deleted = 1
		pollAnswer.Options = ""
		tR := sqlx.TxWrapper(ctx, m.Dao.DB, func(tx *sqlx.Tx, result *sqlx.StoreResult) {
			_, _, err = m.PollAnswerVotersDAO.InsertOrUpdateTx(tx, pollAnswer)
			if err != nil {
				result.Err = err
				return
			}
			_, result.Err = m.PollsDAO.Update(ctx, cMap, pollId)
		})
		if tR.Err != nil {
			return nil, tR.Err
		}
	} else {
		var (
			cMap = make(map[string]interface{}, 1)
		)

		if pollAnswer == nil {
			pollAnswer = &dataobject.PollAnswerVotersDO{
				PollId:     pollId,
				VoteUserId: userId,
				Date2:      time.Now().Unix(),
			}
		} else {
			pollAnswer.Date2 = time.Now().Unix()
			pollAnswer.Deleted = 0
		}

		for _, o := range options {
			switch o {
			case "0":
				poll.Voters0 += 1
				cMap["voters0"] = poll.Voters0
				pollAnswer.Option0 = 1
			case "1":
				poll.Voters1 += 1
				cMap["voters1"] = poll.Voters1
				pollAnswer.Option1 = 1
			case "2":
				poll.Voters2 += 1
				cMap["voters2"] = poll.Voters2
				pollAnswer.Option2 = 1
			case "3":
				poll.Voters3 += 1
				cMap["voters3"] = poll.Voters3
				pollAnswer.Option3 = 1
			case "4":
				poll.Voters4 += 1
				cMap["voters4"] = poll.Voters4
				pollAnswer.Option4 = 1
			case "5":
				poll.Voters5 += 1
				cMap["voters5"] = poll.Voters5
				pollAnswer.Option5 = 1
			case "6":
				poll.Voters6 += 1
				cMap["voters6"] = poll.Voters6
				pollAnswer.Option6 = 1
			case "7":
				poll.Voters7 += 1
				cMap["voters7"] = poll.Voters7
				pollAnswer.Option7 = 1
			case "8":
				poll.Voters8 += 1
				cMap["voters8"] = poll.Voters8
				pollAnswer.Option8 = 1
			case "9":
				poll.Voters9 += 1
				cMap["voters9"] = poll.Voters9
				pollAnswer.Option9 = 1
			default:
				return nil, mtproto.ErrInputRequestInvalid
			}
		}

		tR := sqlx.TxWrapper(ctx, m.Dao.DB, func(tx *sqlx.Tx, result *sqlx.StoreResult) {
			_, _, result.Err = m.PollAnswerVotersDAO.InsertOrUpdateTx(tx, pollAnswer)
			if result.Err != nil {
				return
			}
			_, result.Err = m.PollsDAO.Update(ctx, cMap, pollId)
		})
		if tR.Err != nil {
			return nil, tR.Err
		}
	}

RET100:
	return m.makeMediaPoll(ctx, userId, poll, pollAnswer), nil
}

func (m *PollCore) GetPollVoters(ctx context.Context, userId int32, pollId int64, option string, offset string, limit int32) (*mtproto.Messages_VotesList, error) {
	_ = userId

	var (
		offset2 int64 = 0
		voters  []dataobject.PollAnswerVotersDO
		err     error
	)

	votersList := mtproto.MakeTLMessagesVotesList(&mtproto.Messages_VotesList{
		Count:      m.calcTotalVoters(ctx, pollId),
		Votes:      nil,
		Users:      nil,
		NextOffset: nil,
	}).To_Messages_VotesList()

	if offset == "" {
		offset2 = math.MaxInt32
	} else {
		offset2, _ = strconv.ParseInt(offset, 10, 64)
		if offset2 == 0 {
			offset2 = math.MaxInt32
		}
	}
	if option == "" {
		if voters, err = m.Dao.SelectVoters(ctx, pollId, offset2, limit); err != nil {
			log.Errorf("getPollVoters - error: %v", err)
			return nil, err
		}
	} else {
		if voters, err = m.Dao.SelectOptionVoters(ctx, pollId, "option"+option, offset2, limit); err != nil {
			log.Errorf("getPollVoters - error: %v", err)
			return nil, err
		}
	}

	for i := 0; i < len(voters); i++ {
		votersList.Votes = append(votersList.Votes, mtproto.MakeTLMessageUserVoteInputOption(&mtproto.MessageUserVote{
			UserId: voters[i].VoteUserId,
			Date:   int32(voters[i].Date2),
		}).To_MessageUserVote())

		if i+1 == int(limit) {
			votersList.NextOffset = &types.StringValue{Value: strconv.FormatInt(voters[i].Date2, 10)}
		}
	}

	return votersList, nil
}

func (m *PollCore) calcTotalVoters(ctx context.Context, pollId int64) int32 {
	totalVoters := m.Dao.CommonDAO.CalcSize(ctx, "poll_answer_voters", map[string]interface{}{
		"poll_id": pollId,
	})
	return int32(totalVoters)
}
