package rel

import (
	"testing"

	"github.com/arr-ai/frozen"
	"github.com/stretchr/testify/assert"
)

func TestProjectionBasedOnNames(t *testing.T) {
	t.Parallel()

	r := Relation{attrs: NamesSlice{"a", "b", "c"}, p: valueProjector{0, 1, 2}}
	assert.Equal(t, valueProjector{2, 1, 0}, r.projectionBasedOnNames(NamesSlice{"c", "b", "a"}))
	assert.Equal(t, valueProjector{0, 1, 2}, r.projectionBasedOnNames(NamesSlice{"a", "b", "c"}))
	assert.Equal(t, valueProjector{0, 1}, r.projectionBasedOnNames(NamesSlice{"a", "b"}))
	assert.Equal(t, valueProjector{}, r.projectionBasedOnNames(NamesSlice{}))
	assert.Equal(t, valueProjector{0, 0, 2}, r.projectionBasedOnNames(NamesSlice{"a", "a", "c"}))

	r = Relation{attrs: NamesSlice{"a", "b", "c"}, p: valueProjector{2, 0, 1}}
	assert.Equal(t, valueProjector{1, 0, 2}, r.projectionBasedOnNames(NamesSlice{"c", "b", "a"}))
	assert.Equal(t, valueProjector{2, 0, 1}, r.projectionBasedOnNames(NamesSlice{"a", "b", "c"}))
	assert.Equal(t, valueProjector{2, 0}, r.projectionBasedOnNames(NamesSlice{"a", "b"}))
	assert.Equal(t, valueProjector{}, r.projectionBasedOnNames(NamesSlice{}))
	assert.Equal(t, valueProjector{2, 2, 1}, r.projectionBasedOnNames(NamesSlice{"a", "a", "c"}))

	assert.Panics(t, func() { r.projectionBasedOnNames(NamesSlice{"d"}) })
}

func TestRelationString(t *testing.T) {
	t.Parallel()

	r := Relation{
		attrs: NamesSlice{"c", "b", "a"},
		p:     valueProjector{0, 1, 2},
		rows: &positionalRelation{
			set: frozen.NewSet(
				row(1, 1, 2),
				row(1, 2, 3),
				row(2, 1, 2),
			),
		},
	}
	assert.Equal(t, "{|a, b, c| (2, 1, 1), (2, 1, 2), (3, 2, 1)}", r.String())

	r = Relation{
		attrs: NamesSlice{"c", "b", "a"},
		p:     valueProjector{2, 0, 1},
		rows: &positionalRelation{
			set: frozen.NewSet(
				row(1, 1, 2),
				row(1, 2, 3),
				row(2, 1, 2),
			),
		},
	}
	assert.Equal(t, "{|a, b, c| (1, 1, 2), (1, 2, 2), (2, 1, 3)}", r.String())
}
