package form

import (
	"net/url"
	"strconv"
)

type Form struct {
	url.Values
}

func New(values url.Values) *Form {
	return &Form{Values: values}
}

func (this *Form) GetString(key string) string {
	return this.Values.Get(key)
}

func (this *Form) GetInt(key string, defaultValue int) int {
	i, err := strconv.Atoi(this.Values.Get(key))
	if err == nil {
		return i
	}

	return defaultValue
}

func (this *Form) ToStringMap() map[string]string {
	m := make(map[string]string)
	for key := range this.Values {
		m[key] = this.Values.Get(key)
	}
	return m
}
