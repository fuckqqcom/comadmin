package wxd

import (
	"comadmin/model/wx"
	"comadmin/pkg/e"
	"comadmin/tools/utils"
)

/**
job注册，如果第一次id为kong，查询ip是否存在，id存在,且都不为空  注册的时候是在线状态
第一次ip不为空""
*/
func (d Dao) register(job wx.Job) (interface{}, int) {
	//检查id是否存在
	if b, id := d.exist(job.Id, job.IP); b {
		return id, e.Success
	}
	return nil, e.Errors
}

func (d Dao) exist(id, ip string) (bool, string) {
	job := new(wx.Job)
	affect, err := d.engine.Where("id = ?", id).Or("ip = ? ", ip).Get(job)

	if utils.CheckError(err, affect) {
		if job.Id != "" && job.IP != "" {
			return true, job.Id
		} else if job.IP != "" && job.Id == "" {
			//第一次注册
			job.IP = ip
			job.Id = utils.EncodeMd5(ip)
			job.Status = 1 //在线状态
			job.Count = 1
			d.create(job)
			return true, job.Id
		}
		return false, ""
	}
	return false, ""
}

func (d Dao) findJobCount(id, ip string) int {
	job := new(wx.Job)

	affect, err := d.engine.Where(" id = ? ", id).And("ip = ? ", ip).Get(job)
	if utils.CheckError(err, affect) {
		return job.Count
	}
	return 0
}
