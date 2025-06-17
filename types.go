package main

type Controllers struct {
	Controller []Controller `json:"Controllers"`
}

type Controller struct {
	CommandStatus CommandStatus `json:"Command Status"`
	ResponseData ResponseData `json:"Response Data"`
}

type CommandStatus struct {
	CLIVersion      string `json:"CLI Version"`
	OperatingSystem string `json:"Operating system"`
	Controller      int    `json:"Controller"`
	Status          string `json:"Status"`
	Description     string `json:"Description"`
}

type ResponseData struct {
	Basics Basics `json:"Basics"`
	Version Version `json:"Version"`
	Bus Bus `json:"Bus"`
	PendingImagesInFlash PendingImagesInFlash `json:"Pending Images in Flash"`
	Status Status `json:"Status"`
	SupportedAdapterOperations SupportedAdapterOperations `json:"Supported Adapter Operations"`
	SupportedPDOperations SupportedPDOperations `json:"Supported PD Operations"`
	SupportedVDOperations SupportedVDOperations `json:"Supported VD Operations"`
	HwCfg HwCfg `json:"HwCfg"`
	Policies Policies `json:"Policies"`
	Boot Boot `json:"Boot"`
	HighAvailability HighAvailability `json:"High Availability"`
	Defaults Defaults `json:"Defaults"`
	Capabilities Capabilities `json:"Capabilities"`
	ScheduledTasks ScheduledTasks `json:"Scheduled Tasks"`
	SecurityProtocolProperties SecurityProtocolProperties `json:"Security Protocol properties"`
	DriveGroups int `json:"Drive Groups"`
	TOPOLOGY []TOPOLOGY `json:"TOPOLOGY"`
	VirtualDrives int `json:"Virtual Drives"`
	VDLIST []VDLIST `json:"VD LIST"`
	PhysicalDrives int `json:"Physical Drives"`
	PDLIST []PDLIST `json:"PD LIST"`
	Enclosures    int `json:"Enclosures"`
	EnclosureLIST []EnclosureLIST `json:"Enclosure LIST"`
	BBUInfo []BBUInfo `json:"BBU_Info"`
}

type Basics struct {
	Controller                int    `json:"Controller"`
	Model                     string `json:"Model"`
	SerialNumber              string `json:"Serial Number"`
	CurrentControllerDateTime string `json:"Current Controller Date/Time"`
	CurrentSystemDateTime     string `json:"Current System Date/time"`
	SASAddress                string `json:"SAS Address"`
	PCIAddress                string `json:"PCI Address"`
	MfgDate                   string `json:"Mfg Date"`
	ReworkDate                string `json:"Rework Date"`
	RevisionNo                string `json:"Revision No"`
}

type Version struct {
	FirmwarePackageBuild string `json:"Firmware Package Build"`
	FirmwareVersion      string `json:"Firmware Version"`
	BiosVersion          string `json:"Bios Version"`
	WebBIOSVersion       string `json:"WebBIOS Version"`
	CtrlRVersion         string `json:"Ctrl-R Version"`
	PrebootCLIVersion    string `json:"Preboot CLI Version"`
	BootBlockVersion     string `json:"Boot Block Version"`
	DriverName           string `json:"Driver Name"`
	DriverVersion        string `json:"Driver Version"`
}

type Bus struct {
	VendorID        int    `json:"Vendor Id"`
	DeviceID        int    `json:"Device Id"`
	SubVendorID     int    `json:"SubVendor Id"`
	SubDeviceID     int    `json:"SubDevice Id"`
	HostInterface   string `json:"Host Interface"`
	DeviceInterface string `json:"Device Interface"`
	BusNumber       int    `json:"Bus Number"`
	DeviceNumber    int    `json:"Device Number"`
	FunctionNumber  int    `json:"Function Number"`
	DomainID        int    `json:"Domain ID"`
}

type PendingImagesInFlash struct {
	ImageName string `json:"Image name"`
}

