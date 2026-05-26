package pongo2

// tagForNode represents the {% for %} tag.
//
// The for tag loops over each item in a sequence (slice, array, map, or string).
// It provides loop variables through the special "forloop" object.
//
// Basic usage:
//
//	{% for item in items %}
//	    {{ item }}
//	{% endfor %}
//
// Iterating over maps with key and value:
//
//	{% for key, value in myMap %}
//	    {{ key }}: {{ value }}
//	{% endfor %}
//
// Using the empty clause (displayed when the sequence is empty):
//
//	{% for item in items %}
//	    {{ item }}
//	{% empty %}
//	    No items found.
//	{% endfor %}
//
// Using "reversed" to iterate in reverse order:
//
//	{% for item in items reversed %}
//	    {{ item }}
//	{% endfor %}
//
// Using "sorted" to iterate over maps in sorted key order:
//
//	{% for key, value in myMap sorted %}
//	    {{ key }}: {{ value }}
//	{% endfor %}
//
// Loop variables available via forloop:
//   - forloop.Counter: Current iteration (1-indexed)
//   - forloop.Counter0: Current iteration (0-indexed)
//   - forloop.Revcounter: Iterations remaining (1-indexed)
//   - forloop.Revcounter0: Iterations remaining (0-indexed)
//   - forloop.First: True if this is the first iteration
//   - forloop.Last: True if this is the last iteration
//   - forloop.Parentloop: Access parent loop in nested loops
//
// Example with loop variables:
//
//	{% for item in items %}
//	    {% if forloop.First %}<ul>{% endif %}
//	    <li>{{ forloop.Counter }}. {{ item }}</li>
//	    {% if forloop.Last %}</ul>{% endif %}
//	{% endfor %}
type tagForNode struct {
	key             string
	value           string // only for maps: for key, value in map
	objectEvaluator IEvaluator
	reversed        bool
	sorted          bool

	bodyWrapper  *NodeWrapper
	emptyWrapper *NodeWrapper
}

// tagForLoopInformation provides loop metadata accessible via "forloop" variable.
type tagForLoopInformation struct {
	Counter     int
	Counter0    int
	Revcounter  int
	Revcounter0 int
	First       bool
	Last        bool
	Parentloop  *tagForLoopInformation
}

// Execute iterates over the object and renders the body for each item.
// If the object is empty, it renders the empty wrapper (if present).
func (node *tagForNode) Execute(ctx *ExecutionContext, writer TemplateWriter) (forError error) {
	_ = "STUB: not implemented"
	// Backup forloop (as parentloop in public context), key-name and value-name
	return nil
}

// Create loop struct

// Is it a loop in a loop?

// Register loopInfo in public context

// There's something to iterate over (correct type and at least 1 item)

// Update loop infos and public context

// Render elements with updated context

// Nothing to iterate over (maybe wrong type or no items)

// tagForParser parses the {% for %} tag. It supports key/value iteration,
// "in" keyword, and optional "reversed" and "sorted" modifiers.
func tagForParser(doc *Parser, start *Token, arguments *Parser) (INodeTag, error) {
	_ = "STUB: not implemented"
	return *

	// Arguments parsing
	new(INodeTag), nil
}

// Value name is provided

// Body wrapping

// if there's an else in the if-statement, we need the else-Block as well

func init() {
	mustRegisterTag("for", tagForParser)
}
