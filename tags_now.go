package pongo2

// tagNowNode represents the {% now %} tag.
//
// Django difference: Django uses PHP-style format characters (e.g., "Y-m-d");
// pongo2 uses Go's reference time: Mon Jan 2 15:04:05 MST 2006.
//
// The now tag outputs the current date and/or time using Go's time formatting.
// The format string uses Go's reference time: Mon Jan 2 15:04:05 MST 2006.
//
// Basic usage:
//
//	{% now "2006-01-02" %}
//
// Output: "2024-03-15" (current date)
//
// Various format examples:
//
//	{% now "January 2, 2006" %}          {# Output: "March 15, 2024" #}
//	{% now "Mon, 02 Jan 2006 15:04:05" %} {# Output: "Fri, 15 Mar 2024 10:30:45" #}
//	{% now "3:04 PM" %}                   {# Output: "10:30 AM" #}
//	{% now "15:04:05" %}                  {# Output: "10:30:45" #}
//	{% now "Monday" %}                    {# Output: "Friday" #}
//
// Common format patterns:
//   - "2006-01-02": ISO date (YYYY-MM-DD)
//   - "01/02/2006": US date format (MM/DD/YYYY)
//   - "02/01/2006": European date format (DD/MM/YYYY)
//   - "15:04:05": 24-hour time
//   - "3:04 PM": 12-hour time with AM/PM
//   - "Mon Jan 2 15:04:05 2006": Full date and time
//
// Using "fake" for testing (outputs fixed date: Feb 5, 2014 18:31:45 UTC):
//
//	{% now "2006-01-02" fake %}
//
// Output: "2014-02-05"
type tagNowNode struct {
	position *Token
	format   string
	fake     bool
}

// Execute formats and outputs the current time (or a fixed test time if "fake").
func (node *tagNowNode) Execute(ctx *ExecutionContext, writer TemplateWriter) error {
	_ = "STUB: not implemented"
	return nil
}

// tagNowParser parses the {% now %} tag. It requires a format string argument
// and optionally accepts "fake" for deterministic testing output.
func tagNowParser(doc *Parser, start *Token, arguments *Parser) (INodeTag, error) {
	_ = "STUB: not implemented"
	return *new(INodeTag), nil
}

func init() {
	mustRegisterTag("now", tagNowParser)
}
