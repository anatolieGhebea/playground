// source > https://www.build-redis-from-scratch.dev/en
package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	fmt.Println("Listenting on port :6379")

	// create a new server
	l, err := net.Listen("tcp", ":6379")
	if err != nil {
		fmt.Println("Error creating server:", err)
		os.Exit(1)
	}

	aof, err := NewAof("database.aof")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer aof.Close()

	fmt.Println("Checking data to restore...")
	aof.Read(func(value Value) {
		command := strings.ToUpper(value.array[0].bulk)
		args := value.array[1:]

		if len(args) == 3 {
			fmt.Printf("replay {%s} for hash {%s} with key {%s} and value {%v}\n", command, args[0].bulk, args[1].bulk, args[2].bulk)
		} else if len(args) == 2 {
			fmt.Printf("replay %s with key {%s} and value value {%v}\n", command, args[0].bulk, args[1].bulk)
		}

		handler, ok := Handlers[command]
		if !ok {
			fmt.Println("Invalid command: ", command)
			return
		}

		handler(args)
	})
	fmt.Println("Data restore checks done!")

	fmt.Println("Ready for requests.")
	// Listen for connections
	conn, err := l.Accept()
	if err != nil {
		fmt.Println("Error accepting connection:", err)
		os.Exit(1)
	}
	defer conn.Close()

	for {
		// read the message from the client
		resp := NewResp(conn)
		value, err := resp.Read()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		if value.typ != "array" {
			fmt.Println("Invalid request, expected array")
			continue
		}

		if len(value.array) == 0 {
			fmt.Println("Invalid request, expected array length > 0")
			continue
		}

		command := strings.ToUpper(value.array[0].bulk)
		fmt.Println("Got command %s", string(command))
		args := value.array[1:]

		writer := NewWriter(conn)

		handler, ok := Handlers[command]
		if !ok {
			fmt.Println("Invalid command: ", command)
			writer.Write(Value{typ: "string", str: ""})
			continue
		}

		if command == "SET" || command == "HSET" {
			aof.Write(value)
		}

		result := handler(args)
		writer.Write(result)
	}
}
