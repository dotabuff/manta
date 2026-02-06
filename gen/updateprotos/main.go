package main

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

// gameConfig defines the configuration for a game's proto files.
type gameConfig struct {
	// Name of the game (e.g., "dota2", "deadlock")
	Name string

	// ArchiveDir is the directory within the SteamDatabase/Protobufs archive
	// that contains the proto files for this game.
	ArchiveDir string

	// OutputDir is the local directory where proto files are written.
	OutputDir string

	// GoPackage is the go_package option value for the proto files.
	GoPackage string

	// ExcludeFiles is a list of proto file basenames to exclude.
	ExcludeFiles []string
}

var games = []gameConfig{
	{
		Name:       "dota2",
		ArchiveDir: "dota2/",
		OutputDir:  "dota",
		GoPackage:  "github.com/dotabuff/manta/dota;dota",
		ExcludeFiles: []string{
			"gametoolevents.proto",
			"dota_messages_mlbot.proto",
			"dota_gcmessages_common_bot_script.proto",
			"steammessages_base.proto",
			"steammessages_clientserver_login.proto",
		},
	},
}

const archiveURL = "https://github.com/SteamDatabase/Protobufs/archive/master.tar.gz"

func main() {
	gameName := "dota2"
	if len(os.Args) > 1 {
		gameName = os.Args[1]
	}

	var game *gameConfig
	for i := range games {
		if games[i].Name == gameName {
			game = &games[i]
			break
		}
	}
	if game == nil {
		fmt.Fprintf(os.Stderr, "unknown game: %s\n", gameName)
		os.Exit(1)
	}

	fmt.Printf("Updating protos for %s...\n", game.Name)

	// Step 1: Download and extract
	fmt.Println("Downloading archive...")
	if err := downloadAndExtract(game); err != nil {
		fmt.Fprintf(os.Stderr, "download failed: %v\n", err)
		os.Exit(1)
	}

	// Step 2: Exclude files
	fmt.Println("Removing excluded files...")
	if err := excludeFiles(game); err != nil {
		fmt.Fprintf(os.Stderr, "exclude failed: %v\n", err)
		os.Exit(1)
	}

	// Step 3: Clean dangling imports
	fmt.Println("Cleaning dangling imports...")
	if err := cleanDanglingImports(game); err != nil {
		fmt.Fprintf(os.Stderr, "clean imports failed: %v\n", err)
		os.Exit(1)
	}

	// Step 4: Transform proto files
	fmt.Println("Transforming proto files...")
	if err := transformProtos(game); err != nil {
		fmt.Fprintf(os.Stderr, "transform failed: %v\n", err)
		os.Exit(1)
	}

	// Step 5: Fix Go naming collisions in proto files
	fmt.Println("Fixing Go naming collisions...")
	if err := fixGoNameCollisions(game); err != nil {
		fmt.Fprintf(os.Stderr, "fix name collisions failed: %v\n", err)
		os.Exit(1)
	}

	// Step 6: Compile with protoc
	fmt.Println("Compiling proto files...")
	if err := compileProtos(game); err != nil {
		fmt.Fprintf(os.Stderr, "compile failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Done.")
}

func downloadAndExtract(game *gameConfig) error {
	// Remove and recreate output directory
	if err := os.RemoveAll(game.OutputDir); err != nil {
		return fmt.Errorf("remove output dir: %w", err)
	}
	if err := os.MkdirAll(game.OutputDir, 0755); err != nil {
		return fmt.Errorf("create output dir: %w", err)
	}

	resp, err := http.Get(archiveURL)
	if err != nil {
		return fmt.Errorf("http get: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status: %s", resp.Status)
	}

	gz, err := gzip.NewReader(resp.Body)
	if err != nil {
		return fmt.Errorf("gzip reader: %w", err)
	}
	defer gz.Close()

	tr := tar.NewReader(gz)

	// The archive has a top-level directory like "Protobufs-master/"
	// We want files under "Protobufs-master/<ArchiveDir>/*.proto"
	extracted := 0
	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("tar next: %w", err)
		}

		// Strip the top-level directory (e.g., "Protobufs-master/")
		parts := strings.SplitN(hdr.Name, "/", 2)
		if len(parts) < 2 {
			continue
		}
		relPath := parts[1]

		// Only extract files from the game's directory
		if !strings.HasPrefix(relPath, game.ArchiveDir) {
			continue
		}

		// Get the filename relative to the game directory
		fileName := strings.TrimPrefix(relPath, game.ArchiveDir)

		// Skip directories and subdirectories
		if hdr.Typeflag == tar.TypeDir || strings.Contains(fileName, "/") || fileName == "" {
			continue
		}

		// Only extract .proto files
		if !strings.HasSuffix(fileName, ".proto") {
			continue
		}

		outPath := filepath.Join(game.OutputDir, fileName)
		f, err := os.Create(outPath)
		if err != nil {
			return fmt.Errorf("create %s: %w", outPath, err)
		}
		if _, err := io.Copy(f, tr); err != nil {
			f.Close()
			return fmt.Errorf("write %s: %w", outPath, err)
		}
		f.Close()
		extracted++
	}

	fmt.Printf("  Extracted %d proto files\n", extracted)
	return nil
}

