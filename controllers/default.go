package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"math/rand"
	"moneyBag/models"
	"moneyBag/mysql"
	"strings"
	"time"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.TplName = "add.html"
}

func (c *MainController) Index() {
	c.Data["Website"] = "大脸猫日记"
	c.Data["Email"] = "739221814@qq.com"
	c.Data["Title"] = GetMotto()
	c.TplName = "index.tpl"
}

func (c *MainController) Post() {
	pwd := c.GetString("pwd")
	user := c.GetString("user")
	amount := c.GetString("amount")
	incomePay, _ := c.GetInt("incomePay")
	source := c.GetString("source")
	use := c.GetString("use")
	c.Data["Website"] = "返回大脸猫日记"
	c.Data["Title"] = GetMotto()
	c.Data["章鱼哥"] = "添加成功"
	//验证
	if pwd != "901224" {
		c.Data["Mark"] = "密钥错误"
		c.TplName = "index.tpl"
		return
	}
	if amount == "0" {
		c.Data["Mark"] = "金额错误!"
		c.TplName = "index.tpl"
		return
	}
	var mark string
	if incomePay == 1 {
		if source == "" {
			c.Data["Mark"] = "请填写收入来源!"
			c.TplName = "index.tpl"
			return
		}
		mark = "收入"
		use = ""
		c.Data["Mark"] = "哈哈哈哈哈哈哈哈,又收入了一笔!!!"
	} else {
		if use == "" {
			c.Data["Mark"] = "请填写支出用途!"
			c.TplName = "index.tpl"
			return
		}
		c.Data["Mark"] = "555555555,又支出了一笔..."
		mark = "支出"
		source = ""
		amount = "-" + amount
	}
	describe := strings.Trim(c.GetString("describe"), " ")
	isMust, _ := c.GetInt("isMust")
	bag := models.MoneyBag{User: user, Amount: amount, IncomePay: incomePay, Source: source, Use: use, Mark: mark, Describe: describe, IsMust: isMust}
	fmt.Println(bag)
	mysql.Engine.Insert(&bag)
	c.Data["Email"] = "添加成功"
	c.TplName = "index.tpl"
}

//随机名言
func GetMotto() string {
	str := ""
	list := []string{"你知道我最喜欢什么吗？我最喜欢呵护你。",
		"你这个坏人，为什么要害我？为什么要害我这么喜欢你？",
		"我觉得你今天有点怪，怪好看的。",
		"你是可爱的女孩，我就是可爱本人。",
		"你最近一定是变胖了，在我心的分量都变重了。",
		"你知道什么酒最好喝吗？你和我的天长地久。",
		"猪撞树上了，你撞在我的心里了。",
	"你一定是碳酸饮料成精了，要不然为什么我每次见到你都开心得要冒泡呢？",
	"你的眼睛好漂亮，但是我的眼睛更漂亮，因为我的眼睛里有你。",
	"你知道牛肉要怎么吃才最好吃吗？牛肉我喂你吃最好吃。",
	"你的酒窝没有酒，我却醉的像条狗。",
	"我想在老的时候，吻你光光的牙床。",
	"你是甜筒吗？为什么我想舔遍你全身。",
	"想送你很多很多口红，让你每天还我一点点。"}
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(len(list))
	str = list[n]
	return str
}
