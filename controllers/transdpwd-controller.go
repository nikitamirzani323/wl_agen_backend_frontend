package controllers

import (
	"log"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/nikitamirzani323/wl_super_backend_frontend/entities"
)

func Transdpwdhome(c *fiber.Ctx) error {
	hostname := c.Hostname()
	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	client := new(entities.Home)
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	log.Println("Hostname: ", hostname)
	render_page := time.Now()
	axios := resty.New()
	resp, err := axios.R().
		SetResult(responsedefault{}).
		SetAuthToken(token[1]).
		SetError(responseerror{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"client_hostname": hostname,
			"page":            client.Page,
		}).
		Post(PATH + "api/transaksidepowd")
	if err != nil {
		log.Println(err.Error())
	}
	log.Println("Response Info:")
	log.Println("  Error      :", err)
	log.Println("  Status Code:", resp.StatusCode())
	log.Println("  Status     :", resp.Status())
	log.Println("  Proto      :", resp.Proto())
	log.Println("  Time       :", resp.Time())
	log.Println("  Received At:", resp.ReceivedAt())
	log.Println("  Body       :\n", resp)
	log.Println()
	result := resp.Result().(*responsedefault)
	if result.Status == 200 {
		return c.JSON(fiber.Map{
			"status":  result.Status,
			"message": result.Message,
			"record":  result.Record,
			"time":    time.Since(render_page).String(),
		})
	} else {
		result_error := resp.Error().(*responseerror)
		return c.JSON(fiber.Map{
			"status":  result_error.Status,
			"message": result_error.Message,
			"time":    time.Since(render_page).String(),
		})
	}
}
func TransdpwdSave(c *fiber.Ctx) error {
	type payload_transdpwdsave struct {
		Page                string  `json:"page"`
		Sdata               string  `json:"sdata" `
		Transdpwd_id        string  `json:"transdpwd_id" `
		Transdpwd_tipedoc   string  `json:"transdpwd_tipedoc" `
		Transdpwd_idmember  string  `json:"transdpwd_idmember" `
		Transdpwd_bank_in   int     `json:"transdpwd_bank_in" `
		Transdpwd_bank_out  int     `json:"transdpwd_bank_out" `
		Transdpwd_amount    float32 `json:"transdpwd_amount" `
		Transdpwd_ipaddress string  `json:"transdpwd_ipaddress" `
		Transdpwd_note      string  `json:"transdpwd_note" `
		Transdpwd_status    string  `json:"transdpwd_status" `
	}
	hostname := c.Hostname()
	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	client := new(payload_transdpwdsave)
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	log.Println("Hostname: ", hostname)
	render_page := time.Now()
	axios := resty.New()
	resp, err := axios.R().
		SetResult(responsedefault{}).
		SetAuthToken(token[1]).
		SetError(responseerror{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"client_hostname":     hostname,
			"page":                client.Page,
			"sdata":               client.Sdata,
			"transdpwd_id":        client.Transdpwd_id,
			"transdpwd_tipedoc":   client.Transdpwd_tipedoc,
			"transdpwd_idmember":  client.Transdpwd_idmember,
			"transdpwd_bank_in":   client.Transdpwd_bank_in,
			"transdpwd_bank_out":  client.Transdpwd_bank_out,
			"transdpwd_amount":    client.Transdpwd_amount,
			"transdpwd_ipaddress": client.Transdpwd_ipaddress,
			"transdpwd_note":      client.Transdpwd_note,
			"transdpwd_status":    client.Transdpwd_status,
		}).
		Post(PATH + "api/transaksidepowdsave")
	if err != nil {
		log.Println(err.Error())
	}
	log.Println("Response Info:")
	log.Println("  Error      :", err)
	log.Println("  Status Code:", resp.StatusCode())
	log.Println("  Status     :", resp.Status())
	log.Println("  Proto      :", resp.Proto())
	log.Println("  Time       :", resp.Time())
	log.Println("  Received At:", resp.ReceivedAt())
	log.Println("  Body       :\n", resp)
	log.Println()
	result := resp.Result().(*responsedefault)
	if result.Status == 200 {
		return c.JSON(fiber.Map{
			"status":  result.Status,
			"message": result.Message,
			"record":  result.Record,
			"time":    time.Since(render_page).String(),
		})
	} else {
		result_error := resp.Error().(*responseerror)
		return c.JSON(fiber.Map{
			"status":  result_error.Status,
			"message": result_error.Message,
			"time":    time.Since(render_page).String(),
		})
	}
}
func TransdpwdUpdate(c *fiber.Ctx) error {
	type payload_transdpwdupdate struct {
		Page               string `json:"page"`
		Sdata              string `json:"sdata" `
		Transdpwd_id       string `json:"transdpwd_id" `
		Transdpwd_idmember string `json:"transdpwd_idmember" `
		Transdpwd_note     string `json:"transdpwd_note" `
		Transdpwd_status   string `json:"transdpwd_status" `
	}
	hostname := c.Hostname()
	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	client := new(payload_transdpwdupdate)
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	log.Println("Hostname: ", hostname)
	render_page := time.Now()
	axios := resty.New()
	resp, err := axios.R().
		SetResult(responsedefault{}).
		SetAuthToken(token[1]).
		SetError(responseerror{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"client_hostname":    hostname,
			"page":               client.Page,
			"sdata":              client.Sdata,
			"transdpwd_id":       client.Transdpwd_id,
			"transdpwd_idmember": client.Transdpwd_idmember,
			"transdpwd_note":     client.Transdpwd_note,
			"transdpwd_status":   client.Transdpwd_status,
		}).
		Post(PATH + "api/transaksidepowdupdate")
	if err != nil {
		log.Println(err.Error())
	}
	log.Println("Response Info:")
	log.Println("  Error      :", err)
	log.Println("  Status Code:", resp.StatusCode())
	log.Println("  Status     :", resp.Status())
	log.Println("  Proto      :", resp.Proto())
	log.Println("  Time       :", resp.Time())
	log.Println("  Received At:", resp.ReceivedAt())
	log.Println("  Body       :\n", resp)
	log.Println()
	result := resp.Result().(*responsedefault)
	if result.Status == 200 {
		return c.JSON(fiber.Map{
			"status":  result.Status,
			"message": result.Message,
			"record":  result.Record,
			"time":    time.Since(render_page).String(),
		})
	} else {
		result_error := resp.Error().(*responseerror)
		return c.JSON(fiber.Map{
			"status":  result_error.Status,
			"message": result_error.Message,
			"time":    time.Since(render_page).String(),
		})
	}
}