type Status struct {
	ControllerStatus                                    string `json:"Controller Status"`
	MemoryCorrectableErrors                             int    `json:"Memory Correctable Errors"`
	MemoryUncorrectableErrors                           int    `json:"Memory Uncorrectable Errors"`
	ECCBucketCount                                      int    `json:"ECC Bucket Count"`
	AnyOfflineVDCachePreserved                          string `json:"Any Offline VD Cache Preserved"`
	BBUStatus                                           int    `json:"BBU Status"`
	PDFirmwareDownloadInProgress                        string `json:"PD Firmware Download in progress"`
	SupportPDFirmwareDownload                           string `json:"Support PD Firmware Download"`
	LockKeyAssigned                                     string `json:"Lock Key Assigned"`
	FailedToGetLockKeyOnBootup                          string `json:"Failed to get lock key on bootup"`
	LockKeyHasNotBeenBackedUp                           string `json:"Lock key has not been backed up"`
	BiosWasNotDetectedDuringBoot                        string `json:"Bios was not detected during boot"`
	ControllerMustBeRebootedToCompleteSecurityOperation string `json:"Controller must be rebooted to complete security operation"`
	ARollbackOperationIsInProgress                      string `json:"A rollback operation is in progress"`
	AtLeastOnePFKExistsInNVRAM                          string `json:"At least one PFK exists in NVRAM"`
	SSCPolicyIsWB                                       string `json:"SSC Policy is WB"`
	ControllerHasBootedIntoSafeMode                     string `json:"Controller has booted into safe mode"`
	ControllerShutdownRequired                          string `json:"Controller shutdown required"`
	ControllerHasBootedIntoCertificateProvisionMode     string `json:"Controller has booted into certificate provision mode"`
}

