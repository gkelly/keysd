package webrtc

import (
	"sync"

	"github.com/pion/webrtc/v2"
	"github.com/pkg/errors"
)

// Channel.
type Channel interface {
	Label() string
	OnClose(f func())
}

// Message in channel.
type Message interface {
	Data() []byte
}

type message struct {
	webrtc.DataChannelMessage
}

func (m message) Data() []byte {
	return m.DataChannelMessage.Data
}

type SessionDescription = webrtc.SessionDescription

// Client for webrtc.
type Client struct {
	sync.Mutex
	config  webrtc.Configuration
	conn    *webrtc.PeerConnection
	channel *webrtc.DataChannel

	openLn    func(channel Channel)
	closeLn   func(channel Channel)
	messageLn func(msg Message)
}

// NewClient creates webrtc Client.
func NewClient() (*Client, error) {
	config := webrtc.Configuration{
		ICEServers: []webrtc.ICEServer{
			{
				URLs: []string{"stun:stun.l.google.com:19302"},
			},
		},
	}

	c := &Client{
		config:    config,
		openLn:    func(channel Channel) {},
		closeLn:   func(channel Channel) {},
		messageLn: func(msg Message) {},
	}

	return c, nil
}

func (c *Client) newConnection() (*webrtc.PeerConnection, error) {
	conn, err := webrtc.NewPeerConnection(c.config)
	if err != nil {
		return nil, err
	}

	conn.OnICEConnectionStateChange(func(connectionState webrtc.ICEConnectionState) {
		logger.Infof("ICE: %s", connectionState.String())
	})

	conn.OnDataChannel(func(channel *webrtc.DataChannel) {
		c.register(channel)
	})

	return conn, nil
}

func (c *Client) register(channel *webrtc.DataChannel) {
	channel.OnOpen(func() {
		c.onOpen(channel)
	})
	channel.OnClose(func() {
		c.onClose(channel)
	})
	channel.OnMessage(func(m webrtc.DataChannelMessage) {
		c.onMessage(m)
	})
}

func (c *Client) Offer(label string) (*webrtc.SessionDescription, error) {
	c.Lock()
	defer c.Unlock()

	if c.conn != nil {
		return nil, errors.Errorf("connection already exists")
	}
	conn, err := c.newConnection()
	if err != nil {
		return nil, err
	}
	c.conn = conn

	channel, err := conn.CreateDataChannel(label, nil)
	if err != nil {
		return nil, err
	}
	c.register(channel)

	offer, err := conn.CreateOffer(nil)
	if err != nil {
		return nil, err
	}
	if err := conn.SetLocalDescription(offer); err != nil {
		return nil, err
	}

	return &offer, nil
}

func (c *Client) Answer(offer *webrtc.SessionDescription) (*webrtc.SessionDescription, error) {
	c.Lock()
	defer c.Unlock()

	if c.conn != nil {
		return nil, errors.Errorf("connection already exists")
	}
	conn, err := c.newConnection()
	if err != nil {
		return nil, err
	}
	c.conn = conn

	if err := conn.SetRemoteDescription(*offer); err != nil {
		return nil, err
	}
	answer, err := conn.CreateAnswer(nil)
	if err != nil {
		return nil, err
	}
	if err := conn.SetLocalDescription(answer); err != nil {
		return nil, err
	}
	return &answer, nil
}

func (c *Client) SetAnswer(answer *webrtc.SessionDescription) error {
	c.Lock()
	defer c.Unlock()

	if c.conn == nil {
		return errors.Errorf("no connection")
	}
	if err := c.conn.SetRemoteDescription(*answer); err != nil {
		return err
	}
	return nil
}

func (c *Client) Close() {
	c.Lock()
	defer c.Unlock()

	if c.channel != nil {
		if err := c.channel.Close(); err != nil {
			logger.Warningf("Error closing webrtc channel: %v", err)
		}
	}
	if c.conn != nil {
		if err := c.conn.Close(); err != nil {
			logger.Warningf("Error closing webrtc connection: %v", err)
		}
	}
}

func (c *Client) Channel() Channel {
	return c.channel
}

func (c *Client) onOpen(channel *webrtc.DataChannel) {
	c.channel = channel
	c.openLn(channel)
}

func (c *Client) onClose(channel *webrtc.DataChannel) {
	c.closeLn(channel)
}

func (c *Client) onMessage(m webrtc.DataChannelMessage) {
	c.messageLn(&message{m})
}

func (c *Client) OnOpen(f func(Channel)) {
	c.openLn = f
}

func (c *Client) OnClose(f func(Channel)) {
	c.closeLn = f
}

func (c *Client) OnMessage(f func(Message)) {
	c.messageLn = f
}

func (c *Client) Send(data []byte) error {
	if c.channel == nil {
		return errors.Errorf("no channel")
	}
	return c.channel.Send(data)
}