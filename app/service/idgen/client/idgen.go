package idgen

import (
	"context"
	"flag"
	"fmt"

	id_facade "open.chat/app/service/idgen/facade"
	_ "open.chat/app/service/idgen/facade/redis"
	_ "open.chat/app/service/idgen/facade/snowflake"
	"open.chat/pkg/log"
	"open.chat/pkg/util"
)

// ///////////////////////////////////////////////////////////////////////////////////////////
var (
	messageDataNgenId       = "message_data_ngen_"
	messageBoxUpdatesNgenId = "message_box_ngen_"
	channelMessageBoxNgenId = "channel_message_box_ngen_"
	seqUpdatesNgenId        = "seq_updates_ngen_"
	ptsUpdatesNgenId        = "pts_updates_ngen_"
	qtsUpdatesNgenId        = "qts_updates_ngen_"
	channelPtsUpdatesNgenId = "channel_pts_updates_ngen_"
	scheduledMessageNgenId  = "scheduled_ngen_"
	userBlogUpdatesNgenId   = "user_blog_updates_ngen_"
	blogNgenId              = "blog_id_ngen"
	blogPtsUpdatesNgenId    = "blog_pts_updates_ngen_"
)

var (
	status2 = false
)

func init() {
	flag.BoolVar(&status2, "status2", false, "--status2 true")

	if status2 {
		messageDataNgenId = "message_data_ngen2_"
		messageBoxUpdatesNgenId = "message_box_ngen2_"
		channelMessageBoxNgenId = "channel_message_box_ngen2_"
		seqUpdatesNgenId = "seq_updates_ngen2_"
		ptsUpdatesNgenId = "pts_updates_ngen2_"
		qtsUpdatesNgenId = "qts_updates_ngen2_"
		channelPtsUpdatesNgenId = "channel_pts_updates_ngen2_"
		scheduledMessageNgenId = "scheduled_ngen2_"
	}
}

// ///////////////////////////////////////////////////////////////////////////////////////////
// type Instance func() Initializer
var uuidGen id_facade.UUIDGen

func NewUUID() {
	if uuidGen == nil {
		uuidGen, _ = id_facade.NewUUIDGen("snowflake")
	}
}

func GetUUID() (uuid int64) {
	uuid, _ = uuidGen.GetUUID()
	return
}

// ///////////////////////////////////////////////////////////////////////////////////////////
var seqIDGen id_facade.SeqIDGen

func NewSeqIDGen() {
	if seqIDGen == nil {
		var err error
		seqIDGen, err = id_facade.NewSeqIDGen("redis")
		if err != nil {
			panic(fmt.Errorf("seqidGen init error: %v", err))
		}
	}
}

func GetCurrentSeqID(ctx context.Context, key string) (int64, error) {
	return seqIDGen.GetCurrentSeqID(ctx, key)
}

func GetNextSeqID(ctx context.Context, key string) (int64, error) {
	return seqIDGen.GetNextSeqID(ctx, key)
}

func GetNextNSeqID(ctx context.Context, key string, n int) (seq int64, err error) {
	return seqIDGen.GetNextNSeqID(ctx, key, n)
}

func NextMessageBoxId(ctx context.Context, key int32) (seq int64) {
	var err error
	seq, err = seqIDGen.GetNextSeqID(ctx, messageBoxUpdatesNgenId+util.Int32ToString(key))
	if err != nil {
		log.Errorf("genId error: %v", err)
	}
	return
}

func CurrentMessageBoxId(ctx context.Context, key int32) (seq int64) {
	seq, _ = seqIDGen.GetCurrentSeqID(ctx, messageBoxUpdatesNgenId+util.Int32ToString(key))
	return
}

func SetCurrentMessageBoxId(ctx context.Context, key, v int32) {
	seqIDGen.SetCurrentSeqID(ctx, messageBoxUpdatesNgenId+util.Int32ToString(key), int64(v))
}

func NextMessageDataId(ctx context.Context, key int64) (seq int64) {
	seq, _ = seqIDGen.GetNextSeqID(ctx, messageDataNgenId+util.Int64ToString(key))
	return
}

func SetCurrentMessageDataId(ctx context.Context, key int64, v int32) {
	seqIDGen.SetCurrentSeqID(ctx, messageDataNgenId+util.Int64ToString(key), int64(v))
}

func NextChannelMessageBoxId(ctx context.Context, key int32) (seq int64) {
	var err error
	seq, err = seqIDGen.GetNextSeqID(ctx, channelMessageBoxNgenId+util.Int32ToString(key))
	if err != nil {
		log.Errorf("genId error: %v", err)
	}
	return
}

func CurrentChannelMessageBoxId(ctx context.Context, key int32) (seq int64) {
	seq, _ = seqIDGen.GetCurrentSeqID(ctx, channelMessageBoxNgenId+util.Int32ToString(key))
	return
}

func SetCurrentChannelMessageBoxId(ctx context.Context, key, v int32) {
	seqIDGen.SetCurrentSeqID(ctx, channelMessageBoxNgenId+util.Int32ToString(key), int64(v))
}

func NextSeqId(ctx context.Context, key int64) (seq int64) {
	seq, _ = seqIDGen.GetNextSeqID(ctx, seqUpdatesNgenId+util.Int64ToString(key))
	return
}

func CurrentSeqId(ctx context.Context, key int64) (seq int64) {
	var err error
	seq, err = seqIDGen.GetCurrentSeqID(ctx, seqUpdatesNgenId+util.Int64ToString(key))

	if err != nil {
		seq = -1
	}
	return
}

