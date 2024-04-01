package utils

import (
	"errors"
	"fmt"
	"net/url"
)

type TransportType string

const (
	Amqp     = "amqp"
	Http     = "http"
	Internal = "default"

	amqpSchema  = "amqp"
	httpSchema  = "http"
	httpsSchema = "https"
)

var (
	ErrUnknownTransport = errors.New("unsupported callback. available: 'amqp', 'http', 'default'")
)

func DetermineTransportType(callback string) (TransportType, error) {
	isInternal := callback == "" || callback == Internal
	isAmqp := false
	isHttp := false
	if !isInternal {
		u, err := url.Parse(callback)
		if err != nil {
			return "", fmt.Errorf("invalid url: %v", err)
		}
		isAmqp = u.Scheme == amqpSchema
		isHttp = u.Scheme == httpSchema || u.Scheme == httpsSchema
	}
	switch {
	case isInternal:
		return Internal, nil
	case isAmqp:
		return Amqp, nil
	case isHttp:
		return Http, nil
	default:
		return "", ErrUnknownTransport
	}
}
