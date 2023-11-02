package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"github.com/go-faster/errors"
	"github.com/gotd/td/telegram"
	"github.com/gotd/td/telegram/auth"
	"github.com/gotd/td/tg"
	pauth "github.com/khodand/tg-channel-reader/internal/auth"
	plog "github.com/khodand/tg-channel-reader/pkg/log"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"strings"
)

const (
	appID   = 25350329
	appHash = "5306dd90b9a3d8a0ddc34d6f3325a419"
	limit   = 200
)

var (
	channelIDs = []int64{
		1430811252, // Casting_by_magic_bear
		1576937030, // Casting_by_magic_bear
		1113840655, // castings
		1462226648, // castingspb
		1380801420, // model_option
		1252928410, // APTUCTbI
		1334194221, // irinanashutinskaya_casting
		1719385522, // rjurickcasting
		1732421069, // nazmetova_kino
		1249298788, // gogotovacasting
		1828926620, // castelza
	}
	malePatterns   = []string{"парн", "мужч", "парен", "мальч", "актер ", "актёр ", "мужск"}
	femalePatterns = []string{" дев", " жен", "актри"}
)

func main() {
	f, err := os.Create("data.csv")
	if err != nil {
		panic(err)
	}
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()
	if err := run(ctx, f); err != nil {
		panic(err)
	}
}

func run(ctx context.Context, file *os.File) error {
	log := plog.NewLogger(true)
	defer func() { _ = log.Sync() }()

	// Authentication flow handles authentication process, like prompting for code and 2FA password.
	flow := auth.NewFlow(pauth.Terminal{}, auth.SendCodeOptions{})
	client := telegram.NewClient(appID, appHash, telegram.Options{
		Logger: log,
	})

	api := tg.NewClient(client)
	w := csv.NewWriter(file)
	w.Write([]string{"label", "full_text"})

	return client.Run(ctx, func(ctx context.Context) error {
		// Perform auth if no session is available.
		if err := client.Auth().IfNecessary(ctx, flow); err != nil {
			return errors.Wrap(err, "auth")
		}

		// Fetch user info.
		user, err := client.Self(ctx)
		if err != nil {
			return errors.Wrap(err, "call self")
		}
		fmt.Println(user.ID, user.Contact, user.Fake)

		input := make([]tg.InputChannelClass, 0, len(channelIDs))
		for _, id := range channelIDs {
			input = append(input, &tg.InputChannel{ChannelID: id})
		}

		channels, err := api.ChannelsGetChannels(ctx, []tg.InputChannelClass{input[1]})
		if err != nil {
			return err
		}

		chats := channels.GetChats()
		for i := range chats {
			hs, err := api.MessagesGetHistory(ctx, &tg.MessagesGetHistoryRequest{
				Peer:  chats[i].(*tg.Channel).AsInputPeer(),
				Limit: limit,
			})
			if err != nil {
				return err
			}
			history, ok := hs.(*tg.MessagesChannelMessages)
			if !ok {
				log.Error("BAD HISTORY")
				return nil
			}

			for _, record := range history.Messages {
				msg, ok := record.(*tg.Message)
				if !ok {
					log.Warn("BAD MEssage", zap.Stringer("record", record))
					continue
				}
				if err := w.Write([]string{getLabel(msg.Message), msg.Message}); err != nil {
					log.Fatal("error writing record to csv", zap.Error(err))
				}
			}

		}
		w.Flush()
		if err := w.Error(); err != nil {
			log.Fatal("error flush", zap.Error(err))
		}

		return nil
	})
}

func getChannel(ctx context.Context, client *tg.Client, msg tg.NotEmptyMessage) (*tg.Channel, error) {
	ch, ok := msg.GetPeerID().(*tg.PeerChannel)
	if !ok {
		return nil, errors.New("bad peerID")
	}
	channelID := ch.GetChannelID()

	inputChannel := &tg.InputChannel{
		ChannelID:  channelID,
		AccessHash: 0,
	}
	channels, err := client.ChannelsGetChannels(ctx, []tg.InputChannelClass{inputChannel})

	if err != nil {
		return nil, fmt.Errorf("failed to fetch channel: %w", err)
	}

	if len(channels.GetChats()) == 0 {
		return nil, fmt.Errorf("no channels found")
	}

	return channels.GetChats()[0].(*tg.Channel), nil
}

func getLabel(text string) string {
	text = strings.ToLower(text)

	male, female := false, false
	for i := range malePatterns {
		if strings.Contains(text, malePatterns[i]) {
			male = true
			break
		}
	}

	for i := range femalePatterns {
		if strings.Contains(text, femalePatterns[i]) {
			female = true
			break
		}
	}

	switch {
	case male && female:
		return "0"
	case male:
		return "1"
	case female:
		return "2"
	default:
		return "3"
	}
}
