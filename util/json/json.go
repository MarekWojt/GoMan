package json

import (
	"errors"

	"github.com/valyala/fastjson"
)

// JSON is a wrapper for fastjson value
type JSON struct {
	Value *fastjson.Value
}

// GetSubValue gets the sub value of a fastjson value
func (json *JSON) GetSubValue(keys ...string) (*fastjson.Value, error) {
	var subVal *fastjson.Value
	for _, key := range keys {
		subVal = json.Value.Get(key)
		if subVal == nil {
			return nil, errors.New("does not exist")
		}
	}

	return subVal, nil
}

// GetSubJSON gets a sub json from a json
func (json *JSON) GetSubJSON(keys ...string) (JSON, error) {
	subVal, err := json.GetSubValue(keys...)
	if err != nil {
		return JSON{}, err
	}

	return JSON{subVal}, nil
}

// GetString gets a string from a fastjson value
func (json *JSON) GetString(keys ...string) (string, error) {
	subVal, err := json.GetSubValue(keys...)
	if err != nil {
		return "", err
	}

	switch subVal.Type() {
	case fastjson.TypeString:
		return string(subVal.GetStringBytes()), nil
	}

	return "", errors.New("wrong type")
}

// GetBool gets a bool from a fastjson value
func (json *JSON) GetBool(keys ...string) (bool, error) {
	subVal, err := json.GetSubValue(keys...)
	if err != nil {
		return false, err
	}

	switch subVal.Type() {
	case fastjson.TypeTrue:
	case fastjson.TypeFalse:
		return subVal.GetBool(), nil
	}

	return false, errors.New("wrong type")
}

// GetInt gets a Int from a fastjson value
func (json *JSON) GetInt(keys ...string) (int, error) {
	subVal, err := json.GetSubValue(keys...)
	if err != nil {
		return 0, err
	}

	switch subVal.Type() {
	case fastjson.TypeNumber:
		return subVal.GetInt(), nil
	}

	return 0, errors.New("wrong type")
}

// GetFloat gets a Float from a fastjson value
func (json *JSON) GetFloat(keys ...string) (float64, error) {
	subVal, err := json.GetSubValue(keys...)
	if err != nil {
		return 0, err
	}

	switch subVal.Type() {
	case fastjson.TypeNumber:
		return subVal.GetFloat64(), nil
	}

	return 0, errors.New("wrong type")
}

// IsNull returns a boolean if the value is null
func (json *JSON) IsNull(keys []string) (bool, error) {
	subVal, err := json.GetSubValue(keys...)
	if err != nil {
		// Does not exist
		return false, err
	}

	switch subVal.Type() {
	case fastjson.TypeNumber:
		return true, nil
	}

	return false, nil
}
