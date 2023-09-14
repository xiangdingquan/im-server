package model

import (
	"open.chat/mtproto"
)

type ReportReasonType int8

const (
	REASON_UNKNOWN        ReportReasonType = 0 // Unknown
	REASON_SPAM           ReportReasonType = 1 // Report for spam
	REASON_VIOLENCE       ReportReasonType = 2 // Report for violence
	REASON_PORNOGRAPHY    ReportReasonType = 3 // Report for pornography
	REASON_OTHER          ReportReasonType = 4 // Other
	REASON_COPYRIGHT      ReportReasonType = 5 // Report for copyrighted content
	REASON_CHILD_ABUSED   ReportReasonType = 6 // Report for child abuse
	REASON_GEO_IRRELEVANT ReportReasonType = 7 // Report an irrelevant geogroup
)

const (
	ACCOUNTS_reportPeer          = 0
	MESSAGES_reportSpam          = 1
	MESSAGES_report              = 2
	MESSAGES_reportEncryptedSpam = 3
	CHANNELS_reportSpam          = 4
)

func (i ReportReasonType) String() (s string) {
	switch i {
	case REASON_SPAM:
		s = "inputReportReasonSpam#58dbcab8 = ReportReason"
	case REASON_VIOLENCE:
		s = "inputReportReasonPornography#2e59d922 = ReportReason"
	case REASON_PORNOGRAPHY:
		s = "inputReportReasonOther#e1746d0a text:string = ReportReason"
	case REASON_OTHER:
		s = "inputReportReasonOther#e1746d0a text:string = ReportReason"
	case REASON_COPYRIGHT:
		s = "inputReportReasonCopyright#9b89f93a = ReportReason;"
	case REASON_CHILD_ABUSED:
		s = "inputReportReasonChildAbuse#adf44ee3 = ReportReason;"
	case REASON_GEO_IRRELEVANT:
		s = "inputReportReasonGeoIrrelevant#dbd4feed = ReportReason;"
	default:
		s = "unknown"
	}
	return
}

func FromReportReason(reason *mtproto.ReportReason) (i ReportReasonType, text string) {
	switch reason.PredicateName {
	case mtproto.Predicate_inputReportReasonSpam:
		i = REASON_SPAM
	case mtproto.Predicate_inputReportReasonViolence:
		i = REASON_VIOLENCE
	case mtproto.Predicate_inputReportReasonPornography:
		i = REASON_PORNOGRAPHY
	case mtproto.Predicate_inputReportReasonChildAbuse:
		i = REASON_CHILD_ABUSED
	case mtproto.Predicate_inputReportReasonOther:
		i = REASON_OTHER
		text = reason.Text
	case mtproto.Predicate_inputReportReasonCopyright:
		i = REASON_COPYRIGHT
	case mtproto.Predicate_inputReportReasonGeoIrrelevant:
		i = REASON_GEO_IRRELEVANT
	default:
		i = REASON_UNKNOWN
	}
	return
}

func (i ReportReasonType) ToReportReason(text string) (reason *mtproto.ReportReason) {
	switch i {
	case REASON_SPAM:
		reason = mtproto.MakeTLInputReportReasonSpam(&mtproto.ReportReason{
			Constructor: mtproto.CRC32_inputReportReasonSpam,
		}).To_ReportReason()
	case REASON_VIOLENCE:
		reason = mtproto.MakeTLInputReportReasonViolence(&mtproto.ReportReason{
			Constructor: mtproto.CRC32_inputReportReasonViolence,
		}).To_ReportReason()
	case REASON_PORNOGRAPHY:
		reason = mtproto.MakeTLInputReportReasonPornography(&mtproto.ReportReason{
			Constructor: mtproto.CRC32_inputReportReasonPornography,
		}).To_ReportReason()
	case REASON_CHILD_ABUSED:
		reason = mtproto.MakeTLInputReportReasonChildAbuse(&mtproto.ReportReason{
			Constructor: mtproto.CRC32_inputReportReasonChildAbuse,
		}).To_ReportReason()
	case REASON_OTHER:
		reason = mtproto.MakeTLInputReportReasonOther(&mtproto.ReportReason{
			Constructor: mtproto.CRC32_inputReportReasonOther,
			Text:        text,
		}).To_ReportReason()
	case REASON_COPYRIGHT:
		reason = mtproto.MakeTLInputReportReasonCopyright(&mtproto.ReportReason{
			Constructor: mtproto.CRC32_inputReportReasonCopyright,
		}).To_ReportReason()
	case REASON_GEO_IRRELEVANT:
		reason = mtproto.MakeTLInputReportReasonCopyright(&mtproto.ReportReason{
			Constructor: mtproto.CRC32_inputReportReasonGeoIrrelevant,
		}).To_ReportReason()
	default:
		reason = mtproto.MakeTLInputReportReasonOther(&mtproto.ReportReason{
			Constructor: mtproto.CRC32_inputReportReasonOther,
		}).To_ReportReason()
	}
	return
}
