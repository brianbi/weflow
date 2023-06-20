package service

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/wegoteam/weflow/pkg/common/entity"
	"reflect"
	"testing"
)

type TsetParam struct {
	Param1 string
	Param2 int64
	Param3 float64
	Param4 []string
}

func TestGetParamType(t *testing.T) {

	var instTaskParamMap = make(map[string]interface{})
	instTaskParamMap["testparam1"] = 1
	instTaskParamMap["testparam2"] = 1.22222222222222222222222222222222222222
	instTaskParamMap["testparam3"] = 22222222222222222
	instTaskParamMap["testparam4"] = "testparam4"
	instTaskParamMap["testparam5"] = 1.2222

	var slice = make([]string, 0)
	var slice2 = make([]TsetParam, 0)
	slice = append(slice, "test1")

	instTaskParamMap["testparam6"] = slice

	var tsetParam = &TsetParam{
		Param1: "test1",
		Param2: 1,
		Param3: 22222222222,
		Param4: slice,
	}
	instTaskParamMap["testparam7"] = tsetParam

	slice2 = append(slice2, *tsetParam)
	instTaskParamMap["testparam8"] = slice2
	for key, val := range instTaskParamMap {
		paramType := GetParamType(val)
		t := reflect.TypeOf(val).String()

		kind := reflect.ValueOf(val).Kind()
		hlog.Infof("val 的类型是 %v kind=%v", t, kind)
		hlog.Infof("key=%v  val=%v   valType=%v", key, val, paramType)
	}

}

func TestGetInstTaskParam(t *testing.T) {
	instTaskParamMap := GetInstTaskParam("421397709668421")
	hlog.Infof("instTaskParamMap是 %v", instTaskParamMap)
}

func TestGetModelVersion(t *testing.T) {
	modelVersion := GetModelVersion("420915317174341", "1681335332954505235")
	hlog.Infof("GetModelVersion= %v", modelVersion)

	modelVersionList := GetModelVersionList("420915317174341", "1681335332954505235")
	hlog.Infof("GetModelVersionList= %v", modelVersionList)

	modelVersion2 := GetEnableModelVersion("420915317174341")
	hlog.Infof("GetEnableModelVersion= %v", modelVersion2)
}

func TestGetInstNodeUserTask(t *testing.T) {
	instNodeUserTask := GetInstNodeUserTask("424136865722437")
	hlog.Infof("instNodeUserTask= %v", instNodeUserTask)
}

func TestGetExecNodeTaskMap(t *testing.T) {
	execNodeTaskMap := GetExecNodeTaskMap("425247133954117")
	hlog.Infof("execNodeTaskMap= %v", execNodeTaskMap)
}

func TestGetTodoUserTask(t *testing.T) {
	param := &entity.UserTaskQueryBO{
		PageNum:         2,
		PageSize:        1,
		InstStatus:      2,
		CreateTimeStart: "2020-8-5 13:14:15",
		CreateTimeEnd:   "2024-8-5 13:14:15",
	}
	userTask := GetTodoUserTasks("547")
	page := PageTodoUserTasks("547", param)
	hlog.Infof("userTask= %v", userTask)
	hlog.Infof("page= %v", page)
}

func TestGetDoneUserTask(t *testing.T) {
	userTask := GetDoneUserTasks("547")
	hlog.Infof("userTask= %v", userTask)
}

func TestGetInitiatingInstTask(t *testing.T) {
	initiatingInstTasks := GetInitiatingInstTasks("547")
	hlog.Infof("initiatingInstTasks= %v", initiatingInstTasks)
}

func TestGetDraftInstTask(t *testing.T) {
	draftInstTask := GetDraftInstTask("547")
	hlog.Infof("draftInstTask= %v", draftInstTask)
}

func TestGetModelList(t *testing.T) {
	modelList := GetModelList()
	hlog.Infof("modelList= %v", modelList)
}
