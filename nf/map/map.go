package main

import (
	"fmt"
)

func main() {
	//初始化
	editorSign := map[string]bool{"liteIDE": true, "Notepad": false}
	//更新或增加
	// editorSign["Vim"] = true

	//查询
	sign1, ok := editorSign["Vim"]

	if ok {
		fmt.Println("Vim存在", sign1)
	} else {
		fmt.Println("Vim不存在")
	}

	//删除
	// delete(editorSign, "Vim")

}
