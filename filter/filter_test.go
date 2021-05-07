package filter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilter_Apply(t *testing.T) {
	assert := assert.New(t)

	cases := []struct {
		available []string
		excluded  []string
		src       []string
		expect    []string
	}{
		{
			arr(),
			arr(),
			arr("a", "a.b.c", "a.d"),
			arr(),
		},
		{
			arr("*"),
			arr("*"),
			arr("a", "a.b.c", "a.d"),
			arr(),
		},
		{
			arr("a"),
			arr("b"),
			arr("a", "a.b.c", "a.d"),
			arr("a", "a.b.c", "a.d"),
		},
		{
			arr("a.b"),
			arr(),
			arr("a", "a.b.c", "a.d"),
			arr("a.b.c"),
		},
		{
			arr("*"),
			arr("a.b"),
			arr("a", "a.b.c", "a.d"),
			arr("a", "a.d"),
		},
	}

	for _, c := range cases {
		filter := NewFilter(c.available, c.excluded)
		result := filter.Apply(c.src)
		assert.EqualValues(c.expect, result)
	}
}

func TestCheckPath(t *testing.T) {
	type path struct {
		attribute string
		rule      string
	}
	m := map[path]bool{
		{
			attribute: "$$citizen_relativesescredentials.1000003.login.rel_tp_cd",
			rule:      "*",
		}: true,

		{
			attribute: "$$escredentials.1000003.ad consequat non fugiat quis.access_level",
			rule:      "$$$$escredentials.1000003.login",
		}: false,

		{
			attribute: "$$mdm_id.citizen_relatives.nulla minim ea incididunt et.start_dt",
			rule:      "$$mdm_id",
		}: true,

		{
			attribute: "$$escredentials.1000003.login.etalon_id",
			rule:      "$$escredentials.1000003.login",
		}: true,

		{
			attribute: "$$escredentials.1000003",
			rule:      "$$escredentials.1000003.user.1000003.ad consequat non fugiat quis.access_level",
		}: false,

		{
			attribute: "$$users.mdm_id.is_confirmed_offline",
			rule:      "$$mdm_id",
		}: false,

		{
			attribute: "$$citizen_relativesescredentials.1000003.login.rel_tp_cd",
			rule:      "$$mdm_id",
		}: false,
	}
	var answer bool
	for path, value := range m {
		answer = MatchPath(path.attribute, path.rule)
		if answer != value {
			t.Error("For\n", path.attribute,
				"\n", path.rule, "\nexpected",
				value, "\ngot", answer)
		}
	}
}

func arr(ss ...string) []string {
	if ss == nil {
		return []string{}
	}
	return ss
}
