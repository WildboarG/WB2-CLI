package command

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/urfave/cli/v2"
	"github.com/wb2-cli/config"
	"github.com/wb2-cli/utils"
)

func Init() *cli.Command {
	return &cli.Command{
		Name:   "init",
		Usage:  "Initialize a  WB2 SDK PATH",
		Action: doInit,
	}
}

func doInit(ctx *cli.Context) error {
	fmt.Println("WB2 CLI")
	fmt.Println("Initializing WB2 SDK PATH")
	sdk := inputDefaultSDKPath()
	conf := &config.Config{
		SDK: sdk,
	}

	pathToWrite, err := config.ExistingConfig(config.GetLocations())

	var writeErr error

	if err == config.ErrNoneSet {
		pathToWrite = inputConfigLocation()
		writeErr = config.WriteConfig(pathToWrite, conf)
	} else {
		writeErr = config.WriteConfig(pathToWrite, conf)
	}
	if writeErr == nil {
		fmt.Println("Written config to:", pathToWrite)
	} else {
		fmt.Println("Something Error: ", writeErr)
	}
	return nil
}

func inputConfigLocation() string {
	locations := config.GetLocations()

	if len(locations) == 1 {
		return locations[0]
	}

	for {
		fmt.Println("Where to put the config file?")
		for i, location := range locations {
			fmt.Println(fmt.Sprintf("%d. %s", i+1, location))
		}
		value := inputString("Enter a number: ")
		pn()

		choice, err := strconv.Atoi(value)
		if err != nil {
			continue
		}
		if choice > 0 && choice <= len(locations) {
			indexedChoice := choice - 1
			return locations[indexedChoice]
		}
	}
}

func inputString(text string) string {
	fmt.Print(text)
	reader := bufio.NewReader(os.Stdin)
	readString, err := reader.ReadString('\n')
	if err != nil {
		utils.Exit1With(err)
	}
	return strings.TrimSpace(readString)
}

func inputDefaultSDKPath() string {
	defaultSDKPathStr := inputString("Enter the path to your WB2 SDK (explame: /home/user/Ai-Thinker-WB2/):")
	return defaultSDKPathStr
}

func pn() {
	fmt.Println()
}
