package centrifugo

import (
	"github.com/gorilla/websocket"
	"sync"
)

type Client struct {
	sync.RWMutex

	URL       string
	conn      *websocket.Conn
	connected bool
}

func NewClient(url string) *Client {
	return &Client{URL: url}
}

func (c *Client) Connect() error {
	c.Lock()
	defer c.Unlock()
	conn, _, err := websocket.DefaultDialer.Dial(c.URL, nil)
	if err != nil {
		return err
	}
	c.conn = conn
	c.connected = true
	return nil
}

func (c *Client) Connected() bool {
	c.Lock()
	defer c.Unlock()
	return c.connected
}

func (c *Client) Close() {
	c.Lock()
	defer c.Unlock()
	if c.conn != nil {
		c.conn.Close()
		c.conn = nil
	}
	c.connected = false
}
