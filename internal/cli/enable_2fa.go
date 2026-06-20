package cli

import (
	"context"
	"fmt"

	"github.com/shubhamranswal/ciphergate/internal/auth"
	"github.com/shubhamranswal/ciphergate/internal/mfa"
	"github.com/shubhamranswal/ciphergate/internal/user"
)

func Enable2FA(
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

	if authCtx.User.MFAEnabled {

		fmt.Println(
			"❌ MFA already enabled",
		)

		return
	}

	secret, err := mfaService.GenerateSecret(
		authCtx.User.Username,
	)

	if err != nil {

		fmt.Printf(
			"❌ %v\n",
			err,
		)

		return
	}

	fmt.Println(
		"\n🔐 MFA Setup",
	)

	fmt.Println(
		"Add the following secret to Google Authenticator:",
	)

	fmt.Printf(
		"\n%s\n\n",
		secret,
	)

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

	if !mfaService.Validate(
		code,
		secret,
	) {

		fmt.Println(
			"❌ Invalid code",
		)

		return
	}

	authCtx.User.MFAEnabled = true
	authCtx.User.MFASecret = &secret

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
		"\n✅ MFA enabled successfully",
	)
}
