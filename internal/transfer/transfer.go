package transfer

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"

	"github.com/p1xart/music-lib/internal/entity"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func NewRouter() {
	jsonHandler := slog.NewJSONHandler(os.Stderr, nil)
    log := slog.New(jsonHandler)

	client := &http.Client{}
	router = gin.Default()

	router.GET("/new", func(c *gin.Context) {
		var	songResp entity.SongRequest
		if err := c.ShouldBindJSON(&songResp); err != nil {
			log.Info("failed to bind json", "error", err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		serverUrl, exists := os.LookupEnv("TOKEN")
		if !exists {
			log.Error("server url is not found in env")
			c.JSON(http.StatusInternalServerError, nil)
			return
		}

		var buf bytes.Buffer
		err := json.NewEncoder(&buf).Encode(songResp)
		if err != nil {
			log.Error("failed to encode json", "error", err.Error())
			c.JSON(http.StatusInternalServerError, nil)
			return
		}

		req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/info", serverUrl), &buf)
		if err != nil {
			log.Error("failed to build request", "error", err.Error())
			c.JSON(http.StatusInternalServerError, nil)
			return
		}

		res, err := client.Do(req)
		if err != nil {
			log.Error("failed to run request", "error", err.Error())
			c.JSON(http.StatusInternalServerError, nil)
			return
		}

		resBytes, err := io.ReadAll(res.Body)
		if err != nil {
			log.Error("failed to read response", "error", err.Error())
			c.JSON(http.StatusInternalServerError, nil)
			return
		}

		var songReq entity.SongResponse
		err = json.Unmarshal(resBytes, &songReq)
		if err != nil {
			log.Error("failed to bind json", "error", err.Error())
			c.JSON(http.StatusInternalServerError, nil)
			return
		}
	})

	router.Run()
}