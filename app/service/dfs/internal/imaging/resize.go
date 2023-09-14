package imaging

import (
	"bytes"
	"image"
	"strings"

	"github.com/disintegration/imaging"

	"open.chat/pkg/bytes2"
	"open.chat/pkg/log"
)

const (
	szOriginalType = "0"
	szSmallType    = "s"
	szMediumType   = "m"
	szXLargeType   = "x"
	szYLargeType   = "y"
	szAType        = "a"
	szBType        = "b"
	szCType        = "c"
)
const (
	szOriginalSize = 0
	szSmallSize    = 90
	szMediumSize   = 320
	szXLargeSize   = 800
	szYLargeSize   = 1280
	szASize        = 160
	szBSize        = 320
	szCSize        = 640

	szAIndex = 4
)

var sizeList = []int{
	szOriginalSize,
	szSmallSize,
	szMediumSize,
	szXLargeSize,
	szYLargeSize,
	szASize,
	szBSize,
	szCSize,
}

func GetSizeType(idx int) string {
	switch idx {
	case 0:
		return szOriginalType
	case 1:
		return szSmallType
	case 2:
		return szMediumType
	case 3:
		return szXLargeType
	case 4:
		return szYLargeType
	case 5:
		return szAType
	case 6:
		return szBType
	case 7:
		return szCType
	}

	return ""
}

type resizeInfo struct {
	isWidth bool
	size    int
}

func makeResizeInfo(img image.Image) resizeInfo {
	w := img.Bounds().Dx()
	h := img.Bounds().Dy()

	if w >= h {
		return resizeInfo{
			isWidth: true,
			size:    w,
		}
	} else {
		return resizeInfo{
			isWidth: false,
			size:    h,
		}
	}
}

func getImageFormat(extName string) (int, error) {
	formats := map[string]imaging.Format{
		".jpg":  imaging.JPEG,
		".jpeg": imaging.JPEG,
		".png":  imaging.PNG,
		".tif":  imaging.TIFF,
		".tiff": imaging.TIFF,
		".bmp":  imaging.BMP,
		".gif":  imaging.GIF,
	}

	ext := strings.ToLower(extName)
	f, ok := formats[ext]
	if !ok {
		return -1, imaging.ErrUnsupportedFormat
	}

	return int(f), nil
}

func ReSizeImage(rb []byte, extName string, isABC bool, cb func(szType int, w, h int32, b []byte) error) (err error) {
	var (
		img image.Image
		f   int
	)

	img, err = imaging.Decode(bytes.NewReader(rb))
	if err != nil {
		log.Errorf("decode r(%d) error: %v", len(rb), err)
		return
	}
	imgSz := makeResizeInfo(img)
	willBreak := false
	for i, sz := range sizeList {
		if willBreak {
			break
		}
		if i != 0 {
			if isABC {
				if i <= szAIndex {
					continue
				}
			} else {
				if i > szAIndex {
					continue
				}
			}
		}

		if i == 0 {
			err = cb(i, int32(img.Bounds().Dx()), int32(img.Bounds().Dy()), rb)
			if err != nil {
				return
			}
			continue
		}

		rsz := sz
		if imgSz.size < sz {
			if isABC {
				if i+1 < szAIndex {
					rsz = imgSz.size
				} else {
					break
				}
			} else {
				if i+1 < len(sizeList) {
					rsz = imgSz.size
				} else {
					break
				}
			}
			willBreak = true
		}

		var dst *image.NRGBA
		if imgSz.isWidth {
			dst = imaging.Resize(img, rsz, 0, imaging.Lanczos)
		} else {
			dst = imaging.Resize(img, 0, rsz, imaging.Lanczos)
		}

		f, err = getImageFormat(extName)
		if err != nil {
			log.Error(err.Error())
			return
		}

		o := bytes2.NewBuffer(make([]byte, 0, len(rb)))
		if f == int(imaging.JPEG) {
			err = imaging.Encode(o, dst, imaging.JPEG)
		} else {
			err = imaging.Encode(o, dst, imaging.Format(f))
		}

		if err != nil {
			log.Error(err.Error())
			return
		}
		err = cb(i, int32(dst.Bounds().Dx()), int32(dst.Bounds().Dy()), o.Bytes())
		if err != nil {
			return
		}
	}
	return
}
