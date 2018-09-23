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
		Type: "SELECT",
		Value: Value{
			Type:   "keyword", // TODO: what to put here?
			String: "select",
		},
	},
	"for": Token{
		ID:   9,
		Type: "FOR",
		Value: Value{
			Type:   "keyword", // TODO: what to put here?
			String: "for",
		},
	},
	"if": Token{
		ID:   9,
		Type: "IF",
		Value: Value{
			Type:   "keyword", // TODO: what to put here?
			String: "if",
		},
	},
	"in": Token{
		ID:   9,
		Type: "IN",
		Value: Value{
			Type:   "keyword", // TODO: what to put here?
			String: "in",
		},
	},
	"of": Token{
		ID:   9,
		Type: "OF",
		Value: Value{
			Type:   "keyword", // TODO: what to put here?
			String: "of",
		},
	},
	"function": Token{
		ID:   9,
		Type: "FUNCTION",
		Value: Value{
			Type:   "keyword", // TODO: what to put here?
			String: "function",
		},
	},
	"func": Token{
		ID:   9,
		Type: "FUNC",
		Value: Value{
			Type:   "keyword", // TODO: what to put here?
			String: "func",
		},
	},
	"fn": Token{
		ID:   9,
		Type: "FN",
		Value: Value{
			Type:   "keyword", // TODO: what to put here?
			String: "fn",
		},
	},
	"return": Token{
		ID:   9,
		Type: "RETURN",
		Value: Value{
			Type:   "keyword", // TODO: what to put here?
			String: "return",
		},
	},
	"onexit": Token{
		ID:   9,
		Type: "ONEXIT",
		Value: Value{
			Type: "keyword", // TODO: what to put here?
			// String: OnExit,
			String: "onexit",
		},
	},
	"onreturn": Token{
		ID:   9,
		Type: "ONRETURN",
		Value: Value{
			Type: "keyword", // TODO: what to put here?
			// String: OnReturn,
			String: "onreturn",
		},
	},
	"onleave": Token{
		ID:   9,
		Type: "ONLEAVE",
		Value: Value{
			Type: "keyword", // TODO: what to put here?
			// String: OnLeave,
			String: "onleave",
		},
	},
	"defer": Token{
		ID:   9,
		Type: "DEFER",
		Value: Value{
			Type: "keyword", // TODO: what to put here?
			// String: Defer,
			String: "defer",
		},
	},
}
