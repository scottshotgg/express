# grammar test and shit:

program       :=  <stmt>+

stmt          :=  <decl_stmt> |
                  <assign_stmt>

decl_stmt     :=  <var>                 |
                  <var> <assign_op> <expression> |
                  <ident> <init_op> <expression>

assign_stmt   :=  <ident> <assign_op> <expression>

expression    :=  <term> <sec_op> <term> |
                  <term>

term          :=  <factor> <pri_op> <term> |
                  <factor>

factor        :=  <l_paren> <expression> <r_paren> |
                  <literal> |
                  <ident>

var           :=  <type> <ident>

type          :=  `int` | `bool` | `char` | `float` | `string` | `var`

ident         :=  [A-z]+[A-z0-9]*

assign_op     :=  `=`

init_op       :=  `:=`

literal       :=  <var_lit>

var_lit       :=  <int_lit> | <bool_lit> | <char_lit> | <float_lit> | <string_lit>

sec_op        :=  `+` | `-`

pri_op        :=  `*` | `/`

l_paren       :=  `(`

r_paren       :=  `)`

int_lit       :=  [A-f0-9]+

bool_lit      :=  `false` | `true`

char_lit      :=  `'` [A-z0-9]+ `'`

float_lit     :=  [0-9]+ `.` [0-9]+

string_lit    := `"` [A-z0-9]+ `"`
