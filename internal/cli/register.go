package cli

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"

	"golang.org/x/term"

	"github.com/shubhamranswal/ciphergate/internal/user"
)

func readPassword(
	prompt string,
) (string, error) {

	fmt.Print(prompt)

	passwordBytes, err := term.ReadPassword(
		int(os.Stdin.Fd()),
	)

	fmt.Println()

	if err != nil {
		return "", err
	}

	return strings.TrimSpace(
		string(passwordBytes),
	), nil
}

func Register(
	userService *user.Service,
) {

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("\n📝 User Registration")
	fmt.Println("Type 'cancel' at any prompt to return.\n")

	var username string

	for {

		fmt.Print("Username: ")

		input, _ := reader.ReadString('\n')

		username = strings.TrimSpace(input)

		if strings.EqualFold(
			username,
			"cancel",
		) {

			fmt.Println(
				"\nRegistration cancelled.",
			)

			return
		}

		if err := userService.ValidateUsername(
			username,
		); err != nil {

			fmt.Printf(
				"❌ %v\n\n",
				err,
			)

			continue
		}

		available, err := userService.UsernameAvailable(
			context.Background(),
			username,
		)

		if err != nil {

			fmt.Printf(
				"❌ %v\n",
				err,
			)

			return
		}

		if !available {

			fmt.Println(
				"❌ Username already exists. Try another.\n",
			)

			continue
		}

		break
	}

	var password string

	for {

		password, err := readPassword(
			"Password: ",
		)

		if err != nil {
			fmt.Printf(
				"❌ %v\n",
				err,
			)
			return
		}

		if strings.EqualFold(
			password,
			"cancel",
		) {

			fmt.Println(
				"\nRegistration cancelled.",
			)

			return
		}

		if err := userService.ValidatePassword(
			password,
		); err != nil {

			fmt.Printf(
				"❌ %v\n\n",
				err,
			)

			continue
		}

		confirm, err := readPassword(
			"Confirm Password: ",
		)

		if err != nil {
			fmt.Printf(
				"❌ %v\n",
				err,
			)
			return
		}

		if strings.EqualFold(
			confirm,
			"cancel",
		) {

			fmt.Println(
				"\nRegistration cancelled.",
			)

			return
		}

		if password != confirm {

			fmt.Println(
				"❌ Passwords do not match. Try again.\n",
			)

			continue
		}

		break
	}

	err := userService.Register(
		context.Background(),
		username,
		password,
	)

	if err != nil {

		fmt.Printf(
			"❌ %v\n",
			err,
		)

		return
	}

	fmt.Println(
		"\n✅ User registered successfully!",
	)

}
