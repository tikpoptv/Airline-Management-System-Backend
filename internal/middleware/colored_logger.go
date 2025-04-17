package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/fatih/color"
	"github.com/labstack/echo/v4"
)

func ColoredLoggerMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		start := time.Now()
		err := next(c)
		stop := time.Now()

		method := c.Request().Method
		uri := c.Request().RequestURI
		status := c.Response().Status
		latency := stop.Sub(start)

		// ⏱ latency string (auto unit)
		latencyStr := latency.String()
		if latency < time.Millisecond {
			latencyStr = fmt.Sprintf("%dµs", latency.Microseconds())
		} else {
			latencyStr = fmt.Sprintf("%dms", latency.Milliseconds())
		}

		// 🎨 Method color
		methodColor := color.New(color.FgCyan).SprintfFunc()
		switch method {
		case "POST":
			methodColor = color.New(color.FgHiBlue).SprintfFunc()
		case "DELETE":
			methodColor = color.New(color.FgMagenta).SprintfFunc()
		case "PUT":
			methodColor = color.New(color.FgHiCyan).SprintfFunc()
		}

		// 🎨 Status color
		statusIcon := "✅"
		statusColor := color.New(color.FgGreen).SprintfFunc()
		if status >= 400 && status < 500 {
			statusIcon = "🚫"
			statusColor = color.New(color.FgYellow).SprintfFunc()
		} else if status >= 500 {
			statusIcon = "❌"
			statusColor = color.New(color.FgRed).SprintfFunc()
		}

		// 🧠 Error message
		errMsg := ""
		if err != nil {
			errMsg = color.New(color.FgRed).Sprintf(" | ERROR: %v", err)
		}

		// 🕓 Time
		timestamp := time.Now().Format("2006-01-02 15:04:05")

		// 🖨 Final log
		fmt.Fprintf(os.Stdout,
			"%s [%s] %-7s %-25s %s %s ⏱ %s%s\n",
			statusIcon,
			timestamp,
			methodColor(method),
			uri,
			statusColor(fmt.Sprintf("%d %s", status, http.StatusText(status))),
			statusIcon,
			latencyStr,
			errMsg,
		)

		return err
	}
}
