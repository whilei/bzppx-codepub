package models

import (
	"bzppx-codepub/app/utils"

	"github.com/snail007/go-activerecord/mysql"
)

const Table_Task_Name = "task"

type Task struct {
}

var TaskModel = Task{}

// 根据 task_id 获取任务
func (t *Task) GetTaskByTaskId(taskId string) (tasks map[string]string, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_Task_Name).Where(map[string]interface{}{
		"task_id": taskId,
	}))
	if err != nil {
		return
	}
	tasks = rs.Row()
	return
}

// 根据 user_id 获取任务
func (t *Task) GetTasksByUserId(userId string) (tasks []map[string]string, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_Task_Name).Where(map[string]interface{}{
		"user_id": userId,
	}))
	if err != nil {
		return
	}
	tasks = rs.Rows()
	return
}

// 根据 task_ids 获取任务
func (t *Task) GetTaskByTaskIds(taskIds []string) (tasks []map[string]string, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_Task_Name).Where(map[string]interface{}{
		"task_id": taskIds,
	}))
	if err != nil {
		return
	}
	tasks = rs.Rows()
	return
}

// 插入一条任务
func (l *Task) Insert(task map[string]interface{}) (id int64, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Exec(db.AR().Insert(Table_Task_Name, task))
	if err != nil {
		return
	}
	id = rs.LastInsertId
	return
}

// 通过project_id和task_id查找task
func (l *Task) GetTaskByProjectIdsAndTaskIds(projectIds, taskIds []string) (task []map[string]string, err error) {
	db := G.DB()
	where := make(map[string]interface{})
	where["task_id"] = taskIds
	var rs *mysql.ResultSet
	if len(projectIds) > 0 {
		where["project_id"] = projectIds
	}
	rs, err = db.Query(db.AR().From(Table_Task_Name).Where(where))
	if err != nil {
		return
	}
	task = rs.Rows()
	return
}

func (l *Task) GetTaskByProjectId(projectId string, limit, number int) (tasks []map[string]string, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_Task_Name).Where(map[string]interface{}{
		"project_id": projectId,
	}).Limit(limit, number).OrderBy("task_id", "DESC"))

	if err != nil {
		return
	}
	tasks = rs.Rows()
	return
}

func (l *Task) GetTaskByProjectIdAndUserId(projectId, userId string, limit, number int) (tasks []map[string]string, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_Task_Name).Where(map[string]interface{}{
		"project_id": projectId,
		"user_id":    userId,
	}).Limit(limit, number).OrderBy("task_id", "DESC"))

	if err != nil {
		return
	}
	tasks = rs.Rows()
	return
}

func (l *Task) CountTaskByProjectId(projectId string) (count int64, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().Select("count(*) as total").From(Table_Task_Name).Where(map[string]interface{}{
		"project_id": projectId,
	}))

	if err != nil {
		return
	}
	count = utils.NewConvert().StringToInt64(rs.Value("total"))
	return
}

func (l *Task) CountTaskByProjectIdAndUserId(projectId, userId string) (count int64, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().Select("count(*) as total").From(Table_Task_Name).Where(map[string]interface{}{
		"project_id": projectId,
		"user_id":    userId,
	}))

	if err != nil {
		return
	}
	count = utils.NewConvert().StringToInt64(rs.Value("total"))
	return
}

func (l *Task) GetTaskByProjectIdNoLimit(projectId string) (tasks []map[string]string, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_Task_Name).Where(map[string]interface{}{
		"project_id": projectId,
	}).OrderBy("task_id", "DESC"))

	if err != nil {
		return
	}
	tasks = rs.Rows()
	return
}

func (l *Task) GetTasksByLimit(limit, number int) (tasks []map[string]string, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_Task_Name).Limit(limit, number).OrderBy("task_id", "DESC"))
	if err != nil {
		return
	}
	tasks = rs.Rows()
	return
}

func (l *Task) GetTasksByUserIdsAndProjectIdsAndLimit(userName, projectName string, userIds, projectIds []string, limit, number int) (tasks []map[string]string, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	where := make(map[string]interface{})
	if userName != "" {
		where["user_id"] = userIds
	}
	if projectName != "" {
		where["project_id"] = projectIds
	}
	rs, err = db.Query(db.AR().From(Table_Task_Name).Where(where).Limit(limit, number).OrderBy("task_id", "DESC"))

	if err != nil {
		return
	}
	tasks = rs.Rows()
	return
}

//获取task的数量
func (l *Task) CountTask() (count int64, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().Select("count(*) as total").From(Table_Task_Name))

	if err != nil {
		return
	}
	count = utils.NewConvert().StringToInt64(rs.Value("total"))
	return
}

//通过 user_id 和 project_id 获取task的数量
func (l *Task) CountTaskByUserIdsAndProjectIds(userName, projectName string, userIds, projectIds []string) (count int64, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	where := make(map[string]interface{})
	if userName != "" {
		where["user_id"] = userIds
	}
	if projectName != "" {
		where["project_id"] = projectIds
	}
	rs, err = db.Query(db.AR().Select("count(*) as total").From(Table_Task_Name).Where(where))
	if err != nil {
		return
	}
	count = utils.NewConvert().StringToInt64(rs.Value("total"))
	return
}

func (l *Task) GetAllTask() (tasks []map[string]string, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	rs, err = db.Query(db.AR().From(Table_Task_Name))
	if err != nil {
		return
	}
	tasks = rs.Rows()
	return
}

func (t *Task) GetProjectIdsOrderByCountProjectLimit(limit int) (tasks []map[string]string, err error) {
	db := G.DB()
	var rs *mysql.ResultSet
	sql := db.AR().Select("project_id, count('project_id') as total").
		From(Table_Task_Name).
		GroupBy("project_id").
		OrderBy("total", "DESC").
		Limit(0, limit)
	rs, err = db.Query(sql)
	if err != nil {
		return
	}
	tasks = rs.Rows()
	return
}

// 根据创建时间获取 task 数量
func (l *Task) CountTaskAndUserByCreateTime(startTime int64, endTime int64) (total int64, userTotal int64, err error) {

	db := G.DB()
	var rs *mysql.ResultSet
	sql := db.AR().Select("count(*) as total, count(distinct user_id) as user_total").
		From(Table_Task_Name).
		Where(map[string]interface{}{
			"create_time >= ": startTime,
			"create_time < ":  endTime,
		})
	rs, err = db.Query(sql)
	if err != nil {
		return
	}
	res := rs.Row()
	if len(res) > 0 {
		total = utils.NewConvert().StringToInt64(res["total"])
		userTotal = utils.NewConvert().StringToInt64(res["user_total"])
	}
	return
}
