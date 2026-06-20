package cli

import (
	"context"
	"fmt"

	"github.com/shubhamranswal/ciphergate/internal/auth"
	"github.com/shubhamranswal/ciphergate/internal/session"
)

func ValidateSession(authCtx *auth.Context, sessionService *session.Service) bool {

	if !authCtx.IsAuthenticated() {
		return false
	}

	_, err := sessionService.Validate(context.Background(), authCtx.Session.ID)

	if err != nil {
		_ = sessionService.Deactivate(context.Background(), authCtx.Session.ID)
		authCtx.Logout()
		fmt.Println("❌ Session expired. Please login again.")
		return false
	}

	return true
}
