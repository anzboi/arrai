// AUTOGENERATED. DO NOT EDIT.
package syntax

import (
	"strings"

	"github.com/arr-ai/wbnf/wbnf"
)

func unfakeBackquote(s string) string {
	return strings.ReplaceAll(s, "‵", "`")
}

var arraiParsers = wbnf.MustCompile(unfakeBackquote(`
expr   -> C* amp="&"* @ C* arrow=(
              nest |
              unnest |
              ARROW @ |
              binding="->" C* "\\" C* IDENT C* %%bind C* @ |
              binding="->" C* %%bind @
          )* C*
        > C* @:binop=("with" | "without") C*
        > C* @:binop="||" C*
        > C* @:binop="&&" C*
        > C* @:compare=/{!?(?:<:|<>?=?|>=?|=)} C*
        > C* @ if=("if" t=expr ("else" f=expr)?)* C*
        > C* @:binop=/{\+\+|[+|]|-%?} C*
        > C* @:binop=/{&~|&|~~?|[-<][-&][->]} C*
        > C* @:binop=/{//|[*/%]|\\} C*
        > C* @:rbinop="^" C*
        > C* unop=/{:>|=>|>>|[-+!*^]}* @ C*
        > C* @:binop=">>>" C*
        > C* @ count="count"? C* touch? C*
        > C* (get | @) tail=(
              get
            | call=("("
                  arg=(
                      expr (":" end=expr? (":" step=expr)?)?
                      |     ":" end=expr  (":" step=expr)?
                  ):",",
              ")")
          )* C*
        > C* "{" C* rel=(names tuple=("(" v=@:",", ")"):",",?) "}" C*
        | C* "{" C* set=(elt=@:",",?) "}" C*
        | C* "{" C* dict=((key=@ ":" value=@):",",?) "}" C*
        | C* cond=("cond" "(" (key=@ ":" value=@):",",? ("*" ":" f=expr ","?)? ")") C*
        | C* cond=(("(" control_var=expr ")" | IDENT)? C* "cond" "(" (key=(("(" @:"," ")") | @) ":" value=@):",",? ("*" ":" f=expr ","?)? ")") C*
        | C* "[" C* array=(item=@:",",?) "]" C*
        | C* "{:" C* embed=(grammar=@ ":" subgrammar=%%ast) ":}" C*
        | C* op="\\\\" @ C*
        | C* fn="\\" IDENT @ C*
        | C* "//" pkg=( "{" dot="."? PKGPATH "}" | std=IDENT?)
        | C* "(" tuple=(pairs=(name? ":" v=@):",",?) ")" C*
        | C* "(" @ ")" C*
        | C* let=("let" C* IDENT C* "=" C* @ %%bind C* ";" C* @) C*
        | C* xstr C*
        | C* IDENT C*
        | C* STR C*
        | C* NUM C*;
nest   -> C* "nest" names IDENT C*;
unnest -> C* "unnest" IDENT C*;
touch  -> C* ("->*" ("&"? IDENT | STR))+ "(" expr:"," ","? ")" C*;
get    -> C* dot="." ("&"? IDENT | STR | "*") C*;
names  -> C* "|" C* IDENT:"," C* "|" C*;
name   -> C* IDENT C* | C* STR C*;
xstr   -> C* quote=/{\$"\s*} part=( sexpr | fragment=/{(?: \\. | \$[^{"] | [^\\"$] )+} )* '"' C*
        | C* quote=/{\$'\s*} part=( sexpr | fragment=/{(?: \\. | \$[^{'] | [^\\'$] )+} )* "'" C*
        | C* quote=/{\$‵\s*} part=( sexpr | fragment=/{(?: ‵‵  | \$[^{‵] | [^‵  $] )+} )* "‵" C*;
sexpr  -> "${"
          C* expr C*
          control=/{ (?: : [-+#*\.\_0-9a-z]* (?: : (?: \\. | [^\\:}] )* ){0,2} )? }
          close=/{\}\s*};

ARROW  -> /{:>|=>|>>|orderby|order|where|sum|max|mean|median|min};
IDENT  -> /{ \. | [$@A-Za-z_][0-9$@A-Za-z_]* };
PKGPATH -> /{ (?: \\ | [^\\}] )* };
STR    -> /{ " (?: \\. | [^\\"] )* "
           | ' (?: \\. | [^\\'] )* '
           | ‵ (?: ‵‵  | [^‵  ] )* ‵
           };
NUM    -> /{ (?: \d+(?:\.\d*)? | \.\d+ ) (?: [Ee][-+]?\d+ )? };
C      -> /{ # .* $ };

.wrapRE -> /{\s*()\s*};

`), nil)
