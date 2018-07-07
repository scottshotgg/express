package parse_test

import (
	"fmt"
	"testing"

	"github.com/scottshotgg/ExpressRedo/parse"
	"github.com/scottshotgg/ExpressRedo/token"
)

// // // GetKeyword ...
// func (p *Parser) GetKeyword() (token.Value, error) {
// 	fmt.Println("GetKeyword")

// 	switch p.NextToken.Value.String {
// 	// TODO: this needs to be reworked
// 	case token.For:
// 		fmt.Println("formap", p.DeclarationMap)
// 		fmt.Println("found a for loop22")
// 		temp := *m
// 		meta := &temp
// 		meta.LastMeta = m
// 		meta.DeclarationMap = map[string]token.Value{}
// 		meta.Shift()

// 		value, err := meta.GetStatement()
// 		if err != nil {
// 			return token.Value{}, err
// 		}
// 		fmt.Println("value11", value)
// 		fmt.Println("last", meta.LastToken)
// 		fmt.Println("current", meta.CurrentToken)
// 		fmt.Println("next", meta.NextToken)

// 		value2, err := meta.GetExpression()
// 		if err != nil {
// 			return token.Value{}, err
// 		}
// 		fmt.Println("value22", value2)
// 		fmt.Println("last", meta.LastToken)
// 		fmt.Println("current", meta.CurrentToken)
// 		fmt.Println("next", meta.NextToken)
// 		// p.Shift()

// 		value3, err := meta.GetExpression()
// 		if err != nil {
// 			return token.Value{}, err
// 		}
// 		fmt.Println("value3", value3)

// 		stepAmount, err := p.SubOperands(value3, value)
// 		if err != nil {
// 			return token.Value{}, err
// 		}
// 		fmt.Println("step", stepAmount)

// 		// Need to open up the block
// 		// we might try doing something
// 		// where the new meta stuff is in the function
// 		// block, err := meta.CheckBlock()
// 		// if err != nil {
// 		// 	return token.Value{}, err
// 		// }
// 		// fmt.Println("block", block)
// 		// os.Exit(9)
// 		meta.Shift()
// 		fmt.Println(meta.NextToken)
// 		// block, err := Semantic([]token.Token{
// 		// 	meta.NextToken,
// 		// })
// 		block, err := meta.CheckBlock()
// 		if err != nil {
// 			return token.Value{}, err
// 		}
// 		fmt.Println("body", block)

// 		// Swap the scopes back when the for loop is out of execution
// 		meta.DeclarationMap = p.DeclarationMap
// 		meta.InheritedMap = p.InheritedMap
// 		*m = *meta

// 		// fmt.Println("value2againboi", value2)
// 		mapThing := value2.OpMap.(map[string]interface{})
// 		fmt.Println("mapThing", mapThing)

// 		return token.Value{
// 			Type: token.For,
// 			True: map[string]token.Value{
// 				"start": value,
// 				"end":   value2,
// 				"step":  stepAmount,
// 				"body":  block.True.([]token.Value)[0],
// 				"check": token.Value{
// 					String: value2.String,
// 				},
// 			},
// 		}, nil

// 	case "if":
// 		fmt.Println("m", m)
// 		temp := *m
// 		meta := &temp
// 		meta.LastMeta = m
// 		meta.InheritedMap = temp.DeclarationMap
// 		meta.DeclarationMap = map[string]token.Value{}
// 		fmt.Println("m2", m)
// 		meta.Shift()
// 		fmt.Println("declaredMap", p.DeclarationMap)
// 		fmt.Println("inheritedMap", p.InheritedMap)

// 		value2, err := meta.GetExpression()
// 		if err != nil {
// 			return token.Value{}, err
// 		}
// 		fmt.Println("value22", value2)
// 		fmt.Println("last", meta.LastToken)
// 		fmt.Println("current", meta.CurrentToken)
// 		fmt.Printf("next %+v\n", meta.NextToken)

// 		fmt.Println(meta.NextToken.Value.True)
// 		// block, err := Semantic([]token.Token{
// 		// 	meta.NextToken,
// 		// })
// 		meta.DeclarationMap = p.DeclarationMap
// 		block, err := meta.CheckBlock()
// 		if err != nil {
// 			return token.Value{}, err
// 		}
// 		fmt.Println("body2", block)

// 		// Swap the scopes back when the for loop is out of execution
// 		meta.DeclarationMap = p.DeclarationMap
// 		meta.InheritedMap = p.InheritedMap
// 		*m = *meta

// 		return token.Value{
// 			Type:   token.If,
// 			String: value2.String,
// 			True: map[string]token.Value{
// 				"body": block.True.([]token.Value)[0],
// 				"check": token.Value{
// 					String: value2.String,
// 				},
// 			},
// 			// True: // body would go here
// 		}, nil

// 		os.Exit(9)
// 		// TODO: there would be some composition of blocks here and shit

// 	default:
// 		fmt.Println("keyword not recognized", p.NextToken)
// 		os.Exit(9)
// 	}

// 	return token.Value{}, nil
// }

func TestSemantic(t *testing.T) {
	fmt.Println("TestSemantic")

	TestSyntactic(t)

	values, err := parse.New(tokens).Semantic()
	if err != nil {
		fmt.Println("semanticErr", err)
		t.Fail()
		return
	}

	token.PrintValues(values, "\t")
	fmt.Println()
}
