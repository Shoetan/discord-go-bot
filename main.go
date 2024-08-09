package main

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	

	"github.com/bwmarrin/discordgo"
)

func main(){

	// get discord bot token

	token := GetEnv("TOKEN")

	//create a session 
	session, err := discordgo.New("Bot " + token)

	if err != nil {
		fmt.Println(err.Error())
	}

	session.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate){
		//check the the person sending the message is not the bot
		if m.Author.ID == s.State.User.ID {
			return
		}

		switch {
		case strings.Contains(m.Content, "Hello"):
			s.ChannelMessageSend(m.ChannelID, "How are you😊 ")
		case strings.Contains(m.Content, "What do you think?"):
			s.ChannelMessageSend(m.ChannelID, "About what? ")
		case strings.Contains(m.Content, "Bye"):
			s.ChannelMessageSend(m.ChannelID, "Adios 😉")
						
		}
	})

	session.Identify.Intents = discordgo.IntentsAllWithoutPrivileged

	err = session.Open()

	if err != nil {
		fmt.Println(err.Error())
	}

	defer session.Close()


	fmt.Println("The bot is online")


	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	


}


