package model

import (
	"regexp"
	"strings"

	"open.chat/mtproto"
)

func IsStickerDocument(document *mtproto.Document) bool {
	for _, attribute := range document.GetAttributes() {
		if attribute.PredicateName == mtproto.Predicate_documentAttributeSticker {
			return true
		}
	}
	return false
}

func IsMaskDocument(document *mtproto.Document) bool {
	for _, attribute := range document.GetAttributes() {
		if attribute.PredicateName == mtproto.Predicate_documentAttributeSticker && attribute.Mask {
			return true
		}
	}
	return false
}

func IsVoiceDocument(document *mtproto.Document) bool {
	for _, attribute := range document.GetAttributes() {
		if attribute.PredicateName == mtproto.Predicate_documentAttributeAudio {
			return attribute.Voice
		}
	}
	return false
}

func IsMusicDocument(document *mtproto.Document) bool {
	for _, attribute := range document.GetAttributes() {
		if attribute.PredicateName == mtproto.Predicate_documentAttributeAudio {
			return !attribute.Voice
		}

		if document.MimeType != "" {
			mime := strings.ToLower(document.MimeType)
			switch mime {
			case "audio/flac":
				return true
			case "audio/ogg":
				return true
			case "audio/opus":
				return true
			case "audio/x-opus+ogg":
				return true
			case "application/octet-stream":
				dFileName := getDocumentFileName(document)
				if strings.HasSuffix(dFileName, ".opus") {
					return true
				}
			}
		}
	}
	return false
}

func fixFileName(fileName string) (n string) {
	if fileName != "" {
		re, _ := regexp.Compile("[\u0001-\u001f<>:\"/\\\\|?*\u007f]+")
		n = strings.TrimSpace(re.ReplaceAllString(fileName, ""))
	}
	return
}

func getDocumentFileName(document *mtproto.Document) (fileName string) {
	for _, attribute := range document.GetAttributes() {
		if attribute.PredicateName == mtproto.Predicate_documentAttributeFilename {
			fileName = attribute.FileName
		}
	}
	return fixFileName(fileName)
}

func IsVideoDocument(document *mtproto.Document) bool {
	if document == nil {
		return false
	}

	var (
		isAnimated = false
		isVideo    = false
		width      int32
		height     int32
	)
	for _, attribute := range document.GetAttributes() {
		if attribute.PredicateName == mtproto.Predicate_documentAttributeVideo {
			if attribute.RoundMessage {
				return false
			}
			isVideo = true
			width = attribute.W
			height = attribute.H
		} else if attribute.PredicateName == mtproto.Predicate_documentAttributeAnimated {
			isAnimated = true
		}
	}

	if isAnimated && (width > 1280 || height > 1280) {
		isAnimated = false
	}
	if !isVideo && "video/x-matroska" == document.MimeType {
		isVideo = true
	}
	return isVideo && !isAnimated
}

func IsDocumentHasThumb(document *mtproto.Document) bool {
	if document == nil || len(document.Thumbs) > 0 {
		return false
	}

	for _, photoSize := range document.Thumbs {
		if photoSize != nil &&
			photoSize.PredicateName != mtproto.Predicate_photoSizeEmpty &&
			photoSize.Location != nil && photoSize.Location.PredicateName != mtproto.Predicate_fileLocationUnavailable {
			return true
		}

	}
	return false
}

func IsGifDocument(document *mtproto.Document) bool {
	return document != nil &&
		len(document.Thumbs) > 0 &&
		(document.MimeType == "image/gif" || IsNewGifDocument(document))
}

func IsRoundVideoDocument(document *mtproto.Document) bool {
	if document == nil || document.MimeType != "video/mp4" {
		return false
	}

	var (
		width  int32
		height int32
		round  bool
	)

	for _, attribute := range document.GetAttributes() {
		if attribute.PredicateName == mtproto.Predicate_documentAttributeVideo {
			width = attribute.W
			height = attribute.H
			round = attribute.RoundMessage
		}
	}
	if round && width <= 1280 && height <= 1280 {
		return true
	}
	return false
}

func IsNewGifDocument(document *mtproto.Document) bool {
	if document == nil || document.MimeType != "video/mp4" {
		return false
	}

	var (
		width    int32
		height   int32
		animated bool
	)

	for _, attribute := range document.GetAttributes() {
		if attribute.PredicateName == mtproto.Predicate_documentAttributeAnimated {
			animated = true
		} else if attribute.PredicateName == mtproto.Predicate_documentAttributeVideo {
			width = attribute.W
			height = attribute.H
		}
	}
	if animated && width <= 1280 && height <= 1280 {
		return true
	}
	return false
}
