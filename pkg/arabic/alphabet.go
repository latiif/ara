package arabic

// Harf holds the Arabic character with its different representation forms (glyphs).
type Harf struct {
	Unicode, Isolated, Initial, Medial, Final rune
}

// Vowels (Tashkeel) characters.
var (
	fatha     = '\u064e'
	fathataan = '\u064b'
	damma     = '\u064f'
	dammataan = '\u064c'
	kasra     = '\u0650'
	kasrataan = '\u064d'
	shadda    = '\u0651'
	sukun     = '\u0652'
)

// Arabic Alphabet using the new Harf type.
var (
	alefHamzaAbove = Harf{ // أ
		Unicode:  '\u0623',
		Isolated: '\ufe83',
		Initial:  '\u0623',
		Medial:   '\ufe84',
		Final:    '\ufe84'}

	alef = Harf{ // ا
		Unicode:  '\u0627',
		Isolated: '\ufe8d',
		Initial:  '\u0627',
		Medial:   '\ufe8e',
		Final:    '\ufe8e'}

	alefMaddaAbove = Harf{ // آ
		Unicode:  '\u0622',
		Isolated: '\ufe81',
		Initial:  '\u0622',
		Medial:   '\ufe82',
		Final:    '\ufe82'}

	hamza = Harf{ // ء
		Unicode:  '\u0621',
		Isolated: '\ufe80',
		Initial:  '\u0621',
		Medial:   '\u0621',
		Final:    '\u0621'}

	wawHamzaAbove = Harf{ // ؤ
		Unicode:  '\u0624',
		Isolated: '\ufe85',
		Initial:  '\u0624',
		Medial:   '\ufe86',
		Final:    '\ufe86'}

	alefHamzaBelow = Harf{ // أ
		Unicode:  '\u0625',
		Isolated: '\ufe87',
		Initial:  '\u0625',
		Medial:   '\ufe88',
		Final:    '\ufe88'}

	yehHamzaAbove = Harf{ // ئ
		Unicode:  '\u0626',
		Isolated: '\ufe89',
		Initial:  '\ufe8b',
		Medial:   '\ufe8c',
		Final:    '\ufe8a'}

	beh = Harf{ // ب
		Unicode:  '\u0628',
		Isolated: '\ufe8f',
		Initial:  '\ufe91',
		Medial:   '\ufe92',
		Final:    '\ufe90'}

	teh = Harf{ // ت
		Unicode:  '\u062A',
		Isolated: '\ufe95',
		Initial:  '\ufe97',
		Medial:   '\ufe98',
		Final:    '\ufe96'}

	tehMarbuta = Harf{ // ة
		Unicode:  '\u0629',
		Isolated: '\ufe93',
		Initial:  '\u0629',
		Medial:   '\u0629',
		Final:    '\ufe94'}

	theh = Harf{ // ث
		Unicode:  '\u062b',
		Isolated: '\ufe99',
		Initial:  '\ufe9b',
		Medial:   '\ufe9c',
		Final:    '\ufe9a'}

	jeem = Harf{ // ج
		Unicode:  '\u062c',
		Isolated: '\ufe9d',
		Initial:  '\ufe9f',
		Medial:   '\ufea0',
		Final:    '\ufe9e'}

	hah = Harf{ // ح
		Unicode:  '\u062d',
		Isolated: '\ufea1',
		Initial:  '\ufea3',
		Medial:   '\ufea4',
		Final:    '\ufea2'}

	khah = Harf{ // خ
		Unicode:  '\u062e',
		Isolated: '\ufea5',
		Initial:  '\ufea7',
		Medial:   '\ufea8',
		Final:    '\ufea6'}

	dal = Harf{ // د
		Unicode:  '\u062f',
		Isolated: '\ufea9',
		Initial:  '\u062f',
		Medial:   '\ufeaa',
		Final:    '\ufeaa'}

	thal = Harf{ // ذ
		Unicode:  '\u0630',
		Isolated: '\ufeab',
		Initial:  '\u0630',
		Medial:   '\ufeac',
		Final:    '\ufeac'}

	reh = Harf{ // ر
		Unicode:  '\u0631',
		Isolated: '\ufead',
		Initial:  '\u0631',
		Medial:   '\ufeae',
		Final:    '\ufeae'}

	zain = Harf{ // ز
		Unicode:  '\u0632',
		Isolated: '\ufeaf',
		Initial:  '\u0632',
		Medial:   '\ufeb0',
		Final:    '\ufeb0'}

	seen = Harf{ // س
		Unicode:  '\u0633',
		Isolated: '\ufeb1',
		Initial:  '\ufeb3',
		Medial:   '\ufeb4',
		Final:    '\ufeb2'}

	sheen = Harf{ // ش
		Unicode:  '\u0634',
		Isolated: '\ufeb5',
		Initial:  '\ufeb7',
		Medial:   '\ufeb8',
		Final:    '\ufeb6'}

	sad = Harf{ // ص
		Unicode:  '\u0635',
		Isolated: '\ufeb9',
		Initial:  '\ufebb',
		Medial:   '\ufebc',
		Final:    '\ufeba'}

	dad = Harf{ // ض
		Unicode:  '\u0636',
		Isolated: '\ufebd',
		Initial:  '\ufebf',
		Medial:   '\ufec0',
		Final:    '\ufebe'}

	tah = Harf{ // ط
		Unicode:  '\u0637',
		Isolated: '\ufec1',
		Initial:  '\ufec3',
		Medial:   '\ufec4',
		Final:    '\ufec2'}

	zah = Harf{ // ظ
		Unicode:  '\u0638',
		Isolated: '\ufec5',
		Initial:  '\ufec7',
		Medial:   '\ufec8',
		Final:    '\ufec6'}

	ayn = Harf{ // ع
		Unicode:  '\u0639',
		Isolated: '\ufec9',
		Initial:  '\ufecb',
		Medial:   '\ufecc',
		Final:    '\ufeca'}

	ghayn = Harf{ // غ
		Unicode:  '\u063a',
		Isolated: '\ufecd',
		Initial:  '\ufecf',
		Medial:   '\ufed0',
		Final:    '\ufece'}

	feh = Harf{ // ف
		Unicode:  '\u0641',
		Isolated: '\ufed1',
		Initial:  '\ufed3',
		Medial:   '\ufed4',
		Final:    '\ufed2'}

	qaf = Harf{ // ق
		Unicode:  '\u0642',
		Isolated: '\ufed5',
		Initial:  '\ufed7',
		Medial:   '\ufed8',
		Final:    '\ufed6'}

	kaf = Harf{ // ك
		Unicode:  '\u0643',
		Isolated: '\ufed9',
		Initial:  '\ufedb',
		Medial:   '\ufedc',
		Final:    '\ufeda'}

	lam = Harf{ // ل
		Unicode:  '\u0644',
		Isolated: '\ufedd',
		Initial:  '\ufedf',
		Medial:   '\ufee0',
		Final:    '\ufede'}

	meem = Harf{ // م
		Unicode:  '\u0645',
		Isolated: '\ufee1',
		Initial:  '\ufee3',
		Medial:   '\ufee4',
		Final:    '\ufee2'}

	noon = Harf{ // ن
		Unicode:  '\u0646',
		Isolated: '\ufee5',
		Initial:  '\ufee7',
		Medial:   '\ufee8',
		Final:    '\ufee6'}

	heh = Harf{ // ه
		Unicode:  '\u0647',
		Isolated: '\ufee9',
		Initial:  '\ufeeb',
		Medial:   '\ufeec',
		Final:    '\ufeea'}

	waw = Harf{ // و
		Unicode:  '\u0648',
		Isolated: '\ufeed',
		Initial:  '\u0648',
		Medial:   '\ufeee',
		Final:    '\ufeee'}

	yeh = Harf{ // ي
		Unicode:  '\u064a',
		Isolated: '\ufef1',
		Initial:  '\ufef3',
		Medial:   '\ufef4',
		Final:    '\ufef2'}

	alefMaksura = Harf{ // ى
		Unicode:  '\u0649',
		Isolated: '\ufeef',
		Initial:  '\u0649',
		Medial:   '\ufef0',
		Final:    '\ufef0'}

	kasheeda = Harf{ // ـ
		Unicode:  '\u0640',
		Isolated: '\u0640',
		Initial:  '\u0640',
		Medial:   '\u0640',
		Final:    '\u0640'}

	lamAlef = Harf{ // لا
		Unicode:  '\ufefb',
		Isolated: '\ufefb',
		Initial:  '\ufefb',
		Medial:   '\ufefc',
		Final:    '\ufefc'}

	lamAlefHamzaAbove = Harf{ // ﻷ
		Unicode:  '\ufef7',
		Isolated: '\ufef7',
		Initial:  '\ufef7',
		Medial:   '\ufef8',
		Final:    '\ufef8'}

	lamAlefHamzaBelow = Harf{ // لإ
		Unicode:  '\ufef9',
		Isolated: '\ufef9',
		Initial:  '\ufef9',
		Medial:   '\ufefa',
		Final:    '\ufefa'}

	lamAlefMaddah = Harf{ // لآ
		Unicode:  '\ufef5',
		Isolated: '\ufef5',
		Initial:  '\ufef5',
		Medial:   '\ufef6',
		Final:    '\ufef6'}
)

