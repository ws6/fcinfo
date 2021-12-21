package fcinfo

import (
	"testing"
)

func _TestParseNext2k(t *testing.T) {
	folder := `\\ussd-prd-isi07\ils_metagenomics\devel\200810_VH00257_1_AAAG7F3M5`
	folderInfo, err := CheckFlowcellRunFolder(folder)
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Logf(`%+v`, folderInfo)

	info, err := ParseFlowcellRunFolder(folder, false)
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Logf(`%+v`, info)
}

var runParamXml = `<?xml version="1.0"?>
<RunParameters xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
  <Surface>Both</Surface>
  <ReadType>PairedEnd</ReadType>
  <Side>B</Side>
  <Read1NumberOfCycles>151</Read1NumberOfCycles>
  <Read2NumberOfCycles>151</Read2NumberOfCycles>
  <IndexRead1NumberOfCycles>8</IndexRead1NumberOfCycles>
  <IndexRead2NumberOfCycles>8</IndexRead2NumberOfCycles>
  <PlannedRead1Cycles>151</PlannedRead1Cycles>
  <PlannedRead2Cycles>151</PlannedRead2Cycles>
  <PlannedIndex1ReadCycles>8</PlannedIndex1ReadCycles>
  <PlannedIndex2ReadCycles>8</PlannedIndex2ReadCycles>
  <RunNumber>21</RunNumber>
  <RtaVersion>v3.3.3</RtaVersion>
  <RecipeVersion>1.4.0</RecipeVersion>
  <ExperimentName>Functional3B</ExperimentName>
  <RfidsInfo>
    <FlowCellSerialBarcode>H5NTFDSXX</FlowCellSerialBarcode>
    <FlowCellPartNumber>A</FlowCellPartNumber>
    <FlowCellLotNumber>20264129</FlowCellLotNumber>
    <FlowCellExpirationdate>05/04/2019 00:00:00</FlowCellExpirationdate>
    <FlowCellStartDate>07/16/2018 08:10:00</FlowCellStartDate>
    <FlowCellNumberOfReuseRemaining>1</FlowCellNumberOfReuseRemaining>
    <FlowCellSupportedModes>HTWashOnly;S4</FlowCellSupportedModes>
    <FlowCellMode>S4</FlowCellMode>
    <FlowCellRssi>3</FlowCellRssi>
    <LibraryTubeSerialBarcode>NV0027558-LIB</LibraryTubeSerialBarcode>
    <LibraryTubeSupportedModes>Universal</LibraryTubeSupportedModes>
    <LibraryTubePartNumber />
    <LibraryTubeLotNumber>1712180001</LibraryTubeLotNumber>
    <LibraryTubeExpirationdate>12/31/2169 00:00:00</LibraryTubeExpirationdate>
    <LibraryTubeStartDate>07/16/2018 08:10:00</LibraryTubeStartDate>
    <LibraryTubeRssi>4</LibraryTubeRssi>
    <SbsSerialBarcode>NV2123706-RGSBS</SbsSerialBarcode>
    <SbsSupportedModes>S4</SbsSupportedModes>
    <SbsPartNumber>C</SbsPartNumber>
    <SbsLotNumber>20249426</SbsLotNumber>
    <SbsExpirationdate>03/20/2019 00:00:00</SbsExpirationdate>
    <SbsStartDate>07/16/2018 08:10:00</SbsStartDate>
    <SbsCycleKit>330</SbsCycleKit>
    <SbsNumberOfCyclesRemaining>12</SbsNumberOfCyclesRemaining>
    <SbsNumberOfCyclesSupported>330</SbsNumberOfCyclesSupported>
    <SbsRssi>4</SbsRssi>
    <ClusterSerialBarcode>NV2151437-RGCPE</ClusterSerialBarcode>
    <ClusterSupportedModes>S4</ClusterSupportedModes>
    <ClusterPartNumber>C</ClusterPartNumber>
    <ClusterLotNumber>20263483</ClusterLotNumber>
    <ClusterExpirationdate>05/10/2019 00:00:00</ClusterExpirationdate>
    <ClusterStartDate>07/16/2018 08:10:00</ClusterStartDate>
    <ClusterCycleKit>330</ClusterCycleKit>
    <ClusterNumberOfCyclesRemaining>12</ClusterNumberOfCyclesRemaining>
    <ClusterRssi>5</ClusterRssi>
    <BufferSerialBarcode>NV2115958-BUFFR</BufferSerialBarcode>
    <BufferSupportedModes>S4</BufferSupportedModes>
    <BufferPartNumber>1</BufferPartNumber>
    <BufferLotNumber>20243203</BufferLotNumber>
    <BufferExpirationdate>03/02/2019 00:00:00</BufferExpirationdate>
    <BufferNumberOfCyclesRemaining>12</BufferNumberOfCyclesRemaining>
    <BufferStartDate>07/16/2018 08:10:00</BufferStartDate>
    <BufferRssi>3</BufferRssi>
  </RfidsInfo>
  <RecipeFilePath>C:\Illumina\NovaSeq Control Software\Recipe</RecipeFilePath>
  <SamplesheetFile>\\ukhx-prd-isi01\upload_GEL_dev\Samplesheets\H5NTFDSXX\SampleSheet.csv</SamplesheetFile>
  <UsedLimsSetup>false</UsedLimsSetup>
  <CeLinuxRunFolder>/ilmn/outputfolder/180716_A00494_0021_BH5NTFDSXX/</CeLinuxRunFolder>
  <RtaRawRunFolder>/ilmn/outputfolder</RtaRawRunFolder>
  <CeMountRunFolder>Z:\outputfolder\180716_A00494_0021_BH5NTFDSXX\</CeMountRunFolder>
  <OutputRunFolder />
  <SbcRunFolder>C:\Illumina\NovaSeqTemp\180716_A00494_0021_BH5NTFDSXX\</SbcRunFolder>
  <PreRunFolder>C:\Illumina\NovaSeqTemp\RunSetupLogs\A00494_2018-07-16__08_08_47_SideB</PreRunFolder>
  <RunStartDate>180716</RunStartDate>
  <RunId>180716_A00494_0021_BH5NTFDSXX</RunId>
  <UseBaseSpace>false</UseBaseSpace>
  <UseBaseSpaceMonitoringAndStorage>true</UseBaseSpaceMonitoringAndStorage>
  <BaseSpaceRunId>38038</BaseSpaceRunId>
  <Autocenter>true</Autocenter>
  <BiDirectionalScanning>true</BiDirectionalScanning>
  <UseCustomRecipe>false</UseCustomRecipe>
  <UseCustomRead1Primer>false</UseCustomRead1Primer>
  <UseCustomRead2Primer>false</UseCustomRead2Primer>
  <UseCustomIndexRead1Primer>false</UseCustomIndexRead1Primer>
  <IsRehyb>false</IsRehyb>
  <InstrumentName>A00494</InstrumentName>
  <PlatformType>HighThroughput</PlatformType>
  <Application>NovaSeq Control Software</Application>
  <ApplicationVersion>1.4.0</ApplicationVersion>
  <Build>61</Build>
  <FirmwareVersions>
    <FirmwareVersion>
      <Board>Fluidics Board</Board>
      <Version>novaseq_flu@NovaSeq_1.22.1</Version>
    </FirmwareVersion>
    <FirmwareVersion>
      <Board>Left Buffer Interface Board</Board>
      <Version>novaseq_bim@NovaSeq_1.22.1</Version>
    </FirmwareVersion>
    <FirmwareVersion>
      <Board>Right Buffer Interface Board</Board>
      <Version>novaseq_bim@NovaSeq_1.22.1</Version>
    </FirmwareVersion>
    <FirmwareVersion>
      <Board>Chassis Module Board</Board>
      <Version>novaseq_chm@NovaSeq_1.22.1</Version>
    </FirmwareVersion>
    <FirmwareVersion>
      <Board>Camera Interface Board</Board>
      <Version>novaseq_cib_2@NovaSeq_1.22.1</Version>
    </FirmwareVersion>
    <FirmwareVersion>
      <Board>Focus Interface Board</Board>
      <Version>novaseq_fib_2@NovaSeq_1.22.1</Version>
    </FirmwareVersion>
    <FirmwareVersion>
      <Board>Flow Cell Holder Board</Board>
      <Version>novaseq_fch_2@NovaSeq_1.22.1</Version>
    </FirmwareVersion>
    <FirmwareVersion>
      <Board>Left Reagent Chiller Board</Board>
      <Version>novaseq_rca@NovaSeq_1.22.1</Version>
    </FirmwareVersion>
    <FirmwareVersion>
      <Board>Right Reagent Chiller Board</Board>
      <Version>novaseq_rca@NovaSeq_1.22.1</Version>
    </FirmwareVersion>
    <FirmwareVersion>
      <Board>System Thermal Board</Board>
      <Version>novaseq_syst@NovaSeq_1.22.1</Version>
    </FirmwareVersion>
  </FirmwareVersions>
  <SendIlluminaHealthData>true</SendIlluminaHealthData>
  <UcsRunId>14C92CBF7127D689</UcsRunId>
  <UcsVersion>1.4.2.927</UcsVersion>
  <WorkflowType>NovaSeqXp</WorkflowType>
</RunParameters>`

func TestParseRunParamsXml(t *testing.T) {
	ret, err := ParseRunParamsXML(runParamXml)
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Logf(`%+v`, ret)
}
