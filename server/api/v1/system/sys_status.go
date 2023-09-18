package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/gin-gonic/gin"
)

type SysStatusApi struct{}

func (Api *SysStatusApi) UploadSystemStatus(c *gin.Context) {
	var status system.SysStatus
	err := c.ShouldBindJSON(&status)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = sysStatusService.UploadSystemStatus(status, "")
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithMessage("系统数据成功上传", c)
}

// API提交数据给云端
/*
func (Api *SysStatusApi) SendSystemStatusViaAPI(c *gin.Context) {
	var Status system.SysStatus
	err := c.ShouldBindJSON(&Status)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// API请求
	ApiUrl := "https://..."
	status, err := json.Marshal(Status)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	res, err := http.Post(ApiUrl, "application/json", bytes.NewBuffer(status))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error sending request"})
		return
	}
	defer res.Body.Close()

	// 处理响应
	if res.StatusCode == http.StatusOK {
		// 读取和处理响应数据，可能需要使用JSON解析
		responseData, err := ioutil.ReadAll(res.Body)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error reading response"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"response": string(responseData)})
	} else {
		c.JSON(res.StatusCode, gin.H{"error": "Request failed"})
	}
}
*/
