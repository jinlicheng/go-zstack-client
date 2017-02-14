package zstack

const (
	P2P_EXCHANGE                 = "P2P"
	API_SERVICE_ID               = "zstack.message.api.portal"
	BROADCAST_EXCHANGE           = "BROADCAST"
	QUEUE_PREFIX                 = "zstack.go.message.%s"
	API_EVENT_QUEUE_PREFIX       = "zstck.go.api.event.%s"
	API_EVENT_QUEUE_BINDING_KEY  = "key.event.API.API_EVENT"
	CANONICAL_EVENT_QUEUE_PREFIX = "zstck.go.canonical.event.%s"
	CANONICAL_EVENT_BINDING_KEY  = "key.event.LOCAL.canonicalEvent"
)

type ZStackMQClient struct {
	baseURL string // The base URL of the API
	port    uint16
	user    string
	passwd  string
	timeout int64 // Max waiting timeout in seconds for async jobs to finish; defaults to 300 seconds
}

type Connection struct {
}

type CloudBus struct {
}
