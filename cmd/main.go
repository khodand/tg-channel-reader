package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strings"

	"github.com/go-faster/errors"
	"github.com/gotd/td/telegram"
	"github.com/gotd/td/telegram/auth"
	"github.com/gotd/td/telegram/message"
	"github.com/gotd/td/telegram/updates"
	updhook "github.com/gotd/td/telegram/updates/hook"
	"github.com/gotd/td/tg"
	"go.uber.org/zap"

	pauth "github.com/khodand/tg-channel-reader/internal/auth"
	plog "github.com/khodand/tg-channel-reader/pkg/log"
)

const (
	appID   = 25350329
	appHash = "5306dd90b9a3d8a0ddc34d6f3325a419"
)

var (
	//channelsToRead = []string{
	//	"kastings",
	//	"fancypeolechannel",
	//	"Casting_by_magic_bear",
	//	"model_option",
	//	"APTUCTbI",
	//	"castgcd",
	//	"castingspb",
	//	"justcasting",
	//	"topcasting",
	//	"castings",
	//	"DenisovaAristovaCasting",
	//	"irinanashutinskaya_casting",
	//	"primepeople",
	//	"gogotovacasting",
	//	"artdashikopashiko",
	//	"nazmetova_kino",
	//	"rjurickcasting",
	//	"castelza",
	//	"kondrcast",
	//	"castingDAr",
	//	"tatazalinyan_cast",
	//	"castingTest",
	//}
	malePatterns   = []string{"парн", "мужч", "парен", "мальч", "актер ", "актёр ", "мужск", "юнош"}
	femalePatterns = []string{" деву", " женщ", "актри", " дево", " жен", " женс", " дев"}
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()
	if err := run(ctx); err != nil {
		panic(err)
	}
}

func run(ctx context.Context) error {
	log := plog.NewLogger(true)
	defer func() { _ = log.Sync() }()

	d := tg.NewUpdateDispatcher()
	gaps := updates.New(updates.Config{
		Handler: d,
		Logger:  log.Named("gaps"),
	})

	// Authentication flow handles authentication process, like prompting for code and 2FA password.
	flow := auth.NewFlow(pauth.Terminal{}, auth.SendCodeOptions{})
	client := telegram.NewClient(appID, appHash, telegram.Options{
		Logger:        log,
		UpdateHandler: gaps,
		Middlewares: []telegram.Middleware{
			updhook.UpdateHook(gaps.Handle),
		},
	})

	api := tg.NewClient(client)
	// Helper for sending messages.
	sender := message.NewSender(api)
	maleCastingChannel := sender.Resolve("maleCasting")
	femaleCastingChannel := sender.Resolve("womenCastings")

	// Setup message update handlers.
	d.OnNewChannelMessage(func(ctx context.Context, e tg.Entities, update *tg.UpdateNewChannelMessage) error {
		msg, ok := update.GetMessage().(*tg.Message)
		if !ok || msg.Out {
			log.Info("bad message")
			return nil
		}

		//_, ok := msg.GetEntities()
		//if !ok || msg.Out {
		//	log.Info("bad message")
		//	return nil
		//}

		channel, err := getChannel(ctx, api, msg)
		if err != nil {
			log.Error("get channel", zap.Error(err))
			return err
		}

		text := strings.ToLower(msg.GetMessage())

		for i := range malePatterns {
			if strings.Contains(text, malePatterns[i]) {
				_, err = maleCastingChannel.ForwardMessages(channel.AsInputPeer(), msg).Send(ctx)
				if err != nil {
					log.Error("send message", zap.Error(err))
				}
				break
			}
		}

		for i := range femalePatterns {
			if strings.Contains(text, femalePatterns[i]) {
				_, err = femaleCastingChannel.ForwardMessages(channel.AsInputPeer(), msg).Send(ctx)
				if err != nil {
					log.Error("send message", zap.Error(err))
				}
				break
			}
		}

		return nil
	})

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

		return gaps.Run(ctx, client.API(), user.ID, updates.AuthOptions{
			OnStart: func(ctx context.Context) {
				log.Info("Gaps started")
			},
		})
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
