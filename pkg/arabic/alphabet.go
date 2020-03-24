package arabic

// Harf holds the Arabic character with its different representation forms (glyphs).
type Harf struct {
	Unicode, Isolated, Initial, Medial, Final rune
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
		Unicode:  '\u0623',
		Isolated: '\ufe83',
		Initial:  '\u0623',
		Medial:   '\ufe84',
		Final:    '\ufe84'}

	ALEF = Harf{ // ا
		Unicode:  '\u0627',
		Isolated: '\ufe8d',
		Initial:  '\u0627',
		Medial:   '\ufe8e',
		Final:    '\ufe8e'}

	ALEF_MADDA_ABOVE = Harf{ // آ
		Unicode:  '\u0622',
		Isolated: '\ufe81',
		Initial:  '\u0622',
		Medial:   '\ufe82',
		Final:    '\ufe82'}

	HAMZA = Harf{ // ء
		Unicode:  '\u0621',
		Isolated: '\ufe80',
		Initial:  '\u0621',
		Medial:   '\u0621',
		Final:    '\u0621'}

	WAW_HAMZA_ABOVE = Harf{ // ؤ
		Unicode:  '\u0624',
		Isolated: '\ufe85',
		Initial:  '\u0624',
		Medial:   '\ufe86',
		Final:    '\ufe86'}

	ALEF_HAMZA_BELOW = Harf{ // أ
		Unicode:  '\u0625',
		Isolated: '\ufe87',
		Initial:  '\u0625',
		Medial:   '\ufe88',
		Final:    '\ufe88'}

	YEH_HAMZA_ABOVE = Harf{ // ئ
		Unicode:  '\u0626',
		Isolated: '\ufe89',
		Initial:  '\ufe8b',
		Medial:   '\ufe8c',
		Final:    '\ufe8a'}

	BEH = Harf{ // ب
		Unicode:  '\u0628',
		Isolated: '\ufe8f',
		Initial:  '\ufe91',
		Medial:   '\ufe92',
		Final:    '\ufe90'}

	TEH = Harf{ // ت
		Unicode:  '\u062A',
		Isolated: '\ufe95',
		Initial:  '\ufe97',
		Medial:   '\ufe98',
		Final:    '\ufe96'}

	TEH_MARBUTA = Harf{ // ة
		Unicode:  '\u0629',
		Isolated: '\ufe93',
		Initial:  '\u0629',
		Medial:   '\u0629',
		Final:    '\ufe94'}

	THEH = Harf{ // ث
		Unicode:  '\u062b',
		Isolated: '\ufe99',
		Initial:  '\ufe9b',
		Medial:   '\ufe9c',
		Final:    '\ufe9a'}

	JEEM = Harf{ // ج
		Unicode:  '\u062c',
		Isolated: '\ufe9d',
		Initial:  '\ufe9f',
		Medial:   '\ufea0',
		Final:    '\ufe9e'}

	HAH = Harf{ // ح
		Unicode:  '\u062d',
		Isolated: '\ufea1',
		Initial:  '\ufea3',
		Medial:   '\ufea4',
		Final:    '\ufea2'}

	KHAH = Harf{ // خ
		Unicode:  '\u062e',
		Isolated: '\ufea5',
		Initial:  '\ufea7',
		Medial:   '\ufea8',
		Final:    '\ufea6'}

	DAL = Harf{ // د
		Unicode:  '\u062f',
		Isolated: '\ufea9',
		Initial:  '\u062f',
		Medial:   '\ufeaa',
		Final:    '\ufeaa'}

	THAL = Harf{ // ذ
		Unicode:  '\u0630',
		Isolated: '\ufeab',
		Initial:  '\u0630',
		Medial:   '\ufeac',
		Final:    '\ufeac'}

	REH = Harf{ // ر
		Unicode:  '\u0631',
		Isolated: '\ufead',
		Initial:  '\u0631',
		Medial:   '\ufeae',
		Final:    '\ufeae'}

	ZAIN = Harf{ // ز
		Unicode:  '\u0632',
		Isolated: '\ufeaf',
		Initial:  '\u0632',
		Medial:   '\ufeb0',
		Final:    '\ufeb0'}

	SEEN = Harf{ // س
		Unicode:  '\u0633',
		Isolated: '\ufeb1',
		Initial:  '\ufeb3',
		Medial:   '\ufeb4',
		Final:    '\ufeb2'}

	SHEEN = Harf{ // ش
		Unicode:  '\u0634',
		Isolated: '\ufeb5',
		Initial:  '\ufeb7',
		Medial:   '\ufeb8',
		Final:    '\ufeb6'}

	SAD = Harf{ // ص
		Unicode:  '\u0635',
		Isolated: '\ufeb9',
		Initial:  '\ufebb',
		Medial:   '\ufebc',
		Final:    '\ufeba'}

	DAD = Harf{ // ض
		Unicode:  '\u0636',
		Isolated: '\ufebd',
		Initial:  '\ufebf',
		Medial:   '\ufec0',
		Final:    '\ufebe'}

	TAH = Harf{ // ط
		Unicode:  '\u0637',
		Isolated: '\ufec1',
		Initial:  '\ufec3',
		Medial:   '\ufec4',
		Final:    '\ufec2'}

	ZAH = Harf{ // ظ
		Unicode:  '\u0638',
		Isolated: '\ufec5',
		Initial:  '\ufec7',
		Medial:   '\ufec8',
		Final:    '\ufec6'}

	AIN = Harf{ // ع
		Unicode:  '\u0639',
		Isolated: '\ufec9',
		Initial:  '\ufecb',
		Medial:   '\ufecc',
		Final:    '\ufeca'}

	GHAIN = Harf{ // غ
		Unicode:  '\u063a',
		Isolated: '\ufecd',
		Initial:  '\ufecf',
		Medial:   '\ufed0',
		Final:    '\ufece'}

	FEH = Harf{ // ف
		Unicode:  '\u0641',
		Isolated: '\ufed1',
		Initial:  '\ufed3',
		Medial:   '\ufed4',
		Final:    '\ufed2'}

	QAF = Harf{ // ق
		Unicode:  '\u0642',
		Isolated: '\ufed5',
		Initial:  '\ufed7',
		Medial:   '\ufed8',
		Final:    '\ufed6'}

	KAF = Harf{ // ك
		Unicode:  '\u0643',
		Isolated: '\ufed9',
		Initial:  '\ufedb',
		Medial:   '\ufedc',
		Final:    '\ufeda'}

	LAM = Harf{ // ل
		Unicode:  '\u0644',
		Isolated: '\ufedd',
		Initial:  '\ufedf',
		Medial:   '\ufee0',
		Final:    '\ufede'}

	MEEM = Harf{ // م
		Unicode:  '\u0645',
		Isolated: '\ufee1',
		Initial:  '\ufee3',
		Medial:   '\ufee4',
		Final:    '\ufee2'}

	NOON = Harf{ // ن
		Unicode:  '\u0646',
		Isolated: '\ufee5',
		Initial:  '\ufee7',
		Medial:   '\ufee8',
		Final:    '\ufee6'}

	HEH = Harf{ // ه
		Unicode:  '\u0647',
		Isolated: '\ufee9',
		Initial:  '\ufeeb',
		Medial:   '\ufeec',
		Final:    '\ufeea'}

	WAW = Harf{ // و
		Unicode:  '\u0648',
		Isolated: '\ufeed',
		Initial:  '\u0648',
		Medial:   '\ufeee',
		Final:    '\ufeee'}

	YEH = Harf{ // ي
		Unicode:  '\u064a',
		Isolated: '\ufef1',
		Initial:  '\ufef3',
		Medial:   '\ufef4',
		Final:    '\ufef2'}

	ALEF_MAKSURA = Harf{ // ى
		Unicode:  '\u0649',
		Isolated: '\ufeef',
		Initial:  '\u0649',
		Medial:   '\ufef0',
		Final:    '\ufef0'}

	TATWEEL = Harf{ // ـ
		Unicode:  '\u0640',
		Isolated: '\u0640',
		Initial:  '\u0640',
		Medial:   '\u0640',
		Final:    '\u0640'}

	LAM_ALEF = Harf{ // لا
		Unicode:  '\ufefb',
		Isolated: '\ufefb',
		Initial:  '\ufefb',
		Medial:   '\ufefc',
		Final:    '\ufefc'}

	LAM_ALEF_HAMZA_ABOVE = Harf{ // ﻷ
		Unicode:  '\ufef7',
		Isolated: '\ufef7',
		Initial:  '\ufef7',
		Medial:   '\ufef8',
		Final:    '\ufef8'}

	LAM_ALEF_HAMZA_BELOW = Harf{ // لإ
		Unicode:  '\ufef9',
		Isolated: '\ufef9',
		Initial:  '\ufef9',
		Medial:   '\ufefa',
		Final:    '\ufefa'}

	LAM_ALEF_MADDAH = Harf{ // لآ
		Unicode:  '\ufef5',
		Isolated: '\ufef5',
		Initial:  '\ufef5',
		Medial:   '\ufef6',
		Final:    '\ufef6'}
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
