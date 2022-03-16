package constant

const (
	// NacosDefaultGroup
	NacosDefaultGroup = "SEATA_GROUP"
	// NacosDefaultDataID
	NacosDefaultDataID = "starfish"
	// NacosKey
	NacosKey = "nacos"
	// FileKey
	FileKey = "file"

	Etcdv3Key                = "etcdv3"
	Etcdv3RegistryPrefix     = "etcdv3-starfish-" // according to starfish java version
	Etcdv3LeaseRenewInterval = 5               // according to starfish java version
	Etcdv3LeaseTtl           = 10              // according to starfish java version
	Etcdv3LeaseTtlCritical   = 6               // according to starfish java version
)
