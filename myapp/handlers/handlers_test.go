package handlers

import (
	"net/http"
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

	//bodyText, err := io.ReadAll(resp.Body)
	//if err != nil {
	//	t.Fatal(err)
	//}
	//if !strings.Contains(string(bodyText), "awesome") {
	//	cel.TakeScreenShot(ts.URL+"/", "HomeTest", 1500, 1000)
	//	t.Error("did not find awesome")
	//}
}

func TestClicker(t *testing.T) {
	// Could not test any of this as calls to any functions in `dusk.go` are flagged by Windows Security as suspicious.
	// `Operation did not complete successfully because the file contains a virus or potentially unwanted software.`

	//routes := getRoutes()
	//ts := httptest.NewTLSServer(routes)
	//defer ts.Close()

	//page := cel.FetchPage(ts.URL + "/tester")
	//outputElement := cel.SelectElementByID(page, "output")
	//button := cel.SelectElementByID(page, "clicker")

	//testHTML, _ := outputElement.HTML()
	//if strings.Contains(testHTML, "Clicked the button") {
	//	t.Errorf("found text that should not be there ")
	//}

	//button.MustClick()
	//testHTML, _ = outputElement.HTML()
	//if !strings.Contains(testHTML, "Clicked the button") {
	//	t.Errorf("did not find text that should be there ")
	//}
}

func TestHome2(t *testing.T) {
	req, _ := http.NewRequest("GET", "/", nil)
	ctx := getCtx(req)
	req = req.WithContext(ctx)

	rr := httptest.NewRecorder()
	cel.Session.Put(ctx, "test_key", "Hello, world.")

	h := http.HandlerFunc(testHandlers.Home)
	h.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("for home page, expected status %d but got %d", http.StatusOK, rr.Code)
	}

	if cel.Session.GetString(ctx, "test_key") != "Hello, world." {
		t.Error("did not get correct value from session")
	}
}
