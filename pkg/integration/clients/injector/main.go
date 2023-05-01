package main

import (
	"fmt"
	"os"

	"github.com/BSteffaniak/lazyaws/pkg/app"
	"github.com/BSteffaniak/lazyaws/pkg/app/daemon"
	"github.com/BSteffaniak/lazyaws/pkg/integration/components"
	"github.com/BSteffaniak/lazyaws/pkg/integration/tests"
	integrationTypes "github.com/BSteffaniak/lazyaws/pkg/integration/types"
)

// The purpose of this program is to run lazyaws with an integration test passed in.
// We could have done the check on TEST_NAME in the root main.go but
// that would mean lazyaws would be depending on integration test code which
// would bloat the binary.

// You should not invoke this program directly. Instead you should go through
// go run cmd/integration_test/main.go

func main() {
	dummyBuildInfo := &app.BuildInfo{
		Commit:      "",
		Date:        "",
		Version:     "",
		BuildSource: "integration test",
	}

	integrationTest := getIntegrationTest()

	app.Start(dummyBuildInfo, integrationTest)
}

func getIntegrationTest() integrationTypes.IntegrationTest {
	if daemon.InDaemonMode() {
		// if we've invoked lazyaws as a daemon from within lazyaws,
		// we don't want to pass a test to the rest of the code.
		return nil
	}

	integrationTestName := os.Getenv(components.TEST_NAME_ENV_VAR)
	if integrationTestName == "" {
		panic(fmt.Sprintf(
			"expected %s environment variable to be set, given that we're running an integration test",
			components.TEST_NAME_ENV_VAR,
		))
	}

	allTests := tests.GetTests()
	for _, candidateTest := range allTests {
		if candidateTest.Name() == integrationTestName {
			return candidateTest
		}
	}

	panic("Could not find integration test with name: " + integrationTestName)
}
