package saml

import (
	"testing"
	"time"

	"gotest.tools/assert"
	is "gotest.tools/assert/cmp"
)

func TestRelaxedTimeFormat(t *testing.T) {
	rt := time.Date(1981, 02, 03, 14, 15, 16, 17, time.UTC)
	assert.Check(t, is.Equal("1981-02-03T14:15:16Z", RelaxedTime(rt).String()))

	buf, err := RelaxedTime(rt).MarshalText()
	assert.Check(t, err)
	assert.Check(t, is.Equal("1981-02-03T14:15:16Z", string(buf)))

	loc, err := time.LoadLocation("America/New_York")
	assert.Check(t, err)
	rt = time.Date(1981, 02, 03, 9, 15, 16, 17, loc)

	assert.Check(t, is.Equal("1981-02-03T14:15:16Z", RelaxedTime(rt).String()))
	buf, err = RelaxedTime(rt).MarshalText()
	assert.Check(t, err)
	assert.Check(t, is.Equal("1981-02-03T14:15:16Z", string(buf)))
}
