package main

import "github.com/tidwall/gjson"

const json = `{"name":{"first":"Janet","last":"Prichard"},"age":47}`

func main() {

	jsonstr1 := `{"richtext":{"data":{"items":[{"desc":"想拍出明星范可没有明星的气质怎么办！学学人家明星的出门装扮好伐？狗仔360度偷拍也能保持应有的气质，究竟是怎么办到的？一起学习下。明星出门最重要的是什么？装神秘（bi）！眼睛无神的时候，戴个墨镜。脸上没化妆的时候，戴个口罩，只露出迷人的眼睛！必要时戴个帽子，不仅可以耍帅，还可以遮盖下飞机上酣睡后凌乱的发型。看了上面这些图，你知道怎么提升拍照时的气质了吗？对！就是遮脸！最常见的就是戴个口罩，美丽的大眼睛，很迷人吧？没有口罩的时候，用手挡住半边脸，神秘感倍增。如果你有秀发，利用下，撩人！敷面膜也能让你气质倍增，慵懒的表情配上滑嫩的皮肤，你是个懂得生活的女孩子！你学会怎么拍神秘的蒙面照了吗？点击阅读","end":0,"image":{"key":"d2ed467715d06ae4b80b13e4eaff3eeb","name":"timg (4).jpg","url":"http://nos-yx.netease.com/yixinpublic/pr_0ok_9bi3l3z9b8felxt-jw==_1487148473_212641347"},"linkurl":"http://show.yixin.im/task/taskPage.html?sharesid=f6facd47763e9ae911cefbb036779c64&shareversion=1","start":0,"subTitle":"想拍出明星范可没有明星的气质怎么办！学学人家明星的出门装扮好伐？狗仔360度偷拍也能保持应有的气质，究竟是怎么办到的？一起学习下。明星出门最重要的是什么？装神秘（bi）！眼睛无神的时候，戴个墨镜。脸上没化妆的时候，戴个口罩，只露出迷人的眼睛！必要时戴个帽子，不仅可以耍帅，还可以遮盖下飞机上酣睡后凌乱的发型。看了上面这些图，你知道怎么提升拍照时的气质了吗？对！就是遮脸！最常见的就是戴个口罩，美丽的大眼睛，很迷人吧？没有口罩的时候，用手挡住半边脸，神秘感倍增。如果你有秀发，利用下，撩人！敷面膜也能让你气质倍增，慵懒的表情配上滑嫩的皮肤，你是个懂得生活的女孩子！你学会怎么拍神秘的蒙面照了吗？点击阅读","subsubtype":0,"title":"没有明星的命，但可以有明星的病"}]},"subtype":101}}`
	jsonstr2 := `{"images":[{"md5":"83b25df0edbddcde2769f1855b06b423","size":107241,"url":"http://nos-yx.netease.com/yixinpublic/pr_1qd4af7j14_nfcfpqplwaw==_1489977060_238390862"}],"text":"趁着这个莺飞草长，阳光正好的时节，小伙伴们记得抽空走出城市的喧嚣，郊游、赏花、放风筝，尽情享受四海八荒的浓情春意吧！"}`
	jsonstr3 := `{"audio":{"duration":3189,"md5":"7ac7585a546a171b39a1be95c43d84b0","name":"60f9d8f4-16a6-4eca-b8e6-5e51c39f38691.aac","size":11021,"url":"http://nos-yx.netease.com/yixinpublic/pr_zwx0ir6blam67ccfxn8urw==_1495775574_233545677"},"location":{"coordinate":"28.228272,112.938888","title":"长沙市"},"text":"你好","version":1}`
	jsonstr4 := `{"richtext":{"data":{"items":[{"desc":"惊爆！！！在光棍节这种虐狗的节日到来之际，有一份充满恶意的榜单华丽丽出炉了。咱们look look~此榜单一出，各路好汉躺枪无数！诸位男侠女侠纷纷表示终于找到自己单身的原因了！————阵亡分割线————大侠你忧郁的气质已经出卖了你，这本辟邪剑谱拿去，练了你就有女票了兵哥哥不要难过！你还有基友们！老师，下课留我做作业吧真的菇凉~男朋友是被你红牌罚下了吗这位仁兄，回头看看这榜单，你真的中枪了真的恩，我们的征途，是星（huo）辰(gai)大(dan)海(shen)活捉一只文艺单身汪，打包带走熊孩子一边凉快去……叔叔阿姨还单着呢明天，一个人买单号电影票，一个人玩连连看（消灭一对是一对），做个安静的美男","end":0,"image":{"key":"293634e25fc4c886c957203dc8c9734d","name":"6a3fecd2-f565-4ff1-ba03-943170e5a301 (1).jpg","url":"http://nos-yx.netease.com/yixinpublic/pr_6cntqgf6tlv8yr6eybbbxg==_1447149809_77211016"},"start":0,"subTitle":"惊爆！！！在光棍节这种虐狗的节日到来之际，有一份充满恶意的榜单华丽丽出炉了。咱们look look~此榜单一出，各路好汉躺枪无数！诸位男侠女侠纷纷表示终于找到自己单身的原因了！————阵亡分割线————大侠你忧郁的气质已经出卖了你，这本辟邪剑谱拿去，练了你就有女票了兵哥哥不要难过！你还有基友们！老师，下课留我做作业吧真的菇凉~男朋友是被你红牌罚下了吗这位仁兄，回头看看这榜单，你真的中枪了真的恩，我们的征途，是星（huo）辰(gai)大(dan)海(shen)活捉一只文艺单身汪，打包带走熊孩子一边凉快去……叔叔阿姨还单着呢明天，一个人买单号电影票，一个人玩连连看（消灭一对是一对），做个安静的美男","subsubtype":0,"title":"汪！解开你单身多年的谜团！"}]},"subtype":101}}`

	ParseJson(jsonstr3)
	ParseJson(jsonstr1)
	ParseJson(jsonstr2)
	ParseJson(jsonstr4)

}

func ParseJson(content string) {

	audio := gjson.Get(content, "audio")
	location := gjson.Get(content, "location")
	text := gjson.Get(content, "text")
	images := gjson.Get(content, "images")
	richtext:=gjson.Get(content,"richtext.data.items")


	/*
	richtext:=gjson.Get(content,"richtext")*/

	if audio.Exists() {
		duration := audio.Get("duration").String()
		url := audio.Get("url").String()
		name := audio.Get("name").String()
		println(duration, url, name)
	}

	if location.Exists() {
		coordinate := location.Get("coordinate").String()
		println("纬度", coordinate)
	}

	if text.Exists() {
		textstr := text.String()
		println(textstr)
	}

	if images.Exists() {
		re := images.Array()
		for _, v := range re {
			re := v.Get("url").String()
			println("图片url", re)
		}
	}

	if richtext.Exists(){
		re:= richtext.Array()
		for _, v := range re {
			re := v.Get("desc").String()
			re2:=v.Get("image.url").String()
			println("desc", re)
			println("富文本内置图片URL",re2)
		}
	}
}