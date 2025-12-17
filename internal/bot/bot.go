package bot

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/Lunarisnia/device-finder/internal/pingy"
	"github.com/Lunarisnia/device-finder/internal/tinycli"
	"github.com/bwmarrin/discordgo"
)

type phone struct {
	connected bool
	lastSeen  time.Time
}

var status phone

func sendNotification(s *discordgo.Session, discordId string, message string) error {
	channel, err := s.UserChannelCreate(discordId)
	if err != nil {
		return err
	}

	_, err = s.ChannelMessageSend(channel.ID, message)
	if err != nil {
		return err
	}

	return nil
}

func Run(ctx tinycli.Context) error {
	discordId := ctx.Argument("--target")
	deviceAddr := ctx.Argument("--ip")
	thres := ctx.Argument("--threshold")
	if thres == "" {
		thres = "5"
	}
	thresInt, err := strconv.Atoi(thres)
	if err != nil {
		return err
	}
	threshold := float64(thresInt)
	token := os.Getenv("BOT_TOKEN")
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		return err
	}

	go func() {
		err = dg.Open()
		if err != nil {
			fmt.Println("error opening connection,", err)
		}
	}()

	status.connected = true
	status.lastSeen = time.Now()
	for {
		fmt.Println("Starting Notification sequence")
		receivedPacket, err := pingy.Ping(deviceAddr)
		if err != nil {
			log.Println(err)
		}
		if receivedPacket == 0 && status.connected && time.Since(status.lastSeen).Minutes() >= threshold {
			status.connected = false
			sendNotification(dg, discordId, fmt.Sprintf("Your boyfriend is away from home! (Left at: %s)", status.lastSeen.Format(time.RFC850)))
		} else if receivedPacket > 0 && !status.connected {
			status.connected = true
			status.lastSeen = time.Now()
			sendNotification(dg, discordId, fmt.Sprintf("Your boyfriend is home! (Arrived at: %s)", status.lastSeen.Format(time.RFC850)))
		}
		fmt.Println("Sleeping for 5 seconds")
		time.Sleep(5 * time.Second)
	}
}
