package parse

import "github.com/scottshotgg/express/token"

// FIXME: move this to its own file
func (p *Parser) LessThanOperands(left, right token.Value) (token.Value, error) {
	// FIXME: this only works for ints right now
	// Need to put a type on this

	//fmt.Printf("LessThanOperands %+v %+v\n", left, right)

	return token.Value{
		True: left.True.(int) < right.True.(int),
	}, nil
}
