package cli

import (
	"context"
	"fmt"

	"github.com/shubhamranswal/ciphergate/internal/auth"
	"github.com/shubhamranswal/ciphergate/internal/user"
)

func Login(
	userService *user.Service,
	authCtx *auth.Context,
) {

	if authCtx.IsAuthenticated() {

		fmt.Println(
			"❌ Already logged in",
		)

		return
	}

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

	authCtx.Login(
		user,
		sessionObj,
	)

	fmt.Printf(
		"\n✅ Welcome %s\n",
		user.Username,
	)

	fmt.Printf(
		"Session expires: %s\n",
		sessionObj.ExpiresAt.Format(
			"2006-01-02 15:04:05",
		),
	)
}
