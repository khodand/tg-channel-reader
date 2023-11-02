package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

func main() {
	// ./yt-dlp -f "ba[ext=m4a][abr<200]" --print after_move:filepath -o "songs/%(title)s-%(id)s.%(ext)s" -x --audio-format opus ZEciML8dnCE
	// ./yt-dlp -f "ba[ext=m4a][abr<200]" -q --print after_move:filepath -o "songs/%(title)s-%(id)s.%(ext)s" 5ldE1OE9jk
	// ./yt-dlp -f "ba[ext=m4a][abr<200]" -q --print after_move:filepath -o "songs/%(title)s-%(id)s.%(ext)s" --audio-format ogg ZEciML8dnCE
	ctx, _ := context.WithTimeout(context.Background(), 300*time.Second)
	cmd := exec.CommandContext(ctx, "./yt-dlp",
		"-f", "ba[ext=m4a][abr<200]",
		"-q",
		"--print", "after_move:filepath",
		"-o", "songs/%(title)s-%(id)s.%(ext)s",
		//"--embed-metadata",
		//"--max-filesize", "49M",
		"--",
		"3XtMX3FH_f4")
	fmt.Println(cmd.String())
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("output", output)
		fmt.Println("err", err)
		return
	}

	fmt.Println("output", string(output))
	fmt.Println("err", err)
	//
	//path := strings.TrimSuffix(string(output), "\n")
	//fmt.Println("out", path)
	//fmt.Println(songDuration(path))
	//
	//names := strings.Split(path, "/")
	//name := names[len(names)-1]
	//cutFile(10, 10, path, "songs/part2_"+name)
	//cancel()
}

func songDuration(path string) int {
	cmd := exec.Command("ffprobe",
		"-i", path,
		"-show_entries", "format=duration",
		"-v", "quiet",
		"-of", "json")
	fmt.Println(cmd.String())
	output, err := cmd.Output()
	fmt.Println(err)
	var out ffprobe
	err = json.Unmarshal(output, &out)
	if err != nil {
		fmt.Println(err)
	}
	v, _ := strconv.Atoi(strings.Split(out.Format.Duration, ".")[0])
	return v + 1
}

type ffprobe struct {
	Format struct {
		Duration string `json:"duration"`
	} `json:"format"`
}

func cutFile(from, duration int64, path, newName string) {
	//  ffmpeg -ss 00:00:00 -i "doin' time - lana del rey [edit audio]-Rt3Uy_Q3Rfw.m4a" -t 00:00:10 -c copy out1.m4a
	cmd := exec.Command("ffmpeg",
		"-ss", time.Time{}.Add(time.Duration(from)*time.Second).Format("15:04:05"),
		"-i", path,
		"-t", time.Time{}.Add(time.Duration(duration)*time.Second).Format("15:04:05"),
		"-v", "quiet",
		"-c", "copy", newName)
	fmt.Println(cmd.String())
	output, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("out", string(output))
}
