package zeaburpack

import (
	"github.com/salamer/zbpack/internal/golang"
	"github.com/salamer/zbpack/internal/nodejs"
	"github.com/salamer/zbpack/internal/python"
	"github.com/salamer/zbpack/internal/rust"
	"github.com/salamer/zbpack/internal/static"
	"github.com/salamer/zbpack/pkg/plan"
)

// SupportedIdentifiers returns all supported identifiers
// note that they are in the order of priority
func SupportedIdentifiers(config plan.ImmutableProjectConfiguration) []plan.Identifier {
	identifiers := []plan.Identifier{
		// dart.NewIdentifier(),
		// php.NewIdentifier(),
		// ruby.NewIdentifier(),
		// bun.NewIdentifier(),
		python.NewIdentifier(),
		nodejs.NewIdentifier(),
		golang.NewIdentifier(),
		// java.NewIdentifier(),
		// deno.NewIdentifier(),
		rust.NewIdentifier(),
		// dotnet.NewIdentifier(),
		// elixir.NewIdentifier(),
		// gleam.NewIdentifier(),
		// swift.NewIdentifier(),
		static.NewIdentifier(),
	}

	// if !plan.Cast(config.Get("ignore_nix"), plan.ToWeakBoolE).TakeOr(false) {
	// 	identifiers = append([]plan.Identifier{nix.NewIdentifier()}, identifiers...)
	// }

	// if ignore_dockerfile in config is true, or ZBPACK_IGNORE_DOCKERFILE is true, ignore dockerfile
	// if !plan.Cast(config.Get("ignore_dockerfile"), plan.ToWeakBoolE).TakeOr(false) {
	// identifiers = append([]plan.Identifier{dockerfile.NewIdentifier()}, identifiers...)
	// }

	return identifiers
}