func excludeFiles(game *gameConfig) error {
	for _, name := range game.ExcludeFiles {
		path := filepath.Join(game.OutputDir, name)
		if err := os.Remove(path); err != nil && !os.IsNotExist(err) {
			return fmt.Errorf("remove %s: %w", path, err)
		}
	}

	// Also remove any subdirectories (e.g., tensorflow/)
	entries, err := os.ReadDir(game.OutputDir)
	if err != nil {
		return err
	}
	for _, e := range entries {
		if e.IsDir() {
			path := filepath.Join(game.OutputDir, e.Name())
			if err := os.RemoveAll(path); err != nil {
				return fmt.Errorf("remove dir %s: %w", path, err)
			}
		}
	}

	return nil
}

func cleanDanglingImports(game *gameConfig) error {
	// Build a set of existing proto files
	entries, err := os.ReadDir(game.OutputDir)
	if err != nil {
		return err
	}
	existing := make(map[string]bool)
	for _, e := range entries {
		if strings.HasSuffix(e.Name(), ".proto") {
			existing[e.Name()] = true
		}
	}

	importRe := regexp.MustCompile(`^\s*import\s+"([^"]+)"\s*;`)

	for _, e := range entries {
		if !strings.HasSuffix(e.Name(), ".proto") {
			continue
		}
		path := filepath.Join(game.OutputDir, e.Name())
		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		lines := strings.Split(string(data), "\n")
		var out []string
		changed := false

		for _, line := range lines {
			if m := importRe.FindStringSubmatch(line); m != nil {
				importFile := m[1]
				// Keep google/protobuf imports (they're provided by protoc)
				// Remove imports that reference files we don't have
				if !strings.HasPrefix(importFile, "google/protobuf/") && !existing[importFile] {
					fmt.Printf("  %s: removed dangling import %q\n", e.Name(), importFile)
					changed = true
					continue
				}
			}
			out = append(out, line)
		}

		if changed {
			if err := os.WriteFile(path, []byte(strings.Join(out, "\n")), 0644); err != nil {
				return err
			}
		}
	}

	return nil
}

