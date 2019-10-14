package goarabic

// Harf holds the Arabic character with its different representation forms (glyphs).
type Harf struct {
	Unicode, Isolated, Beginning, Medium, Final rune
}

// Vowels (Tashkeel) characters.
var (
	FATHA    = '\u064e'
	FATHATAN = '\u064b'
	DAMMA    = '\u064f'
	DAMMATAN = '\u064c'
	KASRA    = '\u0650'
	KASRATAN = '\u064d'
	SHADDA   = '\u0651'
	SUKUN    = '\u0652'
)

// Arabic Alphabet using the new Harf type.
var (
	ALEF_HAMZA_ABOVE = Harf{ // أ
		Unicode:   '\u0623',
		Isolated:  '\ufe83',
		Beginning: '\u0623',
		Medium:    '\ufe84',
		Final:     '\ufe84'}

	ALEF = Harf{ // ا
		Unicode:   '\u0627',
		Isolated:  '\ufe8d',
		Beginning: '\u0627',
		Medium:    '\ufe8e',
		Final:     '\ufe8e'}

	ALEF_MADDA_ABOVE = Harf{ // آ
		Unicode:   '\u0622',
		Isolated:  '\ufe81',
		Beginning: '\u0622',
		Medium:    '\ufe82',
		Final:     '\ufe82'}

	HAMZA = Harf{ // ء
		Unicode:   '\u0621',
		Isolated:  '\ufe80',
		Beginning: '\u0621',
		Medium:    '\u0621',
		Final:     '\u0621'}

	WAW_HAMZA_ABOVE = Harf{ // ؤ
		Unicode:   '\u0624',
		Isolated:  '\ufe85',
		Beginning: '\u0624',
		Medium:    '\ufe86',
		Final:     '\ufe86'}

	ALEF_HAMZA_BELOW = Harf{ // أ
		Unicode:   '\u0625',
		Isolated:  '\ufe87',
		Beginning: '\u0625',
		Medium:    '\ufe88',
		Final:     '\ufe88'}

	YEH_HAMZA_ABOVE = Harf{ // ئ
		Unicode:   '\u0626',
		Isolated:  '\ufe89',
		Beginning: '\ufe8b',
		Medium:    '\ufe8c',
		Final:     '\ufe8a'}

	BEH = Harf{ // ب
		Unicode:   '\u0628',
		Isolated:  '\ufe8f',
		Beginning: '\ufe91',
		Medium:    '\ufe92',
		Final:     '\ufe90'}

	TEH = Harf{ // ت
		Unicode:   '\u062A',
		Isolated:  '\ufe95',
		Beginning: '\ufe97',
		Medium:    '\ufe98',
		Final:     '\ufe96'}

	TEH_MARBUTA = Harf{ // ة
		Unicode:   '\u0629',
		Isolated:  '\ufe93',
		Beginning: '\u0629',
		Medium:    '\u0629',
		Final:     '\ufe94'}

	THEH = Harf{ // ث
		Unicode:   '\u062b',
		Isolated:  '\ufe99',
		Beginning: '\ufe9b',
		Medium:    '\ufe9c',
		Final:     '\ufe9a'}

	JEEM = Harf{ // ج
		Unicode:   '\u062c',
		Isolated:  '\ufe9d',
		Beginning: '\ufe9f',
		Medium:    '\ufea0',
		Final:     '\ufe9e'}

	HAH = Harf{ // ح
		Unicode:   '\u062d',
		Isolated:  '\ufea1',
		Beginning: '\ufea3',
		Medium:    '\ufea4',
		Final:     '\ufea2'}

	KHAH = Harf{ // خ
		Unicode:   '\u062e',
		Isolated:  '\ufea5',
		Beginning: '\ufea7',
		Medium:    '\ufea8',
		Final:     '\ufea6'}

	DAL = Harf{ // د
		Unicode:   '\u062f',
		Isolated:  '\ufea9',
		Beginning: '\u062f',
		Medium:    '\ufeaa',
		Final:     '\ufeaa'}

	THAL = Harf{ // ذ
		Unicode:   '\u0630',
		Isolated:  '\ufeab',
		Beginning: '\u0630',
		Medium:    '\ufeac',
		Final:     '\ufeac'}

	REH = Harf{ // ر
		Unicode:   '\u0631',
		Isolated:  '\ufead',
		Beginning: '\u0631',
		Medium:    '\ufeae',
		Final:     '\ufeae'}

	ZAIN = Harf{ // ز
		Unicode:   '\u0632',
		Isolated:  '\ufeaf',
		Beginning: '\u0632',
		Medium:    '\ufeb0',
		Final:     '\ufeb0'}

	SEEN = Harf{ // س
		Unicode:   '\u0633',
		Isolated:  '\ufeb1',
		Beginning: '\ufeb3',
		Medium:    '\ufeb4',
		Final:     '\ufeb2'}

	SHEEN = Harf{ // ش
		Unicode:   '\u0634',
		Isolated:  '\ufeb5',
		Beginning: '\ufeb7',
		Medium:    '\ufeb8',
		Final:     '\ufeb6'}

	SAD = Harf{ // ص
		Unicode:   '\u0635',
		Isolated:  '\ufeb9',
		Beginning: '\ufebb',
		Medium:    '\ufebc',
		Final:     '\ufeba'}

	DAD = Harf{ // ض
		Unicode:   '\u0636',
		Isolated:  '\ufebd',
		Beginning: '\ufebf',
		Medium:    '\ufec0',
		Final:     '\ufebe'}

	TAH = Harf{ // ط
		Unicode:   '\u0637',
		Isolated:  '\ufec1',
		Beginning: '\ufec3',
		Medium:    '\ufec4',
		Final:     '\ufec2'}

	ZAH = Harf{ // ظ
		Unicode:   '\u0638',
		Isolated:  '\ufec5',
		Beginning: '\ufec7',
		Medium:    '\ufec8',
		Final:     '\ufec6'}

	AIN = Harf{ // ع
		Unicode:   '\u0639',
		Isolated:  '\ufec9',
		Beginning: '\ufecb',
		Medium:    '\ufecc',
		Final:     '\ufeca'}

	GHAIN = Harf{ // غ
		Unicode:   '\u063a',
		Isolated:  '\ufecd',
		Beginning: '\ufecf',
		Medium:    '\ufed0',
		Final:     '\ufece'}

	FEH = Harf{ // ف
		Unicode:   '\u0641',
		Isolated:  '\ufed1',
		Beginning: '\ufed3',
		Medium:    '\ufed4',
		Final:     '\ufed2'}

	QAF = Harf{ // ق
		Unicode:   '\u0642',
		Isolated:  '\ufed5',
		Beginning: '\ufed7',
		Medium:    '\ufed8',
		Final:     '\ufed6'}

	KAF = Harf{ // ك
		Unicode:   '\u0643',
		Isolated:  '\ufed9',
		Beginning: '\ufedb',
		Medium:    '\ufedc',
		Final:     '\ufeda'}

	LAM = Harf{ // ل
		Unicode:   '\u0644',
		Isolated:  '\ufedd',
		Beginning: '\ufedf',
		Medium:    '\ufee0',
		Final:     '\ufede'}

	MEEM = Harf{ // م
		Unicode:   '\u0645',
		Isolated:  '\ufee1',
		Beginning: '\ufee3',
		Medium:    '\ufee4',
		Final:     '\ufee2'}

	NOON = Harf{ // ن
		Unicode:   '\u0646',
		Isolated:  '\ufee5',
		Beginning: '\ufee7',
		Medium:    '\ufee8',
		Final:     '\ufee6'}

	HEH = Harf{ // ه
		Unicode:   '\u0647',
		Isolated:  '\ufee9',
		Beginning: '\ufeeb',
		Medium:    '\ufeec',
		Final:     '\ufeea'}

	WAW = Harf{ // و
		Unicode:   '\u0648',
		Isolated:  '\ufeed',
		Beginning: '\u0648',
		Medium:    '\ufeee',
		Final:     '\ufeee'}

	YEH = Harf{ // ي
		Unicode:   '\u064a',
		Isolated:  '\ufef1',
		Beginning: '\ufef3',
		Medium:    '\ufef4',
		Final:     '\ufef2'}

	ALEF_MAKSURA = Harf{ // ى
		Unicode:   '\u0649',
		Isolated:  '\ufeef',
		Beginning: '\u0649',
		Medium:    '\ufef0',
		Final:     '\ufef0'}

	TATWEEL = Harf{ // ـ
		Unicode:   '\u0640',
		Isolated:  '\u0640',
		Beginning: '\u0640',
		Medium:    '\u0640',
		Final:     '\u0640'}

	LAM_ALEF = Harf{ // لا
		Unicode:   '\ufefb',
		Isolated:  '\ufefb',
		Beginning: '\ufefb',
		Medium:    '\ufefc',
		Final:     '\ufefc'}

	LAM_ALEF_HAMZA_ABOVE = Harf{ // ﻷ
		Unicode:   '\ufef7',
		Isolated:  '\ufef7',
		Beginning: '\ufef7',
		Medium:    '\ufef8',
		Final:     '\ufef8'}

	LAM_ALEF_HAMZA_BELOW = Harf{ // لإ
		Unicode:   '\ufef9',
		Isolated:  '\ufef9',
		Beginning: '\ufef9',
		Medium:    '\ufefa',
		Final:     '\ufefa'}

	LAM_ALEF_MADDAH = Harf{ // لآ
		Unicode:   '\ufef5',
		Isolated:  '\ufef5',
		Beginning: '\ufef5',
		Medium:    '\ufef6',
		Final:     '\ufef6'}
)

