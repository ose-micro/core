package tracing

type Config struct {
	Endpoint    string  `mapstructure:"endpoint"`
	ServiceName string  `mapstructure:"serviceName"`
	SampleRatio float64 `mapstructure:"sampleRatio"`
}
