package spammer

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/df-mc/atomic"
	"github.com/mgutz/ansi"
	"github.com/opbteam/spammessage/data"
	"github.com/opbteam/spammessage/util"
	"github.com/sandertv/gophertunnel/minecraft"
	"github.com/sandertv/gophertunnel/minecraft/protocol/login"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

type MessageData struct {
	Message string
	Address string
	Kill    atomic.Bool
	Conn    *minecraft.Conn
	Log     *log.Logger
}

func (m *MessageData) Run() {
	conn, err := minecraft.Dialer{
		ClientData:  login.ClientData{},
		TokenSource: data.TokenSrc,
	}.Dial("raknet", m.Address)

	m.Conn = conn

	if err != nil {
		m.Log.Println("XBL: Error dialing", err)
	}

	err = conn.DoSpawn()

	if err != nil {
		m.Log.Println("XBL: Error spawning", err)
		return
	}

	m.Log.Println("XBL: " + conn.IdentityData().DisplayName + " Spawned on " + m.Address)

	s := bufio.NewScanner(os.Stdin)

	go func() {
		for {
			pk, _ := conn.ReadPacket()
			if err != nil {
				m.Log.Println("XBL: Error reading packet", err)
				return
			}
			if text, ok := pk.(*packet.Text); ok && !text.NeedsTranslation {
				fmt.Println(util.MinecraftToAscii(text.Message) + ansi.Reset)
			}
			conn.WritePacket(&packet.Text{TextType: packet.TextTypeChat, Message: m.Message})
		}
	}()

	for s.Scan() {
		if text := s.Text(); text == ":q" {
			m.Kill.Store(true)
			m.Conn.Close()
			fmt.Println("XBL: Stopped")
		}
	}
}
