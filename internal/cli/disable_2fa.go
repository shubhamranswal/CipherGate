package cli

import (
	"context"
	"fmt"

	"github.com/shubhamranswal/ciphergate/internal/auth"
	"github.com/shubhamranswal/ciphergate/internal/mfa"
	"github.com/shubhamranswal/ciphergate/internal/user"
)

func Disable2FA(
	authCtx *auth.Context,
	userService *user.Service,
	mfaService *mfa.Service,
) {

	if !authCtx.IsAuthenticated() {

		fmt.Println(
			"❌ Not logged in",
		)

		return
	}

	if !authCtx.User.MFAEnabled {

		fmt.Println(
			"❌ MFA is not enabled",
		)

		return
	}

	code, err := readInput(
		"Enter current TOTP code: ",
	)

	if err != nil {

		fmt.Printf(
			"❌ %v\n",
			err,
		)

		return
	}

	if !mfaService.Verify(
		code,
		authCtx.User.MFASecret,
	) {

		fmt.Println(
			"❌ Invalid MFA code",
		)

		return
	}

	authCtx.User.MFAEnabled = false
	authCtx.User.MFASecret = nil

	err = userService.Update(
		context.Background(),
		authCtx.User,
	)

	if err != nil {

		fmt.Printf(
			"❌ %v\n",
			err,
		)

		return
	}

	fmt.Println(
		"✅ MFA disabled successfully",
	)
}
