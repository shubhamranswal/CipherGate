package cli

import (
	"fmt"

	"github.com/shubhamranswal/ciphergate/internal/auth"
)

func WhoAmI(authCtx *auth.Context) {
	if !authCtx.IsAuthenticated() {
		fmt.Println("❌ Not logged in")
		return
	}

	mfaStatus := "Disabled"
	if authCtx.User.MFAEnabled {
		mfaStatus = "Enabled"
	}

	fmt.Println("\n👤 User Details")
	fmt.Println("────────────────────────────────────────")
	fmt.Printf("> Username           : %s\n", authCtx.User.Username)
	fmt.Printf("> Registration Date  : %s UTC\n", authCtx.User.CreatedAt.Format("2006-01-02 15:04:05"))
	fmt.Printf("> MFA Status         : %s\n", mfaStatus)

	if authCtx.User.LastLogin != nil {
		fmt.Printf(
			"> Last Login         : %s UTC\n",
			authCtx.User.LastLogin.Format(
				"2006-01-02 15:04:05",
			),
		)
	}

	fmt.Printf("> Session Status     : Active\n")
	fmt.Printf("> Session Expires    : %s UTC\n", authCtx.Session.ExpiresAt.Format("2006-01-02 15:04:05"))
}
