package ezmgo

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type op string

const (
	eqOp     = op("$eq")
	neOp     = op("$ne")
	gtOp     = op("$gt")
	gteOp    = op("$gte")
	ltOp     = op("$lt")
	lteOp    = op("$lte")
	inOp     = op("$in")
	ninOp    = op("$nin")
	existsOp = op("$exists")
	regexOp  = op("$regex")
)

type Condition struct {
	op        op
	value     interface{}
	fieldName string
}

func Eq[T any](fieldName string, p *T) *Condition {
	if p == nil {
		return nil
	}

	return &Condition{op: eqOp, value: *p, fieldName: fieldName}
}

func Ne[T any](fieldName string, p *T) *Condition {
	if p == nil {
		return nil
	}
	return &Condition{op: neOp, value: *p, fieldName: fieldName}
}

func Gt[T any](fieldName string, p *T) *Condition {
	if p == nil {
		return nil
	}
	return &Condition{op: gtOp, value: *p, fieldName: fieldName}
}

func Gte[T any](fieldName string, p *T) *Condition {
	if p == nil {
		return nil
	}
	return &Condition{op: gteOp, value: *p, fieldName: fieldName}
}

func Lt[T any](fieldName string, p *T) *Condition {
	if p == nil {
		return nil
	}
	return &Condition{op: ltOp, value: *p, fieldName: fieldName}
}

func Lte[T any](fieldName string, p *T) *Condition {
	if p == nil {
		return nil
	}
	return &Condition{op: lteOp, value: *p, fieldName: fieldName}
}

func In[T any](fieldName string, p *T) *Condition {
	if p == nil {
		return nil
	}
	return &Condition{op: inOp, value: *p, fieldName: fieldName}
}

func NIn[T any](fieldName string, p *T) *Condition {
	if p == nil {
		return nil
	}
	return &Condition{op: ninOp, value: *p, fieldName: fieldName}
}

func Exist[T any](fieldName string, p *T) *Condition {
	if p == nil {
		return nil
	}
	return &Condition{op: existsOp, value: *p, fieldName: fieldName}
}

func Contains[T any](fieldName string, p *T) *Condition {
	if p == nil {
		return nil
	}
	return &Condition{op: regexOp, value: *p, fieldName: fieldName}
}

func getArrayOfM(cond []*Condition) []bson.M {
	ar := make([]bson.M, 0, len(cond))
	for _, c := range cond {
		ar = append(ar, bson.M{c.fieldName: getM(c)})
	}
	return ar
}

func getM(c *Condition) bson.M {
	if c.op == regexOp {
		m := bson.M{string(c.op): primitive.Regex{Pattern: c.value.(string), Options: "i"}}
		return m
	}
	m := bson.M{string(c.op): c.value}
	return m
}

func mergeM(m1, m2 bson.M) bson.M {
	m := bson.M{}

	for k, v := range m1 {
		m[k] = v
	}

	for k, v := range m2 {
		m[k] = v
	}
	return m
}