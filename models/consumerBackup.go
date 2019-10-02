package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type ConsumerBackup struct {
	Id                int       `json:"id"`
	ConsumerId        string    `json:"consumer_id"`
	Topic             string    `json:"topic"`
	ConsumerGroup     string    `json:"consumer_group"`
	CallbackUrl       string    `json:"callback_url"`
	Running           bool      `json:"running"`
	Ip                string    `json:"ip"`
	CreatedTime       time.Time `json:"created_time"`
	RestartedTime     time.Time `json:"restarted_time"`
	ConsumeTemplateId int       `json:"consume_template_id"`
}

// 注册orm
func init() {
	orm.RegisterModel(new(ConsumerBackup))
}

// 通过consumerId获取单条备份记录
func (this *ConsumerBackup) GetConsumerBackupByConsumerId(consumerId string) (data ConsumerBackup, err error) {
	o := orm.NewOrm()
	err = o.QueryTable(this).Filter("consumer_id", consumerId).One(&data)
	return
}

// 创建一条新ConsumerBackup数据
func (this *ConsumerBackup) Insert() (int64, error) {
	return orm.NewOrm().Insert(this)
}

// 查询所有的consumerBackUp列表
func (this *ConsumerBackup) GetConsumerBackUpList() (list []*ConsumerBackup, err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable(this).All(&list)
	return
}

// 删除指定的consumerBackUp项
func (this *ConsumerBackup) DeleteByConsumerId(consumerId string) (err error) {
	o := orm.NewOrm()
	_, err = o.QueryTable(this).Filter("consumer_id", consumerId).Delete()
	return
}

// 更新单项数据
func (this *ConsumerBackup) Update() (err error) {
	o := orm.NewOrm()
	_, err = o.Update(this)
	return
}

// 通过template获取所有的存活的消费者
func (this *ConsumerBackup) GetAllLiveConsumerInfoByTemplateId(templateId int) (data []*ConsumerBackup) {
	o := orm.NewOrm()
	o.QueryTable(this).Filter("consume_template_id", templateId).Filter("running", true).All(&data)
	return
}

// 判断消费者模板开启的消费者是否都已经关闭
func (this *ConsumerBackup) IsConsumerShutdownOfTemplateId(templateId int) bool {
	return len(this.GetAllLiveConsumerInfoByTemplateId(templateId)) <= 0
}
