package tailfile

import (
	"dzug/models"
	"fmt"
	"go.uber.org/zap"
)

// tailTask 的管理者
type tailTaskMgr struct {
	tailTaskMap      map[string]*tailTask    // 所有的tailTask任务
	collectEntryList []*models.LogConfig     // 所有配置项
	confChan         chan []models.LogConfig // 等待新配置的通道
}

var ttMgr *tailTaskMgr

// Init 初始化
func Init(allConf []*models.LogConfig) (err error) {
	// allConf里面存了若干个日志的收集项,
	// 针对每一个日志收集项创建一个对应的tailObj
	ttMgr = &tailTaskMgr{
		tailTaskMap:      make(map[string]*tailTask, 20),
		collectEntryList: allConf,
		confChan:         make(chan []models.LogConfig), // 做一个阻塞channel
	}
	for _, logconf := range allConf {
		tt := newTailTask(logconf.Path, logconf.Topic) // 创建一个日志收集任务
		err = tt.Init()                                // 去打开日志文件准备读
		if err != nil {
			//zap.L().Error("create tailObj for path:%s"+fmt.Sprintf("%s", logconf.Path)+" failed, err:", zap.Error(err))
			fmt.Printf("create tailObj for path:%s failed, err: %s", logconf.Path, err)
			continue
		}
		//zap.L().Info("create a tail task for path:" + fmt.Sprintf("%s", logconf.Path) + "  success")
		fmt.Printf("create a tail task for path: %s success", logconf.Path)
		ttMgr.tailTaskMap[tt.path] = tt // 把创建的这个tailTask任务登记在册,方便后续管理
		// 起一个后台的goroutine去收集日志
		go tt.run()
	}
	go ttMgr.watch() // 在后台等新的配置来
	return
}

// 派一个小弟等着新配置来,
func (t *tailTaskMgr) watch() {
	for {
		//派一个小弟等着新配置来,
		newConf := <-t.confChan // 取到值说明新的配置来啦
		// 新配置来了之后应该管理一下我之前启动的那些tailTask
		//zap.L().Info("get new conf from etcd, conf:" + fmt.Sprintf("%s", newConf) + ", start manage tailTask...")
		fmt.Printf("get new conf from etcd, conf: %s  start manage tailTask...", newConf)

		for _, logconf := range newConf {
			// 1. 原来已经存在的任务就不用动
			if t.isExist(logconf) {
				continue
			}
			// 2. 原来没有的我要新创建一个taiTask任务   （和上面init 的步骤一样）
			tt := newTailTask(logconf.Path, logconf.Topic) // 创建一个日志收集任务
			err := tt.Init()                               // 去打开日志文件准备读
			if err != nil {
				//zap.L().Error("create tailObj for path:"+fmt.Sprintf("%s", logconf.Path)+" failed, err:", zap.Error(err))
				fmt.Printf("create tailObj for path:"+fmt.Sprintf("%s", logconf.Path)+" failed, err:", zap.Error(err))
				continue
			}
			//zap.L().Info("create a tail task for path:" + fmt.Sprintf("%s", logconf.Path) + " success")
			fmt.Printf("create a tail task for path:" + fmt.Sprintf("%s", logconf.Path) + " success")
			t.tailTaskMap[tt.path] = tt // 把创建的这个tailTask任务登记在册,方便后续管理
			// 起一个后台的goroutine去收集日志
			go tt.run()
		}
		// 3. 原来有的，现在没有的要tailTask停掉

		// 找出tailTaskMap中存在,但是newConf不存在的那些tailTask,把它们都停掉
		for key, task := range t.tailTaskMap {
			var found bool
			for _, c := range newConf {
				if key == c.Path {
					found = true
					break
				}
			}
			if !found {
				// 这个tailTask要停掉了
				//zap.L().Info("the task collect path:" + fmt.Sprintf("%s", task.path) + " need to stop.")
				fmt.Printf("the task collect path:" + fmt.Sprintf("%s", task.path) + " need to stop.")
				delete(t.tailTaskMap, key) // 从管理类中删掉
				task.cancel()
			}
		}

	}
}

// 判断tailTaskMap中是否存在该收集项
func (t *tailTaskMgr) isExist(conf models.LogConfig) bool {
	_, ok := t.tailTaskMap[conf.Path]
	return ok
}

func SendNewConf(newConf []models.LogConfig) {
	ttMgr.confChan <- newConf
}
