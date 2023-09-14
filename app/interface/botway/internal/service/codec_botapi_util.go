package service

import (
	"encoding/binary"
	"math"
	"math/rand"
	"open.chat/pkg/log"

	"open.chat/app/interface/botway/botapi"
	"open.chat/mtproto"
	"open.chat/pkg/math2"
)

func MakePeer(chatId botapi.ChatId2) *mtproto.InputPeer {
	switch t := chatId.(type) {
	case botapi.ChannelUsername:
		return mtproto.MakeTLInputPeerUsername(&mtproto.InputPeer{
			Username: string(t),
		}).To_InputPeer()
	case botapi.ChatID:
		idType, id := t.ToChatIdTypeId()
		switch idType {
		case int32(botapi.ChatIdPrivate):
			return mtproto.MakeTLInputPeerUser(&mtproto.InputPeer{
				UserId: id,
			}).To_InputPeer()
		case int32(botapi.ChatIdGroup):
			return mtproto.MakeTLInputPeerChat(&mtproto.InputPeer{
				ChatId: id,
			}).To_InputPeer()
		case int32(botapi.ChatIdSuperGroup):
			return mtproto.MakeTLInputPeerChannel(&mtproto.InputPeer{
				ChannelId: id,
			}).To_InputPeer()
		case int32(botapi.ChatIdChannel):
			return mtproto.MakeTLInputPeerChannel(&mtproto.InputPeer{
				ChannelId: id,
			}).To_InputPeer()
		case int32(botapi.ChatIdChannelPrivate):
			return mtproto.MakeTLInputPeerChannel(&mtproto.InputPeer{
				ChannelId: id,
			}).To_InputPeer()
		default:
			log.Errorf("invalid idType")
			return nil
		}
	default:
		log.Errorf("invalid chatId")
		return nil
	}
}

func ToInputFile(
	iFile *botapi.InputFileUpload,
	cb func(fileId int64, filePart int32, bytes []byte) error) *mtproto.InputFile {
	file := mtproto.MakeTLInputFile(&mtproto.InputFile{
		Id:          rand.Int63(),
		Parts:       0,
		Name:        iFile.FileName,
		Md5Checksum: "",
	}).To_InputFile()

	// fileId := rand.Int63()
	for i := 0; i < len(iFile.FileUpload); i = i + 512*1024 {
		if i+512*1024 > len(iFile.FileUpload) {
			file.Parts = int32(i + 1)
			cb(file.Id, int32(i), iFile.FileUpload[i:len(iFile.FileUpload)-i])
		} else {
			cb(file.Id, int32(i), iFile.FileUpload[i:i+512*1024])
		}
	}

	return file
}

func getWaveform2(sampleBuffer []byte) []byte {
	var (
		resultSamples = 100
		samples       = make([]uint16, 100)
		sampleIndex   = 0
		peakSample    = uint16(0)
		sampleRate    = math2.Max(1, len(sampleBuffer)/(2*resultSamples))
		index         = 0
	)

	for i := 0; i < len(sampleBuffer)/2; i++ {
		sample := uint16(math.Abs(float64(int16(binary.BigEndian.Uint16(sampleBuffer[2*i:])))))
		if sample > peakSample {
			peakSample = sample
		}
		if sampleIndex%sampleRate == 0 {
			if index < resultSamples {
				samples[index] = peakSample
				index++
			}
			peakSample = 0
		}
		sampleIndex += 1
	}

	var sumSamples int64 = 0
	for i := 0; i < resultSamples; i++ {
		sumSamples += int64(samples[i])
	}
	peak := uint16(float64(sumSamples) * 1.8 / float64(resultSamples))
	if peak < 2500 {
		peak = 2500
	}

	for i := 0; i < resultSamples; i++ {
		sample := samples[i]
		if sample > peak {
			samples[i] = peak
		}
	}

	bitstreamLength := uint32(resultSamples*5/8 + 1)
	bytes := make([]byte, bitstreamLength+4)
	for i := 0; i < resultSamples; i++ {
		value := math2.Min(31, int(math.Abs(float64(samples[i])*31)/float64(peak)))
		setBits(bytes, i*5, value&31)
	}

	return bytes
}

func setBits(bytes []byte, bitOffset int, value int) {
	bytes = bytes[bitOffset/8:]
	bitOffset = bitOffset % 8
	v := binary.BigEndian.Uint32(bytes)
	v |= uint32(value) << uint32(bitOffset)
	binary.BigEndian.PutUint32(bytes, v)
}
