package cli

import (
	"context"
	"fmt"

	"github.com/shubhamranswal/ciphergate/internal/user"
)

func Login(
	userService *user.Service,
) {

	username, err := readInput(
		"Username: ",
	)

	if err != nil {
		fmt.Printf(
			"❌ %v\n",
			err,
		)
		return
	}

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

	user, sessionObj, err := userService.Login(
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

	fmt.Printf(
		"\n✅ Welcome %s\n",
		user.Username,
	)

	fmt.Printf(
		"Session ID: %s\n",
		sessionObj.ID,
	)
}
