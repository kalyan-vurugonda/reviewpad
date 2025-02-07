// Code generated by goyacc -l -o lang/aladino/parser.go -p Aladino lang/aladino/parser.y. DO NOT EDIT.
// Copyright 2022 Explore.dev Unipessoal Lda. All Rights Reserved.
// Use of this source code is governed by a license that can be
// found in the LICENSE file.

package aladino

import __yyfmt__ "fmt"

var base int

func setAST(l AladinoLexer, root Expr) {
	l.(*AladinoLex).ast = root
}

type AladinoSymType struct {
	yys     int
	str     string
	int     int
	ast     Expr
	astList []Expr
	bool    bool
}

const TIMESTAMP = 57346
const RELATIVETIMESTAMP = 57347
const IDENTIFIER = 57348
const STRINGLITERAL = 57349
const TK_CMPOP = 57350
const NUMBER = 57351
const TRUE = 57352
const FALSE = 57353
const TK_OR = 57354
const TK_AND = 57355
const TK_EQ = 57356
const TK_NEQ = 57357
const TK_NOT = 57358

var AladinoToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"TIMESTAMP",
	"RELATIVETIMESTAMP",
	"IDENTIFIER",
	"STRINGLITERAL",
	"TK_CMPOP",
	"NUMBER",
	"TRUE",
	"FALSE",
	"TK_OR",
	"TK_AND",
	"TK_EQ",
	"TK_NEQ",
	"TK_NOT",
	"'('",
	"')'",
	"'['",
	"']'",
	"'$'",
	"','",
}

var AladinoStatenames = [...]string{}

const AladinoEofCode = 1
const AladinoErrCode = 2
const AladinoInitialStackSize = 16

/*  start  of  programs  */

var AladinoExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const AladinoPrivate = 57344

const AladinoLast = 69

var AladinoAct = [...]int{
	20, 5, 6, 29, 8, 34, 7, 11, 12, 31,
	22, 1, 0, 3, 4, 17, 9, 0, 10, 14,
	13, 15, 16, 21, 2, 0, 0, 18, 19, 30,
	0, 32, 33, 0, 0, 0, 0, 23, 24, 25,
	26, 27, 17, 0, 0, 0, 14, 13, 15, 16,
	17, 0, 28, 17, 14, 13, 15, 16, 13, 15,
	16, 17, 0, 0, 0, 0, 0, 15, 16,
}

var AladinoPact = [...]int{
	-3, -1000, 42, -3, -3, -1000, -1000, -1000, -1000, -3,
	4, -1000, -1000, -3, -3, -3, -3, -3, -1000, 34,
	-17, 7, -8, 53, 45, -1000, -1000, -1000, -1000, -1000,
	-3, -3, -1000, -13, -1000,
}

var AladinoPgo = [...]int{
	0, 23, 0, 11,
}

var AladinoR1 = [...]int{
	0, 3, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 2, 2,
	2,
}

var AladinoR2 = [...]int{
	0, 1, 2, 3, 3, 3, 3, 3, 3, 1,
	1, 1, 1, 3, 2, 1, 1, 5, 3, 1,
	0,
}

var AladinoChk = [...]int{
	-1000, -3, -1, 16, 17, 4, 5, 9, 7, 19,
	21, 10, 11, 13, 12, 14, 15, 8, -1, -1,
	-2, -1, 6, -1, -1, -1, -1, -1, 18, 20,
	22, 17, -2, -2, 18,
}

var AladinoDef = [...]int{
	0, -2, 1, 0, 0, 9, 10, 11, 12, 20,
	0, 15, 16, 0, 0, 0, 0, 0, 2, 0,
	0, 19, 14, 3, 4, 5, 6, 7, 8, 13,
	20, 20, 18, 0, 17,
}

var AladinoTok1 = [...]int{
	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 21, 3, 3, 3,
	17, 18, 3, 3, 22, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 19, 3, 20,
}

