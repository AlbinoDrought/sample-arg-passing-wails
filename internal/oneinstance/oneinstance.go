package oneinstance

import (
	"encoding/json"
	"io"
	"os"
)

func Ensure() (<-chan string, error) {
	ln, err := listen()
	if err != nil {
		// Something is already listening. Send it a message.
		saySomethingToTheOwningInstance()
		return nil, err
	}

	c := make(chan string)

	go func() {
		defer close(c)
		defer ln.Close()
		for {
			// Something is sending us a message,
			conn, err := ln.Accept()
			if err != nil {
				return
			}
			data, _ := io.ReadAll(conn)
			conn.Close()
			if data == nil {
				data = []byte{}
			}
			// relay it to the main thread
			c <- string(data)
		}
	}()

	return c, nil
}

func saySomethingToTheOwningInstance() {
	conn, err := dial()
	if err != nil {
		return
	}
	defer conn.Close()

	// Send whatever you want. We'll send the args our app was opened with.
	args, err := json.Marshal(os.Args)
	if err != nil {
		return
	}

	conn.Write(args)
}
