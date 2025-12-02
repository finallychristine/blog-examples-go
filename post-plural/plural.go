package post_plural

import (
	"math"

	"golang.org/x/text/feature/plural"
	"golang.org/x/text/language"
)

// Form is the possible plural forms.
//
// CLDR provides multiple forms, but I've found that most use cases
// are perfectly okay using the minimal amount of one or other.
//
// See https://cldr.unicode.org/index/cldr-spec/plural-rules
type Form string

const (
	FormOther Form = "other"
	FormOne   Form = "one"
)

func GetForm(lang language.Tag, count float64, precision int) Form {
	// example: count = 123.45678, precision = 3
	exp10 := math.Pow10(precision)       // 10^3 =  1000
	intpart := int(count)                // intpart = 123
	decpartf := count - float64(intpart) // decpartf = 0.4567
	decparti := int(decpartf * exp10)    // decparti = int(0.4567 * 10^3) = int(456.78) = 456

	// https://pkg.go.dev/golang.org/x/text@v0.31.0/feature/plural#Rules.MatchPlural
	//
	// This hard-codes "Cardinal" as the plural matches as that is what is used most of the time.
	// Cardinal is for quantities, e.g. "20 cents".
	// Ordinal is for position, e.g. "take the third right turn"
	//
	// We could make this more accurate if we wanted to calculate number of trailing zeroes.
	// But assuming there are no trailing zeroes proves to be close enough in the majority of cases.
	form := plural.Cardinal.MatchPlural(
		lang,

		intpart, // i  integer digits of n.

		precision, // v  number of visible fraction digits in n, with trailing zeros.
		precision, // w  number of visible fraction digits in n, without trailing zeros.

		decparti, // f  visible fractional digits in n, with trailing zeros (f = t * 10^(v-w))
		decparti, // t  visible fractional digits in n, without trailing zeros.
	)

	if form == plural.One {
		return FormOne
	}

	return FormOther
}
