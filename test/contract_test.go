package test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const baseURL = "http://localhost:9010"

func TestMain(m *testing.M) {
	cmd := setup()

	result := m.Run()

	teardown(cmd)

	os.Exit(result)
}

func TestInventoryItemLifecycle_ShouldCreateRetrieveUpdateAndDelete(t *testing.T) {
	// Test create
	resp := postJSON(t, "/inventory", `{
		"Name": "Cool Runnings (1994)",
		"Location": "AD12"
	}`)
	assertCreated(t, resp)

	// Test read
	id := extractString(t, resp)
	resp = get(t, "/inventory/"+id)
	body := extractString(t, resp)
	expected := fmt.Sprintf(`{"id":%s,"name":"Cool Runnings (1994)","location":"AD12","available":true}`, id)
	assert.Equal(t, body, expected)

	// Test update
	resp = putJSON(t, "/inventory/"+id, `{
		"Name": "Cool Runnings (1994) UPDATED",
		"Location": "AD12 UPDATED"
	}`)
	assertNoContent(t, resp)

	// Test read... for update
	resp = get(t, "/inventory/"+id)
	body = extractString(t, resp)
	expected = fmt.Sprintf(`{"id":%s,"name":"Cool Runnings (1994) UPDATED","location":"AD12 UPDATED","available":true}`, id)
	assert.Equal(t, body, expected)

	// Test delete
	resp = delete(t, "/inventory/"+id)
	assertNoContent(t, resp)
	// TODO: Need to make sure we get appropriate 404's
}

func delete(t *testing.T, path string) *http.Response {
	req, err := http.NewRequest(http.MethodDelete, baseURL+path, nil)
	if err != nil {
		assert.NoError(t, err)
	}
	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		assert.NoError(t, err)
	}
	return resp
}

func putJSON(t *testing.T, path string, body string) *http.Response {
	reader := strings.NewReader(body)
	req, err := http.NewRequest(http.MethodPut, baseURL+path, reader)
	if err != nil {
		assert.NoError(t, err)
	}
	req.Header.Set("Content-Type", "application/json")
	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		assert.NoError(t, err)
	}
	return resp
}

func postJSON(t *testing.T, path string, body string) *http.Response {
	reader := strings.NewReader(body)
	resp, err := http.Post(baseURL+path, "application/json", reader)
	if err != nil {
		assert.NoError(t, err)
	}
	return resp
}

func get(t *testing.T, path string) *http.Response {
	resp, err := http.Get(baseURL + path)
	if err != nil {
		assert.NoError(t, err)
	}
	return resp
}

func setup() *exec.Cmd {
	cmd := exec.Command("matchstick-video")
	cmd.Env = []string{
		"PORT=9010",
		"DB_DRIVER=sqlite3",
		"MIGRATION_SOURCE=file://../migrations",
	}
	if err := cmd.Start(); err != nil {
		panic(err)
	}

	// Wait one second for app to start...
	time.Sleep(time.Second * 1)

	return cmd
}

func teardown(cmd *exec.Cmd) {
	if err := cmd.Process.Kill(); err != nil {
		panic(err)
	}
}

func assertCreated(t *testing.T, resp *http.Response) {
	assert.Equal(t, 201, resp.StatusCode, "expected Created")
}

func assertNoContent(t *testing.T, resp *http.Response) {
	assert.Equal(t, 204, resp.StatusCode, "expected No Content")
}

func extractString(t *testing.T, resp *http.Response) string {
	bytes, err := ioutil.ReadAll(resp.Body)
	assert.NoError(t, err)
	return string(bytes)
}
