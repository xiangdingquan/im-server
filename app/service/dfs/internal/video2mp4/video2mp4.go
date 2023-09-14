package video2mp4

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

func GetVideoFirstFrame(filename string) (*bytes.Buffer, error) {
	cmd := exec.Command("ffmpeg", "-i", filename, "-vframes", "1", "-f", "singlejpeg", "-")
	buf := new(bytes.Buffer)
	cmd.Stdout = buf
	buf2 := new(bytes.Buffer)
	cmd.Stderr = buf2
	if err := cmd.Run(); err != nil {
		return nil, err
	}
	return buf, nil
}

func Video2Mp4Convert(filePath string) (string, error) {
	videoPath := filePath + ".mp4"
	fmt.Println(videoPath)
	if _, err := os.Stat(videoPath); err == nil {
		return videoPath, nil
	}

	fmt.Println(filePath, ", ", videoPath)
	o, err := exec.Command(
		"ffmpeg",
		"-i", filePath,
		"-movflags", "+faststart",
		"-pix_fmt", "yuv420p",
		"-vf", "scale=320:-2",
		"-c:v", "libx264",
		videoPath,
	).CombinedOutput()
	fmt.Println(string(o), ", ", err)
	if err != nil {
		fmt.Println(string(o), ", ", err)
		return "", err
	}
	return videoPath, nil
}
