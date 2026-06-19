package cli

import (
	"fmt"
	"strings"

	"github.com/chzyer/readline"

	"github.com/shubhamranswal/ciphergate/internal/auth"
	"github.com/shubhamranswal/ciphergate/internal/session"
	"github.com/shubhamranswal/ciphergate/internal/user"
)

func getPrompt(
	authCtx *auth.Context,
) string {

	if authCtx.IsAuthenticated() {

		return fmt.Sprintf(
			"ciphergate(%s)> ",
			authCtx.User.Username,
		)
	}

	return "ciphergate> "
}

func Run(
	userService *user.Service,
	sessionService *session.Service,
	authCtx *auth.Context,
) {

	completer := readline.NewPrefixCompleter(
		readline.PcItem("register"),
		readline.PcItem("login"),
		readline.PcItem("whoami"),
		readline.PcItem("logout"),
		readline.PcItem("enable-2fa"),
		readline.PcItem("disable-2fa"),
		readline.PcItem("help"),
		readline.PcItem("exit"),
	)

	rl, err := readline.NewEx(
		&readline.Config{
			Prompt:          getPrompt(authCtx),
			HistoryFile:     ".ciphergate_history",
			AutoComplete:    completer,
			InterruptPrompt: "^C",
			EOFPrompt:       "exit",
		},
	)

	if err != nil {
		panic(err)
	}

	defer rl.Close()
	printGuestHelp()

	for {

		fmt.Println()

		rl.SetPrompt(
			getPrompt(authCtx),
		)

		input, err := rl.Readline()

		if err == readline.ErrInterrupt {

			if len(input) == 0 {

				fmt.Println(
					"\n👋 Goodbye!",
				)

				return
			}

			continue
		}

		if err != nil {

			fmt.Printf(
				"❌ %v\n",
				err,
			)

			continue
		}

		command := strings.TrimSpace(
			strings.ToLower(input),
		)

		if authCtx.IsAuthenticated() {

			switch command {

			case "whoami":
				WhoAmI(
					authCtx,
				)

			case "logout":
				Logout(
					authCtx,
					sessionService,
				)

			case "enable-2fa":
				fmt.Println(
					"🚧 Coming soon",
				)

			case "disable-2fa":
				fmt.Println(
					"🚧 Coming soon",
				)

			case "help":
				printUserHelp()

			case "":
				continue

			default:
				fmt.Printf(
					"❌ Unknown command: %s\n",
					command,
				)

				fmt.Println(
					"💡 Type 'help' to view available commands.",
				)
			}

		} else {

			switch command {

			case "register":
				Register(
					userService,
				)

			case "login":
				Login(
					userService,
					authCtx,
				)

			case "help":
				printGuestHelp()

			case "exit":
				fmt.Println(
					"👋 Goodbye!",
				)
				return

			case "":
				continue

			default:
				fmt.Printf(
					"❌ Unknown command: %s\n",
					command,
				)

				fmt.Println(
					"💡 Type 'help' to view available commands.",
				)
			}
		}
	}
}

func printGuestHelp() {

	fmt.Println("\n📚 CipherGate Commands")
	fmt.Println("────────────────────────────────────")

	fmt.Println("> register      Create a new user account")
	fmt.Println("> login         Authenticate with username and password")
	fmt.Println("> help          Show available commands")
	fmt.Println("> exit          Exit CipherGate")

	fmt.Println("\n💡 Tip: Register an account before logging in.")
}

func printUserHelp() {

	fmt.Println("\n📚 CipherGate Commands")
	fmt.Println("────────────────────────────────────")

	fmt.Println("> whoami        Show current user details")
	fmt.Println("> enable-2fa    Enable TOTP multi-factor authentication")
	fmt.Println("> disable-2fa   Disable TOTP multi-factor authentication")
	fmt.Println("> logout        End current session")
	fmt.Println("> help          Show available commands")

	fmt.Println("\n🔒 You are authenticated.")
}
