package testmanagerd

import (
	"github.com/danielpaulus/go-ios/ios"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Helper function to create mock data and parse the .xctestrun file
func createAndParseXCTestRunFile(t *testing.T) xCTestRunData {
	// Arrange: Create a temporary .xctestrun file with mock data
	tempFile, err := os.CreateTemp("", "testfile*.xctestrun")
	assert.NoError(t, err, "Failed to create temp file")
	defer os.Remove(tempFile.Name()) // Cleanup after test

	//language=xml
	xcTestRunFileFormatVersion1 := `
		<?xml version="1.0" encoding="UTF-8"?>
		<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
		<plist version="1.0">
			<dict>
				<key>RunnerTests</key>
				<dict>
					<key>BlueprintName</key>
					<string>RunnerTests</string>
					<key>BlueprintProviderName</key>
					<string>Runner</string>
					<key>BlueprintProviderRelativePath</key>
					<string>Runner.xcodeproj</string>
					<key>BundleIdentifiersForCrashReportEmphasis</key>
					<array>
						<string>com.example.myApp</string>
						<string>com.example.myApp.RunnerTests</string>
					</array>
					<key>CommandLineArguments</key>
					<array/>
					<key>DefaultTestExecutionTimeAllowance</key>
					<integer>600</integer>
					<key>DependentProductPaths</key>
					<array>
						<string>__TESTROOT__/Release-iphoneos/Runner.app</string>
						<string>__TESTROOT__/Release-iphoneos/Runner.app/PlugIns/RunnerTests.xctest</string>
					</array>
					<key>DiagnosticCollectionPolicy</key>
					<integer>1</integer>
					<key>EnvironmentVariables</key>
					<dict>
						<key>APP_DISTRIBUTOR_ID_OVERRIDE</key>
						<string>com.apple.AppStore</string>
						<key>OS_ACTIVITY_DT_MODE</key>
						<string>YES</string>
						<key>SQLITE_ENABLE_THREAD_ASSERTIONS</key>
						<string>1</string>
						<key>TERM</key>
						<string>dumb</string>
					</dict>
					<key>IsAppHostedTestBundle</key>
					<true/>
					<key>ParallelizationEnabled</key>
					<true/>
					<key>PreferredScreenCaptureFormat</key>
					<string>screenRecording</string>
					<key>ProductModuleName</key>
					<string>RunnerTests</string>
					<key>RunOrder</key>
					<integer>0</integer>
					<key>SystemAttachmentLifetime</key>
					<string>deleteOnSuccess</string>
					<key>TestBundlePath</key>
					<string>__TESTHOST__/PlugIns/RunnerTests.xctest</string>
					<key>TestHostBundleIdentifier</key>
					<string>com.example.myApp</string>
					<key>TestHostPath</key>
					<string>__TESTROOT__/Release-iphoneos/Runner.app</string>
					<key>TestLanguage</key>
					<string></string>
					<key>TestRegion</key>
					<string></string>
					<key>TestTimeoutsEnabled</key>
					<false/>
					<key>TestingEnvironmentVariables</key>
					<dict>
						<key>DYLD_INSERT_LIBRARIES</key>
						<string>__TESTHOST__/Frameworks/libXCTestBundleInject.dylib</string>
						<key>XCInjectBundleInto</key>
						<string>unused</string>
						<key>Test</key>
						<string>xyz</string>
					</dict>
					<key>ToolchainsSettingValue</key>
					<array/>
					<key>UserAttachmentLifetime</key>
					<string>deleteOnSuccess</string>
					<key>OnlyTestIdentifiers</key>
					<array>
						<string>TestClass1/testMethod1</string>
						<string>TestClass2/testMethod1</string>
					</array>
					<key>SkipTestIdentifiers</key>
					<array>
						<string>TestClass1/testMethod2</string>
						<string>TestClass2/testMethod2</string>
					</array>
					<key>IsUITestBundle</key>
					<true/>
				</dict>
				<key>__xctestrun_metadata__</key>
				<dict>
					<key>ContainerInfo</key>
					<dict>
						<key>ContainerName</key>
						<string>Runner</string>
						<key>SchemeName</key>
						<string>Runner</string>
					</dict>
					<key>FormatVersion</key>
					<integer>1</integer>
				</dict>
			</dict>
		</plist>
	`
	_, err = tempFile.WriteString(xcTestRunFileFormatVersion1)
	assert.NoError(t, err, "Failed to write mock data to temp file")
	tempFile.Close()

	// Act: Use the codec to parse the temp file
	xcTestRunData, err := parseFile(tempFile.Name())

	// Assert: Verify the parsed data
	assert.NoError(t, err, "Failed to parse .xctestrun file")
	assert.NotNil(t, xcTestRunData, "Parsed data should not be nil")

	return xcTestRunData
}

// Helper function to create mock data and parse the .xctestrun file
func createAndParseXCTestRunFileVersion2(t *testing.T) xCTestRunData {
	// Arrange: Create a temporary .xctestrun file with mock data
	tempFile, err := os.CreateTemp("", "testfile*.xctestrun")
	assert.NoError(t, err, "Failed to create temp file")
	defer os.Remove(tempFile.Name()) // Cleanup after test

	// language=xml
	xcTestRunFileFormatVersion2 := `
		<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
	<key>CodeCoverageBuildableInfos</key>
	<array>
		<dict>
			<key>Architectures</key>
			<array>
				<string>arm64</string>
			</array>
			<key>BuildableIdentifier</key>
			<string>506AF5D12D429D9E008E829B:primary</string>
			<key>IncludeInReport</key>
			<true/>
			<key>IsStatic</key>
			<false/>
			<key>Name</key>
			<string>FakeApp3.app</string>
			<key>ProductPaths</key>
			<array>
				<string>__TESTROOT__/Debug-iphoneos/FakeApp3.app/FakeApp3</string>
			</array>
			<key>SourceFiles</key>
			<array>
				<string>ContentView.swift</string>
				<string>FakeApp3App.swift</string>
			</array>
			<key>SourceFilesCommonPathPrefix</key>
			<string>/Users/simonschafer/git/FakeApp3/FakeApp3/FakeApp3/</string>
			<key>Toolchains</key>
			<array>
				<string>com.apple.dt.toolchain.XcodeDefault</string>
			</array>
		</dict>
		<dict>
			<key>Architectures</key>
			<array>
				<string>arm64</string>
			</array>
			<key>BuildableIdentifier</key>
			<string>506AF5E12D429DA0008E829B:primary</string>
			<key>IncludeInReport</key>
			<true/>
			<key>IsStatic</key>
			<false/>
			<key>Name</key>
			<string>FakeApp3Tests.xctest</string>
			<key>ProductPaths</key>
			<array>
				<string>__TESTROOT__/Debug-iphoneos/FakeApp3.app/PlugIns/FakeApp3Tests.xctest/FakeApp3Tests</string>
			</array>
			<key>SourceFiles</key>
			<array>
				<string>/Users/simonschafer/git/FakeApp3/FakeApp3/FakeApp3Tests/FakeApp3Tests.swift</string>
			</array>
			<key>Toolchains</key>
			<array>
				<string>com.apple.dt.toolchain.XcodeDefault</string>
			</array>
		</dict>
		<dict>
			<key>Architectures</key>
			<array>
				<string>arm64</string>
			</array>
			<key>BuildableIdentifier</key>
			<string>506AF5EB2D429DA0008E829B:primary</string>
			<key>IncludeInReport</key>
			<true/>
			<key>IsStatic</key>
			<false/>
			<key>Name</key>
			<string>FakeApp3UITests.xctest</string>
			<key>ProductPaths</key>
			<array>
				<string>__TESTROOT__/Debug-iphoneos/FakeApp3UITests-Runner.app/PlugIns/FakeApp3UITests.xctest/FakeApp3UITests</string>
			</array>
			<key>SourceFiles</key>
			<array>
				<string>.swift</string>
				<string>LaunchTests.swift</string>
			</array>
			<key>SourceFilesCommonPathPrefix</key>
			<string>/Users/simonschafer/git/FakeApp3/FakeApp3/FakeApp3UITests/FakeApp3UITests</string>
			<key>Toolchains</key>
			<array>
				<string>com.apple.dt.toolchain.XcodeDefault</string>
			</array>
		</dict>
	</array>
	<key>ContainerInfo</key>
	<dict>
		<key>ContainerName</key>
		<string>FakeApp3</string>
		<key>SchemeName</key>
		<string>FakeApp3</string>
	</dict>
	<key>TestConfigurations</key>
	<array>
		<dict>
			<key>Name</key>
			<string>Test Scheme Action</string>
			<key>TestTargets</key>
			<array>
				<dict>
					<key>BlueprintName</key>
					<string>FakeApp3Tests</string>
					<key>BlueprintProviderName</key>
					<string>FakeApp3</string>
					<key>BlueprintProviderRelativePath</key>
					<string>FakeApp3.xcodeproj</string>
					<key>BundleIdentifiersForCrashReportEmphasis</key>
					<array>
						<string>saucelabs.FakeApp3</string>
						<string>saucelabs.FakeApp3Tests</string>
						<string>saucelabs.FakeApp3UITests</string>
					</array>
					<key>ClangProfileDataDirectoryPath</key>
					<string>__DERIVEDDATA__/Build/ProfileData</string>
					<key>CommandLineArguments</key>
					<array/>
					<key>DefaultTestExecutionTimeAllowance</key>
					<integer>600</integer>
					<key>DependentProductPaths</key>
					<array>
						<string>__TESTROOT__/Debug-iphoneos/FakeApp3.app</string>
						<string>__TESTROOT__/Debug-iphoneos/FakeApp3.app/PlugIns/FakeApp3Tests.xctest</string>
						<string>__TESTROOT__/Debug-iphoneos/FakeApp3UITests-Runner.app</string>
						<string>__TESTROOT__/Debug-iphoneos/FakeApp3UITests-Runner.app/PlugIns/FakeApp3UITests.xctest</string>
					</array>
					<key>DiagnosticCollectionPolicy</key>
					<integer>1</integer>
					<key>EnvironmentVariables</key>
					<dict>
						<key>APP_DISTRIBUTOR_ID_OVERRIDE</key>
						<string>com.apple.AppStore</string>
						<key>OS_ACTIVITY_DT_MODE</key>
						<string>YES</string>
						<key>SQLITE_ENABLE_THREAD_ASSERTIONS</key>
						<string>1</string>
						<key>TERM</key>
						<string>dumb</string>
					</dict>
					<key>IsAppHostedTestBundle</key>
					<true/>
					<key>ParallelizationEnabled</key>
					<true/>
					<key>PreferredScreenCaptureFormat</key>
					<string>screenRecording</string>
					<key>ProductModuleName</key>
					<string>FakeApp3Tests</string>
					<key>SystemAttachmentLifetime</key>
					<string>deleteOnSuccess</string>
					<key>TestBundlePath</key>
					<string>__TESTHOST__/PlugIns/FakeApp3Tests.xctest</string>
					<key>TestHostBundleIdentifier</key>
					<string>saucelabs.FakeApp3</string>
					<key>TestHostPath</key>
					<string>__TESTROOT__/Debug-iphoneos/FakeApp3.app</string>
					<key>TestLanguage</key>
					<string></string>
					<key>TestRegion</key>
					<string></string>
					<key>TestTimeoutsEnabled</key>
					<false/>
					<key>TestingEnvironmentVariables</key>
					<dict>
						<key>DYLD_INSERT_LIBRARIES</key>
						<string>__TESTHOST__/Frameworks/libXCTestBundleInject.dylib</string>
						<key>XCInjectBundleInto</key>
						<string>unused</string>
					</dict>
					<key>ToolchainsSettingValue</key>
					<array/>
					<key>UserAttachmentLifetime</key>
					<string>deleteOnSuccess</string>
				</dict>
				<dict>
					<key>BlueprintName</key>
					<string>FakeApp3UITests</string>
					<key>BlueprintProviderName</key>
					<string>FakeApp3</string>
					<key>BlueprintProviderRelativePath</key>
					<string>FakeApp3.xcodeproj</string>
					<key>BundleIdentifiersForCrashReportEmphasis</key>
					<array>
						<string>saucelabs.FakeApp3</string>
						<string>saucelabs.FakeApp3Tests</string>
						<string>saucelabs.FakeApp3UITests</string>
					</array>
					<key>ClangProfileDataDirectoryPath</key>
					<string>__DERIVEDDATA__/Build/ProfileData</string>
					<key>CommandLineArguments</key>
					<array/>
					<key>DefaultTestExecutionTimeAllowance</key>
					<integer>600</integer>
					<key>DependentProductPaths</key>
					<array>
						<string>__TESTROOT__/Debug-iphoneos/FakeApp3.app</string>
						<string>__TESTROOT__/Debug-iphoneos/FakeApp3.app/PlugIns/FakeApp3Tests.xctest</string>
						<string>__TESTROOT__/Debug-iphoneos/FakeApp3UITests-Runner.app</string>
						<string>__TESTROOT__/Debug-iphoneos/FakeApp3UITests-Runner.app/PlugIns/FakeApp3UITests.xctest</string>
					</array>
					<key>DiagnosticCollectionPolicy</key>
					<integer>1</integer>
					<key>EnvironmentVariables</key>
					<dict>
						<key>APP_DISTRIBUTOR_ID_OVERRIDE</key>
						<string>com.apple.AppStore</string>
						<key>OS_ACTIVITY_DT_MODE</key>
						<string>YES</string>
						<key>SQLITE_ENABLE_THREAD_ASSERTIONS</key>
						<string>1</string>
						<key>TERM</key>
						<string>dumb</string>
					</dict>
					<key>IsUITestBundle</key>
					<true/>
					<key>IsXCTRunnerHostedTestBundle</key>
					<true/>
					<key>ParallelizationEnabled</key>
					<true/>
					<key>PreferredScreenCaptureFormat</key>
					<string>screenRecording</string>
					<key>ProductModuleName</key>
					<string>FakeApp3UITests</string>
					<key>SystemAttachmentLifetime</key>
					<string>deleteOnSuccess</string>
					<key>TestBundlePath</key>
					<string>__TESTHOST__/PlugIns/FakeApp3UITests.xctest</string>
					<key>TestHostBundleIdentifier</key>
					<string>saucelabs.FakeApp3UITests.xctrunner</string>
					<key>TestHostPath</key>
					<string>__TESTROOT__/Debug-iphoneos/FakeApp3UITests-Runner.app</string>
					<key>TestLanguage</key>
					<string></string>
					<key>TestRegion</key>
					<string></string>
					<key>TestTimeoutsEnabled</key>
					<false/>
					<key>TestingEnvironmentVariables</key>
					<dict/>
					<key>ToolchainsSettingValue</key>
					<array/>
					<key>UITargetAppCommandLineArguments</key>
					<array/>
					<key>UITargetAppEnvironmentVariables</key>
					<dict>
						<key>APP_DISTRIBUTOR_ID_OVERRIDE</key>
						<string>com.apple.AppStore</string>
					</dict>
					<key>UITargetAppPath</key>
					<string>__TESTROOT__/Debug-iphoneos/FakeApp3.app</string>
					<key>UserAttachmentLifetime</key>
					<string>deleteOnSuccess</string>
				</dict>
			</array>
		</dict>
	</array>
	<key>TestPlan</key>
	<dict>
		<key>IsDefault</key>
		<true/>
		<key>Name</key>
		<string>FakeApp3</string>
	</dict>
	<key>__xctestrun_metadata__</key>
	<dict>
		<key>FormatVersion</key>
		<integer>2</integer>
	</dict>
</dict>
</plist>

	`
	_, err = tempFile.WriteString(xcTestRunFileFormatVersion2)
	assert.NoError(t, err, "Failed to write mock data to temp file")
	tempFile.Close()

	// Act: Use the codec to parse the temp file
	xcTestRunData, err := parseFile(tempFile.Name())

	// Assert: Verify the parsed data
	assert.NoError(t, err, "Failed to parse .xctestrun file")
	assert.NotNil(t, xcTestRunData, "Parsed data should not be nil")

	return xcTestRunData
}

func TestTestHostBundleIdentifier2(t *testing.T) {
	xcTestRunData := createAndParseXCTestRunFileVersion2(t)
	assert.Equal(t, "com.example.myApp", xcTestRunData.TestConfig.TestHostBundleIdentifier, "TestHostBundleIdentifier mismatch")
}

func TestTestHostBundleIdentifier(t *testing.T) {
	xcTestRunData := createAndParseXCTestRunFile(t)
	assert.Equal(t, "com.example.myApp", xcTestRunData.TestConfig.TestHostBundleIdentifier, "TestHostBundleIdentifier mismatch")
}

func TestTestBundlePath(t *testing.T) {
	xcTestRunData := createAndParseXCTestRunFile(t)
	assert.Equal(t, "__TESTHOST__/PlugIns/RunnerTests.xctest", xcTestRunData.TestConfig.TestBundlePath, "TestBundlePath mismatch")
}

func TestEnvironmentVariables(t *testing.T) {
	xcTestRunData := createAndParseXCTestRunFile(t)
	assert.Equal(t, map[string]any{
		"APP_DISTRIBUTOR_ID_OVERRIDE":     "com.apple.AppStore",
		"OS_ACTIVITY_DT_MODE":             "YES",
		"SQLITE_ENABLE_THREAD_ASSERTIONS": "1",
		"TERM":                            "dumb",
	}, xcTestRunData.TestConfig.EnvironmentVariables, "EnvironmentVariables mismatch")
}

func TestTestingEnvironmentVariables(t *testing.T) {
	xcTestRunData := createAndParseXCTestRunFile(t)
	assert.Equal(t, map[string]any{
		"DYLD_INSERT_LIBRARIES": "__TESTHOST__/Frameworks/libXCTestBundleInject.dylib",
		"XCInjectBundleInto":    "unused",
		"Test":                  "xyz",
	}, xcTestRunData.TestConfig.TestingEnvironmentVariables, "TestingEnvironmentVariables mismatch")
}

func TestCommandLineArguments(t *testing.T) {
	xcTestRunData := createAndParseXCTestRunFile(t)
	assert.Equal(t, []string{}, xcTestRunData.TestConfig.CommandLineArguments, "CommandLineArguments mismatch")
}

func TestOnlyTestIdentifiers(t *testing.T) {
	xcTestRunData := createAndParseXCTestRunFile(t)
	assert.Equal(t, []string{
		"TestClass1/testMethod1",
		"TestClass2/testMethod1",
	}, xcTestRunData.TestConfig.OnlyTestIdentifiers, "OnlyTestIdentifiers mismatch")
}

func TestSkipTestIdentifiers(t *testing.T) {
	xcTestRunData := createAndParseXCTestRunFile(t)
	assert.Equal(t, []string{
		"TestClass1/testMethod2",
		"TestClass2/testMethod2",
	}, xcTestRunData.TestConfig.SkipTestIdentifiers, "SkipTestIdentifiers mismatch")
}

func TestFormatVersion(t *testing.T) {
	xcTestRunData := createAndParseXCTestRunFile(t)
	assert.Equal(t, 1, xcTestRunData.XCTestRunMetadata.FormatVersion, "FormatVersion mismatch")
}

func TestIsUITestBundle(t *testing.T) {
	xcTestRunData := createAndParseXCTestRunFile(t)
	assert.Equal(t, true, xcTestRunData.TestConfig.IsUITestBundle, "IsUITestBundle mismatch")
}

func TestParseXCTestRunNotSupportedForFormatVersionOtherThanOne(t *testing.T) {
	// Arrange: Create a temporary .xctestrun file with mock data
	tempFile, err := os.CreateTemp("", "testfile*.xctestrun")
	assert.NoError(t, err, "Failed to create temp file")
	defer os.Remove(tempFile.Name()) // Cleanup after test

	xcTestRunFileFormatVersion2 := `
		<?xml version="1.0" encoding="UTF-8"?>
		<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
		<plist version="1.0">
		<dict>
			<key>__xctestrun_metadata__</key>
			<dict>
				<key>FormatVersion</key>
				<integer>2</integer>
			</dict>
		</dict>
		</plist>
	`
	_, err = tempFile.WriteString(xcTestRunFileFormatVersion2)
	assert.NoError(t, err, "Failed to write mock data to temp file")
	tempFile.Close()

	// Act: Use the codec to parse the temp file
	_, err = parseFile(tempFile.Name())

	// Assert the Error Message
	assert.Equal(t, "go-ios currently only supports .xctestrun files in formatVersion 1: The formatVersion of your xctestrun file is 2, feel free to open an issue in https://github.com/danielpaulus/go-ios/issues to add support", err.Error(), "Error Message mismatch")
}

// Helper function to create testConfig from parsed mock data
func createTestConfigFromParsedMockData(t *testing.T) (TestConfig, ios.DeviceEntry, *TestListener) {
	// Arrange: Create parsed XCTestRunData using the helper function
	xcTestRunData := createAndParseXCTestRunFile(t)

	// Mock dependencies
	mockDevice := ios.DeviceEntry{
		DeviceID: 8110,
	}
	mockListener := &TestListener{}

	// Act: Convert XCTestRunData to TestConfig
	testConfig, err := xcTestRunData.buildTestConfig(mockDevice, mockListener)

	// Assert: Validate the returned TestConfig
	assert.NoError(t, err, "Error converting to TestConfig")

	return testConfig, mockDevice, mockListener
}

func TestConfigTestRunnerBundleId(t *testing.T) {
	testConfig, _, _ := createTestConfigFromParsedMockData(t)
	assert.Equal(t, "com.example.myApp", testConfig.TestRunnerBundleId, "TestRunnerBundleId mismatch")
}

func TestConfigXctestConfigName(t *testing.T) {
	testConfig, _, _ := createTestConfigFromParsedMockData(t)
	assert.Equal(t, "RunnerTests.xctest", testConfig.XctestConfigName, "XctestConfigName mismatch")
}

func TestConfigCommandLineArguments(t *testing.T) {
	testConfig, _, _ := createTestConfigFromParsedMockData(t)
	assert.Equal(t, []string{}, testConfig.Args, "data mismatch")
}

func TestConfigEnvironmentVariables(t *testing.T) {
	testConfig, _, _ := createTestConfigFromParsedMockData(t)
	assert.Equal(t, map[string]any{
		"APP_DISTRIBUTOR_ID_OVERRIDE":     "com.apple.AppStore",
		"OS_ACTIVITY_DT_MODE":             "YES",
		"SQLITE_ENABLE_THREAD_ASSERTIONS": "1",
		"TERM":                            "dumb",
		"DYLD_INSERT_LIBRARIES":           "__TESTHOST__/Frameworks/libXCTestBundleInject.dylib",
		"XCInjectBundleInto":              "unused",
		"Test":                            "xyz",
	}, testConfig.Env, "EnvironmentVariables mismatch")
}

func TestConfigTestsToRun(t *testing.T) {
	testConfig, _, _ := createTestConfigFromParsedMockData(t)
	assert.Equal(t, []string{
		"TestClass1/testMethod1",
		"TestClass2/testMethod1",
	}, testConfig.TestsToRun, "TestsToRun mismatch")
}

func TestConfigTestsToSkip(t *testing.T) {
	testConfig, _, _ := createTestConfigFromParsedMockData(t)
	assert.Equal(t, []string{
		"TestClass1/testMethod2",
		"TestClass2/testMethod2",
	}, testConfig.TestsToSkip, "TestsToSkip mismatch")
}

func TestConfigXcTest(t *testing.T) {
	testConfig, _, _ := createTestConfigFromParsedMockData(t)
	assert.Equal(t, false, testConfig.XcTest, "XcTest mismatch")
}

func TestConfigDevice(t *testing.T) {
	testConfig, mockDevice, _ := createTestConfigFromParsedMockData(t)
	assert.Equal(t, mockDevice, testConfig.Device, "Device mismatch")
}

func TestConfigListener(t *testing.T) {
	testConfig, _, mockListener := createTestConfigFromParsedMockData(t)
	assert.Equal(t, mockListener, testConfig.Listener, "Listener mismatch")
}
