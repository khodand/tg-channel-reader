package telegram

import "strings"

const (
	EntityTypeLink    = "link"
	EntityTypePlain   = "plain"
	EntityTypeHashtag = "hashtag"
)

type Data struct {
	Name     string      `json:"name"`
	Type     string      `json:"type"`
	ID       int         `json:"id"`
	Messages []MessageTg `json:"messages"`
}

type MessageTg struct {
	ID            int          `json:"id"`
	Type          string       `json:"type"`
	Date          string       `json:"date"`
	Text          []TextEntity `json:"text_entities"`
	From          string       `json:"from,omitempty"`
	FromId        string       `json:"from_id,omitempty"`
	ForwardedFrom string       `json:"forwarded_from"`
	MessageId     int          `json:"message_id,omitempty"`
}

type MessageFull struct {
	ID               int          `json:"id"`
	Type             string       `json:"type"`
	Date             string       `json:"date"`
	DateUnixtime     string       `json:"date_unixtime"`
	Actor            string       `json:"actor,omitempty"`
	ActorID          string       `json:"actor_id,omitempty"`
	Action           string       `json:"action,omitempty"`
	Text             []TextEntity `json:"text_entities"`
	Photo            string       `json:"photo,omitempty"`
	Width            int          `json:"width,omitempty"`
	Height           int          `json:"height,omitempty"`
	Edited           string       `json:"edited,omitempty"`
	EditedUnixtime   string       `json:"edited_unixtime,omitempty"`
	From             string       `json:"from,omitempty"`
	FromId           string       `json:"from_id,omitempty"`
	ForwardedFrom    string       `json:"forwarded_from"`
	MessageId        int          `json:"message_id,omitempty"`
	ReplyToMessageId int          `json:"reply_to_message_id,omitempty"`
	Poll             struct {
		Question    string `json:"question"`
		Closed      bool   `json:"closed"`
		TotalVoters int    `json:"total_voters"`
		Answers     []struct {
			Text   string `json:"text"`
			Voters int    `json:"voters"`
			Chosen bool   `json:"chosen"`
		} `json:"answers"`
	} `json:"poll,omitempty"`
	File            string `json:"file,omitempty"`
	Thumbnail       string `json:"thumbnail,omitempty"`
	MediaType       string `json:"media_type,omitempty"`
	MimeType        string `json:"mime_type,omitempty"`
	DurationSeconds int    `json:"duration_seconds,omitempty"`
}

type TextEntity struct {
	Type string `json:"type"`
	Text string `json:"text"`
	Href string `json:"href,omitempty"`
}

type Message struct {
	ID            int      `json:"id"`
	Date          string   `json:"date"`
	Text          string   `json:"text_entities"`
	Hashtags      []string `json:"hashtags,omitempty"`
	From          string   `json:"from,omitempty"`
	FromId        string   `json:"from_id,omitempty"`
	ForwardedFrom string   `json:"forwarded_from"`
	MessageId     int      `json:"message_id,omitempty"`
}

func (m *MessageTg) ToMessage() *Message {
	result := &Message{
		ID:            m.ID,
		Date:          m.Date,
		Hashtags:      make([]string, 0, 5),
		From:          "",
		FromId:        "",
		ForwardedFrom: "",
		MessageId:     0,
	}

	sb := strings.Builder{}
	for i := range m.Text {
		t := m.Text[i]
		switch t.Type {
		case EntityTypeLink:
			// pass
		case EntityTypeHashtag:
			result.Hashtags = append(result.Hashtags, t.Text)
		default:
			sb.WriteString(t.Text)
			sb.WriteString(" ")
		}
	}
	result.Text = sb.String()
	return result
}
