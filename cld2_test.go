package cld2

import "testing"

func TestLanguageNameFromCode_withUnknown_success(t *testing.T) {
	name := LanguageNameFromCode(UNKNOWN_LANGUAGE)
	if UNKNOWN_LANGUAGE_NAME != name {
		t.Fatalf("Expected unknown language, found %s.", name)
	}
}

func TestLanguageNameFromCode_withEn_success(t *testing.T) {
	name := LanguageNameFromCode("en")
	if "ENGLISH" != name {
		t.Fatalf("Expected ENGLISH, found %s.", name)
	}
}

func TestDetectLang_withPersian_success(t *testing.T) {
	lang := DetectLang(`رُم پایتخت کشور ایتالیا، بزرگترین و پرجمعیت‌ترین شهر این کشور با ۲۶۴۹۷۲۴ سکن
	ه و همچنین مرکز ناحیهٔ لاتزیو می‌باشد. هم چنین رم با مساحت ۱۳۶۲۸۷ `)
	if "fa" != lang {
		t.Fatalf("Expected 'fa', found %s.", lang)
	}
}