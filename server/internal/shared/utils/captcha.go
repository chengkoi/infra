package utils

import (
	"github.com/mojocn/base64Captcha"
)

// CaptchaResult 验证码结果
type CaptchaResult struct {
	CaptchaID string `json:"captcha_id"`
	Image     string `json:"image"`
}

// GenerateCaptcha 生成数字验证码
func GenerateCaptcha() (*CaptchaResult, error) {
	driver := base64Captcha.NewDriverDigit(80, 240, 4, 0.7, 80)

	id, b64s, _, err := base64Captcha.NewCaptcha(driver, base64Captcha.DefaultMemStore).Generate()
	if err != nil {
		return nil, err
	}

	return &CaptchaResult{
		CaptchaID: id,
		Image:     b64s,
	}, nil
}

// VerifyCaptcha 校验验证码
func VerifyCaptcha(captchaID, answer string) bool {
	if captchaID == "" || answer == "" {
		return false
	}
	return base64Captcha.DefaultMemStore.Verify(captchaID, answer, true)
}