func SetCurrentSeqId(ctx context.Context, key int64, v int32) {
	seqIDGen.SetCurrentSeqID(ctx, seqUpdatesNgenId+util.Int64ToString(key), int64(v))
}

func NextPtsId(ctx context.Context, key int32) (seq int64) {
	seq, _ = seqIDGen.GetNextSeqID(ctx, ptsUpdatesNgenId+util.Int32ToString(key))
	log.Debugf("NextPtsId: %d", seq)
	return
}

func NextNPtsId(ctx context.Context, key int32, n int) (seq int64) {
	seq, _ = seqIDGen.GetNextNSeqID(ctx, ptsUpdatesNgenId+util.Int32ToString(key), n)
	return
}

func CurrentPtsId(ctx context.Context, key int32) (seq int64) {
	seq, _ = seqIDGen.GetCurrentSeqID(ctx, ptsUpdatesNgenId+util.Int32ToString(key))
	return
}

func SetCurrentPtsId(ctx context.Context, key, v int32) {
	seqIDGen.SetCurrentSeqID(ctx, ptsUpdatesNgenId+util.Int32ToString(key), int64(v))
}

func NextQtsId(ctx context.Context, key int64) (seq int64) {
	seq, _ = seqIDGen.GetNextSeqID(ctx, qtsUpdatesNgenId+util.Int64ToString(key))
	return
}

func CurrentQtsId(ctx context.Context, key int64) (seq int64) {
	seq, _ = seqIDGen.GetCurrentSeqID(ctx, qtsUpdatesNgenId+util.Int64ToString(key))
	return
}

func SetCurrentQtsId(ctx context.Context, key int64, v int32) {
	seqIDGen.SetCurrentSeqID(ctx, qtsUpdatesNgenId+util.Int64ToString(key), int64(v))
}

func NextChannelPtsId(ctx context.Context, key int32) (seq int64) {
	seq, _ = seqIDGen.GetNextSeqID(ctx, channelPtsUpdatesNgenId+util.Int32ToString(key))
	return
}

func NextChannelNPtsId(ctx context.Context, key int32, n int) (seq int64) {
	seq, _ = seqIDGen.GetNextNSeqID(ctx, channelPtsUpdatesNgenId+util.Int32ToString(key), n)
	return
}

func CurrentChannelPtsId(ctx context.Context, key int32) (seq int64) {
	seq, _ = seqIDGen.GetCurrentSeqID(ctx, channelPtsUpdatesNgenId+util.Int32ToString(key))
	return
}

func SetCurrentChannelPtsId(ctx context.Context, key, v int32) {
	seqIDGen.SetCurrentSeqID(ctx, channelPtsUpdatesNgenId+util.Int32ToString(key), int64(v))
}

func NextScheduledMessageBoxId(ctx context.Context, key int64) (seq int64) {
	var err error
	seq, err = seqIDGen.GetNextSeqID(ctx, scheduledMessageNgenId+util.Int64ToString(key))
	if err != nil {
		log.Errorf("genId error: %v", err)
	}
	return
}

func GetNextPhoneNumber(ctx context.Context, prefix string) (string, error) {
	return seqIDGen.GetNextPhoneNumber(ctx, prefix)
}

func NextUserBlogId(ctx context.Context, key int32) (seq int64) {
	seq, _ = seqIDGen.GetNextSeqID(ctx, userBlogUpdatesNgenId+util.Int32ToString(key))
	return
}

func NextUserBlogNId(ctx context.Context, key int32, n int) (seq int64) {
	seq, _ = seqIDGen.GetNextNSeqID(ctx, userBlogUpdatesNgenId+util.Int32ToString(key), n)
	return
}

func CurrentUserBlogId(ctx context.Context, key int32) (seq int64) {
	seq, _ = seqIDGen.GetCurrentSeqID(ctx, userBlogUpdatesNgenId+util.Int32ToString(key))
	return
}

func SetCurrentUserBlogId(ctx context.Context, key, v int32) {
	seqIDGen.SetCurrentSeqID(ctx, userBlogUpdatesNgenId+util.Int32ToString(key), int64(v))
}

func NextBlogId(ctx context.Context) (seq int64) {
	seq, _ = seqIDGen.GetNextSeqID(ctx, blogNgenId)
	return
}

func NextBlogNId(ctx context.Context, n int) (seq int64) {
	seq, _ = seqIDGen.GetNextNSeqID(ctx, blogNgenId, n)
	return
}

func CurrentBlogId(ctx context.Context) (seq int64) {
	seq, _ = seqIDGen.GetCurrentSeqID(ctx, blogNgenId)
	return
}

func SetCurrentBlogId(ctx context.Context, v int32) {
	seqIDGen.SetCurrentSeqID(ctx, blogNgenId, int64(v))
}

func NextBlogPtsId(ctx context.Context, key int32) (seq int64) {
	seq, _ = seqIDGen.GetNextSeqID(ctx, blogPtsUpdatesNgenId+util.Int32ToString(key))
	return
}

func NextBlogNPtsId(ctx context.Context, key int32, n int) (seq int64) {
	seq, _ = seqIDGen.GetNextNSeqID(ctx, blogPtsUpdatesNgenId+util.Int32ToString(key), n)
	return
}

func CurrentBlogPtsId(ctx context.Context, key int32) (seq int64) {
	seq, _ = seqIDGen.GetCurrentSeqID(ctx, blogPtsUpdatesNgenId+util.Int32ToString(key))
	return
}

func SetCurrentBlogPtsId(ctx context.Context, key, v int32) {
	seqIDGen.SetCurrentSeqID(ctx, blogPtsUpdatesNgenId+util.Int32ToString(key), int64(v))
}
