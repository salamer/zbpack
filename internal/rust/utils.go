package rust

import (
	"fmt"
	"os"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/salamer/zbpack/internal/utils"
	"github.com/spf13/afero"
)

func ReadFile(path string) (string, error) {

	data, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("error reading %s: %w", path, err)
	}

	return strings.ReplaceAll(string(data), "\r\n", "\n"), nil
}

func ReadTOML(fs afero.Fs, path string, v interface{}) error {
	data, err := utils.ReadFileToUTF8(fs, path)
	if err != nil {
		return err
	}

	return toml.Unmarshal([]byte(data), v)
}

// parseCargoTOML parses a Cargo.toml file
func parseCargoTOML(
	fs afero.Fs, path string,
) (*CargoTOML, error) {
	var cargoToml *CargoTOML

	if err := ReadTOML(fs, path, &cargoToml); err != nil {
		return nil, err
	}

	return cargoToml, nil
}

// See https://doc.rust-lang.org/cargo/reference/manifest.html
type CargoTOML struct {
	Package           PackageInfo     `toml:"package"`
	Dependencies      map[string]any  `toml:"dependencies,omitempty"`
	DevDependencies   map[string]any  `toml:"dev-dependencies,omitempty"`
	BuildDependencies map[string]any  `toml:"build-dependencies,omitempty"`
	Lib               LibConfig       `toml:"lib,omitempty"`
	Bin               []BinConfig     `toml:"bin,omitempty"`
	Workspace         WorkspaceConfig `toml:"workspace,omitempty"`
}

type PackageInfo struct {
	Name          string   `toml:"name"`
	Version       string   `toml:"version"`
	RustVersion   string   `toml:"rust-version,omitempty"`
	Authors       []string `toml:"authors,omitempty"`
	Edition       string   `toml:"edition,omitempty"`
	Description   string   `toml:"description,omitempty"`
	License       string   `toml:"license,omitempty"`
	Repository    string   `toml:"repository,omitempty"`
	Homepage      string   `toml:"homepage,omitempty"`
	Documentation string   `toml:"documentation,omitempty"`
	Keywords      []string `toml:"keywords,omitempty"`
	Categories    []string `toml:"categories,omitempty"`
	Readme        string   `toml:"readme,omitempty"`
	Exclude       []string `toml:"exclude,omitempty"`
	Include       []string `toml:"include,omitempty"`
	Publish       bool     `toml:"publish,omitempty"`
	DefaultRun    string   `toml:"default-run,omitempty"`
}

type Dependency struct {
	Version         string   `toml:"version,omitempty"`
	Path            string   `toml:"path,omitempty"`
	Git             string   `toml:"git,omitempty"`
	Branch          string   `toml:"branch,omitempty"`
	Tag             string   `toml:"tag,omitempty"`
	Rev             string   `toml:"rev,omitempty"`
	Features        []string `toml:"features,omitempty"`
	Optional        bool     `toml:"optional,omitempty"`
	DefaultFeatures bool     `toml:"default-features,omitempty"`
	Package         string   `toml:"package,omitempty"`
}

type LibConfig struct {
	Name      string   `toml:"name,omitempty"`
	Path      string   `toml:"path,omitempty"`
	CrateType []string `toml:"crate-type,omitempty"`
	ProcMacro bool     `toml:"proc-macro,omitempty"`
	Harness   bool     `toml:"harness,omitempty"`
	Test      bool     `toml:"test,omitempty"`
	DocTest   bool     `toml:"doctest,omitempty"`
	Bench     bool     `toml:"bench,omitempty"`
	Doc       bool     `toml:"doc,omitempty"`
	Plugin    bool     `toml:"plugin,omitempty"`
}

type BinConfig struct {
	Name     string `toml:"name,omitempty"`
	Path     string `toml:"path,omitempty"`
	Test     bool   `toml:"test,omitempty"`
	Bench    bool   `toml:"bench,omitempty"`
	Doc      bool   `toml:"doc,omitempty"`
	Plugin   bool   `toml:"plugin,omitempty"`
	Harness  bool   `toml:"harness,omitempty"`
	Required bool   `toml:"required,omitempty"`
}

type Profile struct {
	Opt              string            `toml:"opt-level,omitempty"`
	Debug            bool              `toml:"debug,omitempty"`
	Rpath            bool              `toml:"rpath,omitempty"`
	LtoFlags         []string          `toml:"lto,omitempty"`
	Debug_assertions bool              `toml:"debug-assertions,omitempty"`
	Codegen          map[string]string `toml:"codegen-units,omitempty"`
	Panic            string            `toml:"panic,omitempty"`
	Incremental      bool              `toml:"incremental,omitempty"`
	Overflow_checks  bool              `toml:"overflow-checks,omitempty"`
}

type WorkspaceConfig struct {
	Members        []string `toml:"members,omitempty"`
	ExcludeMembers []string `toml:"exclude,omitempty"`
	DefaultMembers []string `toml:"default-members,omitempty"`
	Resolver       string   `toml:"resolver,omitempty"`
}

type RustToolchain struct {
	// The toolchain specification
	Toolchain ToolchainSpec `toml:"toolchain"`
}

type ToolchainSpec struct {
	Channel    string   `toml:"channel"`
	Version    string   `toml:"version,omitempty"`
	Components []string `toml:"components,omitempty"`
	Targets    []string `toml:"targets,omitempty"`
	Profile    string   `toml:"profile,omitempty"`
}
