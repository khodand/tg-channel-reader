package youtube

import (
	"testing"
)

func TestIDFromURL(t *testing.T) {
	testCases := []struct {
		url string
		id  string
	}{
		{
			url: "https://www.youtube.com/live/oOMyJGlb6jw?feature=share",
			id:  "oOMyJGlb6jw",
		},
		{
			url: "https://youtu.be/qAo6fz2dTGw",
			id:  "qAo6fz2dTGw",
		},
		{
			url: "https://www.youtube.com/watch?v=K1PesAm4fHg",
			id:  "K1PesAm4fHg",
		},
		{
			url: "https://www.youtube.com/watch?v=3XtMX3FH_f4&t=48s",
			id:  "3XtMX3FH_f4",
		},
		{
			url: "https://music.youtube.com/watch?v=x5ldE1OE9jk&feature=share",
			id:  "x5ldE1OE9jk",
		},
		{
			url: "http://www.youtube.com/live/oOMyJGlb6jw?feature=share",
			id:  "oOMyJGlb6jw",
		},
		{
			url: "www.youtube.com/live/oOMyJGlb6jw?feature=share",
			id:  "oOMyJGlb6jw",
		},
		{
			url: "youtube.com/live/oOMyJGlb6jw?feature=share",
			id:  "oOMyJGlb6jw",
		},
		{
			url: "https://m.youtube.com/watch?v=0zM3nApSvMg",
			id:  "0zM3nApSvMg",
		},
		{
			url: "https://youtube.com/v/dQw4w9WgXcQ",
			id:  "dQw4w9WgXcQ",
		},
		{
			url: "https://youtube.com/v/-wtIMTCHWuI",
			id:  "-wtIMTCHWuI",
		},
		{
			url: "https://www.youtube-nocookie.com/embed/lalOy8Mbfdc?rel=0",
			id:  "lalOy8Mbfdc",
		},
	}

	for i := range testCases {
		id := IDFromURL(testCases[i].url)
		if id != testCases[i].id {
			t.Errorf("expected:%s actual:%s", testCases[i].id, id)
		}
	}
}

