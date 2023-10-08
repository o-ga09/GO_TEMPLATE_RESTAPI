package middleware

import (
	"context"
	"encoding/json"
	"io"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slog"
)

type ServerLogJsonHandler struct {
	slog.Handler
	w io.Writer
	indent int
}

type ServerLogJsonOptions struct {
	SlogOpts slog.HandlerOptions
	Indent int
}

type RequestInfo struct {
	status int
	contents_length int64
	method, path, sourceIP, quesy, user_agent, errors string
	elapsed time.Duration
}

func Cors() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins: []string{
			"*",
		},
		AllowMethods: []string{
			"POST",
			"GET",
			"OPTIONS",
		},
		AllowHeaders: []string{
			"Content-Type",
		},
		AllowCredentials: false,
		MaxAge: 24 * time.Hour,
	})
}

func Logger(l *slog.Logger) gin.HandlerFunc {
	start := time.Now()
	return func(c *gin.Context) {
		r := &RequestInfo{
			c.Writer.Status(),
			c.Request.ContentLength,
			c.Request.Method,
			c.Request.URL.Path,
			c.ClientIP(),
			c.Request.URL.RawQuery,
			c.Request.UserAgent(),
			c.Errors.ByType(gin.ErrorTypePrivate).String(),
			time.Since(start),
		}
		l.Info("host","RequestInfo",r.LogValue())
	}
}

func NewServerLogJsonHandler(w io.Writer, opts ServerLogJsonOptions) *ServerLogJsonHandler {
	return &ServerLogJsonHandler{
		Handler: slog.NewJSONHandler(w,&opts.SlogOpts),
		w: w,
		indent: opts.Indent,
	}
}

func (r *RequestInfo) LogValue() slog.Value {
	return slog.GroupValue(
		slog.Int("status",r.status),
		slog.Int64("contents_length",r.contents_length),
		slog.String("method",r.method),
		slog.String("path",r.path),
		slog.String("sourceIP",r.sourceIP),
		slog.String("query",r.quesy),
		slog.String("user_agent",r.user_agent),
		slog.String("errors",r.errors),
		slog.Duration("elapsed",r.elapsed),
	)
}

func (h *ServerLogJsonHandler) Handle(_ context.Context, r slog.Record) error {
	fields := make(map[string]any,r.NumAttrs())
	fields["time"] = r.Time
	fields["level"] = r.Level
	fields["msg"] = r.Message

	r.Attrs(func(a slog.Attr) bool {
		addFields(fields,a)
		return true
	})

	b, err := json.MarshalIndent(fields,"",strings.Repeat(" ",h.indent))
	if err != nil {
		return err
	}

	h.w.Write(b)

	return nil
}

func addFields(fields map[string]any, a slog.Attr) {
	value := a.Value.Any()
	if _, ok := value.([]slog.Attr); !ok {
		fields[a.Key] = value
		return
	}

	attrs := value.([]slog.Attr)
	innerFields := make(map[string]any, len(attrs))
	for _, attr := range attrs {
		addFields(innerFields,attr)
	}
	fields[a.Key] = innerFields
}