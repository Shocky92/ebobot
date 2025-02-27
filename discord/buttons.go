package discord

import "github.com/bwmarrin/discordgo"

type button struct {
	Label    string
	CustomID string
}

var buttons = []button{
	{Label: "badumtss", CustomID: "badumtss"},
	{Label: "feron", CustomID: "feron"},
	{Label: "goblin", CustomID: "goblin"},
	{Label: "happyment", CustomID: "happyment"},
	{Label: "pizda", CustomID: "pizda"},
	{Label: "kertcoin", CustomID: "kertcoin"},
	{Label: "omegalul", CustomID: "omegalul"},
	{Label: "pidaras", CustomID: "pidaras"},
	{Label: "who", CustomID: "who"},
}

func createButton(b button) discordgo.MessageComponent {
	return discordgo.Button{
		Label:    b.Label,
		Style:    discordgo.SuccessButton,
		Disabled: false,
		CustomID: b.CustomID,
	}
}

func createButtonRow(buttons []button) discordgo.MessageComponent {
	return discordgo.ActionsRow{
		Components: func() []discordgo.MessageComponent {
			var components []discordgo.MessageComponent
			for _, b := range buttons {
				components = append(components, createButton(b))
			}
			return components
		}(),
	}
}
