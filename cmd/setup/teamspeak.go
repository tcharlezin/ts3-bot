package setup

import (
	"github.com/multiplay/go-ts3"
	"log"
	"os"
	"strconv"
	"time"
)

func SetupTeamspeak() *ts3.Client {

	var counts int
	counts = 0

	var c *ts3.Client
	var err error

	for {
		c, err = ts3.NewClient(os.Getenv("BOT_HOST"))
		if err != nil {
			log.Println("[TeamSpeak Setup] Teamspeak not yet ready...")
			counts++
		} else {
			log.Println("[TeamSpeak Setup] Connected to Teamspeak")
			break
		}

		if counts > 10 {
			log.Println(err)
			log.Fatal("[TeamSpeak Setup] Can't connect to Teamspeak!")
			return nil
		}

		log.Println("[TeamSpeak Setup] Backing off for five seconds")
		time.Sleep(5 * time.Second)
	}

	if err := c.Login(os.Getenv("BOT_USERNAME"), os.Getenv("BOT_PASSWORD")); err != nil {
		log.Println("c.Login")
		log.Fatal(err)
	}

	err = c.Use(1)
	if err != nil {
		log.Println("c.Use")
		log.Fatal(err)
	}

	port, _ := strconv.Atoi(os.Getenv("BOT_PORT"))
	err = c.UsePort(port)
	if err != nil {
		log.Println("c.UsePort")
		log.Fatal(err)
	}

	err = c.Register(ts3.TextChannelEvents)
	if err != nil {
		log.Println("c.Register")
		log.Fatal(err)
	}

	_ = c.SetNick(os.Getenv("BOT_NICK"))

	return c
}
