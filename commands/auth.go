package commands

import (
	"fmt"
	"log"

	"github.com/github/hub/github"
)

var (
	cmdAuth = &Command{
		Run:          printHelp,
		GitExtension: false,
		Usage:        `auth login [DOMAIN]`,
		Long: `Authenticate to GitHub.

## Commands:

	* _login_:
		Authenticate with github.com or GitHub enterprise.
`,
	}

	cmdLogin = &Command{
		Key: "login",
		Run: login,
	}
)

func init() {
	cmdAuth.Use(cmdLogin)
	CmdRunner.Use(cmdAuth)
}

func login(command *Command, args *Args) {
	// TODO
	fmt.Printf("Command: %v\n Args: %v\n", command, args)

	// TODO: Host is an arg
	config := github.CurrentConfig()
	host, err := config.PromptForHost("github.com")
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	fmt.Printf("Host: %v\n", host)
}
