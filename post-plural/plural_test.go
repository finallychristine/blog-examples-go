package post_plural

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/text/language"
)

func Test_GetForm(t *testing.T) {
	/* Be aware this version of go (1.25.4) is using cldr32, which might not match
	what is shown in the URLs below */

	/* English: https://www.unicode.org/cldr/charts/48/supplemental/language_plural_rules.html#en */

	// "1" is other in english
	assertForm(t, "en-US", 1.0, 0, FormOne)
	// "1.1" is other in english
	assertForm(t, "en-US", 1.1, 1, FormOther)
	// "1.00" is other, perhaps unintuitively.
	assertForm(t, "en-US", 1.0, 2, FormOther)
	// "2" is other
	assertForm(t, "en-US", 2.0, 0, FormOther)

	/* Russian: https://www.unicode.org/cldr/charts/48/supplemental/language_plural_rules.html#ru */
	assertForm(t, "ru", 1, 0, FormOne)
	assertForm(t, "ru", 1, 1, FormOther)
	assertForm(t, "ru", 21, 0, FormOne)
	// Note that our choice to only do "one" or "other" has implications for languages like russian,
	// as Russian has a "few" type of pluralization
	assertForm(t, "ru", 2, 0, FormOther)

	/** Japanese: https://www.unicode.org/cldr/charts/48/supplemental/language_plural_rules.html#ja */
	// The concept of a "one" pluralization does not exist in Japanese
	assertForm(t, "ja", 1, 0, FormOther)
	assertForm(t, "ja", 1, 3, FormOther)
	assertForm(t, "ja", 10, 3, FormOther)
}

func assertForm(t *testing.T, tagStr string, count float64, precision int, expected Form) {
	t.Helper()

	tag, err := language.Parse(tagStr)
	require.NoError(t, err)

	actual := GetForm(tag, count, precision)
	assert.Equal(t, expected, actual, "language %s", tagStr)
}
