package mtproto

import (
	"github.com/go-kratos/kratos/pkg/ecode"
	"github.com/gogo/protobuf/types"
)

// //////////////////////////////////////////////////////////////////////////////
var (
	BoolTrue       = MakeTLBoolTrue(nil).To_Bool()
	BoolFalse      = MakeTLBoolFalse(nil).To_Bool()
	UpdatesTooLong = MakeTLUpdatesTooLong(nil).To_Updates()
)

func ToBool(b bool) *Bool {
	if b {
		return BoolTrue
	} else {
		return BoolFalse
	}
}

func FromBool(b *Bool) bool {
	return Predicate_boolTrue == b.GetPredicateName()
}

// /////////////////////////////////////////////////////////////////////////////////////////
func MakeFlagsInt32(v int32) *types.Int32Value {
	if v == 0 {
		return nil
	} else {
		return &types.Int32Value{Value: v}
	}
}

func MakeFlagsInt64(v int64) *types.Int64Value {
	if v == 0 {
		return nil
	} else {
		return &types.Int64Value{Value: v}
	}
}

func MakeFlagsString(v string) *types.StringValue {
	if v == "" {
		return nil
	} else {
		return &types.StringValue{Value: v}
	}
}

// /////////////////////////////////////////////////////////////////////////////////////////
func (m *TLRpcError) ToError() *ecode.Status {
	return ecode.Error(ecode.New(int(m.GetErrorCode())), m.GetErrorMessage())
}

// /////////////////////////////////////////////////////////////////////////////////////////
type MessageEntitySlice []*MessageEntity

// sort
func (m MessageEntitySlice) Len() int {
	return len(m)
}
func (m MessageEntitySlice) Swap(i, j int) {
	m[j], m[i] = m[i], m[j]
}
func (m MessageEntitySlice) Less(i, j int) bool {
	// if date[i] == date[j]
	return m[i].Offset < m[j].Offset
}
