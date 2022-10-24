package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmznt4345/sa-65-example/entity"
	"github.com/jmznt4345/sa-65-example/service"
	"golang.org/x/crypto/bcrypt"
)

// LoginPayload login body
type LoginPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// SignUpPayload signup body
type SignUpPayload struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Phonenumber string `json:"Phonenumber"`

	RoleID     uint `json:"RoleID"`
	GenderID   uint `json:"GenderID"`
	Educational_backgroundID uint `json:"Educational_backgroundID"`
}

// LoginResponse token response
type LoginResponse struct {
	Token string `json:"token"`
	ID    uint   `json:"id"`
}

// POST /login
func Login(c *gin.Context) {
	var payload LoginPayload
	var user entity.User

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// ค้นหา user ด้วย email ที่ผู้ใช้กรอกเข้ามา
	if err := entity.DB().Raw("SELECT * FROM users WHERE email = ?", payload.Email).Scan(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// ตรวจสอบรหัสผ่าน
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "password is incerrect"})
		return
	}

	// กำหนดค่า SecretKey, Issuer และระยะเวลาหมดอายุของ Token สามารถกำหนดเองได้
	// SecretKey ใช้สำหรับการ sign ข้อความเพื่อบอกว่าข้อความมาจากตัวเราแน่นอน
	// Issuer เป็น unique id ที่เอาไว้ระบุตัว client
	// ExpirationHours เป็นเวลาหมดอายุของ token

	jwtWrapper := service.JwtWrapper{
		SecretKey:       "SvNQpBN8y3qlVrsGAYYWoJJk56LtzFHx",
		Issuer:          "AuthService",
		ExpirationHours: 24,
	}

	signedToken, err := jwtWrapper.GenerateToken(user.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error signing token"})
		return
	}

	tokenResponse := LoginResponse{
		Token: signedToken,
		ID:    user.ID,
	}

	c.JSON(http.StatusOK, gin.H{"data": tokenResponse})
}

// POST /create
func CreateUser(c *gin.Context) {
	var user entity.User
    var gender entity.Gender
    var role entity.Role
    var educational_background entity.Educational_background
	var payload SignUpPayload

    // ผลลัพธ์ที่ได้จากขั้นตอนที่ 8 จะถูก bind เข้าตัวแปร watchVideo
    if err := c.ShouldBindJSON(&user); err != nil {
    	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    	return
    }

    // 9: ค้นหา gender ด้วย id
    if tx := entity.DB().Where("id = ?", user.GenderID).First(&gender); tx.RowsAffected == 0 {
    	c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
    	return
    }

    // 10: ค้นหา role ด้วย id
    if tx := entity.DB().Where("id = ?", user.RoleID).First(&role); tx.RowsAffected == 0 {
    	c.JSON(http.StatusBadRequest, gin.H{"error": "role not found"})
    	return
    }

    // 11: ค้นหา Educational_background ด้วย id
    if tx := entity.DB().Where("id = ?", user.Educational_backgroundID).First(&educational_background); tx.RowsAffected == 0 {
    	c.JSON(http.StatusBadRequest, gin.H{"error": "educational_background not found"})
    	return
    }

	 // เข้ารหัสลับรหัสผ่านที่ผู้ใช้กรอกก่อนบันทึกลงฐานข้อมูล
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), 14)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error hashing password"})
		return
	}
    //12: สร้าง WatchVideo
    wv := entity.User{
    	Role:      					role,            // โยงความสัมพันธ์กับ Entity Role
    	Gender:             		gender,               // โยงความสัมพันธ์กับ Entity Gender
    	Educational_background:     educational_background,           // โยงความสัมพันธ์กับ Entity Educational_background

		Name:				payload.Name,
		Email: 				payload.Email,
		Password:			string(hashPassword),
		Phonenumber:		payload.Phonenumber,
    }

    // 13: บันทึก
    if err := entity.DB().Create(&wv).Error; err != nil {
    	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    	return
    }
    c.JSON(http.StatusOK, gin.H{"data": wv})
	// var payload SignUpPayload
	// var user entity.User

	// if err := c.ShouldBindJSON(&payload); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }

	// // เข้ารหัสลับรหัสผ่านที่ผู้ใช้กรอกก่อนบันทึกลงฐานข้อมูล
	// hashPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), 14)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "error hashing password"})
	// 	return
	// }

	// user.Name = payload.Name
	// user.Email = payload.Email
	// user.Password = string(hashPassword)
	// user.Phonenumber = payload.Phonenumber

	// user.RoleID = &payload.RoleID
	// user.GenderID = &payload.GenderID
	// user.Educational_backgroundID = &payload.Educational_backgroundID

	// if err := entity.DB().Create(&user).Error; err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }

	// c.JSON(http.StatusCreated, gin.H{"data": user})
}
