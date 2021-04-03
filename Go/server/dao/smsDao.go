package dao

import (
	"ChatRoom/Go/server/database"
	"log"
)

/*   mysql */
type MySqlSmsDao struct {
}

const (
	SMS_SET = "insert into sms(user_id, content, created_at) values(?, ?, Now());"
	SMS_GET = "select sms_id, content from sms where user_id = ? and created_at >= subtime(now(), ?) and isnull(deleted_at);"
	SMS_DEL = "update sms set deleted_at = now() where sms_id = ?;"
)

func (msd *MySqlSmsDao) DepositeByID(id int, content string) (err error) {
	err = database.Exec(SMS_SET, id, content)
	return
}

func (msd *MySqlSmsDao) WithdrawByID(id int) ([]string, error) {
	rows, err := database.Query(SMS_GET, id, "12:00:00")
	if err != nil {
		return nil, err
	}
	var (
		smsId   int
		content string
	)

	smsList := make([]string, 0, 20)
	for rows.Next() {
		err = rows.Scan(&smsId, &content)
		if err != nil {
			log.Println(err)
			continue
		}
		err = database.Exec(SMS_DEL, smsId)
		if err != nil {
			log.Println(err)
			continue
		}
		smsList = append(smsList, content)
	}
	return smsList, nil
}

//  redis
/*
type RedisSmsDao struct {
}

func (rsd *RedisSmsDao) DepositeByID(id int, content string) (err error) {
	// 将数据存入mesList[userId]中
	err = cache.RedisLpush("smsList"+strconv.Itoa(id), content)
	if err != nil {
		return err
	}
	// 退出
	return
}

func (rsd *RedisSmsDao) WithdrawByID(id int) (dataSlice []string, err error) {
	// 将数据存入mesList[userId]中
	dataSlice, err = cache.RedisGetList("smsList" + strconv.Itoa(id))
	log.Println(dataSlice)
	if err != nil {
		return
	}

	// 如果留言数量不为零
	if len(dataSlice) != 0 {
		err = cache.RedisDel("smsList" + strconv.Itoa(id))
		if err != nil {
			log.Println(err.Error())
		}
	}

	// 退出
	return
}
*/
// end redis
