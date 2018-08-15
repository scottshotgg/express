package parse

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/pkg/errors"
	"github.com/scottshotgg/express/token"
)

var (
	DefinedTypes = map[string]token.Value{}
)

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
	// p.ProcessedTokens = append(p.ProcessedTokens, p.LastToken)
	p.ProcessedTokens.Push(p.LastToken)
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
	// fmt.Println("DECREMENTING INDEX", p.Index)
	// fmt.Println("last at index", p.source[p.Index-1])
	// fmt.Println("current at index", p.source[p.Index])
	// fmt.Println("next at index", p.source[p.Index+1])
	// Decrement atleast one back
	p.Index--
	// fmt.Println("AFTER decrement", p.Index)
	// fmt.Println("last at index", p.source[p.Index-1])
	// fmt.Println("current at index", p.source[p.Index])
	// fmt.Println("next at index", p.source[p.Index+1])

	for p.Index-1 > 0 {
		if p.source[p.Index-1].Type != token.Whitespace {
			// p.ProcessedTokens = append(p.ProcessedTokens, p.source[p.Index])
			break
		}
		p.Index--
	}
	// fmt.Println("INDEX", p.Index)

	p.NextToken = p.CurrentToken
	p.CurrentToken = p.LastToken

	// p.ProcessedTokens should be changed to use a stack
	// p.LastToken = p.ProcessedTokens[len(p.ProcessedTokens)-1]
	poppedInterface, err := p.ProcessedTokens.Pop()
	if err != nil {
		fmt.Println("gots da pop error srry", err)
		os.Exit(9)
	}

	poppedToken := token.Token{}
	if poppedInterface != nil {
		poppedToken = poppedInterface.(token.Token)
	}

	// fmt.Println("POPPED TOKOEN p.LastToken", poppedToken)
	p.LastToken = poppedToken
}

// ShiftWithWS operates the parses like a 3-bit (3 token) SIPO shift register consuming the tokens until the end of the line
func (p *Parser) ShiftWithWS() {
	// p.ProcessedTokens = append(p.ProcessedTokens, p.LastToken)
	p.ProcessedTokens.Push(p.LastToken)
	p.LastToken = p.CurrentToken
	p.CurrentToken = p.NextToken

	for {
		if p.Index+1 < p.length {
			p.NextToken = p.source[p.Index]
			p.Index++
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
	case "struct":
		vt = STRUCT
	case "set":
		vt = SET
	case "array":
		vt = ARRAY
	case "object":
		vt = OBJECT
	case "function":
		vt = FUNCTION

		// TODO: need a map case here to get the custom struct types

		// default:
		// 	//fmt.Println(vtString)
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
	case STRUCT:
		st = "struct"
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

// Shouldn't this return a token.Value?
func getDefaultValueForType(trueType, actingType string) (interface{}, error) {
	switch trueType {
	case token.IntType:
		return 0, nil

	case token.BoolType:
		return false, nil

	case token.CharType:
		return "", nil

	case token.FloatType:
		return 0.0, nil

	case token.StringType:
		return "", nil

	case token.VarType:
		return []token.Value{}, nil
		// fallthrough

	case token.ObjectType:
		return map[string]interface{}{}, nil

	// case token.FunctionType:

	// case token.ArrayType:
	// 	return

	case token.StructType:
		// First we need to check the type map
		value, ok := DefinedTypes[actingType]
		fmt.Println("value, ok", value, ok)
		if !ok {
			fmt.Println("typemap", DefinedTypes)
			return nil, errors.Errorf("Couldn't find struct type in type map: %s", actingType)
		}

		return value, nil

	default:
		// First we need to check the type map
		// value, ok := DefinedTypes[]

		return nil, errors.Errorf("Base not defined for type: %s %s", trueType, actingType)
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

	//fmt.Println("md", md)
	//fmt.Println("metadata", v.Metadata)

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
