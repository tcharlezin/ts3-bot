package teamspeak

import (
	"fmt"
	"github.com/multiplay/go-ts3"
	"time"
)

type Message struct {
	Client  *ts3.Client
	User    *User
	Channel int
	Msg     string
	Private bool
}

func (m *Message) SendMessageToUser(options ...ts3.CmdArg) error {

	var targetMode int
	if m.Private {
		targetMode = 1
	} else {
		targetMode = 2
	}

	options = append(options, ts3.NewArg("targetmode", targetMode), ts3.NewArg("target", m.User.Id), ts3.NewArg("msg", m.Msg))

	time.Sleep(1 * time.Second)

	_, err := m.Client.ExecCmd(ts3.NewCmd("sendtextmessage").WithArgs(options...))

	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func (m *Message) SendMessageChannel(options ...ts3.CmdArg) error {

	options = append(options, ts3.NewArg("targetmode", 2), ts3.NewArg("target", m.Channel), ts3.NewArg("msg", m.Msg))

	time.Sleep(1 * time.Second)

	_, err := m.Client.ExecCmd(ts3.NewCmd("sendtextmessage").WithArgs(options...))

	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func (m *Message) SendMessageGlobal(options ...ts3.CmdArg) error {

	options = append(options, ts3.NewArg("msg", m.Msg))

	time.Sleep(1 * time.Second)

	_, err := m.Client.ExecCmd(ts3.NewCmd("gm").WithArgs(options...))

	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
