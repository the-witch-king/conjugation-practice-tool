package main

import (
	"bufio"
	"errors"
	"fmt"
	"strconv"
	"io/ioutil"
	"os"
    "encoding/json"
)

var CONFIG_FILE string = "config.json"
var stdinScanner bufio.Scanner
var config ToolConfig

func main() {
    configFile, err := os.Open(CONFIG_FILE)
    stdinScanner = *bufio.NewScanner(os.Stdin)

    if errors.Is(err, os.ErrNotExist) {
        fmt.Println("No config file loaded. Creating new config file.")

        var wordsPerGo int
        var err error

        for {
            fmt.Println("How many words per review would you like to do?")
            stdinScanner.Scan()
            wordsPerGo, err = strconv.Atoi(stdinScanner.Text())

            if err != nil {
                fmt.Println("Unable to determine integer from input.")
                continue
            }

            break
        }

        fmt.Printf("Setting words per review to %d.\n", wordsPerGo)
        config = ToolConfig{wordsPerGo}
    } else if err != nil {
        fmt.Printf("Unable to open config file. Error: %v\n", err)
        os.Exit(1)
    } else {
        configBytes, err := ioutil.ReadAll(configFile)
        if err != nil {
            fmt.Println("Error reading config file.")
        }

        config = ParseConfig(configBytes)
    }

    defer configFile.Close()

    fmt.Printf("Config is: %+v", config)

    configBytes, err := json.Marshal(&config)

    if err != nil {
        fmt.Println("Unable to marshal config.")
        os.Exit(5)
    }

    fmt.Printf("Config bytes: %v", string(configBytes))
    writeErr := ioutil.WriteFile(CONFIG_FILE, configBytes, 0644)

    if writeErr != nil {
        fmt.Printf("Unable to save file. Error: %v\n", writeErr)
    }
}
