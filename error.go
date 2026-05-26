package pongo2

// The Error type is being used to address an error during lexing, parsing or
// execution. If you want to return an error object (for example in your own
// tag or filter) fill this object with as much information as you have.
// Make sure "Sender" is always given (if you're returning an error within
// a filter, make Sender equals 'filter:yourfilter'; same goes for tags: 'tag:mytag').
// It's okay if you only fill in ErrorMsg if you don't have any other details at hand.
type Error struct {
	Template  *Template
	Filename  string
	Line      int
	Column    int
	Token     *Token
	Sender    string
	OrigError error
}

// updateFromTokenIfNeeded updates the error with template and token information
// if they haven't been set yet. This helps provide better error location context.
func (e *Error) updateFromTokenIfNeeded(template *Template, t *Token) *Error {
	_ = "STUB: not implemented"
	return nil
}

// updateErrorToken is a helper that updates token info on a *Error if the error
// is of that type, otherwise returns the error as-is.
func updateErrorToken(err error, template *Template, t *Token) error {
	_ = "STUB: not implemented"
	return nil
}

// Unwrap returns the underlying error for use with errors.Is and errors.As.
func (e *Error) Unwrap() error {
	_ = "STUB: not implemented"

	// Returns a nice formatted error string.
	return nil
}

func (e *Error) Error() string { _ = "STUB: not implemented"; return "" }

// RawLine returns the affected line from the original template, if available.
func (e *Error) RawLine() (line string, available bool, outErr error) {
	_ = "STUB: not implemented"
	return "", false, nil
}

// Try to get the file through the template's loader first (supports fs.FS),
// falling back to os.Open for backwards compatibility

// If reader implements io.Closer, ensure we close it
