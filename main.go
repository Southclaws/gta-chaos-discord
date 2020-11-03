package main

import (
	"fmt"
	"os"
	"time"

	"github.com/Southclaws/go-multicloser"
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

type bot struct {
	dg              *discordgo.Session
	pipe            *npipe.PipeConn
	channel         string
	votemsg         *discordgo.Message
	effects         []Effect
	delReactHandler func()
	options         [3]Effect
	counts          map[string]int // map from username to vote count
}

func run() error {
	b := bot{}

	b.channel = os.Getenv("DISCORD_CHANNEL")
	fmt.Println("using channel:", b.channel)

	mc := multicloser.MultiErrorMultiCloser{}
	b.dg = mc.Add(discordgo.New("Bot " + os.Getenv("DISCORD_TOKEN"))).(*discordgo.Session)
	b.pipe = mc.Add(npipe.Dial(PIPE)).(*npipe.PipeConn)
	defer mc.Close()

	// connect to discord
	if err := b.dg.Open(); err != nil {
		return err
	}
	fmt.Println("Connected!")

	b.effects = Effects()

	defer b.deleteVoteMessage()

	last := time.Now()
	tc := time.NewTicker(time.Millisecond * 100)
	for {
		select {
		// TODO: sync votes
		case <-tc.C:
			b.write(fmt.Sprintf("time:%d,%d,%s:", int64(COOLDOWN.Milliseconds())-time.Since(last).Milliseconds(), COOLDOWN.Milliseconds() /*mode:*/, ""), false)

			if time.Since(last) < COOLDOWN {
				continue
			}

			if err := b.handleEffect(); err != nil {
				fmt.Println("failed to handle effect:", err)
			}

			last = time.Now()
		}
	}
}

func (b *bot) handleEffect() (err error) {
	var title string
	if b.votemsg != nil && len(b.counts) > 0 {
		list := []int{0, 0, 0}
		for _, choice := range b.counts {
			list[choice]++
		}

		winner := 0
		for i := range list {
			if i == 0 || list[i] > list[winner] {
				winner = i
			}
		}
		fmt.Println("[votes] Winner:", b.options[winner].Name())

		b.write(fmt.Sprintf(
			"votes:%s;%d;;%s;%d;;%s;%d;;%d",
			b.options[0].ID(), list[0],
			b.options[1].ID(), list[1],
			b.options[2].ID(), list[2],
			winner,
		), true)

		effect := b.options[winner]
		b.write(fmt.Sprintf("%s:N/A:0", EffectToMessage(effect, COOLDOWN)), true)

		b.dg.ChannelMessageDelete(b.channel, b.votemsg.ID)

		title = fmt.Sprintf("Winner: %s (%d votes)\n\nVote for the next effect!", b.options[winner].Name(), list[winner])
	} else {
		idx := rand.Intn(len(b.effects))
		effect := b.effects[idx]
		message := fmt.Sprintf("%s:N/A:0", EffectToMessage(effect, COOLDOWN))

		fmt.Println("writing", message)
		_, err = b.pipe.Write([]byte(message))
		if err != nil {
			return err
		}

		title = "Vote for the next effect!"
	}

	b.options[0] = b.effects[rand.Intn(len(b.effects))]
	b.options[1] = b.effects[rand.Intn(len(b.effects))]
	b.options[2] = b.effects[rand.Intn(len(b.effects))]
	b.counts = make(map[string]int)

	b.votemsg, err = b.dg.ChannelMessageSendEmbed(b.channel, &discordgo.MessageEmbed{
		Type:  discordgo.EmbedTypeRich,
		Title: title,
		// Description: "React to this message to vote for the next effect!",
		Fields: []*discordgo.MessageEmbedField{
			{Name: "A", Value: b.options[0].Name(), Inline: true},
			{Name: "B", Value: b.options[1].Name(), Inline: true},
			{Name: "C", Value: b.options[2].Name(), Inline: true},
		},
	})
	if err != nil {
		return errors.Wrap(err, "failed to send embed")
	}
	if err := b.dg.MessageReactionAdd(b.channel, b.votemsg.ID, "🇦"); err != nil {
		return err
	}
	if err := b.dg.MessageReactionAdd(b.channel, b.votemsg.ID, "🇧"); err != nil {
		return err
	}
	if err := b.dg.MessageReactionAdd(b.channel, b.votemsg.ID, "🇨"); err != nil {
		return err
	}

	if b.delReactHandler != nil {
		b.delReactHandler()
	}
	b.delReactHandler = b.dg.AddHandler(func(s *discordgo.Session, e *discordgo.MessageReactionAdd) {
		if e.MessageID != b.votemsg.ID {
			return
		}

		if _, ok := b.counts[e.UserID]; ok && e.UserID != b.votemsg.Author.ID {
			// user has already voted
			b.dg.MessageReactionRemove(b.channel, b.votemsg.ID, e.MessageReaction.Emoji.Name, e.UserID)
			return
		}

		switch e.MessageReaction.Emoji.Name {
		case "🇦":
			b.counts[e.UserID] = 0
		case "🇧":
			b.counts[e.UserID] = 1
		case "🇨":
			b.counts[e.UserID] = 2
		}
	})

	return nil
}

func (b *bot) deleteVoteMessage() {
	if b.votemsg != nil {
		b.dg.ChannelMessageDelete(b.channel, b.votemsg.ID)
	}
}

func (b *bot) write(message string, print bool) {
	if print {
		fmt.Println("[pipe] writing: '%s'", message)
	}
	_, err := b.pipe.Write([]byte(message))
	if err != nil {
		panic(err)
	}
}

// func main() {
// 	fmt.Println("starting to listen to", PIPE)
// 	ln, err := npipe.Listen(PIPE)
// 	if err != nil {
// 		panic(err)
// 	}
// 	for {
// 		conn, err := b.pipe.Accept()
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
