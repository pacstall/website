package query

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type Key = string
type Enum = []string
type OptionalEnum struct {
	enum         []string
	defaultValue string
}

type QueryResult struct {
	Strings map[Key]string
	Ints    map[Key]int
}

type Query struct {
	req           *http.Request
	requiredStrs  []Key
	optionalStrs  map[Key]string
	requiredInts  []Key
	optionalInts  map[Key]int
	requiredEnums map[Key]Enum
	optionalEnums map[Key]OptionalEnum
	constraints   map[Key]Key
}

func New(r *http.Request) *Query {
	return &Query{
		req:           r,
		requiredStrs:  make([]string, 0),
		optionalStrs:  make(map[string]string),
		requiredInts:  make([]string, 0),
		optionalInts:  make(map[string]int),
		requiredEnums: make(map[string][]string),
		optionalEnums: make(map[string]OptionalEnum),
		constraints:   make(map[string]string),
	}
}

func (q *Query) RequireStr(name string) *Query {
	q.requiredStrs = append(q.requiredStrs, name)
	return q
}

func (q *Query) OptionalStr(name string, defaultVal string) *Query {
	q.optionalStrs[name] = defaultVal
	return q
}

func (q *Query) RequireInt(name string) *Query {
	q.requiredInts = append(q.requiredStrs, name)
	return q
}

func (q *Query) OptionalInt(name string, defaultVal int) *Query {
	q.optionalInts[name] = defaultVal
	return q
}

func (q *Query) RequireEnum(name string, enum Enum) *Query {
	q.requiredEnums[name] = enum
	return q
}

func (q *Query) OptionalEnum(name string, enum Enum, defaultVal string) *Query {
	q.optionalEnums[name] = OptionalEnum{
		enum:         enum,
		defaultValue: defaultVal,
	}
	return q
}

func (q *Query) IfParamExistsRequireOther(first string, second string) *Query {
	q.constraints[first] = second
	return q
}

func (q *Query) MustComeTogheter(params []Key) *Query {
	for idx, param := range params[:len(params)-1] {
		q.IfParamExistsRequireOther(param, params[idx+1])
	}
	return q
}

func (q *Query) Parse() (*QueryResult, error) {
	stringParams := make(map[Key]string)
	intParams := make(map[Key]int)

	vals := q.req.URL.Query()

	for _, requiredStr := range q.requiredStrs {
		if !vals.Has(requiredStr) {
			return nil, fmt.Errorf("missing required query parameter '%v'", requiredStr)
		}

		stringParams[requiredStr] = vals.Get(requiredStr)
	}

	for _, requiredInt := range q.requiredInts {
		if !vals.Has(requiredInt) {
			return nil, fmt.Errorf("missing required query parameter '%v'", requiredInt)
		}

		value, err := strconv.ParseInt(vals.Get(requiredInt), 10, 32)
		if err != nil {
			return nil, fmt.Errorf("required query parameter '%v' is not int", requiredInt)
		}

		intParams[requiredInt] = int(value)
	}

	for requiredEnum, enumValues := range q.requiredEnums {
		if !vals.Has(requiredEnum) {
			return nil, fmt.Errorf("missing required query parameter '%v'", requiredEnum)
		}

		value := vals.Get(requiredEnum)
		found := false
		for _, enumValue := range enumValues {
			if strings.Compare(enumValue, value) == 0 {
				found = true
				break
			}
		}

		if !found {
			return nil, fmt.Errorf("required query parameter '%v' has value '%v' but expected one of [%v]", requiredEnum, value, enumValues)
		}

		stringParams[requiredEnum] = value
	}

	for optionalStr, defaultValue := range q.optionalStrs {
		value := vals.Get(optionalStr)
		if !vals.Has(optionalStr) {
			value = defaultValue
		}

		stringParams[optionalStr] = value
	}

	for optionalInt, defaultValue := range q.optionalInts {
		value64, _ := strconv.ParseInt(vals.Get(optionalInt), 10, 32)
		value := int(value64)

		if !vals.Has(optionalInt) {
			value = defaultValue
		}

		intParams[optionalInt] = value
	}

	for optionalEnum, details := range q.optionalEnums {
		if !vals.Has(optionalEnum) {
			stringParams[optionalEnum] = details.defaultValue
			continue
		}

		value := vals.Get(optionalEnum)
		for _, enumValue := range details.enum {
			if strings.Compare(enumValue, value) == 0 {
				stringParams[optionalEnum] = value
				continue
			}
		}
	}

	for key, required := range q.constraints {
		_, okStr := stringParams[key]
		_, okInt := intParams[key]

		if !okStr && !okInt {
			continue
		}

		_, okStr = stringParams[required]
		_, okInt = intParams[required]

		if !okStr && !okInt {
			return nil, fmt.Errorf("constraint error: param '%v' exists and requires '%v' but does not exit", key, required)
		}
	}

	return &QueryResult{Strings: stringParams, Ints: intParams}, nil
}