type SupportedAdapterOperations struct {
	RebuildRate                               string `json:"Rebuild Rate"`
	CCRate                                    string `json:"CC Rate"`
	BGIRate                                   string `json:"BGI Rate "`
	ReconstructRate                           string `json:"Reconstruct Rate"`
	PatrolReadRate                            string `json:"Patrol Read Rate"`
	AlarmControl                              string `json:"Alarm Control"`
	ClusterSupport                            string `json:"Cluster Support"`
	BBU                                       string `json:"BBU"`
	Spanning                                  string `json:"Spanning"`
	DedicatedHotSpare                         string `json:"Dedicated Hot Spare"`
	RevertibleHotSpares                       string `json:"Revertible Hot Spares"`
	ForeignConfigImport                       string `json:"Foreign Config Import"`
	SelfDiagnostic                            string `json:"Self Diagnostic"`
	AllowMixedRedundancyOnArray               string `json:"Allow Mixed Redundancy on Array"`
	GlobalHotSpares                           string `json:"Global Hot Spares"`
	DenySCSIPassthrough                       string `json:"Deny SCSI Passthrough"`
	DenySMPPassthrough                        string `json:"Deny SMP Passthrough"`
	DenySTPPassthrough                        string `json:"Deny STP Passthrough"`
	SupportMoreThan8Phys                      string `json:"Support more than 8 Phys"`
	FWAndEventTimeInGMT                       string `json:"FW and Event Time in GMT"`
	SupportEnhancedForeignImport              string `json:"Support Enhanced Foreign Import"`
	SupportEnclosureEnumeration               string `json:"Support Enclosure Enumeration"`
	SupportAllowedOperations                  string `json:"Support Allowed Operations"`
	AbortCCOnError                            string `json:"Abort CC on Error"`
	SupportMultipath                          string `json:"Support Multipath"`
	SupportOddEvenDriveCountInRAID1E          string `json:"Support Odd & Even Drive count in RAID1E"`
	SupportSecurity                           string `json:"Support Security"`
	SupportConfigPageModel                    string `json:"Support Config Page Model"`
	SupportTheOCEWithoutAddingDrives          string `json:"Support the OCE without adding drives"`
	SupportEKM                                string `json:"Support EKM"`
	SnapshotEnabled                           string `json:"Snapshot Enabled"`
	SupportPFK                                string `json:"Support PFK"`
	SupportPI                                 string `json:"Support PI"`
	SupportLdBBMInfo                          string `json:"Support Ld BBM Info"`
	SupportShieldState                        string `json:"Support Shield State"`
	BlockSSDWriteDiskCacheChange              string `json:"Block SSD Write Disk Cache Change"`
	SupportSuspendResumeBGOps                 string `json:"Support Suspend Resume BG ops"`
	SupportEmergencySpares                    string `json:"Support Emergency Spares"`
	SupportSetLinkSpeed                       string `json:"Support Set Link Speed"`
	SupportBootTimePFKChange                  string `json:"Support Boot Time PFK Change"`
	SupportSystemPD                           string `json:"Support SystemPD"`
	DisableOnlinePFKChange                    string `json:"Disable Online PFK Change"`
	SupportPerfTuning                         string `json:"Support Perf Tuning"`
	SupportSSDPatrolRead                      string `json:"Support SSD PatrolRead"`
	RealTimeScheduler                         string `json:"Real Time Scheduler"`
	SupportResetNow                           string `json:"Support Reset Now"`
	SupportEmulatedDrives                     string `json:"Support Emulated Drives"`
	HeadlessMode                              string `json:"Headless Mode"`
	DedicatedHotSparesLimited                 string `json:"Dedicated HotSpares Limited"`
	PointInTimeProgress                       string `json:"Point In Time Progress"`
	ExtendedLD                                string `json:"Extended LD"`
	SupportUnevenSpan                         string `json:"Support Uneven span "`
	SupportConfigAutoBalance                  string `json:"Support Config Auto Balance"`
	SupportMaintenanceMode                    string `json:"Support Maintenance Mode"`
	SupportDiagnosticResults                  string `json:"Support Diagnostic results"`
	SupportExtEnclosure                       string `json:"Support Ext Enclosure"`
	SupportSesmonitoring                      string `json:"Support Sesmonitoring"`
	SupportSecurityonJBOD                     string `json:"Support SecurityonJBOD"`
	SupportForceFlash                         string `json:"Support ForceFlash"`
	SupportDisableImmediateIO                 string `json:"Support DisableImmediateIO"`
	SupportLargeIOSupport                     string `json:"Support LargeIOSupport"`
	SupportDrvActivityLEDSetting              string `json:"Support DrvActivityLEDSetting"`
	SupportFlushWriteVerify                   string `json:"Support FlushWriteVerify"`
	SupportCPLDUpdate                         string `json:"Support CPLDUpdate"`
	SupportForceTo512E                        string `json:"Support ForceTo512e"`
	SupportDiscardCacheDuringLDDelete         string `json:"Support discardCacheDuringLDDelete"`
	SupportJBODWriteCache                     string `json:"Support JBOD Write cache"`
	SupportLargeQDSupport                     string `json:"Support Large QD Support"`
	SupportCtrlInfoExtended                   string `json:"Support Ctrl Info Extended"`
	SupportIButtonLess                        string `json:"Support IButton less"`
	SupportAESEncryptionAlgorithm             string `json:"Support AES Encryption Algorithm"`
	SupportEncryptedMFC                       string `json:"Support Encrypted MFC"`
	SupportSnapdump                           string `json:"Support Snapdump"`
	SupportForcePersonalityChange             string `json:"Support Force Personality Change"`
	SupportDualFwImage                        string `json:"Support Dual Fw Image"`
	SupportPSOCUpdate                         string `json:"Support PSOC Update"`
	SupportSecureBoot                         string `json:"Support Secure Boot"`
	SupportDebugQueue                         string `json:"Support Debug Queue"`
	SupportLeastLatencyMode                   string `json:"Support Least Latency Mode"`
	SupportOnDemandSnapdump                   string `json:"Support OnDemand Snapdump"`
	SupportClearSnapdump                      string `json:"Support Clear Snapdump"`
	SupportPHYCurrentSpeed                    string `json:"Support PHY current speed"`
	SupportLaneCurrentSpeed                   string `json:"Support Lane current speed"`
	SupportNVMeWidth                          string `json:"Support NVMe Width"`
	SupportLaneDeviceType                     string `json:"Support Lane DeviceType"`
	SupportExtendedDrivePerformanceMonitoring string `json:"Support Extended Drive performance Monitoring"`
	SupportNVMeRepair                         string `json:"Support NVMe Repair"`
	SupportPlatformSecurity                   string `json:"Support Platform Security"`
	SupportNoneModeParams                     string `json:"Support None Mode Params"`
	SupportExtendedControllerProperty         string `json:"Support Extended Controller Property"`
	SupportSmartPollIntervalForDirectAttached string `json:"Support Smart Poll Interval for DirectAttached"`
	SupportWriteJournalPinning                string `json:"Support Write Journal Pinning"`
	SupportSMPPassthruWithPortNumber          string `json:"Support SMP Passthru with Port Number"`
	SupportSnapDumpPrebootTraceBufferToggle   string `json:"Support SnapDump Preboot Trace Buffer Toggle"`
	SupportParityReadCacheBypass              string `json:"Support Parity Read Cache Bypass"`
	SupportNVMeInitErrorDeviceConnectorIndex  string `json:"Support NVMe Init Error Device ConnectorIndex"`
	SupportVolatileKey                        string `json:"Support VolatileKey"`
	SupportPSOCPartInformation                string `json:"Support PSOC Part Information"`
}

