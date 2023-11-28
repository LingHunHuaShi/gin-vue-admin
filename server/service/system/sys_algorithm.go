package system

import (
	"errors"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"gorm.io/gorm"
	"log"
	"os"
	"time"
)

/*
type SysAlgorithm struct {
	global.GVA_MODEL
	AlgorithmID      uint      `json:"algorithmID" gorm:"default:0;primaryKey;autoIncrement;index;comment:算法ID"`
	AlgorithmName    string    `json:"algorithmName" gorm:"comment:算法名"`
	AlgorithmVersion string    `json:"algorithmVersion" gorm:"comment:算法版本"`
	Description      string    `json:"description" gorm:"comment:算法描述"`
	UpdateDate       time.Time `json:"updateDate" gorm:"comment:更新时间"`
	Size             float32   `json:"size" gorm:"comment:模型文件大小"`
	// DownloadLink     string    `json:"downloadLink" gorm:"comment:模型下载链接"`
	MD5              string    `json:"MD5" gorm:"comment:MD5校验值"`
}
*/

type AlgorithmService struct{}

// CreateAlgorithm 算法注册
// 参数 algorithm
// 返回值 被添加的算法 错误信息
func (algorithmService *AlgorithmService) CreateAlgorithm(algorithm system.SysAlgorithm) (algorithmInter system.SysAlgorithm, err error) {
	err = global.GVA_DB.Create(&algorithm).Error
	if err != nil {
		return algorithm, errors.New("创建算法失败")
	}
	return algorithm, nil
}

// DeleteAlgorithm 根据算法ID删除算法及其相关权限信息
// @param id uint 算法ID
// @return err error 删除操作的错误，如果删除失败则返回错误信息，否则为nil
func (algorithmService *AlgorithmService) DeleteAlgorithm(algorithmID uint) (err error) {
	var algorithm system.SysAlgorithm
	err = global.GVA_DB.Where("ID = ?", algorithmID).Delete(&algorithm).Error
	if err != nil {
		return errors.New("删除失败")
	}
	return err
}

// FindAlgorithmById FindAlgorithmByID 根据算法ID查找算法信息
// @param algorithmID uint 要查找的算法ID
// @return algorithmInter *system.SysAlgorithm 查找到的算法信息，如果找到则为非空指针，否则为nil
// @return err error 查找操作的错误，如果算法不存在或查找失败则返回错误信息，否则为nil
func (algorithmService *AlgorithmService) FindAlgorithmById(algorithmID uint) (algorithmInter *system.SysAlgorithm, err error) {
	var algorithm system.SysAlgorithm
	err = global.GVA_DB.Where("ID = ?", algorithmID).First(&algorithm).Error
	if err != nil {
		return &algorithm, errors.New("算法不存在")
	}
	return &algorithm, err
}

// UpdateAlgorithm 更新算法信息
// @param algorithmID uint 要更新的算法ID
// @param newAlgorithm system.SysAlgorithm 新的算法信息
// @return updatedAlgorithmInter *system.SysAlgorithm 更新后的算法信息，如果更新成功则为非空指针，否则为nil
// @return err error 更新操作的错误，如果更新失败则返回错误信息，否则为nil
func (algorithmService *AlgorithmService) UpdateAlgorithm(algorithmID uint, newAlgorithm system.SysAlgorithm) (updatedAlgorithmInter *system.SysAlgorithm, err error) {
	var algorithm system.SysAlgorithm

	// 查找要更新的算法
	err = global.GVA_DB.Where("ID = ?", algorithmID).First(&algorithm).Error
	if err != nil {
		return &algorithm, errors.New("算法不存在")
	}

	// 更新算法信息
	err = global.GVA_DB.Model(&algorithm).Updates(newAlgorithm).Error
	if err != nil {
		return &algorithm, errors.New("更新失败")
	}

	return &algorithm, nil
}

// QueryAllAlgorithm  查询全部算法信息
// @return algorithms []system.SysAlgorithm 查找到的全部算法信息，如果没有算法则返回空切片
// @return err error 查找操作的错误，如果查找失败则返回错误信息，否则为nil
func (algorithmService *AlgorithmService) QueryAllAlgorithm() (algorithms []system.SysAlgorithm, err error) {
	err = global.GVA_DB.Find(&algorithms).Error
	if err != nil {
		return nil, err
	}
	return algorithms, nil
}

// GetAlgorithms  获取云端的算法
// @return error
func (algorithmService *AlgorithmService) GetAlgorithms(algorithms []system.SysAlgorithm) error {
	for _, algorithm := range algorithms {
		var tmp system.SysAlgorithm
		err := global.GVA_DB.Where("ID = ?", algorithm.ID).First(&tmp)

		if err.Error != nil && !errors.Is(err.Error, gorm.ErrRecordNotFound) {
			return err.Error
		}

		if err.Error == gorm.ErrRecordNotFound {
			err := global.GVA_DB.Create(&algorithm)
			if err.Error != nil {
				return err.Error
			}
		} else {
			tmp.AlgorithmName = algorithm.AlgorithmName
			tmp.AlgorithmVersion = algorithm.AlgorithmVersion
			tmp.Description = algorithm.Description
			tmp.UpdateDate = time.Now()
			tmp.MD5 = algorithm.MD5
			tmp.Size = algorithm.Size
			err := global.GVA_DB.Save(&tmp)
			if err.Error != nil {
				return err.Error
			}
		}
	}

	return nil
}

// DownloadAlgorithms  获取云端的算法
// @return error
func (algorithmService *AlgorithmService) DownloadAlgorithms(algorithmID uint) error {
	var algorithm system.SysAlgorithm
	err := global.GVA_DB.First(&algorithm, algorithmID).Error
	if err != nil {
		log.Println("数据库查询失败.")
		return err
	}

	var path string = "/data/yolo/" + algorithm.AlgorithmName

	if _, err := os.Stat(path); os.IsNotExist(err) {
		// 文件夹不存在则创建
		err := os.MkdirAll(path, 0755)
		if err != nil {
			fmt.Println("Error creating directory:", err)
			return err
		}
		fmt.Println("Directory created:", path)
	} else {
		// 文件夹存在 删除创建该文件夹
		err := os.RemoveAll(path)
		if err != nil {
			fmt.Println("Error removing directory:", err)
			return err
		}

		err = os.MkdirAll(path, 0755)
		if err != nil {
			fmt.Println("Error creating directory:", err)
			return err
		}
		fmt.Println("Directory deleted and recreated:", path)
	}

	algorithm.StoragePath = path
	// TO DO DOWNLOAD THE ALGORITHM FORM THE CLOUD.
	// ...
	// ...
	// ...

	algorithm.Downloaded = true
	err = global.GVA_DB.Updates(&algorithm).Error
	if err != nil {
		log.Println("数据库更新失败")
		return err
	}
	return nil
}

// downloadhandle
// do something
// func downloadHandle()
