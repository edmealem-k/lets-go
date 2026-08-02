package custom

// imports and exports
// Basically, any value (like a variable or function) can be exported and visible
// from other packages if they have been defined with an upper case identifier.
var (
	value int = 10 // will not be exported
	Value int = 20 // will be exported
)
