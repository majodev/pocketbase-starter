package main

import (
	"net/http"
	"testing"

	"github.com/pocketbase/pocketbase/tests"
	"github.com/stretchr/testify/require"
)

// https://pocketbase.io/docs/go-testing/
//
// Feel free to start the test pocketbase app with the embedded test data with:
// make test-server
//
// The test database is checked into the repository under the following path:
const testDataDir = "./test_pb_data"

func TestHelloEndpoint(t *testing.T) {
	// setup the test ApiScenario app instance
	setupTestApp := func(t testing.TB) *tests.TestApp {
		testApp, err := tests.NewTestApp(testDataDir)
		require.NoError(t, err)

		// no need to cleanup since scenario.Test() will do that for us
		// defer testApp.Cleanup()

		err = bindAppHooks(testApp)
		require.NoError(t, err)

		return testApp
	}

	scenarios := []tests.ApiScenario{
		{
			Name:            "index page",
			Method:          http.MethodGet,
			URL:             "/",
			ExpectedStatus:  200,
			ExpectedContent: []string{"github.com/majodev/pocketbase-starter"},
			TestAppFactory:  setupTestApp,
		},
	}

	for _, scenario := range scenarios {
		scenario.Test(t)
	}
}
