package main

import (
	"flag"
	"fmt"
	"image-processor/config"
	"image-processor/lib"
	"io/ioutil"
	"os"
	"strconv"
	"sync"
	"time"

	vips "github.com/davidbyttow/govips/v2"
	"github.com/panjf2000/ants/v2"
	"github.com/sirupsen/logrus"
)

var params = &vips.ExportParams{}
var wg sync.WaitGroup

func init() {
	vips.Startup(nil)
}

func main() {
	start := time.Now()
	readFlag()

	// set vips config
	config.C.OutputType = "webp"
	params.Format = vips.ImageTypeWEBP // TODO: 待優化
	params.Quality = config.C.Quality
	params.Lossless = config.C.Lossless
	params.Effort = config.C.ReductionEffort

	// create output directory
	config.C.OutputDirPath = config.C.OutputDirPath +
		start.Format("2006-01-02_15:04:05") + "_" +
		strconv.Itoa(config.C.Limit) + "_" +
		strconv.Itoa(params.Quality) + "_" +
		strconv.Itoa(params.Effort) + "/"
	os.MkdirAll(config.C.OutputDirPath, os.ModePerm)

	// new goroutine pool
	pool, _ := ants.NewPool(config.C.Limit, ants.WithPreAlloc(true))
	defer func() {
		// 釋放pool
		pool.Release()
		// 清除緩存
		vips.ClearCache()
		// 功能關閉
		vips.Shutdown()
	}()

	// get all file path
	paths := lib.GetInputImgPaths(config.C.InputDirPath)
	for _, v := range paths {
		wg.Add(1)
		pool.Submit(vipsCompress(v))
	}

	wg.Wait()

	end := time.Now()
	fmt.Println("end: ", end)
	fmt.Println("total cost: ", end.Sub(start))
}

// readFlag -
func readFlag() {
	// ------ file ------
	// 設定讀檔的資料夾位置
	flag.StringVar(&config.C.InputDirPath, "in", "./image/in/", "set input directory")
	// 設定寫檔的資料夾位置
	flag.StringVar(&config.C.OutputDirPath, "out", "./image/out/", "set output directory")

	// ------fcompress config ------
	// 壓縮品質(default: 75)
	flag.IntVar(&config.C.Quality, "q", 70, "set compress quality")
	// 設定cpu努力級別(default: 2)
	flag.IntVar(&config.C.ReductionEffort, "r", 2, "set compress cpu reduction effort")
	// 設定是否失真(default: false - 失真)
	flag.BoolVar(&config.C.Lossless, "l", false, "set compress lossless")

	// ------ efficacy ------
	// goroutine limit(default: 4)
	flag.IntVar(&config.C.Limit, "g", 4, "set goroutine limit")

	flag.Parse()
}

// vipsCompress -
func vipsCompress(file *lib.FileInfo) func() {
	return func() {
		defer wg.Done()

		imgFile, err := vips.NewImageFromFile(file.Path)
		if err != nil {
			logrus.Infof("new image failed: %s", err.Error())
			return
		}

		// 將圖片垂直旋轉並重置 EXIF 方向標籤
		if err = imgFile.AutoRotate(); err != nil {
			logrus.Infof("auto rotate image failed: %s", err.Error())
			return
		}

		// 輸出圖片
		ibp, _, err := imgFile.Export(params)
		if err != nil {
			logrus.Infof("create webp image failed: %s", err.Error())
			return
		}

		// 寫出圖片
		if err = ioutil.WriteFile(config.C.OutputDirPath+file.Name+"."+config.C.OutputType, ibp, 0644); err != nil {
			logrus.Infof("write image failed: %s", err.Error())
			return
		}

		fmt.Println("OK. ", file.Path)
		imgFile.Close()
	}
}