type SupportedPDOperations struct {
	ForceOnline             string `json:"Force Online"`
	ForceOffline            string `json:"Force Offline"`
	ForceRebuild            string `json:"Force Rebuild"`
	DenyForceFailed         string `json:"Deny Force Failed"`
	DenyForceGoodBad        string `json:"Deny Force Good/Bad"`
	DenyMissingReplace      string `json:"Deny Missing Replace"`
	DenyClear               string `json:"Deny Clear"`
	DenyLocate              string `json:"Deny Locate"`
	SupportPowerState       string `json:"Support Power State"`
	SetPowerStateForCfg     string `json:"Set Power State For Cfg"`
	SupportT10PowerState    string `json:"Support T10 Power State"`
	SupportTemperature      string `json:"Support Temperature"`
	NCQ                     string `json:"NCQ"`
	SupportMaxRateSATA      string `json:"Support Max Rate SATA"`
	SupportDegradedMedia    string `json:"Support Degraded Media"`
	SupportParallelFWUpdate string `json:"Support Parallel FW Update"`
	SupportDriveCryptoErase string `json:"Support Drive Crypto Erase"`
	SupportSSDWearGauge     string `json:"Support SSD Wear Gauge"`
	SupportSanitize         string `json:"Support Sanitize"`
	SupportExtendedSanitize string `json:"Support Extended Sanitize"`
}

type SupportedVDOperations struct {
	ReadPolicy                          string `json:"Read Policy"`
	WritePolicy                         string `json:"Write Policy"`
	IOPolicy                            string `json:"IO Policy"`
	AccessPolicy                        string `json:"Access Policy"`
	DiskCachePolicy                     string `json:"Disk Cache Policy"`
	Reconstruction                      string `json:"Reconstruction"`
	DenyLocate                          string `json:"Deny Locate"`
	DenyCC                              string `json:"Deny CC"`
	AllowCtrlEncryption                 string `json:"Allow Ctrl Encryption"`
	EnableLDBBM                         string `json:"Enable LDBBM"`
	SupportFastPath                     string `json:"Support FastPath"`
	PerformanceMetrics                  string `json:"Performance Metrics"`
	PowerSavings                        string `json:"Power Savings"`
	SupportPowersaveMaxWithCache        string `json:"Support Powersave Max With Cache"`
	SupportBreakmirror                  string `json:"Support Breakmirror"`
	SupportSSCWriteBack                 string `json:"Support SSC WriteBack"`
	SupportSSCAssociation               string `json:"Support SSC Association"`
	SupportVDHide                       string `json:"Support VD Hide"`
	SupportVDCachebypass                string `json:"Support VD Cachebypass"`
	SupportVDDiscardCacheDuringLDDelete string `json:"Support VD discardCacheDuringLDDelete"`
	SupportVDScsiUnmap                  string `json:"Support VD Scsi Unmap"`
}