var alphabet = []Harf{
	ALEF_HAMZA_ABOVE,
	ALEF,
	ALEF_MADDA_ABOVE,
	HAMZA,
	WAW_HAMZA_ABOVE,
	ALEF_HAMZA_BELOW,
	YEH_HAMZA_ABOVE,
	BEH,
	TEH,
	TEH_MARBUTA,
	THEH,
	JEEM,
	HAH,
	KHAH,
	DAL,
	THAL,
	REH,
	ZAIN,
	SEEN,
	SHEEN,
	SAD,
	DAD,
	TAH,
	ZAH,
	AIN,
	GHAIN,
	FEH,
	QAF,
	KAF,
	LAM,
	MEEM,
	NOON,
	HEH,
	WAW,
	YEH,
	ALEF_MAKSURA,
	TATWEEL,
	LAM_ALEF,
	LAM_ALEF_HAMZA_ABOVE,
	LAM_ALEF_HAMZA_BELOW,
	LAM_ALEF_MADDAH,
}

// use map for faster lookups.
var tashkeel = map[rune]bool{FATHA: true, FATHATAN: true, DAMMA: true,
	DAMMATAN: true, KASRA: true, KASRATAN: true,
	SHADDA: true, SUKUN: true}

// use map for faster lookups.
// var special_char = map[rune]bool{"": true, ' ': true, '?': true,
//	'؟': true, '.': true, KASRATAN: true,
//	SHADDA: true, SUKUN: true}

// use map for faster lookups.
var Beginning_after = map[Harf]bool{
	ALEF_HAMZA_ABOVE:     true,
	ALEF_MADDA_ABOVE:     true,
	ALEF:                 true,
	HAMZA:                true,
	WAW_HAMZA_ABOVE:      true,
	ALEF_HAMZA_BELOW:     true,
	TEH_MARBUTA:          true,
	DAL:                  true,
	THAL:                 true,
	REH:                  true,
	ZAIN:                 true,
	WAW:                  true,
	ALEF_MAKSURA:         true,
	LAM_ALEF:             true,
	LAM_ALEF_HAMZA_ABOVE: true,
	LAM_ALEF_HAMZA_BELOW: true,
	LAM_ALEF_MADDAH:      true,
}

var ligatures = map[Harf]map[Harf]Harf{
	LAM: {
		ALEF:             LAM_ALEF,
		ALEF_HAMZA_ABOVE: LAM_ALEF_HAMZA_ABOVE,
		ALEF_HAMZA_BELOW: LAM_ALEF_HAMZA_BELOW,
		ALEF_MADDA_ABOVE: LAM_ALEF_MADDAH,
	},
}
