package testmanagerd

import (
	"bytes"
	"fmt"
	"github.com/danielpaulus/go-ios/ios"
	"github.com/danielpaulus/go-ios/ios/installationproxy"
	"howett.net/plist"
	"io"
	"maps"
	"os"
	"path/filepath"
	"strings"
)

// xctestrunutils provides utilities for parsing `.xctestrun` files with FormatVersion 1.
// It simplifies the extraction of test configurations and metadata into structured objects (`xCTestRunData`),
// enabling efficient setup for iOS test execution.
//
// Features:
// - Parses `.xctestrun` files to extract test metadata and configurations.
// - Supports building `TestConfig` objects for test execution.
//
// Note: Only `.xctestrun` files with `FormatVersion` 1 are supported. For other versions,
// contributions or requests for support can be made in the relevant GitHub repository.

// xCTestRunData represents the structure of an .xctestrun file

// schemeData represents the structure of a scheme-specific test configuration
type xCTestRunVersion2 struct {
	ContainerInfo struct {
		ContainerName string `plist:"ContainerName"`
	} `plist:"ContainerInfo"`
	TestConfigurations []struct {
		TestTargets []schemeData `plist:"TestTargets"`
	} `plist:"TestConfigurations"`
}

type schemeData struct {
	TestHostBundleIdentifier        string
	TestBundlePath                  string
	SkipTestIdentifiers             []string
	OnlyTestIdentifiers             []string
	IsUITestBundle                  bool
	CommandLineArguments            []string
	EnvironmentVariables            map[string]any
	TestingEnvironmentVariables     map[string]any
	UITargetAppEnvironmentVariables map[string]any
	UITargetAppPath                 string
}

func (data schemeData) buildTestConfig(device ios.DeviceEntry, listener *TestListener, allAps []installationproxy.AppInfo) (TestConfig, error) {
	testsToRun := data.OnlyTestIdentifiers
	testsToSkip := data.SkipTestIdentifiers

	testEnv := make(map[string]any)
	var bundleId string

	if data.IsUITestBundle {
		maps.Copy(testEnv, data.EnvironmentVariables)
		maps.Copy(testEnv, data.TestingEnvironmentVariables)
		maps.Copy(testEnv, data.UITargetAppEnvironmentVariables)
		// Only call getBundleID if allAps is provided
		if allAps != nil {
			bundleId, _ = getBundleID(allAps, data.UITargetAppPath)
		}
	}

	// Extract only the file name
	var testBundlePath = filepath.Base(data.TestBundlePath)

	// Build the TestConfig object from parsed data
	testConfig := TestConfig{
		BundleId:           bundleId,
		TestRunnerBundleId: data.TestHostBundleIdentifier,
		XctestConfigName:   testBundlePath,
		Args:               data.CommandLineArguments,
		Env:                testEnv,
		TestsToRun:         testsToRun,
		TestsToSkip:        testsToSkip,
		XcTest:             !data.IsUITestBundle,
		Device:             device,
		Listener:           listener,
	}

	return testConfig, nil
}

// parseFile reads the .xctestrun file and decodes it into a map
func parseFile(filePath string) ([]schemeData, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return []schemeData{}, fmt.Errorf("failed to open xctestrun file: %w", err)
	}
	defer file.Close()
	return decode(file)
}

// decode decodes the binary xctestrun content into the xCTestRunData struct
func decode(r io.Reader) ([]schemeData, error) {
	// Read the entire content once
	xctestrunFileContent, err := io.ReadAll(r)
	if err != nil {
		return []schemeData{}, fmt.Errorf("unable to read xctestrun content: %w", err)
	}

	// First, we only parse the version property of the xctestrun file. The rest of the parsing depends on this version.
	version, err := getFormatVersion(xctestrunFileContent)
	if err != nil {
		return []schemeData{}, err
	}

	switch version {
	case 1:
		return parseXCTestRunFileFormatVersion1(xctestrunFileContent)
	case 2:
		return parseXCTestRunFileFormatVersion2(xctestrunFileContent)
	default:
		return []schemeData{}, fmt.Errorf("the provided .xctestrun format version %d is not supported", version)
	}
}

// Helper method to get the format version of the xctestrun file
func getFormatVersion(xctestrunFileContent []byte) (int, error) {

	type xCTestRunMetadata struct {
		Metadata struct {
			Version int `plist:"FormatVersion"`
		} `plist:"__xctestrun_metadata__"`
	}

	var metadata xCTestRunMetadata
	if _, err := plist.Unmarshal(xctestrunFileContent, &metadata); err != nil {
		return 0, fmt.Errorf("failed to parse format version: %w", err)
	}

	return metadata.Metadata.Version, nil
}

func parseXCTestRunFileFormatVersion1(content []byte) ([]schemeData, error) {
	// xctestrun files in version 1 use a dynamic key for the pListRoot of the TestConfig. As in the 'key' for the TestConfig is the name
	// of the app. This forces us to iterate over the root of the plist, instead of using a static struct to decode the xctestrun file.
	var pListRoot map[string]interface{}
	if _, err := plist.Unmarshal(content, &pListRoot); err != nil {
		return []schemeData{}, fmt.Errorf("failed to unmarshal plist: %w", err)
	}

	for key, value := range pListRoot {
		// Skip the metadata object
		if key == "__xctestrun_metadata__" {
			continue
		}

		// Attempt to convert to schemeData
		schemeMap, ok := value.(map[string]interface{})
		if !ok {
			continue // Skip if not a valid scheme map
		}

		// Parse the scheme into schemeData and update the TestConfig
		var schemeParsed schemeData
		schemeBuf := new(bytes.Buffer)
		encoder := plist.NewEncoder(schemeBuf)
		if err := encoder.Encode(schemeMap); err != nil {
			return []schemeData{}, fmt.Errorf("failed to encode scheme %s: %w", key, err)
		}

		// Decode the plist buffer into schemeData
		decoder := plist.NewDecoder(bytes.NewReader(schemeBuf.Bytes()))
		if err := decoder.Decode(&schemeParsed); err != nil {
			return []schemeData{}, fmt.Errorf("failed to decode scheme %s: %w", key, err)
		}
		return []schemeData{schemeParsed}, nil
	}
	return []schemeData{}, nil
}

func parseXCTestRunFileFormatVersion2(content []byte) ([]schemeData, error) {
	var testConfigs xCTestRunVersion2
	if _, err := plist.Unmarshal(content, &testConfigs); err != nil {
		return []schemeData{}, fmt.Errorf("failed to parse format version: %w", err)
	}

	return testConfigs.TestConfigurations[0].TestTargets, nil
}

func getBundleID(apps []installationproxy.AppInfo, uiTargetAppPath string) (string, error) {
	var appNameWithSuffix = filepath.Base(uiTargetAppPath)
	var uiTargetAppName = strings.TrimSuffix(appNameWithSuffix, ".app")
	for _, app := range apps {
		if app.CFBundleName == uiTargetAppName {
			return app.CFBundleIdentifier, nil
		}
	}
	return "", fmt.Errorf("app %s not found", uiTargetAppName)
}
