const fs = require('fs')

var lexTokens = JSON.parse(fs.readFileSync("test/output/lex/declarations.expr.lex.json"))

function parseType(index) {
  console.log("index", index)

  var nextToken = lexTokens[index+1]
  console.log("nextToken", nextToken)

  // Determine what to do from the next token
  switch () {
    
  }
}

for (var i = 0; i < lexTokens.length; i++) {
  var token = lexTokens[i]
  console.log(token)

  switch (token.Type) {
    case 'TYPE':
      console.log('type')
      parseType(i)

      return
      break

    case 'WS':
      console.log('ws')
      break

    case 'IDENT':
      console.log('ident')
      break

    case 'ASSIGN':
      console.log('assign')
      break

    case 'LITERAL':
      console.log('literal')
      break

    default:
      console.log('default')
  }
  
  console.log()
}
