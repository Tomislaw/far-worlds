package ecs

type Tag struct {
	flags   uint64 // limited to 64 components
	inverse bool
}

func (tag Tag) matches(smallertag Tag) bool {
	res := tag.flags&smallertag.flags == smallertag.flags

	if smallertag.inverse {
		return !res
	}

	return res
}

func (tag *Tag) binaryORInPlace(othertag Tag) *Tag {
	tag.flags |= othertag.flags
	return tag
}

func (tag *Tag) binaryNOTInPlace(othertag Tag) *Tag {
	tag.flags ^= othertag.flags
	return tag
}

func (tag Tag) clone() Tag {
	return tag
}

func (tag Tag) Inverse(values ...bool) Tag {

	clone := tag.clone()
	inverse := true
	if len(values) > 0 {
		inverse = values[0]
	}

	clone.inverse = inverse
	return clone
}