type HwCfg struct {
	ChipRevision                   string `json:"ChipRevision"`
	BatteryFRU                     string `json:"BatteryFRU"`
	FrontEndPortCount              int    `json:"Front End Port Count"`
	BackendPortCount               int    `json:"Backend Port Count"`
	BBU                            string `json:"BBU"`
	Alarm                          string `json:"Alarm"`
	SerialDebugger                 string `json:"Serial Debugger"`
	NVRAMSize                      string `json:"NVRAM Size"`
	FlashSize                      string `json:"Flash Size"`
	OnBoardMemorySize              string `json:"On Board Memory Size"`
	CacheVaultFlashSize            string `json:"CacheVault Flash Size"`
	TPM                            string `json:"TPM"`
	UpgradeKey                     string `json:"Upgrade Key"`
	OnBoardExpander                string `json:"On Board Expander"`
	TemperatureSensorForROC        string `json:"Temperature Sensor for ROC"`
	TemperatureSensorForController string `json:"Temperature Sensor for Controller"`
	UpgradableCPLD                 string `json:"Upgradable CPLD"`
	UpgradablePSOC                 string `json:"Upgradable PSOC"`
	CurrentSizeOfCacheCadeGB       int    `json:"Current Size of CacheCade (GB)"`
	CurrentSizeOfFWCacheMB         int    `json:"Current Size of FW Cache (MB)"`
}

type Policies struct {
	PoliciesTable []PolicyTable `json:"Policies Table"`
	FlushTimeDefault             string `json:"Flush Time(Default)"`
	DriveCoercionMode            string `json:"Drive Coercion Mode"`
	AutoRebuild                  string `json:"Auto Rebuild"`
	BatteryWarning               string `json:"Battery Warning"`
	ECCBucketSize                int    `json:"ECC Bucket Size"`
	ECCBucketLeakRateHrs         int    `json:"ECC Bucket Leak Rate (hrs)"`
	RestoreHotSpareOnInsertion   string `json:"Restore Hot Spare on Insertion"`
	ExposeEnclosureDevices       string `json:"Expose Enclosure Devices"`
	MaintainPDFailHistory        string `json:"Maintain PD Fail History"`
	ReorderHostRequests          string `json:"Reorder Host Requests"`
	AutoDetectBackPlane          string `json:"Auto detect BackPlane"`
	LoadBalanceMode              string `json:"Load Balance Mode"`
	SecurityKeyAssigned          string `json:"Security Key Assigned"`
	DisableOnlineControllerReset string `json:"Disable Online Controller Reset"`
	UseDriveActivityForLocate    string `json:"Use drive activity for locate"`
			
}

type PolicyTable struct {
	Policy  string `json:"Policy"`
	Current string `json:"Current"`
	Default string `json:"Default"`
}

type Boot struct {
	BIOSEnumerateVDs                                  int    `json:"BIOS Enumerate VDs"`
	StopBIOSOnError                                   string `json:"Stop BIOS on Error"`
	DelayDuringPOST                                   int    `json:"Delay during POST"`
	SpinDownMode                                      string `json:"Spin Down Mode"`
	EnableCtrlR                                       string `json:"Enable Ctrl-R"`
	EnableWebBIOS                                     string `json:"Enable Web BIOS"`
	EnablePreBootCLI                                  string `json:"Enable PreBoot CLI"`
	EnableBIOS                                        string `json:"Enable BIOS"`
	MaxDrivesToSpinupAtOneTime                        int    `json:"Max Drives to Spinup at One Time"`
	MaximumNumberOfDirectAttachedDrivesToSpinUpIn1Min int    `json:"Maximum number of direct attached drives to spin up in 1 min"`
	DelayAmongSpinupGroupsSec                         int    `json:"Delay Among Spinup Groups (sec)"`
	AllowBootWithPreservedCache                       string `json:"Allow Boot with Preserved Cache"`
}

