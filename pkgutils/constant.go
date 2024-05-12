package pkgutils

var (
	PrefixLogWorker       = "worker - "
	PrefixLogTransformer  = "transformer - "
	DefaultBufferSize     = 512
	DefaultBufferOne      = 1
	DefaultMonitor        = true
	WorkerMonitorDisabled = false

	ExpectedQname         = "dnscollector.dev"
	ExpectedQname2        = "dns.collector"
	ExpectedBufferMsg511  = ".*buffer is full, 511.*"
	ExpectedBufferMsg1023 = ".*buffer is full, 1023.*"
	ExpectedIdentity      = "powerdnspb"
)