var alphabet = []Harf{
	alefHamzaAbove,
	alef,
	alefMaddaAbove,
	hamza,
	wawHamzaAbove,
	alefHamzaBelow,
	yehHamzaAbove,
	beh,
	teh,
	tehMarbuta,
	theh,
	jeem,
	hah,
	khah,
	dal,
	thal,
	reh,
	zain,
	seen,
	sheen,
	sad,
	dad,
	tah,
	zah,
	ayn,
	ghayn,
	feh,
	qaf,
	kaf,
	lam,
	meem,
	noon,
	heh,
	waw,
	yeh,
	alefMaksura,
	kasheeda,
	lamAlef,
	lamAlefHamzaAbove,
	lamAlefHamzaBelow,
	lamAlefMaddah,
}

// use map for faster lookups.
var tashkeel = map[rune]bool{fatha: true, fathataan: true, damma: true,
	dammataan: true, kasra: true, kasrataan: true,
	shadda: true, sukun: true}

// use map for faster lookups.
// var special_char = map[rune]bool{"": true, ' ': true, '?': true,
//	'؟': true, '.': true, KASRATAN: true,
//	SHADDA: true, SUKUN: true}

// initialAfter use map for faster lookups.
var initialAfter = map[Harf]bool{
	alefHamzaAbove:    true,
	alefMaddaAbove:    true,
	alef:              true,
	hamza:             true,
	wawHamzaAbove:     true,
	alefHamzaBelow:    true,
	tehMarbuta:        true,
	dal:               true,
	thal:              true,
	reh:               true,
	zain:              true,
	waw:               true,
	alefMaksura:       true,
	lamAlef:           true,
	lamAlefHamzaAbove: true,
	lamAlefHamzaBelow: true,
	lamAlefMaddah:     true,
}

var ligatures = map[Harf]map[Harf]Harf{
	lam: {
		alef:           lamAlef,
		alefHamzaAbove: lamAlefHamzaAbove,
		alefHamzaBelow: lamAlefHamzaBelow,
		alefMaddaAbove: lamAlefMaddah,
	},
}
