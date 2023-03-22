package main

import (
	"context"
	"fmt"

	openai "github.com/sashabaranov/go-openai"
)

func main() {
	client := openai.NewClient("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") //API key
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					//Content: "你：一公斤是多少磅？\n马夫:又是这个？一公斤有 2.2 磅。请注意这一点。你:HTML代表什么?\nMarv:谷歌太忙了吗?超文本标记语言。T代表将来尝试提出更好的问题。\n你:第一架飞机是什么时候飞的?\nMarv:17年1903月<>日，威尔伯和奥维尔·赖特进行了首次飞行。我希望他们来把我带走。\n你:生命的意义是什么?\n马夫:我不确定。我会问我的朋友谷歌。\n你:为什么天空是蓝色的？", //文本内容
				    //Content: "我希望你充当 JavaScript console。我将键入命令，您将回复 JavaScript console 应显示的内容。我希望你只回复一个唯一代码块中的终端输出，没有别的。不要写注释。除非我指示你这样做，否则不要键入命令。当我需要用英语告诉你一些事情时，我会通过将文本放在大括号内{像这样}来做到这一点。我的第一个命令是 console.log（“Hello World”）;",
				    Content: "请做一个实体抽取任务，从下面这段话中提取出人名和地名，并用json格式输出：刘亦菲（ Crystal Liu，1987 年 8 月 25 日－[1]）为 华裔美籍的女演员、歌手。出生于湖北武汉，后随母亲移居美国纽约，毕业于北京电影学院表演系 2002 级本科班。2002 年，主演电视剧《金粉世家》进入演艺圈；之后几年相继饰演了《天龙八部》的王语嫣、《仙剑奇侠传》的赵灵儿和《神雕侠侣》的小龙女等知名角色。2005 年签约唱片公司进军歌坛。2008 年凭借好莱坞电影《功夫之王》亮相国际银幕 [4]。2012 年主演电影《 铜雀台》，斩获第五届澳门国际电影节金莲花奖最佳女主角 [5]。2017 年从 迪士尼《花木兰》真人版电影试镜中脱颖而出，成为花木兰的饰演者[3]。",
				},
			},
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return
	}

	fmt.Println(resp.Choices[0].Message.Content)
}
