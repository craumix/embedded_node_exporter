# Embedded node_exporter

The goal of this project is to make a version of [node_exporter](https://github.com/prometheus/node_exporter) that is easily embeddable into another golang application. While this is considered an "anti-pattern" for prometheus exporters as discussed [here](https://github.com/prometheus/node_exporter/issues/864) there are sometimes still use cases where including the exporter is the better solution.

## Usage

The following example creates and registers a new NodeCollector with some basic collectors and **no** logging for the collectors.

```golang
nc, err := exporter.NewNodeCollector(
    nil,
    "cpu",
    "hwmon",
    "meminfo",
    "stat",
    "diskstats",
    "netstat",
    "time",
    "filesystem",
)
if err != nil {
    log.Panic(err)
}
err = prometheus.Register(nc)
if err != nil {
    log.Panic(err)
}
```

The collectors use <https://github.com/go-kit/log> for logging so you can either use that or write a translation to your own logger. A translator for logrus is already included and may be used as an example for further use cases.

```golang
level := logrus.WarnLevel

nc, err := exporter.NewNodeCollector(
    &exporter.LogrusTranslator{
        LogLevel: &level,
    },
    "cpu",
)
```

The `LogLevel` can be optionally specified, if not set will use ur default log level.
