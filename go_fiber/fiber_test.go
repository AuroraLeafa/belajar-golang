package main

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"go_fiber/helper"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var app = fiber.New()

func TestRoutingGet(t *testing.T) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	request := httptest.NewRequest("GET", "/", nil)
	resp, err := app.Test(request)
	assert.Nil(t, err, "Error should be nil")
	assert.Equal(t, 200, resp.StatusCode, "Status should be 200")

	body, _ := io.ReadAll(resp.Body)
	assert.Equal(t, "Hello, World!", string(body))
}

// -------------------------------------- Query params  ----------------------------
func TestCtxQueryParams(t *testing.T) {
	app.Get("/hello", func(c *fiber.Ctx) error {
		name := c.Query("name", "Guest")
		return c.SendString("Hello, " + name)
	})

	//-------------------------------------- Check With Param ----------------------------
	request := httptest.NewRequest("GET", "/hello?name=Reff", nil)
	resp, err := app.Test(request)
	assert.Nil(t, err, "Error should be nil")
	assert.Equal(t, 200, resp.StatusCode, "Status should be 200")

	body, _ := io.ReadAll(resp.Body)
	assert.Equal(t, "Hello, Reff", string(body))

	//-------------------------------------- Check Default ----------------------------
	request = httptest.NewRequest("GET", "/hello?name=", nil)
	resp, err = app.Test(request)
	assert.Nil(t, err, "Error should be nil")
	assert.Equal(t, 200, resp.StatusCode, "Status should be 200")

	body, _ = io.ReadAll(resp.Body)
	assert.Equal(t, "Hello, Guest", string(body))
}

// -------------------------------------- HTTP Request  ----------------------------
func TestHttpRequest(t *testing.T) {
	app.Get("/request", func(c *fiber.Ctx) error {
		firstName := c.Get("firstname")   //Get Header
		lastName := c.Cookies("lastname") //Get Cookies
		return c.SendString("Hello, " + firstName + " " + lastName)
	})

	request := httptest.NewRequest("GET", "/request", nil)
	request.Header.Set("firstname", "Reff")
	request.AddCookie(&http.Cookie{Name: "lastname", Value: "Sugg"})

	resp, err := app.Test(request)
	assert.Nil(t, err, "Error should be nil")
	assert.Equal(t, 200, resp.StatusCode, "Status should be 200")

	body, _ := io.ReadAll(resp.Body)
	assert.Equal(t, "Hello, Reff Sugg", string(body))
}

// -------------------------------------- Route with params  ----------------------------
func TestRouteParam(t *testing.T) {
	app.Get("/users/:userId/orders/:orderId", func(c *fiber.Ctx) error {
		userId := c.Params("userId")
		orderId := c.Params("orderId")
		return c.SendString("Order dari User " + userId + " Adalah " + orderId)
	})

	request := httptest.NewRequest("GET", "/users/1/orders/2", nil)

	resp, err := app.Test(request)
	assert.Nil(t, err, "Error should be nil")
	assert.Equal(t, 200, resp.StatusCode, "Status should be 200")

	body, _ := io.ReadAll(resp.Body)
	assert.Equal(t, "Order dari User 1 Adalah 2", string(body))
}

