package main

import (
	"fmt"
	"os"
	"time"

	"github.com/bwmarrin/discordgo"
	_ "github.com/joho/godotenv/autoload"
	"github.com/natefinch/npipe"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/util/rand"
)

const PIPE = `\\.\pipe\GTATrilogyChaosModPipe`
const COOLDOWN = time.Minute

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	dg, err := discordgo.New("Bot " + os.Getenv("DISCORD_TOKEN"))
	if err != nil {
		return err
	}
	channel := os.Getenv("DISCORD_CHANNEL")
	fmt.Println("using channel:", channel)

	if err := dg.Open(); err != nil {
		return err
	}

	ln, err := npipe.Dial(PIPE)
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected!")

	effects := Effects()

	var discordMessage *discordgo.Message
	defer func() {
		if discordMessage != nil {
			dg.ChannelMessageDelete(channel, discordMessage.ID) // delete on shutdown
		}
	}()
	var remoteReactHandler func()
	var options [3]Effect
	var counts map[string]int // map from username to vote count

	last := time.Now()
	tc := time.NewTicker(time.Millisecond * 100)
	for range tc.C {
		message := fmt.Sprintf("time:%d,%d:-1:N/A:N/A:0", int64(COOLDOWN.Milliseconds())-time.Since(last).Milliseconds(), COOLDOWN.Milliseconds())

		_, err := ln.Write([]byte(message))
		if err != nil {
			return err
		}
		// fmt.Println("writing", message, n, err)

		if time.Since(last) > COOLDOWN {
			var title string
			if discordMessage != nil && len(counts) > 0 {
				list := []int{0, 0, 0}
				for _, choice := range counts {
					list[choice]++
				}

				winner := 0
				for i := range list {
					if i == 0 || list[i] > list[winner] {
						winner = i
					}
				}
				fmt.Println("WINNER!", winner, options[winner].Name())

				message := fmt.Sprintf(
					"votes:%s;%d;;%s;%d;;%s;%d;;%d",
					options[0].ID(), list[0],
					options[1].ID(), list[1],
					options[2].ID(), list[2],
					winner,
				)
				fmt.Println("writing", message)
				_, err := ln.Write([]byte(message))
				if err != nil {
					return err
				}

				effect := options[winner]
				message = fmt.Sprintf("%s:N/A:0", EffectToMessage(effect, COOLDOWN))
				fmt.Println("writing", message)
				_, err = ln.Write([]byte(message))
				if err != nil {
					return err
				}

				dg.ChannelMessageDelete(channel, discordMessage.ID)

				title = fmt.Sprintf("Winner: %s\n\nVote for the next effect!", options[winner].Name())
			} else {
				idx := rand.Intn(len(effects))
				effect := effects[idx]
				message := fmt.Sprintf("%s:N/A:0", EffectToMessage(effect, COOLDOWN))

				fmt.Println("writing", message)
				_, err := ln.Write([]byte(message))
				if err != nil {
					return err
				}

				title = "Vote for the next effect!"
			}

			options[0] = effects[rand.Intn(len(effects))]
			options[1] = effects[rand.Intn(len(effects))]
			options[2] = effects[rand.Intn(len(effects))]
			counts = make(map[string]int)

			discordMessage, err = dg.ChannelMessageSendEmbed(channel, &discordgo.MessageEmbed{
				Type:  discordgo.EmbedTypeRich,
				Title: title,
				// Description: "React to this message to vote for the next effect!",
				Fields: []*discordgo.MessageEmbedField{
					{Name: "A", Value: options[0].Name(), Inline: true},
					{Name: "B", Value: options[1].Name(), Inline: true},
					{Name: "C", Value: options[2].Name(), Inline: true},
				},
			})
			if err != nil {
				return errors.Wrap(err, "failed to send embed")
			}
			if err := dg.MessageReactionAdd(channel, discordMessage.ID, "🇦"); err != nil {
				return err
			}
			if err := dg.MessageReactionAdd(channel, discordMessage.ID, "🇧"); err != nil {
				return err
			}
			if err := dg.MessageReactionAdd(channel, discordMessage.ID, "🇨"); err != nil {
				return err
			}

			if remoteReactHandler != nil {
				remoteReactHandler()
			}
			remoteReactHandler = dg.AddHandler(func(s *discordgo.Session, e *discordgo.MessageReactionAdd) {
				if e.MessageID != discordMessage.ID {
					return
				}

				if _, ok := counts[e.UserID]; ok && e.UserID != discordMessage.Author.ID {
					// user has already voted
					dg.MessageReactionRemove(channel, discordMessage.ID, e.MessageReaction.Emoji.Name, e.UserID)
					return
				}

				switch e.MessageReaction.Emoji.Name {
				case "🇦":
					counts[e.UserID] = 0
				case "🇧":
					counts[e.UserID] = 1
				case "🇨":
					counts[e.UserID] = 2
				}
			})

			last = time.Now()
		}
	}
	return nil
}

// func main() {
// 	fmt.Println("starting to listen to", PIPE)
// 	ln, err := npipe.Listen(PIPE)
// 	if err != nil {
// 		panic(err)
// 	}
// 	for {
// 		conn, err := ln.Accept()
// 		if err != nil {
// 			panic(err)
// 		}
// 		go handleConnection(conn)
// 	}
// }

// func handleConnection(conn net.Conn) {
// 	// fmt.Println("handleConnection", conn.LocalAddr())

// 	// buf := make([]byte, 0, 4096)
// 	tmp := make([]byte, 256)
// 	for {
// 		_, err := conn.Read(tmp)
// 		if err != nil {
// 			if err != io.EOF {
// 				panic(err)
// 			}
// 			break
// 		}
// 		stmp := string(tmp)

// 		if !strings.HasPrefix(stmp, "time") {
// 			fmt.Println(string(tmp))
// 		}
// 	}
// }