// Regex patterns for transformation
var (
	// Matches import of valve_extensions.proto
	valveExtImportRe = regexp.MustCompile(`^\s*import\s+"google/protobuf/valve_extensions\.proto"\s*;`)

	// Matches option lines with custom (parenthesized) option names, e.g.:
	//   option (msgpool_soft_limit) = 256;
	//   option (maximum_size_bytes) = 4096;
	customOptionRe = regexp.MustCompile(`^\s*option\s+\(`)

	// Matches field options in square brackets at the end of a field declaration.
	// We strip ALL field options (including default values) to match the
	// original Makefile behavior and avoid issues with custom Valve options
	// that contain quoted strings, commas, etc.
	fieldOptionsRe = regexp.MustCompile(`\s+\[[^\]]*\]\s*;`)

	// Matches leading-dot type references anywhere in a proto file.
	// Since all files share the same package, any .CMsgFoo, .EEnumBar,
	// or .dotaunitorder_t reference can have its leading dot removed.
	// We match a dot followed by a word character that is preceded by
	// whitespace, comma+space, or opening paren — i.e., a type position.
	// We exclude ".proto" to avoid mangling import statements.
	leadingDotTypeRe = regexp.MustCompile(`([\s,(])\.([A-Za-z])`)
)

func transformProtos(game *gameConfig) error {
	entries, err := os.ReadDir(game.OutputDir)
	if err != nil {
		return err
	}

	header := fmt.Sprintf("syntax = \"proto2\";\n\npackage dota;\noption go_package = \"%s\";\n\n", game.GoPackage)

	for _, e := range entries {
		if !strings.HasSuffix(e.Name(), ".proto") {
			continue
		}
		path := filepath.Join(game.OutputDir, e.Name())
		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		lines := strings.Split(string(data), "\n")
		var out []string

		for _, line := range lines {
			// Remove valve_extensions import
			if valveExtImportRe.MatchString(line) {
				continue
			}

			// Remove custom option lines
			if customOptionRe.MatchString(line) {
				continue
			}

			// Strip all field options from field declarations
			line = stripFieldOptions(line)

			// Fix leading-dot type references
			line = fixLeadingDots(line)

			out = append(out, line)
		}

		content := header + strings.Join(out, "\n")
		if err := os.WriteFile(path, []byte(content), 0644); err != nil {
			return err
		}
	}

	return nil
}

// stripFieldOptions removes all field options in square brackets from a line.
// This matches the original Makefile's `sed 's/\s\[.*\]//g'` behavior.
func stripFieldOptions(line string) string {
	loc := fieldOptionsRe.FindStringIndex(line)
	if loc == nil {
		return line
	}
	return line[:loc[0]] + ";"
}

func fixLeadingDots(line string) string {
	return leadingDotTypeRe.ReplaceAllString(line, "${1}${2}")
}

// fixGoNameCollisions fixes proto definitions that would cause Go naming
// collisions in protoc-gen-go. Specifically, when a oneof field name matches
// a sibling message name with underscores, protoc-gen-go generates conflicting
// Go type names.
//
// For example, message AttributeValue with oneof field "single" of type
// AttributeValue_Single generates a Go wrapper type with the same name as
// the message type CDotaMsgStructuredTooltipProperties_AttributeValue_Single.
//
// We fix this by scanning for such patterns and renaming the oneof field.
func fixGoNameCollisions(game *gameConfig) error {
	entries, err := os.ReadDir(game.OutputDir)
	if err != nil {
		return err
	}

	for _, e := range entries {
		if !strings.HasSuffix(e.Name(), ".proto") {
			continue
		}
		path := filepath.Join(game.OutputDir, e.Name())
		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		content := string(data)
		original := content

		// Fix known collision: AttributeValue oneof fields named "single",
		// "variable", "delta" conflict with messages AttributeValue_Single,
		// AttributeValue_Variable, AttributeValue_Delta.
		// We detect this pattern generally: within a oneof block, if a field
		// has name X and type contains _X (case-insensitive first letter match),
		// rename the field to X_value.
		content = fixOneofCollisions(content)

		if content != original {
			fmt.Printf("  Fixed naming collision in %s\n", e.Name())
			if err := os.WriteFile(path, []byte(content), 0644); err != nil {
				return err
			}
		}
	}

	return nil
}

