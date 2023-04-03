# 一、基本使用



```markdown
1. Go中的没有单独的字符类型，只能用byte，int来存储字符


2. byte只能存储ACSII表中的字符，如果存中文用int

	var c1 int = '杯'
	var c2 int = '子'
	fmt.Printf("%c %c", c1, c2)//安装字符输出
	fmt.Println(c1,c2)//输出unicode编码
	
	
3. Go的默认编码全部统一为UTF-8
```



