package system

import (
	"archive/tar"
	"bytes"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/gin-gonic/gin"
	"github.com/klauspost/compress/gzip"
	"go.uber.org/zap"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

type AlgorithmApi struct{}

// CreateAlgorithm 算法注册
// 参数 algorithm
// 返回值 被添加的算法 错误信息
func (s *AlgorithmApi) CreateAlgorithm(c *gin.Context) {
	var sysAlgorithm system.SysAlgorithm
	err := c.ShouldBindJSON(&sysAlgorithm)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	_, err = algorithmService.CreateAlgorithm(sysAlgorithm)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteAlgorithm 根据算法ID删除算法及其相关权限信息
// @param id uint 算法ID
// @return err error 删除操作的错误，如果删除失败则返回错误信息，否则为nil
func (s *AlgorithmApi) DeleteAlgorithm(c *gin.Context) {
	var algorithm system.SysAlgorithm
	err := c.ShouldBindJSON(&algorithm)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = algorithmService.DeleteAlgorithm(algorithm.ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败!"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功!", c)
}

// FindAlgorithmById FindAlgorithmByID 根据算法ID查找算法信息
// @param algorithmID uint 要查找的算法ID
// @return algorithmInter *system.SysAlgorithm 查找到的算法信息，如果找到则为非空指针，否则为nil
// @return err error 查找操作的错误，如果算法不存在或查找失败则返回错误信息，否则为nil
func (s *AlgorithmApi) FindAlgorithmById(c *gin.Context) {

	var algorithm system.SysAlgorithm
	err := c.ShouldBindJSON(&algorithm)
	if err != nil {
		response.FailWithMessage("查找失败!", c)
		return
	}
	algo, err := algorithmService.FindAlgorithmById(algorithm.ID)
	//algorithmIDStr := c.Param("algorithmID")
	//log.Println(algorithmIDStr)
	//algorithmID, err := strconv.ParseUint(algorithmIDStr, 10, 64)
	//if err != nil {
	//	response.FailWithMessage("算法ID参数错误!", c)
	//	return
	//}
	//
	//algo, err := algorithmService.FindAlgorithmById(uint(algorithmID))

	/*
		algorithmIDStr := c.Param("algorithmID")
		log.Println(algorithmIDStr)
		algorithmID, err := strconv.ParseUint(algorithmIDStr, 10, 64)
		if err != nil {
			response.FailWithMessage("算法ID参数错误!", c)
			return
		}
	*/

	algo, err = algorithmService.FindAlgorithmById(algorithm.ID)

	if err != nil {
		global.GVA_LOG.Error("查找失败!2", zap.Error(err))
		response.FailWithMessage("查找失败!3"+err.Error(), c)
		return
	}
	response.OkWithDetailed(algo, "查找成功!", c)
}

// UpdateAlgorithm 更新算法信息
// @param algorithmID uint 要更新的算法ID
// @param newAlgorithm system.SysAlgorithm 新的算法信息
// @return updatedAlgorithmInter *system.SysAlgorithm 更新后的算法信息，如果更新成功则为非空指针，否则为nil
// @return err error 更新操作的错误，如果更新失败则返回错误信息，否则为nil
func (s *AlgorithmApi) UpdateAlgorithm(c *gin.Context) {
	var algorithm system.SysAlgorithm
	err := c.ShouldBindJSON(&algorithm)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	newAlgo, err := algorithmService.UpdateAlgorithm(algorithm.ID, algorithm)
	if err != nil {
		global.GVA_LOG.Error("更新失败2 GVA_LOG!", zap.Error(err))
		response.FailWithMessage("更新失败2 RESPONSE!", c)
		return
	}
	response.OkWithDetailed(newAlgo, "更新成功!", c)
}

// QueryAllAlgorithm  查询全部算法信息
// @return algorithms []system.SysAlgorithm 查找到的全部算法信息，如果没有算法则返回空切片
// @return err error 查找操作的错误，如果查找失败则返回错误信息，否则为nil
func (s *AlgorithmApi) QueryAllAlgorithm(c *gin.Context) {

	algos, err := algorithmService.QueryAllAlgorithm()
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败!1"+err.Error(), c)
		return
	}
	response.OkWithDetailed(algos, "查询成功!", c)
}

// DownLoadAlgorithms  下载云端的算法
func (s *AlgorithmApi) DownLoadAlgorithms(c *gin.Context) {
	var SyncAlgorithm system.SysSyncAlgorithm
	err := c.ShouldBindJSON(&SyncAlgorithm)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	var url = "http://192.168.x.x:8888/files" + SyncAlgorithm.DownLoadLink
	var path = "/data/yolo/" + SyncAlgorithm.AlgorithmName
	if _, err := os.Stat(path); os.IsNotExist(err) {
		// 如果文件夹不存在，则创建
		err := os.MkdirAll(path, 0755)
		if err != nil {
			log.Println("创建文件夹失败:", err)
			response.FailWithMessage("无法创建路径", c)
			return
		}
	}
	resp, err := http.Get(url)
	if err != nil {
		response.FailWithMessage("无法下载文件", c)
		return
	}
	defer resp.Body.Close()
	var buf bytes.Buffer
	_, err = io.Copy(&buf, resp.Body)
	if err != nil {
		response.FailWithMessage("无法读取下载内容", c)
		return
	}

	gr, err := gzip.NewReader(&buf)
	if err != nil {
		response.FailWithMessage("无法读取下载内容", c)
		return
	}
	defer gr.Close()

	tr := tar.NewReader(gr)

	// 解压缩并保存文件到指定目录
	for {
		header, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			response.FailWithMessage("解压错误", c)
			return
		}

		// 目标文件路径
		targetFile := filepath.Join(path, header.Name)

		// 确保目录已创建
		if header.Typeflag == tar.TypeDir {
			if err := os.MkdirAll(targetFile, os.FileMode(header.Mode)); err != nil {
				response.FailWithMessage("创建目录失败", c)
				return
			}
			continue
		}

		// 创建文件并写入内容
		file, err := os.OpenFile(targetFile, os.O_CREATE|os.O_RDWR, os.FileMode(header.Mode))
		if err != nil {
			response.FailWithMessage("创建文件失败", c)
			return
		}
		defer file.Close()

		if _, err := io.Copy(file, tr); err != nil {
			response.FailWithMessage("写入文件失败", c)
			return
		}
	}

	var algorithm system.SysAlgorithm
	algorithm.AlgorithmVersion = SyncAlgorithm.AlgorithmVersion
	algorithm.AlgorithmName = SyncAlgorithm.AlgorithmName
	algorithm.MD5 = SyncAlgorithm.MD5
	algorithm.Size = SyncAlgorithm.Size
	algorithm.UpdatedAt = SyncAlgorithm.UpdateDate
	algorithm.Description = SyncAlgorithm.Description
	_, err = algorithmService.CreateAlgorithm(algorithm)
	if err != nil {
		response.FailWithMessage("算法创建失败", c)
	}
	response.OkWithMessage("算法下载成功", c)
}
