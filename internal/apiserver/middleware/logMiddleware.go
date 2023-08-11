package middleware

import (
	"fmt"
	"context"
	"io/ioutil"
	"net/http"
	"bytes"
	"github.com/gin-gonic/gin"
	"gitlab.intsig.net/textin-gateway/internal/apiserver/svc"
	"gitlab.intsig.net/textin-gateway/internal/apiserver/types/auth"
	"gitlab.intsig.net/textin-gateway/internal/pkg/optlogs"
)

var logInfoHandlers = map[string]func(*gin.Context) optlogs.OptLogger{
    "/gateway/v1/applications/:id/update":  updateApplicationLogInfo,
    "/gateway/v1/applications/:id/delete":  deleteApplicationLogInfo,
}

func updateApplicationLogInfo(c *gin.Context) optlogs.OptLogger {
	id := c.Param("id")
	req_body := c.Keys["myreq"].(string)
	userId := c.Keys["user"].(auth.UserInfo).ID
	return optlogs.OptLogger{
		Operation:    "update",
		Resource:     fmt.Sprintf(`%s：%s`, "应用ID", id),
		ResourceType: "application",
		UserID:       userId,
		ReqBody:      req_body,
	}
}

func deleteApplicationLogInfo(c *gin.Context) optlogs.OptLogger {
	id := c.Param("id")
	userId := c.Keys["user"].(auth.UserInfo).ID
	return optlogs.OptLogger{
		Operation:    "delete",
		Resource:     fmt.Sprintf(`%s：%s`, "应用ID", id),
		ResourceType: "application",
		UserID:       userId,
	}
}

func LogMiddleware(svcCtx *svc.ServiceContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		// 读取请求体数据
		rawData, _ := c.GetRawData()
		req_body := string(rawData)
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(rawData))
		c.Set("myreq", req_body)
		c.Next()
		path := c.FullPath()
		if c.Request.Method != http.MethodGet {
			// 查找对应的处理函数
			handler, exists := logInfoHandlers[path]
			if exists {
				logInfo := handler(c)
				optlogs.AddOptLogs(context.Background(), svcCtx, logInfo)
			}
		}
	}
}