type HighAvailability struct {
	TopologyType     string `json:"Topology Type"`
	ClusterPermitted string `json:"Cluster Permitted"`
	ClusterActive    string `json:"Cluster Active"`
}

type Defaults struct {
	PhyPolarity                   int    `json:"Phy Polarity"`
	PhyPolaritySplit              int    `json:"Phy PolaritySplit"`
	StripSize                     string `json:"Strip Size"`
	WritePolicy                   string `json:"Write Policy"`
	ReadPolicy                    string `json:"Read Policy"`
	CacheWhenBBUBad               string `json:"Cache When BBU Bad"`
	CachedIO                      string `json:"Cached IO"`
	VDPowerSavePolicy             string `json:"VD PowerSave Policy"`
	DefaultSpinDownTimeMins       int    `json:"Default spin down time (mins)"`
	CoercionMode                  string `json:"Coercion Mode"`
	ZCRConfig                     string `json:"ZCR Config"`
	MaxChainedEnclosures          int    `json:"Max Chained Enclosures"`
	DirectPDMapping               string `json:"Direct PD Mapping"`
	RestoreHotSpareOnInsertion    string `json:"Restore Hot Spare on Insertion"`
	ExposeEnclosureDevices        string `json:"Expose Enclosure Devices"`
	MaintainPDFailHistory         string `json:"Maintain PD Fail History"`
	ZeroBasedEnclosureEnumeration string `json:"Zero Based Enclosure Enumeration"`
	DisablePuncturing             string `json:"Disable Puncturing"`
	EnableLDBBM                   string `json:"EnableLDBBM"`
	DisableHII                    string `json:"DisableHII"`
	UnCertifiedHardDiskDrives     string `json:"Un-Certified Hard Disk Drives"`
	SMARTMode                     string `json:"SMART Mode"`
	EnableLEDHeader               string `json:"Enable LED Header"`
	LEDShowDriveActivity          string `json:"LED Show Drive Activity"`
	DirtyLEDShowsDriveActivity    string `json:"Dirty LED Shows Drive Activity"`
	EnableCrashDump               string `json:"EnableCrashDump"`
	DisableOnlineControllerReset  string `json:"Disable Online Controller Reset"`
	TreatSingleSpanR1EAsR10       string `json:"Treat Single span R1E as R10"`
	PowerSavingOption             string `json:"Power Saving option"`
	TTYLogInFlash                 string `json:"TTY Log In Flash"`
	AutoEnhancedImport            string `json:"Auto Enhanced Import"`
	BreakMirrorRAIDSupport        string `json:"BreakMirror RAID Support"`
	DisableJoinMirror             string `json:"Disable Join Mirror"`
	EnableShieldState             string `json:"Enable Shield State"`
	TimeTakenToDetectCME          string `json:"Time taken to detect CME"`
}

type Capabilities struct {
	SupportedDrives                        string `json:"Supported Drives"`
	RAIDLevelSupported                     string `json:"RAID Level Supported"`
	EnableSystemPD                         string `json:"Enable SystemPD"`
	MixInEnclosure                         string `json:"Mix in Enclosure"`
	MixOfSASSATAOfHDDTypeInVD              string `json:"Mix of SAS/SATA of HDD type in VD"`
	MixOfSASSATAOfSSDTypeInVD              string `json:"Mix of SAS/SATA of SSD type in VD"`
	MixOfSSDHDDInVD                        string `json:"Mix of SSD/HDD in VD"`
	SASDisable                             string `json:"SAS Disable"`
	MaxArmsPerVD                           int    `json:"Max Arms Per VD"`
	MaxSpansPerVD                          int    `json:"Max Spans Per VD"`
	MaxArrays                              int    `json:"Max Arrays"`
	MaxVDPerArray                          int    `json:"Max VD per array"`
	MaxNumberOfVDs                         int    `json:"Max Number of VDs"`
	MaxParallelCommands                    int    `json:"Max Parallel Commands"`
	MaxSGECount                            int    `json:"Max SGE Count"`
	MaxDataTransferSize                    string `json:"Max Data Transfer Size"`
	MaxStripsPerIO                         int    `json:"Max Strips PerIO"`
	MaxConfigurableCacheCadeSizeGB         int    `json:"Max Configurable CacheCade Size(GB)"`
	MaxTransportableDGs                    int    `json:"Max Transportable DGs"`
	EnableSnapdump                         string `json:"Enable Snapdump"`
	EnableSCSIUnmap                        string `json:"Enable SCSI Unmap"`
	ReadCacheBypassEnabledForParityRAIDLDs string `json:"Read cache bypass enabled for Parity RAID LDs"`
	FDEDriveMixSupport                     string `json:"FDE Drive Mix Support"`
	MinStripSize                           string `json:"Min Strip Size"`
	MaxStripSize                           string `json:"Max Strip Size"`
}

