package exporter

import (
	_ "unsafe"

	"github.com/go-kit/log"
	"github.com/prometheus/node_exporter/collector"
	"github.com/prometheus/procfs"
)

//go:linkname collectorState github.com/prometheus/node_exporter/collector.collectorState
var collectorState map[string]*bool

//--path.procfs
//
//procfs mountpoint.
//go:linkname ProcPath github.com/prometheus/node_exporter/collector.procPath
var ProcPath *string

//--path.sysfs
//
//sysfs mountpoint.
//go:linkname SysPath github.com/prometheus/node_exporter/collector.sysPath
var SysPath *string

//--path.rootfs
//
//rootfs mountpoint.
//go:linkname RootfsPath github.com/prometheus/node_exporter/collector.rootfsPath
var RootfsPath *string

//--path.udev.data
//
//udev data path.
//go:linkname UdevDataPath github.com/prometheus/node_exporter/collector.udevDataPath
var UdevDataPath *string

//--collector.filesystem.mount-points-exclude
//
//Regexp of mount points to exclude for filesystem collector.
//go:linkname MountPointsExclude github.com/prometheus/node_exporter/collector.mountPointsExclude
var MountPointsExclude *string

//--collector.filesystem.ignored-mount-points
//
//Regexp of mount points to ignore for filesystem collector.
//go:linkname FsTypesExclude github.com/prometheus/node_exporter/collector.fsTypesExclude
var FsTypesExclude *string

//--collector.netdev.device-exclude
//
//Regexp of net devices to exclude (mutually exclusive to device-include).
//go:linkame NetdevDeviceExclude github.com/prometheus/node_exporter/collector.netdevDeviceExclude
var NetdevDeviceExclude *string

func init() {
	for name := range collectorState {
		collectorState[name] = ref(true)
	}

	ProcPath = ref(procfs.DefaultMountPoint)
	SysPath = ref("/sys")
	RootfsPath = ref("/")
	UdevDataPath = ref("/run/udev/data")
	MountPointsExclude = ref("^/(dev|proc|run/credentials/.+|sys|var/lib/docker/.+|var/lib/containers/storage/.+)($|/)")
	FsTypesExclude = ref("^(autofs|binfmt_misc|bpf|cgroup2?|configfs|debugfs|devpts|devtmpfs|fusectl|hugetlbfs|iso9660|mqueue|nsfs|overlay|proc|procfs|pstore|rpc_pipefs|securityfs|selinuxfs|squashfs|sysfs|tracefs)$")
	NetdevDeviceExclude = ref("")
}

func NewNodeCollector(logger log.Logger, collectors ...string) (*collector.NodeCollector, error) {
	if logger == nil {
		logger = log.NewNopLogger()
	}

	nc, err := collector.NewNodeCollector(logger, collectors...)
	if err != nil {
		return nil, err
	}

	//fmt.Printf("Collectors: %v\n", maps.Keys(nc.Collectors))

	return nc, nil
}

func ref[T any](v T) *T {
	return &v
}
