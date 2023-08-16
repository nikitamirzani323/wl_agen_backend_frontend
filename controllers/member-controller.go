package controllers

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/nikitamirzani323/wl_super_backend_frontend/entities"
)

type responsemember struct {
	Status   int         `json:"status"`
	Message  string      `json:"message"`
	Listbank interface{} `json:"listbank"`
	Record   interface{} `json:"record"`
}

func Memberhome(c *fiber.Ctx) error {
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
		SetResult(responsemember{}).
		SetAuthToken(token[1]).
		SetError(responseerror{}).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"client_hostname": hostname,
			"page":            client.Page,
		}).
		Post(PATH + "api/member")
	if err != nil {
		log.Println(err.Error())
	}
	fmt.Println("Response Info:")
	fmt.Println("  Error      :", err)
	fmt.Println("  Status Code:", resp.StatusCode())
	fmt.Println("  Status     :", resp.Status())
	fmt.Println("  Proto      :", resp.Proto())
	fmt.Println("  Time       :", resp.Time())
	fmt.Println("  Received At:", resp.ReceivedAt())
	fmt.Println("  Body       :\n", resp)
	fmt.Println()
	result := resp.Result().(*responsemember)
	if result.Status == 200 {
		return c.JSON(fiber.Map{
			"status":   result.Status,
			"message":  result.Message,
			"record":   result.Record,
			"listbank": result.Listbank,
			"time":     time.Since(render_page).String(),
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
func MemberSave(c *fiber.Ctx) error {
	type payload_membersave struct {
		Page            string `json:"page"`
		Sdata           string `json:"sdata" `
		Member_id       string `json:"member_id"`
		Member_username string `json:"member_username" `
		Member_password string `json:"member_password" `
		Member_name     string `json:"member_name" `
		Member_phone    string `json:"member_phone" `
		Member_email    string `json:"member_email" `
		Member_status   string `json:"member_status" `
	}
	hostname := c.Hostname()
	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	client := new(payload_membersave)
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
			"sdata":           client.Sdata,
			"member_id":       client.Member_id,
			"member_username": client.Member_username,
			"member_password": client.Member_password,
			"member_name":     client.Member_name,
			"member_phone":    client.Member_phone,
			"member_email":    client.Member_email,
			"member_status":   client.Member_status,
		}).
		Post(PATH + "api/membersave")
	if err != nil {
		log.Println(err.Error())
	}
	fmt.Println("Response Info:")
	fmt.Println("  Error      :", err)
	fmt.Println("  Status Code:", resp.StatusCode())
	fmt.Println("  Status     :", resp.Status())
	fmt.Println("  Proto      :", resp.Proto())
	fmt.Println("  Time       :", resp.Time())
	fmt.Println("  Received At:", resp.ReceivedAt())
	fmt.Println("  Body       :\n", resp)
	fmt.Println()
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
func MemberBankSave(c *fiber.Ctx) error {
	type payload_memberbanksave struct {
		Page                    string `json:"page"`
		Sdata                   string `json:"sdata" `
		Memberbank_idagenmember string `json:"memberbank_idagenmember" `
		Memberbank_idbanktype   string `json:"memberbank_idbanktype" `
		Memberbank_norek        string `json:"memberbank_norek" `
		Memberbank_nmownerbank  string `json:"memberbank_nmownerbank" `
		Memberbank_status       string `json:"memberbank_status" `
	}
	hostname := c.Hostname()
	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	client := new(payload_memberbanksave)
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
			"client_hostname":         hostname,
			"page":                    client.Page,
			"sdata":                   client.Sdata,
			"memberbank_idagenmember": client.Memberbank_idagenmember,
			"memberbank_idbanktype":   client.Memberbank_idbanktype,
			"memberbank_norek":        client.Memberbank_norek,
			"memberbank_nmownerbank":  client.Memberbank_nmownerbank,
			"memberbank_status":       client.Memberbank_status,
		}).
		Post(PATH + "api/memberbanksave")
	if err != nil {
		log.Println(err.Error())
	}
	fmt.Println("Response Info:")
	fmt.Println("  Error      :", err)
	fmt.Println("  Status Code:", resp.StatusCode())
	fmt.Println("  Status     :", resp.Status())
	fmt.Println("  Proto      :", resp.Proto())
	fmt.Println("  Time       :", resp.Time())
	fmt.Println("  Received At:", resp.ReceivedAt())
	fmt.Println("  Body       :\n", resp)
	fmt.Println()
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
func MemberBankDelete(c *fiber.Ctx) error {
	type payload_memberbankdelete struct {
		Page                    string `json:"page"`
		Sdata                   string `json:"sdata" `
		Memberbank_id           int    `json:"memberbank_id" `
		Memberbank_idagenmember string `json:"memberbank_idagenmember" `
	}
	hostname := c.Hostname()
	bearToken := c.Get("Authorization")
	token := strings.Split(bearToken, " ")
	client := new(payload_memberbankdelete)
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
			"client_hostname":         hostname,
			"page":                    client.Page,
			"sdata":                   client.Sdata,
			"memberbank_id":           client.Memberbank_id,
			"memberbank_idagenmember": client.Memberbank_idagenmember,
		}).
		Post(PATH + "api/memberbankdelete")
	if err != nil {
		log.Println(err.Error())
	}
	fmt.Println("Response Info:")
	fmt.Println("  Error      :", err)
	fmt.Println("  Status Code:", resp.StatusCode())
	fmt.Println("  Status     :", resp.Status())
	fmt.Println("  Proto      :", resp.Proto())
	fmt.Println("  Time       :", resp.Time())
	fmt.Println("  Received At:", resp.ReceivedAt())
	fmt.Println("  Body       :\n", resp)
	fmt.Println()
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
