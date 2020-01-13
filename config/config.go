package config

import (
	"log"
	"net"

	"github.com/spf13/viper"
)

// Config holds the global server-wide config.
type Config struct {
	HostIP string `mapstructure:"host_ip"`

	Database Database
	Launcher Launcher
	Sign     Sign
	Channel  Channel
	Entrance Entrance
}

// Database holds the postgres database config.
type Database struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
}

// Launcher holds the launcher server config.
type Launcher struct {
	Port int
}

// Sign holds the sign server config.
type Sign struct {
	Port int
}

// Channel holds the channel server config.
type Channel struct {
	Port int
}

// Entrance holds the entrance server config.
type Entrance struct {
	Port    uint16
	Entries []EntranceServerInfo
}

// EntranceServerInfo represents an entry in the serverlist.
type EntranceServerInfo struct {
	IP     string
	Unk2   uint16
	Type   uint8  // Server type. 0=?, 1=open, 2=cities, 3=newbie, 4=bar
	Season uint8  // Server activity. 0 = green, 1 = orange, 2 = blue
	Unk6   uint8  // Something to do with server recommendation on 0, 3, and 5.
	Name   string // Server name, 66 byte null terminated Shift-JIS(JP) or Big5(TW).

	// 4096(PC, PS3/PS4)?, 8258(PC, PS3/PS4)?, 8192 == nothing?
	// THIS ONLY EXISTS IF Binary8Header.type == "SV2", NOT "SVR"!
	AllowedClientFlags uint32

	Channels []EntranceChannelInfo
}

// EntranceChannelInfo represents an entry in a server's channel list.
type EntranceChannelInfo struct {
	Port           uint16
	MaxPlayers     uint16
	CurrentPlayers uint16
	Unk4           uint16
	Unk5           uint16
	Unk6           uint16
	Unk7           uint16
	Unk8           uint16
	Unk9           uint16
	Unk10          uint16
	Unk11          uint16
	Unk12          uint16
	Unk13          uint16
}

// getOutboundIP4 gets the preferred outbound ip4 of this machine
// From https://stackoverflow.com/a/37382208
func getOutboundIP4() net.IP {
	conn, err := net.Dial("udp4", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP.To4()
}

// LoadConfig loads the given config toml file.
func LoadConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	c := &Config{}
	err = viper.Unmarshal(c)
	if err != nil {
		return nil, err
	}

	if c.HostIP == "" {
		c.HostIP = getOutboundIP4().To4().String()
	}

	return c, nil
}
