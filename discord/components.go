package discord

import "github.com/bwmarrin/discordgo"

var (
	links = map[string]string{
		"badumtss":  "https://cdn.discordapp.com/attachments/209759858856165377/394033551294726154/7445930.jpg",
		"feron":     "http://image.prntscr.com/image/2bf970c168d14558a5a5c94f7fbb2c36.png",
		"goblin":    "https://youtu.be/TfV0QGthMmc",
		"happyment": "https://pp.userapi.com/c830400/v830400244/18f4d1/UszOuoziVEc.jpg",
		"kertcoin":  "https://i.imgur.com/po5nbTF.jpg",
		"omegalul":  "https://youtu.be/5ejsttVNPLA",
		"pidaras":   "https://media.discordapp.net/attachments/209759858856165377/642729886728847370/unknown.png?ex=672b7ada&is=672a295a&hm=d25106434000f23f88a322820270f6215b87552695f17268a49848ae59ba1f3d&=&format=webp&quality=lossless",
		"pizda":     "https://youtu.be/a3SarF-TOMc",
		"who":       "https://media.discordapp.net/attachments/209759858856165377/444549331199066122/who.PNG",
	}
)

func MakeComponentHandlers() map[string]func(*discordgo.Session, *discordgo.InteractionCreate) {
	handlers := make(map[string]func(*discordgo.Session, *discordgo.InteractionCreate))
	for k, v := range links {
		handlers[k] = func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: v,
				},
			})
		}
	}

	return handlers
}
