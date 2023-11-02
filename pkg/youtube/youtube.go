package youtube

import (
	"io"
	"os"
	"regexp"
	"strings"
)

type MusicFile struct {
	name string
	path string
}

type MusicFiles []MusicFile

func NewMusicFile(name, path string) MusicFile {
	return MusicFile{path: path, name: name}
}

func (f *MusicFile) Name() string {
	return f.name
}

func (f *MusicFile) NeedsUpload() bool {
	return true
}

func (f *MusicFile) UploadData() (string, io.Reader, error) {
	open, err := os.Open(f.path)
	return f.name, open, err
}

func (f *MusicFile) SendData() string {
	return ""
}

func (f *MusicFile) Delete() {
	if f.path != "" {
		_ = os.Remove(f.path)
	}
}

func (f MusicFiles) Delete() {
	for i := range f {
		f[i].Delete()
	}
}

func IDFromURL(url string) string {
	url = strings.TrimPrefix(url, `https:`)
	url = strings.TrimPrefix(url, `http:`)
	url = strings.TrimPrefix(url, `//`)
	url = strings.TrimPrefix(url, `www.`)
	url = strings.TrimPrefix(url, `m.`)
	url = strings.TrimPrefix(url, `music.`)
	url = strings.TrimPrefix(url, `youtu.be/`)
	url = strings.TrimPrefix(url, `youtube.com/`)
	url = strings.TrimPrefix(url, `youtube-nocookie.com/`)
	url = strings.TrimPrefix(url, `embed/`)
	url = strings.TrimPrefix(url, `shorts/`)
	url = strings.TrimPrefix(url, `v/`)
	url = strings.TrimPrefix(url, `live/`)
	url = strings.TrimPrefix(url, `watch?`)
	url = strings.TrimPrefix(url, `v=`)
	url = strings.TrimPrefix(url, `e/`)
	url = strings.TrimPrefix(url, `feature=player_embedded&v=`)
	url = strings.TrimPrefix(url, `app=desktop&v=`)
	url = strings.TrimPrefix(url, `attribution_link?a=`)

	url = strings.TrimSuffix(url, "\n")
	url = strings.Split(url, "?")[0]
	url = strings.Split(url, "&")[0]
	url = strings.Split(url, "#")[0]

	match, _ := regexp.Match("^[-_a-zA-Z0-9]+$", []byte(url))
	if !match {
		return ""
	}

	return url
}
