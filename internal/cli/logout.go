package cli

import (
	"context"
	"fmt"

	"github.com/shubhamranswal/ciphergate/internal/auth"
	"github.com/shubhamranswal/ciphergate/internal/session"
)

func Logout(
	authCtx *auth.Context,
	sessionSvc *session.Service,
) {

	if !authCtx.IsAuthenticated() {
		fmt.Println(
			"❌ No active session",
		)
		return
	}

	err := sessionSvc.Deactivate(
		context.Background(),
		authCtx.Session.ID,
	)

	if err != nil {
		fmt.Printf(
			"❌ %v\n",
			err,
		)
		return
	}

	authCtx.Logout()

	fmt.Println(
		"✅ Logged out successfully",
	)
}
