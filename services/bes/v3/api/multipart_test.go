package api

import "testing"

// TestBuildMultipartBodyRejectsNilFile covers the `file == nil` guard. No production caller can
// trigger it — openUploadFile always returns a non-nil *os.File on success — so the guard only
// exists for direct callers and has to be exercised directly.
func TestBuildMultipartBodyRejectsNilFile(t *testing.T) {
	if _, _, err := buildMultipartBody(nil, "file", nil, "plugin.zip", 0); err == nil {
		t.Fatal("expected an error for a nil upload file")
	}
}

// TestGuessContentType covers both fallbacks to the generic binary type: a name with no extension
// at all, and an extension missing from the MIME map. Only the known-extension path was warm
// before, via the plugin upload tests.
func TestGuessContentType(t *testing.T) {
	cases := map[string]string{
		"archive":             "application/octet-stream",
		"plugin.unknown-ext":  "application/octet-stream",
		"plugin.zip":          "application/zip",
		"mapping.json":        "text/json",
		"notes.with.dots.txt": "text/plain",
		".zip":                "application/zip",
	}
	for fileName, want := range cases {
		if got := guessContentType(fileName); got != want {
			t.Fatalf("guessContentType(%q) = %q, want %q", fileName, got, want)
		}
	}
}

// TestSanitizeFilename pins the escaping that keeps a crafted filename from breaking the
// Content-Disposition header or injecting extra multipart fields.
func TestSanitizeFilename(t *testing.T) {
	cases := map[string]string{
		"plain.zip":                  "plain.zip",
		`quote".zip`:                 `quote\".zip`,
		`back\slash.zip`:             `back\\slash.zip`,
		"carriage\rreturn.zip":       "carriagereturn.zip",
		"line\nfeed.zip":             "linefeed.zip",
		"both\r\nbreaks.zip":         "bothbreaks.zip",
		`mix"\` + "\r\n" + `end.zip`: `mix\"\\end.zip`,
	}
	for filename, want := range cases {
		if got := sanitizeFilename(filename); got != want {
			t.Fatalf("sanitizeFilename(%q) = %q, want %q", filename, got, want)
		}
	}
}