var AladinoTok2 = [...]int{
	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16,
}

var AladinoTok3 = [...]int{
	0,
}

var AladinoErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

/*	parser for yacc output	*/

var (
	AladinoDebug        = 0
	AladinoErrorVerbose = false
)

type AladinoLexer interface {
	Lex(lval *AladinoSymType) int
	Error(s string)
}

type AladinoParser interface {
	Parse(AladinoLexer) int
	Lookahead() int
}

type AladinoParserImpl struct {
	lval  AladinoSymType
	stack [AladinoInitialStackSize]AladinoSymType
	char  int
}

func (p *AladinoParserImpl) Lookahead() int {
	return p.char
}

func AladinoNewParser() AladinoParser {
	return &AladinoParserImpl{}
}

const AladinoFlag = -1000

func AladinoTokname(c int) string {
	if c >= 1 && c-1 < len(AladinoToknames) {
		if AladinoToknames[c-1] != "" {
			return AladinoToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func AladinoStatname(s int) string {
	if s >= 0 && s < len(AladinoStatenames) {
		if AladinoStatenames[s] != "" {
			return AladinoStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func AladinoErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !AladinoErrorVerbose {
		return "syntax error"
	}

	for _, e := range AladinoErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + AladinoTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := AladinoPact[state]
	for tok := TOKSTART; tok-1 < len(AladinoToknames); tok++ {
		if n := base + tok; n >= 0 && n < AladinoLast && AladinoChk[AladinoAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if AladinoDef[state] == -2 {
		i := 0
		for AladinoExca[i] != -1 || AladinoExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; AladinoExca[i] >= 0; i += 2 {
			tok := AladinoExca[i]
			if tok < TOKSTART || AladinoExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if AladinoExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += AladinoTokname(tok)
	}
	return res
}

func Aladinolex1(lex AladinoLexer, lval *AladinoSymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = AladinoTok1[0]
		goto out
	}
	if char < len(AladinoTok1) {
		token = AladinoTok1[char]
		goto out
	}
	if char >= AladinoPrivate {
		if char < AladinoPrivate+len(AladinoTok2) {
			token = AladinoTok2[char-AladinoPrivate]
			goto out
		}
	}
	for i := 0; i < len(AladinoTok3); i += 2 {
		token = AladinoTok3[i+0]
		if token == char {
			token = AladinoTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = AladinoTok2[1] /* unknown char */
	}
	if AladinoDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", AladinoTokname(token), uint(char))
	}
	return char, token
}

func AladinoParse(Aladinolex AladinoLexer) int {
	return AladinoNewParser().Parse(Aladinolex)
}

func (Aladinorcvr *AladinoParserImpl) Parse(Aladinolex AladinoLexer) int {
	var Aladinon int
	var AladinoVAL AladinoSymType
	var AladinoDollar []AladinoSymType
	_ = AladinoDollar // silence set and not used
	AladinoS := Aladinorcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	Aladinostate := 0
	Aladinorcvr.char = -1
	Aladinotoken := -1 // Aladinorcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		Aladinostate = -1
		Aladinorcvr.char = -1
		Aladinotoken = -1
	}()
	Aladinop := -1
	goto Aladinostack

ret0:
	return 0

ret1:
	return 1

Aladinostack:
	/* put a state and value onto the stack */
	if AladinoDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", AladinoTokname(Aladinotoken), AladinoStatname(Aladinostate))
	}

	Aladinop++
	if Aladinop >= len(AladinoS) {
		nyys := make([]AladinoSymType, len(AladinoS)*2)
		copy(nyys, AladinoS)
		AladinoS = nyys
	}
	AladinoS[Aladinop] = AladinoVAL
	AladinoS[Aladinop].yys = Aladinostate

Aladinonewstate:
	Aladinon = AladinoPact[Aladinostate]
	if Aladinon <= AladinoFlag {
		goto Aladinodefault /* simple state */
	}
	if Aladinorcvr.char < 0 {
		Aladinorcvr.char, Aladinotoken = Aladinolex1(Aladinolex, &Aladinorcvr.lval)
	}
	Aladinon += Aladinotoken
	if Aladinon < 0 || Aladinon >= AladinoLast {
		goto Aladinodefault
	}
	Aladinon = AladinoAct[Aladinon]
	if AladinoChk[Aladinon] == Aladinotoken { /* valid shift */
		Aladinorcvr.char = -1
		Aladinotoken = -1
		AladinoVAL = Aladinorcvr.lval
		Aladinostate = Aladinon
		if Errflag > 0 {
			Errflag--
		}
		goto Aladinostack
	}

Aladinodefault:
	/* default state action */
	Aladinon = AladinoDef[Aladinostate]
	if Aladinon == -2 {
		if Aladinorcvr.char < 0 {
			Aladinorcvr.char, Aladinotoken = Aladinolex1(Aladinolex, &Aladinorcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if AladinoExca[xi+0] == -1 && AladinoExca[xi+1] == Aladinostate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			Aladinon = AladinoExca[xi+0]
			if Aladinon < 0 || Aladinon == Aladinotoken {
				break
			}
		}
		Aladinon = AladinoExca[xi+1]
		if Aladinon < 0 {
			goto ret0
		}
	}
	if Aladinon == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			Aladinolex.Error(AladinoErrorMessage(Aladinostate, Aladinotoken))
			Nerrs++
			if AladinoDebug >= 1 {
				__yyfmt__.Printf("%s", AladinoStatname(Aladinostate))
				__yyfmt__.Printf(" saw %s\n", AladinoTokname(Aladinotoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for Aladinop >= 0 {
				Aladinon = AladinoPact[AladinoS[Aladinop].yys] + AladinoErrCode
				if Aladinon >= 0 && Aladinon < AladinoLast {
					Aladinostate = AladinoAct[Aladinon] /* simulate a shift of "error" */
					if AladinoChk[Aladinostate] == AladinoErrCode {
						goto Aladinostack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if AladinoDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", AladinoS[Aladinop].yys)
				}
				Aladinop--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if AladinoDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", AladinoTokname(Aladinotoken))
			}
			if Aladinotoken == AladinoEofCode {
				goto ret1
			}
			Aladinorcvr.char = -1
			Aladinotoken = -1
			goto Aladinonewstate /* try again in the same state */
		}
	}

	/* reduction by production Aladinon */
	if AladinoDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", Aladinon, AladinoStatname(Aladinostate))
	}

	Aladinont := Aladinon
	Aladinopt := Aladinop
	_ = Aladinopt // guard against "declared and not used"

	Aladinop -= AladinoR2[Aladinon]
	// Aladinop is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if Aladinop+1 >= len(AladinoS) {
		nyys := make([]AladinoSymType, len(AladinoS)*2)
		copy(nyys, AladinoS)
		AladinoS = nyys
	}
	AladinoVAL = AladinoS[Aladinop+1]

	/* consult goto table to find next state */
	Aladinon = AladinoR1[Aladinon]
	Aladinog := AladinoPgo[Aladinon]
	Aladinoj := Aladinog + AladinoS[Aladinop].yys + 1

	if Aladinoj >= AladinoLast {
		Aladinostate = AladinoAct[Aladinog]
	} else {
		Aladinostate = AladinoAct[Aladinoj]
		if AladinoChk[Aladinostate] != -Aladinon {
			Aladinostate = AladinoAct[Aladinog]
		}
	}
	// dummy call; replaced with literal code
	switch Aladinont {

	case 1:
		AladinoDollar = AladinoS[Aladinopt-1 : Aladinopt+1]
		{
			setAST(Aladinolex, AladinoDollar[1].ast)
		}
	case 2:
		AladinoDollar = AladinoS[Aladinopt-2 : Aladinopt+1]
		{
			AladinoVAL.ast = BuildNotOp(AladinoDollar[2].ast)
		}
	case 3:
		AladinoDollar = AladinoS[Aladinopt-3 : Aladinopt+1]
		{
			AladinoVAL.ast = BuildAndOp(AladinoDollar[1].ast, AladinoDollar[3].ast)
		}
	case 4:
		AladinoDollar = AladinoS[Aladinopt-3 : Aladinopt+1]
		{
			AladinoVAL.ast = BuildOrOp(AladinoDollar[1].ast, AladinoDollar[3].ast)
		}
	case 5:
		AladinoDollar = AladinoS[Aladinopt-3 : Aladinopt+1]
		{
			AladinoVAL.ast = BuildEqOp(AladinoDollar[1].ast, AladinoDollar[3].ast)
		}
	case 6:
		AladinoDollar = AladinoS[Aladinopt-3 : Aladinopt+1]
		{
			AladinoVAL.ast = BuildNeqOp(AladinoDollar[1].ast, AladinoDollar[3].ast)
		}
	case 7:
		AladinoDollar = AladinoS[Aladinopt-3 : Aladinopt+1]
		{
			AladinoVAL.ast = BuildCmpOp(AladinoDollar[1].ast, AladinoDollar[2].str, AladinoDollar[3].ast)
		}
	case 8:
		AladinoDollar = AladinoS[Aladinopt-3 : Aladinopt+1]
		{
			AladinoVAL.ast = AladinoDollar[2].ast
		}
	case 9:
		AladinoDollar = AladinoS[Aladinopt-1 : Aladinopt+1]
		{
			AladinoVAL.ast = BuildTimeConst(AladinoDollar[1].str)
		}
	case 10:
		AladinoDollar = AladinoS[Aladinopt-1 : Aladinopt+1]
		{
			AladinoVAL.ast = BuildRelativeTimeConst(AladinoDollar[1].str)
		}
	case 11:
		AladinoDollar = AladinoS[Aladinopt-1 : Aladinopt+1]
		{
			AladinoVAL.ast = BuildIntConst(AladinoDollar[1].int)
		}
	case 12:
		AladinoDollar = AladinoS[Aladinopt-1 : Aladinopt+1]
		{
			AladinoVAL.ast = BuildStringConst(AladinoDollar[1].str)
		}
	case 13:
		AladinoDollar = AladinoS[Aladinopt-3 : Aladinopt+1]
		{
			AladinoVAL.ast = BuildArray(AladinoDollar[2].astList)
		}
	case 14:
		AladinoDollar = AladinoS[Aladinopt-2 : Aladinopt+1]
		{
			AladinoVAL.ast = BuildVariable(AladinoDollar[2].str)
		}
	case 15:
		AladinoDollar = AladinoS[Aladinopt-1 : Aladinopt+1]
		{
			AladinoVAL.ast = BuildBoolConst(true)
		}
	case 16:
		AladinoDollar = AladinoS[Aladinopt-1 : Aladinopt+1]
		{
			AladinoVAL.ast = BuildBoolConst(false)
		}
	case 17:
		AladinoDollar = AladinoS[Aladinopt-5 : Aladinopt+1]
		{
			AladinoVAL.ast = BuildFunctionCall(BuildVariable(AladinoDollar[2].str), AladinoDollar[4].astList)
		}
	case 18:
		AladinoDollar = AladinoS[Aladinopt-3 : Aladinopt+1]
		{
			AladinoVAL.astList = append([]Expr{AladinoDollar[1].ast}, AladinoDollar[3].astList...)
		}
	case 19:
		AladinoDollar = AladinoS[Aladinopt-1 : Aladinopt+1]
		{
			AladinoVAL.astList = []Expr{AladinoDollar[1].ast}
		}
	case 20:
		AladinoDollar = AladinoS[Aladinopt-0 : Aladinopt+1]
		{
			AladinoVAL.astList = []Expr{}
		}
	}
	goto Aladinostack /* stack new state and value */
}
