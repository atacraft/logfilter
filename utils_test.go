package logfilter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCall(t *testing.T) {
	t.Log("Describe #Call func behaviour")
	{
		type card struct {
			Label  string
			Owner  string
			ID     string `log:"omit"`
			Number string `log:"filtered"`
		}

		c1 := card{
			Label:  "c1",
			Owner:  "manuel",
			ID:     "123456789",
			Number: "12345678901234567",
		}

		t.Log("\tContext: when called with wrong in arg")
		{
			flt := ""
			err1 := Call(nil, &flt)
			assert.NotNil(t, err1, "For wrong in arg error should not be nil")
			flt = ""
			err2 := Call(c1, &flt)
			assert.NotNil(t, err2, "For wrong in arg error should not be nil")
			flt = ""
			err3 := Call(flt, &flt)
			assert.NotNil(t, err3, "For wrong in arg error should not be nil")
		}

		t.Log("\tContext: when called with wrong out arg")
		{
			err := Call(&c1, nil)
			assert.NotNil(t, err, "For wrong out arg error should not be nil")
			err2 := Call(&c1, "")
			assert.NotNil(t, err2, "For wrong out arg error should not be nil")
			flt := 123
			err3 := Call(&c1, flt)
			assert.NotNil(t, err3, "For wrong out arg error should not be nil")
		}

		t.Log("\tContext: when both in and out args are wrong")
		{
			err := Call(nil, nil)
			assert.NotNil(t, err, "For wrong in and out arg error should not be nil")
		}

		t.Log("\tContext: when in struct is not tagged")
		{
			type untaggedCard struct {
				Label  string
				Owner  string
				ID     string
				Number string
			}

			c := untaggedCard{
				Label:  "c1",
				Owner:  "manuel",
				ID:     "123456789",
				Number: "12345678901234567",
			}

			flt := ""
			_ = Call(&c, &flt)
			wanted := "Label: c1; Owner: manuel; ID: 123456789; Number: 12345678901234567;"
			assert.Equal(t, wanted, flt, "Shoud be equal")
		}

		t.Log("\tContext: when omit is used as tag value")
		{
			type omitTaggedCard struct {
				Label  string
				Owner  string `log:"omit"`
				ID     string
				Number string
			}

			c := omitTaggedCard{
				Label:  "c1",
				Owner:  "manuel",
				ID:     "123456789",
				Number: "12345678901234567",
			}

			flt := ""
			_ = Call(&c, &flt)
			wanted := "Label: c1; ID: 123456789; Number: 12345678901234567;"
			assert.Equal(t, wanted, flt, "Shoud be equal")
		}

		t.Log("\tContext: when filtered is used as tag value")
		{
			type filteredTaggedCard struct {
				Label  string `log:"filtered"`
				Owner  string
				ID     string
				Number string
			}

			c := filteredTaggedCard{
				Label:  "c1",
				Owner:  "manuel",
				ID:     "123456789",
				Number: "12345678901234567",
			}

			flt := ""
			_ = Call(&c, &flt)
			wanted := "Label: *************; Owner: manuel; ID: 123456789; Number: 12345678901234567;"
			assert.Equal(t, wanted, flt, "Shoud be equal")
		}

		t.Log("\tContext: when omit and filtered are used as tag values")
		{
			flt := ""
			_ = Call(&c1, &flt)
			wanted := "Label: c1; Owner: manuel; Number: *************;"
			assert.Equal(t, wanted, flt, "Shoud be equal")
		}
	}
}
