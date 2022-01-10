package fcinfo

import (
	"encoding/xml"
)

type NovaSeqRunParameters struct {
	XMLName                  xml.Name `xml:"RunParameters"`
	Text                     string   `xml:",chardata"`
	Xsd                      string   `xml:"xsd,attr"`
	Xsi                      string   `xml:"xsi,attr"`
	Surface                  string   `xml:"Surface"`
	ReadType                 string   `xml:"ReadType"`
	Side                     string   `xml:"Side"`
	Read1NumberOfCycles      string   `xml:"Read1NumberOfCycles"`
	Read2NumberOfCycles      string   `xml:"Read2NumberOfCycles"`
	IndexRead1NumberOfCycles string   `xml:"IndexRead1NumberOfCycles"`
	IndexRead2NumberOfCycles string   `xml:"IndexRead2NumberOfCycles"`
	PlannedRead1Cycles       string   `xml:"PlannedRead1Cycles"`
	PlannedRead2Cycles       string   `xml:"PlannedRead2Cycles"`
	PlannedIndex1ReadCycles  string   `xml:"PlannedIndex1ReadCycles"`
	PlannedIndex2ReadCycles  string   `xml:"PlannedIndex2ReadCycles"`
	RunNumber                string   `xml:"RunNumber"`
	RtaVersion               string   `xml:"RtaVersion"`
	RecipeVersion            string   `xml:"RecipeVersion"`
	ExperimentName           string   `xml:"ExperimentName"`
	RfidsInfo                struct {
		Text                           string `xml:",chardata"`
		FlowCellSerialBarcode          string `xml:"FlowCellSerialBarcode"`
		FlowCellPartNumber             string `xml:"FlowCellPartNumber"`
		FlowCellLotNumber              string `xml:"FlowCellLotNumber"`
		FlowCellExpirationdate         string `xml:"FlowCellExpirationdate"`
		FlowCellStartDate              string `xml:"FlowCellStartDate"`
		FlowCellNumberOfReuseRemaining string `xml:"FlowCellNumberOfReuseRemaining"`
		FlowCellSupportedModes         string `xml:"FlowCellSupportedModes"`
		FlowCellMode                   string `xml:"FlowCellMode"`
		FlowCellConsumableVersion      string `xml:"FlowCellConsumableVersion"`
		FlowCellRssi                   string `xml:"FlowCellRssi"`
		LibraryTubeSerialBarcode       string `xml:"LibraryTubeSerialBarcode"`
		LibraryTubeSupportedModes      string `xml:"LibraryTubeSupportedModes"`
		LibraryTubePartNumber          string `xml:"LibraryTubePartNumber"`
		LibraryTubeLotNumber           string `xml:"LibraryTubeLotNumber"`
		LibraryTubeExpirationdate      string `xml:"LibraryTubeExpirationdate"`
		LibraryTubeStartDate           string `xml:"LibraryTubeStartDate"`
		LibraryTubeRssi                string `xml:"LibraryTubeRssi"`
		SbsSerialBarcode               string `xml:"SbsSerialBarcode"`
		SbsSupportedModes              string `xml:"SbsSupportedModes"`
		SbsPartNumber                  string `xml:"SbsPartNumber"`
		SbsLotNumber                   string `xml:"SbsLotNumber"`
		SbsExpirationdate              string `xml:"SbsExpirationdate"`
		SbsStartDate                   string `xml:"SbsStartDate"`
		SbsCycleKit                    string `xml:"SbsCycleKit"`
		SbsNumberOfCyclesRemaining     string `xml:"SbsNumberOfCyclesRemaining"`
		SbsNumberOfCyclesSupported     string `xml:"SbsNumberOfCyclesSupported"`
		SbsConsumableVersion           string `xml:"SbsConsumableVersion"`
		SbsRssi                        string `xml:"SbsRssi"`
		ClusterSerialBarcode           string `xml:"ClusterSerialBarcode"`
		ClusterSupportedModes          string `xml:"ClusterSupportedModes"`
		ClusterPartNumber              string `xml:"ClusterPartNumber"`
		ClusterLotNumber               string `xml:"ClusterLotNumber"`
		ClusterExpirationdate          string `xml:"ClusterExpirationdate"`
		ClusterStartDate               string `xml:"ClusterStartDate"`
		ClusterCycleKit                string `xml:"ClusterCycleKit"`
		ClusterNumberOfCyclesRemaining string `xml:"ClusterNumberOfCyclesRemaining"`
		ClusterRssi                    string `xml:"ClusterRssi"`
		BufferSerialBarcode            string `xml:"BufferSerialBarcode"`
		BufferSupportedModes           string `xml:"BufferSupportedModes"`
		BufferPartNumber               string `xml:"BufferPartNumber"`
		BufferLotNumber                string `xml:"BufferLotNumber"`
		BufferExpirationdate           string `xml:"BufferExpirationdate"`
		BufferNumberOfCyclesRemaining  string `xml:"BufferNumberOfCyclesRemaining"`
		BufferStartDate                string `xml:"BufferStartDate"`
		BufferRssi                     string `xml:"BufferRssi"`
	} `xml:"RfidsInfo"`
	RecipeFilePath                   string `xml:"RecipeFilePath"`
	SamplesheetFile                  string `xml:"SamplesheetFile"`
	UsedLimsSetup                    string `xml:"UsedLimsSetup"`
	CeLinuxRunFolder                 string `xml:"CeLinuxRunFolder"`
	RtaRawRunFolder                  string `xml:"RtaRawRunFolder"`
	CeMountRunFolder                 string `xml:"CeMountRunFolder"`
	OutputRunFolder                  string `xml:"OutputRunFolder"`
	OutputRootFolder                 string `xml:"OutputRootFolder"`
	SbcRunFolder                     string `xml:"SbcRunFolder"`
	PreRunFolder                     string `xml:"PreRunFolder"`
	RunStartDate                     string `xml:"RunStartDate"`
	RunId                            string `xml:"RunId"`
	UseBaseSpace                     string `xml:"UseBaseSpace"`
	UseBaseSpaceMonitoringAndStorage string `xml:"UseBaseSpaceMonitoringAndStorage"`
	BaseSpaceRunId                   string `xml:"BaseSpaceRunId"`
	BaseSpaceEnvironment             string `xml:"BaseSpaceEnvironment"`
	BaseSpaceDomain                  string `xml:"BaseSpaceDomain"`
	Autocenter                       string `xml:"Autocenter"`
	BiDirectionalScanning            string `xml:"BiDirectionalScanning"`
	UseCustomRecipe                  string `xml:"UseCustomRecipe"`
	UseCustomRead1Primer             string `xml:"UseCustomRead1Primer"`
	UseCustomRead2Primer             string `xml:"UseCustomRead2Primer"`
	UseCustomIndexRead1Primer        string `xml:"UseCustomIndexRead1Primer"`
	IsRehyb                          string `xml:"IsRehyb"`
	InstrumentName                   string `xml:"InstrumentName"`
	PlatformType                     string `xml:"PlatformType"`
	Application                      string `xml:"Application"`
	ApplicationVersion               string `xml:"ApplicationVersion"`
	Build                            string `xml:"Build"`
	FirmwareVersions                 struct {
		Text            string `xml:",chardata"`
		FirmwareVersion []struct {
			Text    string `xml:",chardata"`
			Board   string `xml:"Board"`
			Version string `xml:"Version"`
		} `xml:"FirmwareVersion"`
	} `xml:"FirmwareVersions"`
	SendIlluminaHealthData string `xml:"SendIlluminaHealthData"`
	UcsRunId               string `xml:"UcsRunId"`
	UcsVersion             string `xml:"UcsVersion"`
	RunSetupMode           string `xml:"RunSetupMode"`
	WorkflowType           string `xml:"WorkflowType"`
}
