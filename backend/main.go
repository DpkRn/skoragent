package main

import (
	"fmt"
	"math/rand"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/skor/agentreferal/config"
	"github.com/skor/agentreferal/models"
)

var otpStore = make(map[string]string) // stores phone -> otp

func main() {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "https://yourdomain.com"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	r.LoadHTMLGlob("templates/*")
	db := config.DB()
	r.GET("/api/:agentId", func(c *gin.Context) {
		c.HTML(http.StatusOK, "form.html", gin.H{
			"AgentID": c.Param("agentId"),
		})
	})

	// Send OTP
	r.POST("/api/send-otp", func(c *gin.Context) {
		phone := c.PostForm("phone")
		otp := fmt.Sprintf("%06d", rand.Intn(1000000))
		otpStore[phone] = otp

		fmt.Printf("Sending OTP %s to phone %s\n", otp, phone)
		c.String(http.StatusOK, "OTP sent to your phone.")
	})

	r.POST("/api/confirm", func(c *gin.Context) {
		name := c.PostForm("name")
		phone := c.PostForm("phone")
		agentId := c.PostForm("agentId")
		enteredOtp := c.PostForm("otp")

		storedOtp, exists := otpStore[phone]
		if !exists || storedOtp != enteredOtp {
			c.String(http.StatusBadRequest, "Invalid or expired OTP")
			return
		}
		var record models.SC_AGENT_CUSTOMER_MAPPING
		if err := db.Model(&models.SC_AGENT_CUSTOMER_MAPPING{}).Where("customer_phone=?", phone).First(&record).Error; err != nil {
			db.Create(&models.SC_AGENT_CUSTOMER_MAPPING{Customer_name: name, Customer_phone: phone, Agent_ID: agentId})

			delete(otpStore, phone)
			c.String(http.StatusOK, fmt.Sprintf("Thank you %s! Your interest has been recorded.", name))
			return
		}
		c.String(http.StatusOK, fmt.Sprintf("You have already mapped with %s", record.Agent_ID))

	})

	r.GET("/api/verify/:agentId", func(ctx *gin.Context) {
		fmt.Println("verifying")
		var agent models.SC_AGENT
		if err := db.Model(&models.SC_AGENT{}).Where("agent_id=?", ctx.Param("agentId")).First(&agent).Error; err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"success": false})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"success": true})
	})

	r.Run(":8080")
}
