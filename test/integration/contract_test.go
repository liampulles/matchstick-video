//abuild integration

package integration

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
	// Test update on a non-existant item
	resp := putJSON(t, "/inventory/999", `{
		"Name": "Cool Runnings (1993) UPDATED",
		"Location": "AD12 UPDATED"
	}`)
	assertNotFound(t, resp)
	body := extractString(t, resp)
	expected := fmt.Sprintf(`could not update inventory item - repository find error: cannot execute query - db scan error: entity not found: type=[inventory item]`)
	assert.Equal(t, expected, body)

	// Test delete on a non-existant item
	resp = delete(t, "/inventory/999")
	assertNotFound(t, resp)
	body = extractString(t, resp)
	expected = fmt.Sprintf(`could not delete inventory item - repository delete error: entity not found: type=[inventory item]`)
	assert.Equal(t, expected, body)

	// Test checkout on a non-existant item
	resp = putJSON(t, "/inventory/999/checkout", "")
	assertNotFound(t, resp)
	body = extractString(t, resp)
	expected = fmt.Sprintf(`could not checkout inventory item - repository find error: cannot execute query - db scan error: entity not found: type=[inventory item]`)
	assert.Equal(t, expected, body)

	// Test check in on a non-existant item
	resp = putJSON(t, "/inventory/999/checkin", "")
	assertNotFound(t, resp)
	body = extractString(t, resp)
	expected = fmt.Sprintf(`could not check in inventory item - repository find error: cannot execute query - db scan error: entity not found: type=[inventory item]`)
	assert.Equal(t, expected, body)

	// Test create
	resp = postJSON(t, "/inventory", `{
		"Name": "Cool Runnings (1993)",
		"Location": "AD12"
	}`)
	assertCreated(t, resp)

	// Test read
	id := extractString(t, resp)
	resp = get(t, "/inventory/"+id)
	assertOk(t, resp)
	body = extractString(t, resp)
	expected = fmt.Sprintf(`{"id":%s,"name":"Cool Runnings (1993)","location":"AD12","available":true}`, id)
	assert.Equal(t, expected, body)

	// Test create with same name.. should be constraint violation
	resp = postJSON(t, "/inventory", `{
		"Name": "Cool Runnings (1993)",
		"Location": "CD12"
	}`)
	assertBadRequest(t, resp)
	body = extractString(t, resp)
	expected = fmt.Sprintf(`could not create inventory item - repository create error: cannot execute query - db scan error: uniqueness constraint error: ERROR: duplicate key value violates unique constraint "inventory_item_name_key" (SQLSTATE 23505)`)
	assert.Equal(t, expected, body)

	// Test create with invalid name
	resp = postJSON(t, "/inventory", `{
		"Name": "",
		"Location": "AD70"
	}`)
	assertBadRequest(t, resp)
	body = extractString(t, resp)
	expected = fmt.Sprintf(`could not create inventory item - factory error: validation error: field=[name], problem=[must not be blank]`)
	assert.Equal(t, expected, body)

	// Test read all
	resp = get(t, "/inventory")
	assertOk(t, resp)
	body = extractString(t, resp)
	expected = fmt.Sprintf(`[{"id":%s,"name":"Cool Runnings (1993)"}]`, id)
	assert.Equal(t, expected, body)

	// Test update
	resp = putJSON(t, "/inventory/"+id, `{
		"Name": "Cool Runnings (1993) UPDATED",
		"Location": "AD12 UPDATED"
	}`)
	assertNoContent(t, resp)

	// Test read... for update
	resp = get(t, "/inventory/"+id)
	body = extractString(t, resp)
	assertOk(t, resp)
	expected = fmt.Sprintf(`{"id":%s,"name":"Cool Runnings (1993) UPDATED","location":"AD12 UPDATED","available":true}`, id)
	assert.Equal(t, expected, body)

	// Test checkout
	resp = putJSON(t, "/inventory/"+id+"/checkout", "")
	assertNoContent(t, resp)

	// Test read... for checkout
	resp = get(t, "/inventory/"+id)
	body = extractString(t, resp)
	assertOk(t, resp)
	expected = fmt.Sprintf(`{"id":%s,"name":"Cool Runnings (1993) UPDATED","location":"AD12 UPDATED","available":false}`, id)
	assert.Equal(t, expected, body)

	// Test check in
	resp = putJSON(t, "/inventory/"+id+"/checkin", "")
	assertNoContent(t, resp)

	// Test read... for check in
	resp = get(t, "/inventory/"+id)
	body = extractString(t, resp)
	assertOk(t, resp)
	expected = fmt.Sprintf(`{"id":%s,"name":"Cool Runnings (1993) UPDATED","location":"AD12 UPDATED","available":true}`, id)
	assert.Equal(t, expected, body)

	// Test delete
	resp = delete(t, "/inventory/"+id)
	assertNoContent(t, resp)

	// Test read... for delete
	resp = get(t, "/inventory/"+id)
	body = extractString(t, resp)
	assertNotFound(t, resp)
	expected = fmt.Sprintf(`could not read inventory item - repository find error: cannot execute query - db scan error: entity not found: type=[inventory item]`)
	assert.Equal(t, expected, body)
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
		"MIGRATION_SOURCE=file://../../migrations",
		"DB_USER=integration",
		"DB_PASSWORD=integration",
		"DB_NAME=integration",
		"DB_PORT=5050",
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

func assertOk(t *testing.T, resp *http.Response) {
	assert.Equal(t, 200, resp.StatusCode, "expected OK")
}

func assertCreated(t *testing.T, resp *http.Response) {
	assert.Equal(t, 201, resp.StatusCode, "expected Created")
}

func assertNoContent(t *testing.T, resp *http.Response) {
	assert.Equal(t, 204, resp.StatusCode, "expected No Content")
}

func assertNotFound(t *testing.T, resp *http.Response) {
	assert.Equal(t, 404, resp.StatusCode, "expected Not Found")
}

func assertBadRequest(t *testing.T, resp *http.Response) {
	assert.Equal(t, 400, resp.StatusCode, "expected Bad Request")
}

func extractString(t *testing.T, resp *http.Response) string {
	bytes, err := ioutil.ReadAll(resp.Body)
	assert.NoError(t, err)
	return string(bytes)
}