// -------------------------------------- Request Form  ----------------------------
func TestRequestForm(t *testing.T) {
	app.Post("/hello", func(c *fiber.Ctx) error {
		name := c.FormValue("name")
		return c.SendString("Hello, " + name)
	})

	reader := strings.NewReader("name=Reff")
	request := httptest.NewRequest(http.MethodPost, "/hello", reader)
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := app.Test(request)
	assert.Nil(t, err, "Error should be nil")
	assert.Equal(t, 200, resp.StatusCode, "Status should be 200")

	body, err := io.ReadAll(resp.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Hello, Reff", string(body))
}

// -------------------------------------- Form Upload ----------------------------
// Karena mengunakan unit test maka perlu file sample, disini file sample dimasukkan menggunakan embed

//go:embed source/sample.txt
var sample []byte

func TestFormUpload(t *testing.T) {
	app.Post("/upload", func(c *fiber.Ctx) error { // path adalah /uplaod
		file, err := c.FormFile("file") //Menggunakan method FormFile dengan key adalah "file"
		helper.PanicIfError(err)
		err = c.SaveFile(file, "./target/"+file.Filename) //file terupload akan disave pada folder target
		helper.PanicIfError(err)
		return c.SendString("Upload Success")
	})

	//Dibawah ini untuk melakukan unit test Upload Form
	body := new(bytes.Buffer)                                 //Body harus dalam bentuk binary
	writer := multipart.NewWriter(body)                       //Membuat writer untuk menulis pada file. disimpan pada body
	file, err := writer.CreateFormFile("file", "example.txt") //membuat file baru
	helper.PanicIfError(err)
	file.Write(sample)
	writer.Close() //writer perlu diclose, untuk menghentikan proses menulis file karena sudah tersimpan pada body

	request := httptest.NewRequest(http.MethodPost, "/upload", body)
	request.Header.Set("Content-Type", writer.FormDataContentType())

	resp, err := app.Test(request)
	assert.Nil(t, err, "Error should be nil")
	assert.Equal(t, 200, resp.StatusCode, "Status should be 200")

	bytes, err := io.ReadAll(resp.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Upload Success", string(bytes))
}

// -------------------------------------- Request Body ----------------------------
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func TestRequestBody(t *testing.T) {
	app.Post("/login", func(c *fiber.Ctx) error { // path adalah /login
		body := c.Body()
		request := new(LoginRequest)
		err := json.Unmarshal(body, &request)
		helper.PanicIfError(err)

		return c.SendString("Hello, " + request.Username)
	})

	//Dibawah ini untuk melakukan unit test Request Form
	body := strings.NewReader(`{"username":"Reff", "password":"Secret"}`) //Body harus dalam bentuk JSON
	request := httptest.NewRequest(http.MethodPost, "/login", body)
	request.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(request)
	assert.Nil(t, err, "Error should be nil")
	assert.Equal(t, 200, resp.StatusCode, "Status should be 200")

	bytes, err := io.ReadAll(resp.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Hello, Reff", string(bytes))
}

// -------------------------------------- Request Body ----------------------------
type RegisterUser struct {
	Username string `json:"username" xml:"username" form:"username"`
	Password string `json:"password" xml:"password" form:"password"`
	Name     string `json:"name" xml:"name" form:"name"`
}

func TestBodyParser(t *testing.T) {
	app.Post("/register", func(c *fiber.Ctx) error {
		request := new(RegisterUser)
		err := c.BodyParser(request)
		helper.PanicIfError(err)
		return c.SendString("Register Success, " + request.Username)
	})
}

func TestBodyParserJSON(t *testing.T) {
	TestBodyParser(t)
	//Dibawah ini untuk melakukan unit test Body Parser
	body := strings.NewReader(`{"username":"Reff", "password":"Secret", "name":"Refansyah"}`) //Body harus dalam bentuk JSON
	request := httptest.NewRequest(http.MethodPost, "/register", body)
	request.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(request)
	assert.Nil(t, err, "Error should be nil")
	assert.Equal(t, 200, resp.StatusCode, "Status should be 200")

	bytes, err := io.ReadAll(resp.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Register Success, Reff", string(bytes))
}

func TestBodyParserForm(t *testing.T) {
	TestBodyParser(t)
	//Dibawah ini untuk melakukan unit test Body Parser
	body := strings.NewReader(`username=Reff&password=Secret&name=Refansyah`) //Body dalam bentuk JSON
	request := httptest.NewRequest(http.MethodPost, "/register", body)
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := app.Test(request)
	assert.Nil(t, err, "Error should be nil")
	assert.Equal(t, 200, resp.StatusCode, "Status should be 200")

	bytes, err := io.ReadAll(resp.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Register Success, Reff", string(bytes))
}

func TestBodyParserXML(t *testing.T) {
	TestBodyParser(t)
	//Dibawah ini untuk melakukan unit test Body Parser
	body := strings.NewReader(`
		<RegisterUser>
			<username>Reff</username>
			<password>Secret</password>
			<name>Refansyah</name>
		</RegisterUser>
	`) //Body dalam bentuk XML

	request := httptest.NewRequest(http.MethodPost, "/register", body)
	request.Header.Set("Content-Type", "application/xml")

	resp, err := app.Test(request)
	assert.Nil(t, err, "Error should be nil")
	assert.Equal(t, 200, resp.StatusCode, "Status should be 200")

	bytes, err := io.ReadAll(resp.Body)
	assert.Nil(t, err)
	assert.Equal(t, "Register Success, Reff", string(bytes))
}

// -------------------------------------- Response JSON ----------------------------
func TestResponseJSON(t *testing.T) {
	app.Get("/user", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"username": "Reff",
			"name":     "Refansyah",
		})
	})

	//Dibawah ini untuk melakukan unit test Response JSON
	request := httptest.NewRequest(http.MethodGet, "/user", nil)
	request.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(request)
	assert.Nil(t, err, "Error should be nil")
	assert.Equal(t, 200, resp.StatusCode, "Status should be 200")

	bytes, err := io.ReadAll(resp.Body)
	assert.Nil(t, err)
	assert.Equal(t, `{"name":"Refansyah","username":"Reff"}`, string(bytes))
}

// -------------------------------------- Download Files ----------------------------
func TestDownloadFiles(t *testing.T) {
	app.Get("/download", func(c *fiber.Ctx) error {
		return c.Download("./source/sample.txt", "sample.txt")
	})

	//Dibawah ini untuk melakukan unit test Download File
	request := httptest.NewRequest(http.MethodGet, "/download", nil)

	resp, err := app.Test(request)
	assert.Nil(t, err, "Error should be nil")
	assert.Equal(t, 200, resp.StatusCode, "Status should be 200")
	assert.Equal(t, `attachment; filename="sample.txt"`, resp.Header.Get("Content-Disposition"))

	bytes, err := io.ReadAll(resp.Body)
	assert.Nil(t, err)
	assert.Equal(t, `Sample File for Upload`, string(bytes))
}

// -------------------------------------- Routing Group ----------------------------
func TestRoutingGroup(t *testing.T) {
	HelloWorldHandler := func(c *fiber.Ctx) error {
		return c.SendString("Hello World")
	}

	api := app.Group("/api")
	api.Get("/hello", HelloWorldHandler)
	api.Get("/world", HelloWorldHandler)

	web := app.Group("/web")
	web.Get("/hello", HelloWorldHandler)
	web.Get("/world", HelloWorldHandler)

	//Dibawah ini untuk melakukan unit test Routing Group
	requestWebHello := httptest.NewRequest(http.MethodGet, "/api/hello", nil)

	resp, err := app.Test(requestWebHello)
	assert.Nil(t, err, "Error should be nil")
	assert.Equal(t, 200, resp.StatusCode, "Status should be 200")

	bytes, err := io.ReadAll(resp.Body)
	assert.Nil(t, err)
	assert.Equal(t, `Hello World`, string(bytes))
}

// -------------------------------------- Static ----------------------------
func TestStatic(t *testing.T) {
	app.Static("/public", "./source")

	//Dibawah ini untuk melakukan unit test Static
	requestWebHello := httptest.NewRequest(http.MethodGet, "/public/sample.txt", nil)

	resp, err := app.Test(requestWebHello)
	assert.Nil(t, err, "Error should be nil")
	assert.Equal(t, 200, resp.StatusCode, "Status should be 200")

	bytes, err := io.ReadAll(resp.Body)
	assert.Nil(t, err)
	assert.Equal(t, `Sample File for Upload`, string(bytes))
}
