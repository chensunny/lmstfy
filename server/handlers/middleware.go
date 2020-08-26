package handlers

import (
	"math"
	"net/http"
	"regexp"
	"strings"

	"github.com/bitleak/lmstfy/auth"
	"github.com/bitleak/lmstfy/config"
	"github.com/bitleak/lmstfy/engine"
	"github.com/gin-gonic/gin"
)

func getToken(c *gin.Context) (token string) {
	token = c.GetHeader("X-Token")
	if token == "" {
		token = c.Query("token")
	}
	return
}

// The user token's format is: [{pool}:]{token}
// there is a optional pool prefix, if provided, use that pool; otherwise use the default pool.
func parseToken(rawToken string) (pool, token string) {
	splits := strings.SplitN(rawToken, ":", 2)
	if len(splits) == 2 {
		return splits[0], splits[1]
	}
	return config.DefaultPoolName, rawToken
}

func SetupQueueEngine(c *gin.Context) {
	pool, token := parseToken(getToken(c))
	c.Set("pool", pool)
	c.Set("token", token)
	e := engine.GetEngine(pool)
	if e == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "pool not found"})
		c.Abort()
		return
	}
	c.Set("engine", e)
}

func SetupQueue(c *gin.Context) {
	c.Set("queue", engine.NewQueue(c.Param("queue")))
}

func ValidateToken(c *gin.Context) {
	tk := c.GetString("token")
	if tk == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "token not found"})
		c.Abort()
		return
	}
	tm := auth.GetTokenManager()
	ok, err := tm.Exist(c.GetString("pool"), c.Param("namespace"), tk)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal error"})
		c.Abort()
		return
	}
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
		c.Abort()
		return
	}
}

var paramRegex = regexp.MustCompile("^[-_[:alnum:]]+$")
var queueRegex = regexp.MustCompile(`^[-_[:alnum:]]+(\.(\{low\}|\{normal\}|\{high\}))?$`)
var multiQueuesRegex = regexp.MustCompile("^[-_,[:alnum:]]+$")

// Validate namespace and queue names don't contain any illegal characters
func ValidateParams(c *gin.Context) {
	ns := c.Param("namespace")
	q := c.Param("queue")
	if len(ns) > math.MaxUint8 || len(q) > math.MaxUint8 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "namespace or queue name too long"})
		c.Abort()
		return
	}
	if !paramRegex.MatchString(ns) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "namespace name contains forbidden characters"})
		c.Abort()
		return
	}
	if strings.HasPrefix(q, "_") || !queueRegex.MatchString(q) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "queue name contains forbidden characters"})
		c.Abort()
		return
	}
}

func ValidateMultiConsume(c *gin.Context) {
	ns := c.Param("namespace")
	q := c.Param("queue")
	if !paramRegex.MatchString(ns) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "namespace name contains forbidden characters"})
		c.Abort()
		return
	}
	for _, queue := range strings.Split(q, ",") {
		if queue == "" {
			continue
		}
		if !queueRegex.MatchString(queue) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "queue name contains forbidden characters"})
			c.Abort()
			return
		}
	}
}
