package token

// KeywordMap is a map of all the keywords
var KeywordMap = map[string]Token{
	"let": Token{
		Type: "TYPE",
		Value: Value{
			Type:   "var", // this doesn't create a var
			String: "let",
		},
	},
	"select": Token{
		ID:   9,
		Type: "KEYWORD",
		Value: Value{
			Type:   "keyword", // TODO: what to put here?
			String: "select",
		},
	},
	"for": Token{
		ID:   9,
		Type: "KEYWORD",
		Value: Value{
			Type:   "keyword", // TODO: what to put here?
			String: "for",
		},
	},
	"if": Token{
		ID:   9,
		Type: "KEYWORD",
		Value: Value{
			Type:   "keyword", // TODO: what to put here?
			String: "if",
		},
	},
	"in": Token{
		ID:   9,
		Type: "KEYWORD",
		Value: Value{
			Type:   "keyword", // TODO: what to put here?
			String: "in",
		},
	},
	"of": Token{
		ID:   9,
		Type: "KEYWORD",
		Value: Value{
			Type:   "keyword", // TODO: what to put here?
			String: "of",
		},
	},
	"function": Token{
		ID:   9,
		Type: "KEYWORD",
		Value: Value{
			Type:   "keyword", // TODO: what to put here?
			String: "function",
		},
	},
	"func": Token{
		ID:   9,
		Type: "KEYWORD",
		Value: Value{
			Type:   "keyword", // TODO: what to put here?
			String: "func",
		},
	},
	"return": Token{
		ID:   9,
		Type: "KEYWORD",
		Value: Value{
			Type:   "keyword", // TODO: what to put here?
			String: "return",
		},
	},
}
