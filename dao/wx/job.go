package wxd

import (
	"comadmin/model/wx"
	"comadmin/tools/utils"
)

/**
job注册，如果第一次id为kong，查询ip是否存在，id存在,且都不为空  注册的时候是在线状态
第一次ip不为空""
*/

func (d Dao) findJobCount(id, ip string) int {
	job := new(wx.Job)

	affect, err := d.engine.Where(" id = ? ", id).And("ip = ? ", ip).Get(job)
	if utils.CheckError(err, affect) {
		return job.Count
	}
	return 0
}
