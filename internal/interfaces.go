package internal

import "net/url"

type RequestBody interface {
	Values() url.Values
	Encode() string
}