type ScheduledTasks struct {
	ConsistencyCheckReoccurrence string `json:"Consistency Check Reoccurrence"`
	NextConsistencyCheckLaunch   string `json:"Next Consistency check launch"`
	PatrolReadReoccurrence       string `json:"Patrol Read Reoccurrence"`
	NextPatrolReadLaunch         string `json:"Next Patrol Read launch"`
	BatteryLearnReoccurrence     string `json:"Battery learn Reoccurrence"`
	OEMID                        string `json:"OEMID"`
}

type SecurityProtocolProperties struct {
	SecurityProtocol string `json:"Security Protocol"`
}

type TOPOLOGY struct {
	DG      interface{}   `json:"DG"`
	Arr     interface{} `json:"Arr"`
	Row     interface{} `json:"Row"`
	EIDSlot interface{} `json:"EID:Slot"`
	DID     interface{} `json:"DID"`
	Type    interface{} `json:"Type"`
	State   interface{} `json:"State"`
	BT      interface{} `json:"BT"`
	Size    interface{} `json:"Size"`
	PDC     interface{} `json:"PDC"`
	PI      interface{} `json:"PI"`
	SED     interface{} `json:"SED"`
	DS3     interface{} `json:"DS3"`
	FSpace  interface{} `json:"FSpace"`
	TR      interface{} `json:"TR"`
}

type VDLIST struct {
	DGVD    string `json:"DG/VD"`
	TYPE    string `json:"TYPE"`
	State   string `json:"State"`
	Access  string `json:"Access"`
	Consist string `json:"Consist"`
	Cache   string `json:"Cache"`
	Cac     string `json:"Cac"`
	SCC     string `json:"sCC"`
	Size    string `json:"Size"`
	Name    string `json:"Name"`
}

type PDLIST struct {
	EIDSlt string `json:"EID:Slt"`
	DID    int    `json:"DID"`
	State  string `json:"State"`
	DG     int    `json:"DG"`
	Size   string `json:"Size"`
	Intf   string `json:"Intf"`
	Med    string `json:"Med"`
	SED    string `json:"SED"`
	PI     string `json:"PI"`
	SeSz   string `json:"SeSz"`
	Model  string `json:"Model"`
	Sp     string `json:"Sp"`
	Type   string `json:"Type"`
}

type EnclosureLIST struct {
	EID            int    `json:"EID"`
	State          string `json:"State"`
	Slots          int    `json:"Slots"`
	PD             int    `json:"PD"`
	PS             int    `json:"PS"`
	Fans           int    `json:"Fans"`
	TSs            int    `json:"TSs"`
	Alms           int    `json:"Alms"`
	SIM            int    `json:"SIM"`
	Port           string `json:"Port#"`
	ProdID         string `json:"ProdID"`
	VendorSpecific string `json:"VendorSpecific"`
}

type BBUInfo struct {
	Model         string `json:"Model"`
	State         string `json:"State"`
	RetentionTime string `json:"RetentionTime"`
	Temp          string `json:"Temp"`
	Mode          string `json:"Mode"`
	MfgDate       string `json:"MfgDate"`
}
