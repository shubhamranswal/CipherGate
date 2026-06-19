package cli

import (
	"fmt"

	"github.com/shubhamranswal/ciphergate/internal/auth"
)

func WhoAmI(
	authCtx *auth.Context,
) {

	if !authCtx.IsAuthenticated() {

		fmt.Println(
			"❌ Not logged in",
		)

		return
	}

	fmt.Println("\n👤 User Details")

	fmt.Printf(
		"Username: %s\n",
		authCtx.User.Username,
	)

	fmt.Printf(
		"MFA Enabled: %t\n",
		authCtx.User.MFAEnabled,
	)

	fmt.Printf(
		"Registered: %s\n",
		authCtx.User.CreatedAt.Format(
			"2006-01-02 15:04:05",
		),
	)

	if authCtx.User.LastLogin != nil {

		fmt.Printf(
			"Last Login: %s\n",
			authCtx.User.LastLogin.Format(
				"2006-01-02 15:04:05",
			),
		)
	}

	fmt.Printf(
		"Session Expires: %s\n",
		authCtx.Session.ExpiresAt.Format(
			"2006-01-02 15:04:05",
		),
	)
}
