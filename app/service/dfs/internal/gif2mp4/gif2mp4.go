package gif2mp4

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

func GetVideoFirstFrame(filename string, width, height int) (*bytes.Buffer, error) {
	cmd := exec.Command("ffmpeg", "-i", filename, "-vframes", "1", "-s", fmt.Sprintf("%dx%d", width, height), "-f", "singlejpeg", "-")

	buf := new(bytes.Buffer)

	cmd.Stdout = buf

	buf2 := new(bytes.Buffer)
	cmd.Stderr = buf2

	if err := cmd.Run(); err != nil {
		return nil, err
	}

	return buf, nil
}

func Gif2Mp4Convert(gifPath string, videoExtension string) (string, error) {
	videoPath := gifPath + "." + videoExtension
	fmt.Println(videoPath)
	if _, err := os.Stat(videoPath); err == nil {
		return videoPath, nil
	}

	if videoExtension == "jpg" {
		o, err := exec.Command(
			"convert",
			gifPath+"[0]",
			videoPath,
		).CombinedOutput()
		if err != nil {
			fmt.Println(string(o))
			return "", err
		}
		return videoPath, nil
	} else if videoExtension == "webm" {
		o, err := exec.Command(
			"ffmpeg",
			"-i", gifPath,
			"-y",
			"-b:v", "5M",
			videoPath,
		).CombinedOutput()
		if err != nil {
			fmt.Println(string(o))
			return "", err
		}
		return videoPath, nil
	} else if videoExtension == "mp4" {
		fmt.Println(videoExtension, ", ", gifPath, ", ", videoPath)
		o, err := exec.Command(
			"ffmpeg",
			"-f", "gif",
			"-i", gifPath,
			"-movflags", "+faststart",
			"-pix_fmt", "yuv420p",
			"-vf", "scale=320:-2",
			"-c:v", "libx264",
			"-strict", "experimental",
			"-b:v", "218k",
			"-bufsize", "218k",
			videoPath,
		).CombinedOutput()

		fmt.Println(string(o), ", ", err)
		if err != nil {
			fmt.Println(string(o), ", ", err)
			return "", err
		}
		return videoPath, nil
	}

	return "", nil
}
