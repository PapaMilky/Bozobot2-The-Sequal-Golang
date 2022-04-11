package main

import (
	"Bozo_Bot_2_-_The_Sequel_/utils"
	"context"
	"encoding/json"
	"github.com/diamondburned/arikawa/v3/api"
	"github.com/diamondburned/arikawa/v3/discord"
	"github.com/diamondburned/arikawa/v3/gateway"
	"github.com/diamondburned/arikawa/v3/session"
	"github.com/diamondburned/arikawa/v3/utils/json/option"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

// To run, do `GUILD_ID="GUILD ID" BOT_TOKEN="TOKEN HERE" go run .`

func main() {
	guildID := discord.GuildID(mustSnowflakeEnv("GUILD_ID"))

	token := os.Getenv("BOT_TOKEN")
	if token == "" {
		log.Fatalln("no $BOT_TOKEN given")
	}

	s := session.New("Bot " + token)

	app, err := s.CurrentApplication()
	if err != nil {
		log.Fatalln("failed to get application ID:", err)
	}

	s.AddHandler(func(e *gateway.InteractionCreateEvent) {
		var resp api.InteractionResponse

		//var editResp api.EditInteractionResponseData

		switch data := e.Data.(type) {
		case *discord.CommandInteraction:
			switch data.Name {
			case "buttons":
				resp = api.InteractionResponse{
					Type: api.MessageInteractionWithSource,
					Data: &api.InteractionResponseData{
						Content: option.NewNullableString("This is a message with a button!"),
						Components: discord.ComponentsPtr(
							&discord.ActionRowComponent{
								&discord.ButtonComponent{
									Label:    "Hello World!",
									CustomID: "first_button",
									Emoji:    &discord.ComponentEmoji{Name: "👋"},
									Style:    discord.PrimaryButtonStyle(),
								},
								&discord.ButtonComponent{
									Label:    "Secondary",
									CustomID: "second_button",
									Style:    discord.SecondaryButtonStyle(),
								},
								&discord.ButtonComponent{
									Label:    "Success",
									CustomID: "success_button",
									Style:    discord.SuccessButtonStyle(),
								},
								&discord.ButtonComponent{
									Label:    "Danger",
									CustomID: "danger_button",
									Style:    discord.DangerButtonStyle(),
								},
							},
							// This is automatically put into its own row.
							&discord.ButtonComponent{
								Label: "Link",
								Style: discord.LinkButtonStyle("https://google.com"),
							},
						),
					},
				}
			case "test":
				var embed []discord.Embed
				var Author discord.EmbedAuthor
				var Footer discord.EmbedFooter
				var Fields []discord.EmbedField
				Author.Name = "Bozo Bot 2 - The Sequel"
				Author.Icon = "https://cdn.discordapp.com/avatars/961043988889088020/67faf6d2258d1674cfe186dbfe45574f.webp?size=80"
				Footer.Text = time.Now().Format("02-01-2006 Mon") + " | Thank You For Choosing Bozo Bot 2"
				Fields = append(Fields, discord.EmbedField{Inline: true, Name: "Cunt", Value: "Cock And Balls"})

				embed = append(embed, discord.Embed{Title: "This Is A Test Command For Demo Purposes", Description: "Demo Command", Author: &Author, Footer: &Footer, Fields: Fields})
				resp = api.InteractionResponse{
					Type: api.MessageInteractionWithSource,
					Data: &api.InteractionResponseData{
						Embeds: &embed,
					},
				}
			case "r34":

				var pendingEmbed []discord.Embed
				var Author discord.EmbedAuthor
				var Footer discord.EmbedFooter
				Author.Name = "Bozo Bot 2 - The Sequel"
				Author.Icon = "https://cdn.discordapp.com/avatars/961043988889088020/67faf6d2258d1674cfe186dbfe45574f.webp?size=80"
				Footer.Text = time.Now().Format("02-01-2006 Mon") + " | Thank You For Choosing Bozo Bot 2"

				var pendingEmbedField []discord.EmbedField

				pendingEmbedField = append(pendingEmbedField, discord.EmbedField{Inline: true, Name: "r34 API Request Is Pending", Value: "Thank You For Waiting!"})

				pendingEmbed = append(pendingEmbed, discord.Embed{Title: "RULE 34 BABYY", Description: "r34 Command", Author: &Author, Footer: &Footer, Fields: pendingEmbedField})

				resp = api.InteractionResponse{
					Type: api.MessageInteractionWithSource,
					Data: &api.InteractionResponseData{
						Embeds: &pendingEmbed,
					},
				}

				if err := s.RespondInteraction(e.ID, e.Token, resp); err != nil {
					log.Println("failed to send interaction callback:", err)
				}

				if len(data.Options) != 2 {
					break
				}

				body := utils.R34api(data.Options[0].String(), data.Options[1].String())

				var out []r34
				json.Unmarshal(body, &out)

				if len(out) != 2 {
					break
				}

				//fmt.Println(out)
				var embed []discord.Embed
				var embedImage0 discord.EmbedImage
				embedImage0.URL = out[0].FileUrl
				embedImage0.Width = uint(out[0].Width)
				embedImage0.Height = uint(out[0].Height)
				var embedImage1 discord.EmbedImage
				embedImage1.URL = out[1].FileUrl
				embedImage1.Width = uint(out[1].Width)
				embedImage1.Height = uint(out[1].Height)

				embed = append(embed, discord.Embed{Title: "RULE 34 BABYY", Description: "r34 Command", Author: &Author, Footer: &Footer, Image: &embedImage0})
				embed = append(embed, discord.Embed{Title: "RULE 34 BABYY", Description: "r34 Command", Author: &Author, Footer: &Footer, Image: &embedImage1})

				respData := api.EditInteractionResponseData{
					Embeds: &embed,
				}
				s.EditInteractionResponse(e.AppID, e.Token, respData)
			case "randn":
				rand.Seed(time.Now().UnixNano())
				randMin, _ := data.Options[0].IntValue()
				randMax, _ := data.Options[1].IntValue()
				numb := rand.Intn(int(randMax)-int(randMin)+1) + int(randMin)

				var embed []discord.Embed
				var Author discord.EmbedAuthor
				var Footer discord.EmbedFooter
				Author.Name = "Bozo Bot 2 - The Sequel"
				Author.Icon = "https://cdn.discordapp.com/avatars/961043988889088020/67faf6d2258d1674cfe186dbfe45574f.webp?size=80"
				Footer.Text = time.Now().Format("02-01-2006 Mon") + " | Thank You For Choosing Bozo Bot 2"

				var embedField []discord.EmbedField

				embedField = append(embedField, discord.EmbedField{Inline: true, Name: "Random Number:", Value: strconv.Itoa(numb)})

				embed = append(embed, discord.Embed{Title: "Random Number Generator", Description: "randn Command", Author: &Author, Footer: &Footer, Fields: embedField})

				resp = api.InteractionResponse{
					Type: api.MessageInteractionWithSource,
					Data: &api.InteractionResponseData{
						Embeds: &embed,
					},
				}
			case "randomurban":
				var pendingEmbed []discord.Embed
				var Author discord.EmbedAuthor
				var Footer discord.EmbedFooter
				Author.Name = "Bozo Bot 2 - The Sequel"
				Author.Icon = "https://cdn.discordapp.com/avatars/961043988889088020/67faf6d2258d1674cfe186dbfe45574f.webp?size=80"
				Footer.Text = time.Now().Format("02-01-2006 Mon") + " | Thank You For Choosing Bozo Bot 2"

				var pendingEmbedField []discord.EmbedField

				pendingEmbedField = append(pendingEmbedField, discord.EmbedField{Inline: true, Name: "urban API Request Is Pending", Value: "Thank You For Waiting!"})

				pendingEmbed = append(pendingEmbed, discord.Embed{Title: "URBAN DICTIONARY BABYY", Description: "urban Command", Author: &Author, Footer: &Footer, Fields: pendingEmbedField})

				resp = api.InteractionResponse{
					Type: api.MessageInteractionWithSource,
					Data: &api.InteractionResponseData{
						Embeds: &pendingEmbed,
					},
				}

				if err := s.RespondInteraction(e.ID, e.Token, resp); err != nil {
					log.Println("failed to send interaction callback:", err)
				}

				body := utils.RandomUrban()

				var out Urban
				json.Unmarshal(body, &out)

				f := len(out.List)
				i := f - 1

				var embed []discord.Embed
				var newembedField []discord.EmbedField

				for f != 0 {
					newembedField = append(newembedField, discord.EmbedField{Inline: true, Name: out.List[i].Word, Value: strings.Replace(strings.Replace(out.List[i].Definition, "]", "", -1), "[", "", -1)})
					//fmt.Println(out.List[i].Word)
					f--
					i--
				}

				embed = append(embed, discord.Embed{Title: "URBAN DICTIONARY BABYY", Description: "urban Command", Author: &Author, Footer: &Footer, Fields: newembedField})

				respData := api.InteractionResponseData{
					Embeds: &embed,
				}
				_, err := s.CreateInteractionFollowup(e.AppID, e.Token, respData)
				if err != nil {
					return
				}
			case "urbandict":
				break

			}

			// Send a message with a button back on slash commands.

		case discord.ComponentInteraction:
			resp = api.InteractionResponse{
				Type: api.UpdateMessage,
				Data: &api.InteractionResponseData{
					Content: option.NewNullableString("Custom ID: " + string(data.ID())),
				},
			}
		default:
			log.Printf("unknown interaction type %T", e.Data)
			return
		}

		//if err := s.RespondInteraction(e.ID, e.Token, resp); err != nil {
		//	log.Println("(respint) failed to send interaction callback:", err)
		//}
	})

	s.AddIntents(gateway.IntentGuilds)
	s.AddIntents(gateway.IntentGuildMessages)

	if err := s.Open(context.Background()); err != nil {
		log.Fatalln("failed to open:", err)
	}
	defer s.Close()

	log.Println("Gateway connected. Getting all guild commands.")

	commands, err := s.GuildCommands(app.ID, guildID)
	if err != nil {
		log.Fatalln("failed to get guild commands:", err)
	}

	for _, command := range commands {
		log.Println("Existing command", command.Name, "found.")
	}

	newCommands := []api.CreateCommandData{
		{
			Name:        "buttons",
			Description: "Send an interactable message.",
		},
		{
			Name:        "test",
			Description: "Test command to prove my point",
		},
		{
			Name:        "r34",
			Description: "Get your freak on",
			Options: []discord.CommandOption{
				&discord.StringOption{
					OptionName:   "tag",
					Description:  "Tags",
					Autocomplete: false,
					Required:     true,
				},
				&discord.IntegerOption{
					OptionName:   "page",
					Description:  "Page Number",
					Autocomplete: false,
					Required:     true,
				},
			},
		},
		{
			Name:        "randn",
			Description: "Generates A Random Number From A Given Range",
			Options: []discord.CommandOption{
				&discord.IntegerOption{
					OptionName:  "min",
					Description: "Minimum Number",
					Required:    true,
				},
				&discord.IntegerOption{
					OptionName:  "max",
					Description: "Maximum Number",
					Required:    true,
				},
			},
		},
		{
			Name:        "randomurban",
			Description: "Random Word And Definition From Urban Dictionary",
		},
		{
			Name:        "urbandict",
			Description: "Urban Dictionary Lookup",
			Options: []discord.CommandOption{
				&discord.StringOption{
					OptionName:  "lookup",
					Description: "Word Or Phrase To Lookup On The Dictionary",
					Required:    true,
				},
			},
		},
	}

	log.Println("Creating guild commands...")

	if _, err := s.BulkOverwriteGuildCommands(app.ID, guildID, newCommands); err != nil {
		log.Fatalln("failed to create guild command:", err)
	}

	log.Println("Guild commands created. Bot is ready.")

	// Block forever.
	select {}
}

func mustSnowflakeEnv(env string) discord.Snowflake {
	s, err := discord.ParseSnowflake(os.Getenv(env))
	if err != nil {
		log.Fatalf("Invalid snowflake for $%s: %v", env, err)
	}
	return s
}

type r34 struct {
	PreviewUrl   string `json:"preview_url"`
	SampleUrl    string `json:"sample_url"`
	FileUrl      string `json:"file_url"`
	Directory    int    `json:"directory"`
	Hash         string `json:"hash"`
	Height       int    `json:"height"`
	Id           int    `json:"id"`
	Image        string `json:"image"`
	Change       int    `json:"change"`
	Owner        string `json:"owner"`
	ParentId     int    `json:"parent_id"`
	Rating       string `json:"rating"`
	Sample       int    `json:"sample"`
	SampleHeight int    `json:"sample_height"`
	SampleWidth  int    `json:"sample_width"`
	Score        int    `json:"score"`
	Tags         string `json:"tags"`
	Width        int    `json:"width"`
}

type Urban struct {
	List []List `json:"list"`
}
type List struct {
	Definition  string    `json:"definition"`
	Permalink   string    `json:"permalink"`
	ThumbsUp    int       `json:"thumbs_up"`
	SoundUrls   []string  `json:"sound_urls"`
	Author      string    `json:"author"`
	Word        string    `json:"word"`
	Defid       int       `json:"defid"`
	CurrentVote string    `json:"current_vote"`
	WrittenOn   time.Time `json:"written_on"`
	Example     string    `json:"example"`
	ThumbsDown  int       `json:"thumbs_down"`
}
