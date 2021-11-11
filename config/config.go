package config

// C -
var C Config

// Config -
type Config struct {
	// Quality - 壓縮品質(default: 70)
	Quality int

	// Lossless - true: 無失真(default: false)
	Lossless bool

	// ReductionEffort - 減少文件大小的 CPU 努力級別，整數 0-6（可選，默認4）
	ReductionEffort int

	// Limit - goroutine limit
	Limit int

	// InputDirPath - 輸入圖片的路徑
	InputDirPath string

	// OutputDirPath - 輸出圖片的路徑
	OutputDirPath string

	// OutputType - 輸出檔案類型
	OutputType string
}