// fixOneofCollisions finds oneof fields where the field name (e.g., "single")
// matches the suffix of its type name (e.g., "AttributeValue_Single"), and
// renames the field to avoid the Go naming collision.
func fixOneofCollisions(content string) string {
	lines := strings.Split(content, "\n")
	// Track nested message names by parsing message declarations
	// Build a set of all message names defined at any level
	messageNames := make(map[string]bool)
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if strings.HasPrefix(trimmed, "message ") {
			parts := strings.Fields(trimmed)
			if len(parts) >= 2 {
				messageNames[parts[1]] = true
			}
		}
	}

	// Now find oneof fields where the type ends with _FieldName and
	// that type name is also a declared message
	inOneof := false
	oneofFieldRe := regexp.MustCompile(`^(\s+)(\S+)\s+(\w+)\s*=\s*(\d+)\s*;`)
	for i, line := range lines {
		trimmed := strings.TrimSpace(line)
		if strings.HasPrefix(trimmed, "oneof ") {
			inOneof = true
			continue
		}
		if inOneof && trimmed == "}" {
			inOneof = false
			continue
		}
		if !inOneof {
			continue
		}

		m := oneofFieldRe.FindStringSubmatch(line)
		if m == nil {
			continue
		}
		indent := m[1]
		typeName := m[2]
		fieldName := m[3]
		fieldNum := m[4]

		// Check if the type name ends with _FieldName (case insensitive first char)
		// e.g., type "AttributeValue_Single" with field "single"
		suffix := "_" + strings.ToUpper(fieldName[:1]) + fieldName[1:]
		if strings.HasSuffix(typeName, suffix) {
			// Check that the type (without package prefix) is a known message
			shortType := typeName
			if idx := strings.LastIndex(typeName, "."); idx >= 0 {
				shortType = typeName[idx+1:]
			}
			if messageNames[shortType] {
				newFieldName := fieldName + "_value"
				lines[i] = fmt.Sprintf("%s%s %s = %s;", indent, typeName, newFieldName, fieldNum)
				fmt.Printf("    Renamed oneof field %q to %q (type %s)\n", fieldName, newFieldName, typeName)
			}
		}
	}

	return strings.Join(lines, "\n")
}

func compileProtos(game *gameConfig) error {
	// Find all proto files
	matches, err := filepath.Glob(filepath.Join(game.OutputDir, "*.proto"))
	if err != nil {
		return err
	}
	if len(matches) == 0 {
		return fmt.Errorf("no proto files found in %s", game.OutputDir)
	}

	// Build protoc-gen-go from the correct module to a temp location.
	// This ensures we use google.golang.org/protobuf/cmd/protoc-gen-go
	// (not the deprecated github.com/golang/protobuf version).
	fmt.Println("  Building protoc-gen-go...")
	tmpDir, err := os.MkdirTemp("", "protoc-gen-go-*")
	if err != nil {
		return fmt.Errorf("create temp dir: %w", err)
	}
	defer os.RemoveAll(tmpDir)

	pluginPath := filepath.Join(tmpDir, "protoc-gen-go")
	buildCmd := exec.Command("go", "install", "google.golang.org/protobuf/cmd/protoc-gen-go@latest")
	buildCmd.Env = append(os.Environ(), "GOBIN="+tmpDir)
	buildCmd.Stdout = os.Stdout
	buildCmd.Stderr = os.Stderr
	if err := buildCmd.Run(); err != nil {
		return fmt.Errorf("build protoc-gen-go: %w", err)
	}

	args := []string{
		"-I", game.OutputDir,
		fmt.Sprintf("--plugin=protoc-gen-go=%s", pluginPath),
		fmt.Sprintf("--go_out=paths=source_relative:%s", game.OutputDir),
	}
	args = append(args, matches...)

	cmd := exec.Command("protoc", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Printf("  Compiling %d proto files...\n", len(matches))
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("protoc failed: %w", err)
	}

	return nil
}
