package cli

import (
	"context"
	"errors"
	"fmt"

	"github.com/shubhamranswal/ciphergate/internal/auth"
	"github.com/shubhamranswal/ciphergate/internal/mfa"
	"github.com/shubhamranswal/ciphergate/internal/session"
	"github.com/shubhamranswal/ciphergate/internal/user"
)

func Login(
	userService *user.Service,
	sessionService *session.Service,
	mfaService *mfa.Service,
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

	loggedInUser, sessionObj, err := userService.Login(
		context.Background(),
		username,
		password,
	)

	if errors.Is(
		err,
		user.ErrMFARequired,
	) {

		fmt.Println(
			"\n🔐 MFA Verification Required",
		)

		code, inputErr := readInput(
			"TOTP Code: ",
		)

		if inputErr != nil {

			fmt.Printf(
				"❌ %v\n",
				inputErr,
			)

			return
		}

		if !mfaService.Verify(
			code,
			loggedInUser.MFASecret,
		) {

			fmt.Println(
				"❌ Invalid MFA code",
			)

			return
		}

		sessionObj, err = sessionService.Create(
			context.Background(),
			loggedInUser.ID,
		)

		if err != nil {

			fmt.Printf(
				"❌ %v\n",
				err,
			)

			return
		}

		err = userService.UpdateLastLogin(
			context.Background(),
			loggedInUser,
		)

		if err != nil {

			fmt.Printf(
				"❌ %v\n",
				err,
			)

			return
		}
	}

	if err != nil &&
		!errors.Is(
			err,
			user.ErrMFARequired,
		) {

		fmt.Printf(
			"❌ %v\n",
			err,
		)

		return
	}

	authCtx.Login(
		loggedInUser,
		sessionObj,
	)

	if sessionObj == nil {

		fmt.Println(
			"❌ Session creation failed",
		)

		return
	}

	fmt.Println(
		"\n✅ Login successful",
	)

	WhoAmI(authCtx)
}
