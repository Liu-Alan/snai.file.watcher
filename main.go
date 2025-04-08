package main

import (
	"fmt"
	"log"

	"github.com/fsnotify/fsnotify"
)

func main() {
	// 创建新的文件监听器
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	// 监听的目录（当前目录）
	dirs := []string{"./log", "../data"}

	// 添加要监视的目录
	for _, dir := range dirs {
		err = watcher.Add(dir)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("监听%s目录是否有新文件...\n", dir)
	}

	// 监听事件
	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return
			}
			// 检测到文件新增
			if event.Op&fsnotify.Create == fsnotify.Create {
				fmt.Println("检测到新文件:", event.Name)
			}

		case err, ok := <-watcher.Errors:
			if !ok {
				return
			}
			fmt.Println("错误:", err)
		}
	}
}
