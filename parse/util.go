package parse

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	"github.com/pkg/errors"
	"github.com/scottshotgg/express/token"
)

// // CollectTokens appends an array of tokens passed in to the EndTokens attribute of Meta
// func (m *Meta) CollectTokens(tokens []token.Token) {
// 	m.LastCollectedToken = tokens[len(tokens)-1]
// 	m.EndTokens = append(m.EndTokens, tokens...)
// }

// // CollectToken appends a single token to the EndTokens attribute of Meta
// func (m *Meta) CollectToken(token token.Token) {
// 	m.LastCollectedToken = token
// 	m.EndTokens = append(m.EndTokens, token)
// }

// // RemoveLastCollectedToken removes the last token put into EndTokens
// func (m *Meta) RemoveLastCollectedToken() {
// 	m.LastCollectedToken = m.EndTokens[len(m.EndTokens)-1]
// 	m.EndTokens = m.EndTokens[:len(m.EndTokens)-1]
// }

// // PopLastCollectedToken removes the last token put into EndTokens
// func (m *Meta) PopLastCollectedToken() token.Token {
// 	m.LastCollectedToken = m.EndTokens[len(m.EndTokens)-2]
// 	m.EndTokens = m.EndTokens[:len(m.EndTokens)-1]

// 	return m.EndTokens[len(m.EndTokens)-1]
// }

// // CollectCurrentToken appends the token held in the CurrentToken attribute to the EndTokens array
// func (m *Meta) CollectCurrentToken() {
// 	m.CollectToken(m.CurrentToken)
// }

// // CollectLastToken appends the token held in the LastToken attribute to the EndTokens array
// func (m *Meta) CollectLastToken() {
// 	m.CollectToken(m.LastToken)
// }

// // GetLastToken returns the LastToken attribute
// func (m *Meta) GetLastToken() token.Token {
// 	return m.LastToken
// }

// // PeekLastCollectedToken returns the last token appended to the EndTokens array
// func (m *Meta) PeekLastCollectedToken() token.Token {
// 	return m.LastCollectedToken
// }

// // GetCurrentToken returns the CurrentToken attribute
// func (m *Meta) GetCurrentToken() token.Token {
// 	return m.CurrentToken
// }

// // PeekTokenAtIndex returns the token at that ParseIndex if valid
// func (m *Meta) PeekTokenAtIndex(index int) (token.Token, error) {
// 	if index > -1 && index < p.Length {
// 		return p.source[index], nil
// 	}

// 	return token.Token{}, errors.New("Current parseIndex outside of token range")
// }

func (p *Parser) SaveState() {
	pState := *p
	meta := p.meta

	pState.meta = meta
	p.States = append(p.States, pState)
}

func (p *Parser) PopState() {
	pNew := &p.States[len(p.States)-1]
	pNew.States = append([]Parser{}, p.States[0:len(p.States)-1]...)
	*p = *pNew
}

// Shift operates the parses like a 3-bit (3 token) SIPO shift register consuming the tokens until the end of the line
func (p *Parser) Shift() {
	p.ProcessedTokens = append(p.ProcessedTokens, p.LastToken)
	p.LastToken = p.CurrentToken
	p.CurrentToken = p.NextToken

	for {
		if p.Index < p.length {
			if p.source[p.Index].Type == token.Whitespace {
				// p.ProcessedTokens = append(p.ProcessedTokens, p.source[p.Index])
				p.Index++
				continue
			}

			p.NextToken = p.source[p.Index]
			p.Index++
			return
		}

		p.NextToken = token.Token{}
		return
	}
}

// Unshift operates the parses like a 3-bit (3 token) SIPO shift register consuming the tokens until the end of the line
func (p *Parser) Unshift() {
	p.Index--
	p.NextToken = p.CurrentToken
	p.CurrentToken = p.LastToken
	p.LastToken = p.ProcessedTokens[len(p.ProcessedTokens)-1]
}

// ShiftWithWS operates the parses like a 3-bit (3 token) SIPO shift register consuming the tokens until the end of the line
func (p *Parser) ShiftWithWS() {
	p.LastToken = p.CurrentToken
	p.CurrentToken = p.source[p.Index]

	for {
		if p.Index+1 < p.length {
			p.Index++

			p.NextToken = p.source[p.Index]
			return
		}

		p.NextToken = token.Token{}
		return
	}
}

// TokenToString marshals a token into it's JSON representation
func TokenToString(t token.Token) string {
	jsonToken, err := json.Marshal(t)
	if err != nil {
		return err.Error()
	}

	return string(jsonToken)
}

func variableTypeFromString(vtString string) (vt VariableType) {
	switch vtString {
	case "var":
		vt = VAR
	case "int":
		vt = INT
	case "bool":
		vt = BOOL
	case "string":
		vt = STRING
	case "float":
		vt = FLOAT
	case "BLOCK":
		vt = OBJECT
	case "set":
		vt = SET
	case "array":
		vt = ARRAY
	case "object":
		vt = OBJECT
	case "function":
		vt = FUNCTION

		// default:
		// 	fmt.Println(vtString)
		// 	os.Exit(9)
	}

	return
}

func accessTypeFromString(atString string) (at AccessType) {
	switch atString {
	case "public":
		at = PUBLIC
	case "private":
		at = PRIVATE
	}

	return
}

func VariableTypeString(vt VariableType) (st string) {
	switch vt {
	case VAR:
		st = "var"
	case INT:
		st = "int"
	case FLOAT:
		st = "float"
	case BOOL:
		st = "bool"
	case STRING:
		st = "string"
	case SET:
		st = "set"
	case OBJECT:
		st = "object"
	case ARRAY:
		st = "array"
	case FUNCTION:
		st = "function"
	// case STRINGA:
	// 	st = "string[]"

	default:
		st = ""
	}

	return
}

func AccessTypeString(at AccessType) (st string) {
	switch at {
	case PUBLIC:
		st = "public"
	case PRIVATE:
		st = "private"

	default:
		st = ""
	}

	return
}

func getBaseForType(trueType, actingType string) (interface{}, error) {
	switch trueType {
	case token.IntType:
		return 0, nil

	case token.StringType:
		return "", nil

	case token.BoolType:
		return false, nil

	case token.FloatType:
		return 0.0, nil

	case token.CharType:
		return "", nil

	case token.VarType:
		fallthrough
	case token.StructType:
		fallthrough
	case token.ObjectType:
		return map[string]interface{}{}, nil

	// case token.FunctionType:

	// case token.ArrayType:
	// 	return

	default:
		var shit interface{}
		return shit, errors.Errorf("Base not defined for type: %s %s", trueType, actingType)
	}
}

// func getAccessTypeFromIdentName(string identName) (string, error) {
// 	if len(identName) > 0 {
// 		PRIVATE
// 	}
// 	return "", error
// }

func mapVariableToTokenValue(v *Variable) token.Value {
	md := map[string]interface{}{}
	for k, value := range v.Metadata {
		md[k] = value
	}

	fmt.Println("md", md)
	fmt.Println("metadata", v.Metadata)

	return token.Value{
		Name:   v.Name,
		Type:   VariableTypeString(v.Type),
		Acting: VariableTypeString(v.ActingType),
		True:   v.Value,
		// String: ,
		AccessType: AccessTypeString(v.AccessType),
		Metadata:   md,
	}
}

// FIXME: recode this so it's not so fucky looking
func RandStringBytesMaskImprSrc(n int) string {
	var src = rand.NewSource(time.Now().UnixNano())

	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}
