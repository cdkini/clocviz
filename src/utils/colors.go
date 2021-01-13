package utils

import (
	"fmt"
	"time"

	"encoding/json"
	"math/rand"
)

type RGB struct {
	Red   int
	Green int
	Blue  int
}

func NewRGB(red int, green int, blue int) RGB {
	return RGB{red, green, blue}
}

func (r RGB) String() string {
	return fmt.Sprintf("rgb(%v,%v,%v)", r.Red, r.Green, r.Blue)
}

func (r RGB) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.String())
}

func (r RGB) Gradate(percentage float32) {
	r.Red = int(float32(r.Red) * percentage)
	r.Green = int(float32(r.Green) * percentage)
	r.Blue = int(float32(r.Blue) * percentage)
}

func GetLangColor(lang string) RGB {
	if color, ok := COLORS[lang]; ok {
		return color
	}
	// If color is not in COLORS, we save a random color for future usage.
	rand := getRandomColor()
	COLORS[lang] = rand
	return rand
}

func getRandomColor() RGB {
	rand.Seed(time.Now().UnixNano())
	red := rand.Intn(255)
	green := rand.Intn(255)
	blue := rand.Intn(255)
	return NewRGB(red, green, blue)
}

// Associations between language name and official color
// Source: https://github.com/doda/github-language-colors/blob/master/colors.json
var COLORS = map[string]RGB{
	"ABAP":                  NewRGB(232, 39, 75),
	"AGS Script":            NewRGB(185, 217, 255),
	"AMPL":                  NewRGB(230, 239, 187),
	"ANTLR":                 NewRGB(157, 195, 255),
	"API Blueprint":         NewRGB(42, 204, 168),
	"APL":                   NewRGB(90, 129, 100),
	"ASP":                   NewRGB(106, 64, 253),
	"ATS":                   NewRGB(26, 198, 32),
	"ActionScript":          NewRGB(136, 43, 15),
	"Ada":                   NewRGB(2, 248, 140),
	"Agda":                  NewRGB(49, 86, 101),
	"Alloy":                 NewRGB(100, 200, 0),
	"Arc":                   NewRGB(170, 42, 254),
	"Arduino":               NewRGB(189, 121, 209),
	"AspectJ":               NewRGB(169, 87, 176),
	"Assembly":              NewRGB(110, 76, 19),
	"AutoHotkey":            NewRGB(101, 148, 185),
	"AutoIt":                NewRGB(28, 53, 82),
	"BlitzMax":              NewRGB(205, 100, 0),
	"Boo":                   NewRGB(212, 190, 193),
	"Brainfuck":             NewRGB(47, 37, 48),
	"C Sharp":               NewRGB(23, 134, 0),
	"C":                     NewRGB(85, 85, 85),
	"CSS":                   NewRGB(86, 61, 124),
	"Chapel":                NewRGB(141, 198, 63),
	"Cirru":                 NewRGB(204, 204, 255),
	"Clarion":               NewRGB(219, 144, 30),
	"Clean":                 NewRGB(63, 133, 175),
	"Click":                 NewRGB(228, 230, 243),
	"Clojure":               NewRGB(219, 88, 85),
	"CoffeeScript":          NewRGB(36, 71, 118),
	"ColdFusion CFC":        NewRGB(237, 44, 214),
	"ColdFusion":            NewRGB(237, 44, 214),
	"Common Lisp":           NewRGB(63, 182, 139),
	"Component Pascal":      NewRGB(176, 206, 78),
	"Crystal":               NewRGB(119, 103, 145),
	"D":                     NewRGB(186, 89, 94),
	"DM":                    NewRGB(68, 114, 101),
	"Dart":                  NewRGB(0, 180, 171),
	"Diff":                  NewRGB(136, 221, 221),
	"Dogescript":            NewRGB(204, 167, 96),
	"Dylan":                 NewRGB(108, 97, 110),
	"E":                     NewRGB(204, 206, 53),
	"ECL":                   NewRGB(138, 18, 103),
	"Eagle":                 NewRGB(129, 76, 5),
	"Eiffel":                NewRGB(148, 109, 87),
	"Elixir":                NewRGB(110, 74, 126),
	"Elm":                   NewRGB(96, 181, 204),
	"Emacs Lisp":            NewRGB(192, 101, 219),
	"EmberScript":           NewRGB(255, 244, 243),
	"Erlang":                NewRGB(184, 57, 152),
	"F#":                    NewRGB(184, 69, 252),
	"FLUX":                  NewRGB(136, 204, 255),
	"FORTRAN":               NewRGB(77, 65, 177),
	"Factor":                NewRGB(99, 103, 70),
	"Fancy":                 NewRGB(123, 157, 180),
	"Fantom":                NewRGB(219, 222, 213),
	"Forth":                 NewRGB(52, 23, 8),
	"FreeMarker":            NewRGB(0, 80, 178),
	"Frege":                 NewRGB(0, 202, 254),
	"Game Maker Language":   NewRGB(143, 178, 0),
	"Glyph":                 NewRGB(228, 204, 152),
	"Gnuplot":               NewRGB(240, 169, 240),
	"Go":                    NewRGB(55, 94, 171),
	"Golo":                  NewRGB(136, 86, 42),
	"Gosu":                  NewRGB(130, 147, 127),
	"Grammatical Framework": NewRGB(121, 170, 122),
	"Groovy":                NewRGB(230, 159, 86),
	"HTML":                  NewRGB(228, 75, 35),
	"Handlebars":            NewRGB(1, 169, 214),
	"Harbour":               NewRGB(14, 96, 227),
	"Haskell":               NewRGB(41, 181, 68),
	"Haxe":                  NewRGB(223, 121, 0),
	"Hy":                    NewRGB(119, 144, 178),
	"IDL":                   NewRGB(163, 82, 47),
	"Io":                    NewRGB(169, 24, 141),
	"Ioke":                  NewRGB(7, 129, 147),
	"Isabelle":              NewRGB(254, 254, 0),
	"J":                     NewRGB(158, 237, 255),
	"JFlex":                 NewRGB(219, 202, 0),
	"JSONiq":                NewRGB(64, 212, 126),
	"Java":                  NewRGB(176, 114, 25),
	"JavaScript":            NewRGB(241, 224, 90),
	"Julia":                 NewRGB(162, 112, 186),
	"Jupyter Notebook":      NewRGB(218, 91, 11),
	"KRL":                   NewRGB(40, 67, 31),
	"Kotlin":                NewRGB(241, 142, 51),
	"LFE":                   NewRGB(0, 66, 0),
	"LOLCODE":               NewRGB(204, 153, 0),
	"LSL":                   NewRGB(61, 153, 112),
	"Lasso":                 NewRGB(153, 153, 153),
	"Latte":                 NewRGB(168, 255, 151),
	"Lex":                   NewRGB(219, 202, 0),
	"LiveScript":            NewRGB(73, 152, 134),
	"LookML":                NewRGB(101, 43, 129),
	"Lua":                   NewRGB(0, 0, 128),
	"MAXScript":             NewRGB(0, 166, 166),
	"MTML":                  NewRGB(183, 225, 244),
	"Makefile":              NewRGB(66, 120, 25),
	"Mask":                  NewRGB(249, 119, 50),
	"Matlab":                NewRGB(187, 146, 172),
	"Max":                   NewRGB(196, 167, 156),
	"Mercury":               NewRGB(255, 43, 43),
	"Metal":                 NewRGB(143, 20, 233),
	"Mirah":                 NewRGB(199, 169, 56),
	"NCL":                   NewRGB(40, 67, 31),
	"Nemerle":               NewRGB(61, 60, 110),
	"NetLinx":               NewRGB(10, 160, 255),
	"NetLinx+ERB":           NewRGB(116, 127, 170),
	"NetLogo":               NewRGB(255, 99, 117),
	"NewLisp":               NewRGB(135, 174, 215),
	"Nimrod":                NewRGB(55, 119, 91),
	"Nit":                   NewRGB(0, 153, 23),
	"Nix":                   NewRGB(126, 126, 255),
	"Nu":                    NewRGB(201, 223, 64),
	"OCaml":                 NewRGB(59, 225, 51),
	"Objective-C":           NewRGB(67, 142, 255),
	"Objective-C++":         NewRGB(104, 102, 251),
	"Objective-J":           NewRGB(255, 12, 90),
	"Omgrofl":               NewRGB(202, 187, 255),
	"Opal":                  NewRGB(247, 237, 224),
	"Oxygene":               NewRGB(205, 208, 227),
	"Oz":                    NewRGB(250, 183, 56),
	"PAWN":                  NewRGB(219, 178, 132),
	"PHP":                   NewRGB(79, 93, 149),
	"PLSQL":                 NewRGB(218, 216, 216),
	"Pan":                   NewRGB(204, 0, 0),
	"Papyrus":               NewRGB(102, 0, 204),
	"Parrot":                NewRGB(243, 202, 10),
	"Pascal":                NewRGB(176, 206, 78),
	"Perl":                  NewRGB(2, 152, 195),
	"Perl6":                 NewRGB(0, 0, 251),
	"PigLatin":              NewRGB(252, 215, 222),
	"Pike":                  NewRGB(0, 83, 144),
	"PogoScript":            NewRGB(216, 0, 116),
	"Processing":            NewRGB(0, 150, 216),
	"Prolog":                NewRGB(116, 40, 60),
	"Propeller Spin":        NewRGB(127, 162, 167),
	"Puppet":                NewRGB(48, 43, 109),
	"Pure Data":             NewRGB(145, 222, 121),
	"PureBasic":             NewRGB(90, 105, 134),
	"PureScript":            NewRGB(29, 34, 45),
	"Python":                NewRGB(53, 114, 165),
	"QML":                   NewRGB(68, 165, 28),
	"R":                     NewRGB(25, 140, 231),
	"RAML":                  NewRGB(119, 217, 251),
	"Racket":                NewRGB(34, 34, 143),
	"Ragel in Ruby Host":    NewRGB(157, 82, 0),
	"Rebol":                 NewRGB(53, 138, 91),
	"Red":                   NewRGB(238, 0, 0),
	"Ren'Py":                NewRGB(255, 127, 127),
	"Rouge":                 NewRGB(204, 0, 136),
	"Ruby":                  NewRGB(112, 21, 22),
	"Rust":                  NewRGB(222, 165, 132),
	"SAS":                   NewRGB(179, 73, 54),
	"SQF":                   NewRGB(63, 63, 63),
	"SaltStack":             NewRGB(100, 100, 100),
	"Scala":                 NewRGB(220, 50, 47),
	"Scheme":                NewRGB(30, 74, 236),
	"Self":                  NewRGB(5, 121, 170),
	"Shell":                 NewRGB(137, 224, 81),
	"Shen":                  NewRGB(18, 15, 20),
	"Slash":                 NewRGB(0, 126, 255),
	"Slim":                  NewRGB(255, 143, 119),
	"Smalltalk":             NewRGB(89, 103, 6),
	"SourcePawn":            NewRGB(92, 118, 17),
	"Squirrel":              NewRGB(128, 0, 0),
	"Stan":                  NewRGB(178, 1, 29),
	"Standard ML":           NewRGB(220, 86, 109),
	"SuperCollider":         NewRGB(70, 57, 11),
	"Swift":                 NewRGB(255, 172, 69),
	"SystemVerilog":         NewRGB(218, 225, 194),
	"Tcl":                   NewRGB(228, 204, 152),
	"TeX":                   NewRGB(61, 97, 23),
	"Turing":                NewRGB(69, 247, 21),
	"TypeScript":            NewRGB(43, 116, 137),
	"Unified Parallel C":    NewRGB(78, 54, 23),
	"Unity3D Asset":         NewRGB(171, 105, 161),
	"UnrealScript":          NewRGB(165, 76, 77),
	"VHDL":                  NewRGB(173, 178, 203),
	"Vala":                  NewRGB(251, 229, 205),
	"Verilog":               NewRGB(178, 183, 248),
	"VimL":                  NewRGB(25, 159, 75),
	"Visual Basic":          NewRGB(148, 93, 183),
	"Volt":                  NewRGB(31, 31, 31),
	"Vue":                   NewRGB(44, 62, 80),
	"Web Ontology Language": NewRGB(156, 201, 221),
	"X10":                   NewRGB(75, 107, 239),
	"XC":                    NewRGB(153, 218, 7),
	"XQuery":                NewRGB(82, 50, 231),
	"Zephir":                NewRGB(17, 143, 158),
	"cpp":                   NewRGB(243, 75, 125),
	"eC":                    NewRGB(145, 57, 96),
	"edn":                   NewRGB(219, 88, 85),
	"nesC":                  NewRGB(148, 176, 199),
	"ooc":                   NewRGB(176, 183, 126),
	"wisp":                  NewRGB(117, 130, 209),
	"xBase":                 NewRGB(64, 58, 64),
}
