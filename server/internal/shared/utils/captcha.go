package utils

import (
	"fmt"
	"image/color"
	"sync"
	"time"

	"github.com/mojocn/base64Captcha"
)

// CaptchaConfig 验证码配置
type CaptchaConfig struct {
	Enabled bool
	Mode    string // digit | string | math | chinese
	Length  int
	Width   int
	Height  int
	Expire  int // 秒
}

// CaptchaResult 验证码结果
type CaptchaResult struct {
	CaptchaID string `json:"captcha_id"`
	Image     string `json:"image"`
}

var (
	captchaStore base64Captcha.Store
	storeOnce    sync.Once
)

func initStore(expire int) {
	storeOnce.Do(func() {
		captchaStore = base64Captcha.NewMemoryStore(10240, time.Duration(expire)*time.Second)
	})
}

// ResetStore 重置 store（配置热加载时调用）
func ResetStore() {
	storeOnce = sync.Once{}
}

// NewCaptchaStore 初始化验证码存储
func NewCaptchaStore(expire int) {
	ResetStore()
	initStore(expire)
}

// GenerateCaptcha 根据配置生成验证码
func GenerateCaptcha(cfg CaptchaConfig) (*CaptchaResult, error) {
	initStore(cfg.Expire)

	if cfg.Width <= 0 {
		cfg.Width = 240
	}
	if cfg.Height <= 0 {
		cfg.Height = 80
	}
	if cfg.Length <= 0 {
		cfg.Length = 4
	}

	var driver base64Captcha.Driver
	bgColor := &color.RGBA{R: 240, G: 240, B: 240, A: 255}

	switch cfg.Mode {
	case "string":
		d := base64Captcha.NewDriverString(
			cfg.Height, cfg.Width,
			1, // noiseCount
			base64Captcha.OptionShowHollowLine|base64Captcha.OptionShowSlimeLine,
			cfg.Length,
			"abcdefghjkmnpqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ23456789",
			bgColor,
			base64Captcha.DefaultEmbeddedFonts,
			nil,
		)
		d.ConvertFonts()
		driver = d
	case "math":
		d := base64Captcha.NewDriverMath(
			cfg.Height, cfg.Width,
			1,
			base64Captcha.OptionShowHollowLine|base64Captcha.OptionShowSlimeLine,
			bgColor,
			base64Captcha.DefaultEmbeddedFonts,
			nil,
		)
		d.ConvertFonts()
		driver = d
	case "chinese":
		d := base64Captcha.NewDriverChinese(
			cfg.Height, cfg.Width,
			1,
			base64Captcha.OptionShowHollowLine|base64Captcha.OptionShowSlimeLine,
			cfg.Length,
			"中国你好世界大家和平均平安快乐康健康幸福美的善",
			bgColor,
			base64Captcha.DefaultEmbeddedFonts,
			nil,
		)
		d.ConvertFonts()
		driver = d
	default: // digit
		driver = base64Captcha.NewDriverDigit(cfg.Height, cfg.Width, cfg.Length, 0.7, 60)
	}

	id, b64s, _, err := base64Captcha.NewCaptcha(driver, captchaStore).Generate()
	if err != nil {
		return nil, fmt.Errorf("生成验证码失败: %w", err)
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
	if captchaStore == nil {
		return false
	}
	return captchaStore.Verify(captchaID, answer, true)
}
