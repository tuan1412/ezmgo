package ezmgo

type logicOp int64

const (
	lopAnd logicOp = iota
	lopOr
)

type builder struct {
	conditionGroups map[logicOp][]*Condition
}

func Filter() *builder {
	b := &builder{}
	b.conditionGroups = make(map[logicOp][]*Condition)
	return b
}

func filterNilConditions(conditions []*Condition) []*Condition {
	notNilConditions := make([]*Condition, 0, len(conditions))

	for _, c := range conditions {
		if c != nil {
			notNilConditions = append(notNilConditions, c)
		}
	}
	return notNilConditions
}


func (b *builder) And(c ...*Condition) *builder {
	b.conditionGroups[lopAnd] = filterNilConditions(c)
	return b
}

func (b *builder) Or(c ...*Condition) *builder {
	notNilConditions := filterNilConditions(c)

	if len(notNilConditions) < 2 {
		panic("OR logical operator require minimum two conditions")
	}
	b.conditionGroups[lopOr] = notNilConditions
	return b
}


func (b *builder) Build() bson.M {
	m := bson.M{}
	for lOp, conditions := range b.conditionGroups {
		switch lOp {
		case lopAnd:
			for _, c := range conditions {
				v, _ := m[c.fieldName]
				if oldM, ok := v.(bson.M); ok {
					m[c.fieldName] = mergeM(oldM, getM(c))
				} else {
					m[c.fieldName] = getM(c)
				}
			}

		case lopOr:
			m["$or"] = getArrayOfM(conditions)
		}
	}
	return m
}
