package provider

type merakiResponse interface {
	GetPayload() interface{}
}
