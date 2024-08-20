package handlers

import (
	"net/http/httptest"
	"testing"
)

func TestHome(t *testing.T) {
	routes := getRoutes()
	ts := httptest.NewTLSServer(routes)
	defer ts.Close()

	resp, err := ts.Client().Get(ts.URL + "/")
	if err != nil {
		t.Log(err)
		t.Fatal(err)
	}

	if resp.StatusCode != 200 {
		t.Errorf("for home page, expected status 200 but got %d", resp.StatusCode)
	}

	// Could not test this as it was flagged by Windows Security as suspicious. So commented out...
	//     `Operation did not complete successfully because the file contains a virus or potentially unwanted software.`
	// Solution 1: Disable `leakless` by creating your own *Launcher and set its `leakless` property to false.
	//     [](https://pkg.go.dev/github.com/go-rod/rod@v0.79.0/lib/launcher#Launcher.Leakless)
	// Solution 2: Tell your antivirus to ignore the `leakless` binary.

	//bodyText, err := io.ReadAll(resp.Body)
	//if err != nil {
	//	t.Fatal(err)
	//}
	//if !strings.Contains(string(bodyText), "awesome") {
	//	cel.TakeScreenShot(ts.URL+"/", "HomeTest", 1500, 1000)
	//	t.Error("did not find awesome")
	//}
}
