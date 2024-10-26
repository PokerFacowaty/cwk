package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func setupRouter(pcounter *int) *gin.Engine {
    r := gin.Default()
    trustedProxies := []string{"127.0.0.1"}
    r.SetTrustedProxies(trustedProxies)

    r.GET("/", func (c *gin.Context) {
        c.String(http.StatusOK, strconv.Itoa(*pcounter))
    })

    // Single user since I don't need anything more
    user := os.Getenv("CWK_USERNAME")
    password := os.Getenv("CWK_PASSWORD")

    authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
        user: password,
    }))

    authorized.POST("/add", func(context *gin.Context) {
        *pcounter++
        context.String(http.StatusOK, strconv.Itoa(*pcounter))
    })
    return r
}

func main() {
    // export GIN_MODE=release
    var counter int

    c := make(chan os.Signal, 1)
    signal.Notify(c, os.Interrupt)
    go func() {
        // Receive the signal from the c channel
        <-c
        fmt.Println("\ncwk: Saving the counter to file...")
        os.WriteFile("counter.txt", []byte(strconv.Itoa(counter)), 0666)
        fmt.Println("\ncwk: Saved.")
        os.Exit(0)
    }()

    content, err := os.ReadFile("counter.txt")
    if err != nil {
        log.Fatal(err)
    }

    counter, err = strconv.Atoi(strings.TrimSpace(string(content)))
    if err != nil {
        log.Fatal(err)
    }

    r := setupRouter(&counter)
    r.Run(":8080")
}