func TestAllFormats(t *testing.T) {
	// https://gist.github.com/rodrigoborgesdeoliveira/987683cfbfcc8d800192da1e73adc486
	tests := []string{
		"http://www.youtube.com/watch?v=lalOy8Mbfdc",
		"http://youtube.com/watch?v=lalOy8Mbfdc",
		"http://m.youtube.com/watch?v=lalOy8Mbfdc",
		"https://www.youtube.com/watch?v=lalOy8Mbfdc",
		"https://youtube.com/watch?v=lalOy8Mbfdc",
		"https://m.youtube.com/watch?v=lalOy8Mbfdc",
		"http://www.youtube.com/watch?v=lalOy8Mbfdc&feature=em-uploademail",
		"http://youtube.com/watch?v=lalOy8Mbfdc&feature=em-uploademail",
		"http://m.youtube.com/watch?v=lalOy8Mbfdc&feature=em-uploademail",
		"https://www.youtube.com/watch?v=lalOy8Mbfdc&feature=em-uploademail",
		"https://youtube.com/watch?v=lalOy8Mbfdc&feature=em-uploademail",
		"https://m.youtube.com/watch?v=lalOy8Mbfdc&feature=em-uploademail",
		"http://www.youtube.com/watch?v=lalOy8Mbfdc&feature=feedrec_grec_index",
		"http://youtube.com/watch?v=lalOy8Mbfdc&feature=feedrec_grec_index",
		"http://m.youtube.com/watch?v=lalOy8Mbfdc&feature=feedrec_grec_index",
		"https://www.youtube.com/watch?v=lalOy8Mbfdc&feature=feedrec_grec_index",
		"https://youtube.com/watch?v=lalOy8Mbfdc&feature=feedrec_grec_index",
		"https://m.youtube.com/watch?v=lalOy8Mbfdc&feature=feedrec_grec_index",
		"http://www.youtube.com/watch?v=lalOy8Mbfdc#t=0m10s",
		"http://youtube.com/watch?v=lalOy8Mbfdc#t=0m10s",
		"http://m.youtube.com/watch?v=lalOy8Mbfdc#t=0m10s",
		"https://www.youtube.com/watch?v=lalOy8Mbfdc#t=0m10s",
		"https://youtube.com/watch?v=lalOy8Mbfdc#t=0m10s",
		"https://m.youtube.com/watch?v=lalOy8Mbfdc#t=0m10s",
		"http://www.youtube.com/watch?v=lalOy8Mbfdc&feature=channel",
		"http://youtube.com/watch?v=lalOy8Mbfdc&feature=channel",
		"http://m.youtube.com/watch?v=lalOy8Mbfdc&feature=channel",
		"https://www.youtube.com/watch?v=lalOy8Mbfdc&feature=channel",
		"https://youtube.com/watch?v=lalOy8Mbfdc&feature=channel",
		"https://m.youtube.com/watch?v=lalOy8Mbfdc&feature=channel",
		"http://www.youtube.com/watch?v=lalOy8Mbfdc&playnext_from=TL&videos=osPknwzXEas&feature=sub",
		"http://youtube.com/watch?v=lalOy8Mbfdc&playnext_from=TL&videos=osPknwzXEas&feature=sub",
		"http://m.youtube.com/watch?v=lalOy8Mbfdc&playnext_from=TL&videos=osPknwzXEas&feature=sub",
		"https://www.youtube.com/watch?v=lalOy8Mbfdc&playnext_from=TL&videos=osPknwzXEas&feature=sub",
		"https://youtube.com/watch?v=lalOy8Mbfdc&playnext_from=TL&videos=osPknwzXEas&feature=sub",
		"https://m.youtube.com/watch?v=lalOy8Mbfdc&playnext_from=TL&videos=osPknwzXEas&feature=sub",
		"http://www.youtube.com/watch?v=lalOy8Mbfdc&feature=youtu.be",
		"http://youtube.com/watch?v=lalOy8Mbfdc&feature=youtu.be",
		"http://m.youtube.com/watch?v=lalOy8Mbfdc&feature=youtu.be",
		"https://www.youtube.com/watch?v=lalOy8Mbfdc&feature=youtu.be",
		"https://youtube.com/watch?v=lalOy8Mbfdc&feature=youtu.be",
		"https://m.youtube.com/watch?v=lalOy8Mbfdc&feature=youtu.be",
		"http://www.youtube.com/watch?v=lalOy8Mbfdc&feature=youtube_gdata_player",
		"http://youtube.com/watch?v=lalOy8Mbfdc&feature=youtube_gdata_player",
		"http://m.youtube.com/watch?v=lalOy8Mbfdc&feature=youtube_gdata_player",
		"https://www.youtube.com/watch?v=lalOy8Mbfdc&feature=youtube_gdata_player",
		"https://youtube.com/watch?v=lalOy8Mbfdc&feature=youtube_gdata_player",
		"https://m.youtube.com/watch?v=lalOy8Mbfdc&feature=youtube_gdata_player",
		"http://www.youtube.com/watch?v=lalOy8Mbfdc&list=PLGup6kBfcU7Le5laEaCLgTKtlDcxMqGxZ&index=106&shuffle=2655",
		"http://youtube.com/watch?v=lalOy8Mbfdc&list=PLGup6kBfcU7Le5laEaCLgTKtlDcxMqGxZ&index=106&shuffle=2655",
		"http://m.youtube.com/watch?v=lalOy8Mbfdc&list=PLGup6kBfcU7Le5laEaCLgTKtlDcxMqGxZ&index=106&shuffle=2655",
		"https://www.youtube.com/watch?v=lalOy8Mbfdc&list=PLGup6kBfcU7Le5laEaCLgTKtlDcxMqGxZ&index=106&shuffle=2655",
		"https://youtube.com/watch?v=lalOy8Mbfdc&list=PLGup6kBfcU7Le5laEaCLgTKtlDcxMqGxZ&index=106&shuffle=2655",
		"https://m.youtube.com/watch?v=lalOy8Mbfdc&list=PLGup6kBfcU7Le5laEaCLgTKtlDcxMqGxZ&index=106&shuffle=2655",
		"http://www.youtube.com/watch?feature=player_embedded&v=lalOy8Mbfdc",
		"http://youtube.com/watch?feature=player_embedded&v=lalOy8Mbfdc",
		"http://m.youtube.com/watch?feature=player_embedded&v=lalOy8Mbfdc",
		"https://www.youtube.com/watch?feature=player_embedded&v=lalOy8Mbfdc",
		"https://youtube.com/watch?feature=player_embedded&v=lalOy8Mbfdc",
		"https://m.youtube.com/watch?feature=player_embedded&v=lalOy8Mbfdc",
		"http://www.youtube.com/watch?app=desktop&v=lalOy8Mbfdc",
		"http://youtube.com/watch?app=desktop&v=lalOy8Mbfdc",
		"http://m.youtube.com/watch?app=desktop&v=lalOy8Mbfdc",
		"https://www.youtube.com/watch?app=desktop&v=lalOy8Mbfdc",
		"https://youtube.com/watch?app=desktop&v=lalOy8Mbfdc",
		"https://m.youtube.com/watch?app=desktop&v=lalOy8Mbfdc",
		"http://www.youtube.com/v/lalOy8Mbfdc",
		"http://youtube.com/v/lalOy8Mbfdc",
		"http://m.youtube.com/v/lalOy8Mbfdc",
		"https://www.youtube.com/v/lalOy8Mbfdc",
		"https://youtube.com/v/lalOy8Mbfdc",
		"https://m.youtube.com/v/lalOy8Mbfdc",
		"http://www.youtube.com/v/lalOy8Mbfdc?version=3&autohide=1",
		"http://youtube.com/v/lalOy8Mbfdc?version=3&autohide=1",
		"http://m.youtube.com/v/lalOy8Mbfdc?version=3&autohide=1",
		"https://www.youtube.com/v/lalOy8Mbfdc?version=3&autohide=1",
		"https://youtube.com/v/lalOy8Mbfdc?version=3&autohide=1",
		"https://m.youtube.com/v/lalOy8Mbfdc?version=3&autohide=1",
		"http://www.youtube.com/v/lalOy8Mbfdc?fs=1&hl=en_US&rel=0",
		"http://youtube.com/v/lalOy8Mbfdc?fs=1&hl=en_US&rel=0",
		"http://m.youtube.com/v/lalOy8Mbfdc?fs=1&hl=en_US&rel=0",
		"https://www.youtube.com/v/lalOy8Mbfdc?fs=1&amp;hl=en_US&amp;rel=0",
		"https://www.youtube.com/v/lalOy8Mbfdc?fs=1&hl=en_US&rel=0",
		"https://youtube.com/v/lalOy8Mbfdc?fs=1&hl=en_US&rel=0",
		"https://m.youtube.com/v/lalOy8Mbfdc?fs=1&hl=en_US&rel=0",
		"http://www.youtube.com/v/lalOy8Mbfdc?feature=youtube_gdata_player",
		"http://youtube.com/v/lalOy8Mbfdc?feature=youtube_gdata_player",
		"http://m.youtube.com/v/lalOy8Mbfdc?feature=youtube_gdata_player",
		"https://www.youtube.com/v/lalOy8Mbfdc?feature=youtube_gdata_player",
		"https://youtube.com/v/lalOy8Mbfdc?feature=youtube_gdata_player",
		"https://m.youtube.com/v/lalOy8Mbfdc?feature=youtube_gdata_player",
		"http://youtu.be/lalOy8Mbfdc",
		"https://youtu.be/lalOy8Mbfdc",
		"http://youtu.be/lalOy8Mbfdc?feature=youtube_gdata_player",
		"https://youtu.be/lalOy8Mbfdc?feature=youtube_gdata_player",
		"http://youtu.be/lalOy8Mbfdc?list=PLToa5JuFMsXTNkrLJbRlB--76IAOjRM9b",
		"https://youtu.be/lalOy8Mbfdc?list=PLToa5JuFMsXTNkrLJbRlB--76IAOjRM9b",
		"http://youtu.be/lalOy8Mbfdc&feature=channel",
		"https://youtu.be/lalOy8Mbfdc&feature=channel",
		"http://youtu.be/lalOy8Mbfdc?t=1",
		"http://youtu.be/lalOy8Mbfdc?t=1s",
		"https://youtu.be/lalOy8Mbfdc?t=1",
		"https://youtu.be/lalOy8Mbfdc?t=1s",
		"http://www.youtube.com/attribution_link?a=lalOy8Mbfdc&u=%2Fwatch%3Fv%3DEhxJLojIE_o%26feature%3Dshare",
		"http://youtube.com/attribution_link?a=lalOy8Mbfdc&u=%2Fwatch%3Fv%3DEhxJLojIE_o%26feature%3Dshare",
		"http://m.youtube.com/attribution_link?a=lalOy8Mbfdc&u=%2Fwatch%3Fv%3DEhxJLojIE_o%26feature%3Dshare",
		"https://www.youtube.com/attribution_link?a=lalOy8Mbfdc&u=%2Fwatch%3Fv%3DEhxJLojIE_o%26feature%3Dshare",
		"https://youtube.com/attribution_link?a=lalOy8Mbfdc&u=%2Fwatch%3Fv%3DEhxJLojIE_o%26feature%3Dshare",
		"https://m.youtube.com/attribution_link?a=lalOy8Mbfdc&u=%2Fwatch%3Fv%3DEhxJLojIE_o%26feature%3Dshare",
		"http://www.youtube.com/attribution_link?a=lalOy8Mbfdc&u=/watch%3Fv%3DyZv2daTWRZU%26feature%3Dem-uploademail",
		"http://youtube.com/attribution_link?a=lalOy8Mbfdc&u=/watch%3Fv%3DyZv2daTWRZU%26feature%3Dem-uploademail",
		"http://m.youtube.com/attribution_link?a=lalOy8Mbfdc&u=/watch%3Fv%3DyZv2daTWRZU%26feature%3Dem-uploademail",
		"https://www.youtube.com/attribution_link?a=lalOy8Mbfdc&u=/watch%3Fv%3DyZv2daTWRZU%26feature%3Dem-uploademail",
		"https://youtube.com/attribution_link?a=lalOy8Mbfdc&u=/watch%3Fv%3DyZv2daTWRZU%26feature%3Dem-uploademail",
		"https://m.youtube.com/attribution_link?a=lalOy8Mbfdc&u=/watch%3Fv%3DyZv2daTWRZU%26feature%3Dem-uploademail",
		"http://www.youtube.com/embed/lalOy8Mbfdc",
		"http://youtube.com/embed/lalOy8Mbfdc",
		"http://m.youtube.com/embed/lalOy8Mbfdc",
		"https://www.youtube.com/embed/lalOy8Mbfdc",
		"https://youtube.com/embed/lalOy8Mbfdc",
		"https://m.youtube.com/embed/lalOy8Mbfdc",
		"http://www.youtube.com/embed/lalOy8Mbfdc?rel=0",
		"http://youtube.com/embed/lalOy8Mbfdc?rel=0",
		"http://m.youtube.com/embed/lalOy8Mbfdc?rel=0",
		"https://www.youtube.com/embed/lalOy8Mbfdc?rel=0",
		"https://youtube.com/embed/lalOy8Mbfdc?rel=0",
		"https://m.youtube.com/embed/lalOy8Mbfdc?rel=0",
		"http://www.youtube-nocookie.com/embed/lalOy8Mbfdc?rel=0",
		"https://www.youtube-nocookie.com/embed/lalOy8Mbfdc?rel=0",
		"http://www.youtube.com/e/lalOy8Mbfdc",
		"http://youtube.com/e/lalOy8Mbfdc",
		"http://m.youtube.com/e/lalOy8Mbfdc",
		"https://www.youtube.com/e/lalOy8Mbfdc",
		"https://youtube.com/e/lalOy8Mbfdc",
		"https://m.youtube.com/e/lalOy8Mbfdc",
	}

	for i := range tests {
		id := IDFromURL(tests[i])
		if id != "lalOy8Mbfdc" {
			t.Errorf("url:%s, id:%s", tests[i], id)
		}
	}
}
