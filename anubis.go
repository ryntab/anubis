// Package anubis contains the version number of Anubis.
package anubis

import (
	"runtime/debug"
	"time"
)

func init() {
	bi, ok := debug.ReadBuildInfo()
	if !ok {
		return
	}

	// XXX(Xe): many things in this repo assume that the development version
	// of anubis is `devel` and ReadBuildInfo returns `(devel)`. Shim the gap.
	if bi.Main.Version != "(devel)" {
		Version = bi.Main.Version
	}
}

// Version is the current version of Anubis.
//
// This is set from the Go module runtime version.
var Version = "devel"

// BrandName is the name shown in the challenge page footer. Override this in
// custom forks to keep upstream updates while using your own branding.
var BrandName = "Defibtech"

// BrandURL is the project link shown in the challenge page footer.
var BrandURL = "https://www.defibtech.com/"

// BrandOwnerName is the entity shown after the protected-from text in the footer.
var BrandOwnerName = "Defibtech"

// BrandOwnerURL is the website shown after the protected-from text in the footer.
var BrandOwnerURL = "Defibtech"

// CookieName is the name of the cookie that Anubis uses in order to validate
// access.
var CookieName = "techaro.lol-anubis"

// TestCookieName is the name of the cookie that Anubis uses in order to check
// if cookies are enabled on the client's browser.
var TestCookieName = "techaro.lol-anubis-cookie-verification"

// CookieDefaultExpirationTime is the amount of time before the cookie/JWT expires.
const CookieDefaultExpirationTime = 7 * 24 * time.Hour

// BasePrefix is a global prefix for all Anubis endpoints. Can be emptied to remove the prefix entirely.
var BasePrefix = ""

// PublicUrl is the externally accessible URL for this Anubis instance.
var PublicUrl = ""

// StaticPath is the location where all static Anubis assets are located.
const StaticPath = "/.within.website/x/cmd/anubis/"

// APIPrefix is the location where all Anubis API endpoints are located.
const APIPrefix = "/.within.website/x/cmd/anubis/api/"

// DefaultDifficulty is the default "difficulty" (number of leading zeroes)
// that must be met by the client in order to pass the challenge.
const DefaultDifficulty = 4

// ForcedLanguage is the language being used instead of the one of the request's Accept-Language header
// if being set.
var ForcedLanguage = ""

// UseSimplifiedExplanation can be set to true for using the simplified explanation
var UseSimplifiedExplanation = false
