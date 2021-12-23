package translation_test

import (
	"github.com/Vadakan/holmes89/translation"
	"testing"
)

func TestTranslate(t *testing.T) {
	// Arrange
	tt := []struct {
		Word        string
		Language    string
		Translation string
	}{
		{
			Word:        "hello",
			Language:    "english",
			Translation: "hello",
		},
		{
			Word:        "hello",
			Language:    "german",
			Translation: "hallo",
		},
		{
			Word:        "hello",
			Language:    "finnish",
			Translation: "hei",
		},
		{
			Word:        "hello",
			Language:    "dutch",
			Translation: "",
		},
		{
			Word:        "bye",
			Language:    "german",
			Translation: "",
		},
	}

	for _, test := range tt {
		// Act
		res := translation.Translate(test.Word, test.Language)

		// Assert
		if res != test.Translation {
			t.Errorf(
				`expected "%s" to be "%s" from "%s" but received "%s"`,
				test.Word, test.Language, test.Translation, res)
		}
	}
}

func Translate(word, language string) {
	panic("unimplemented")
}
