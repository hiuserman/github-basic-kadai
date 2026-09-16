package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Entry struct {
	Timestamp time.Time `json:"timestamp"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
}

func saveDiary(entry []Entry) error {
	data, err := json.Marshal(entry) // JSON形式に変換
	if err != nil {
		return fmt.Errorf("JSONの変換に失敗しました: %w", err)
	}

	if err := os.WriteFile("diary.json", data, 0644); err != nil {
		return fmt.Errorf("書き込みに失敗しました: %w", err)
	}
	return nil
}

func loadDiary() ([]Entry, error) {
	data, err := os.ReadFile("diary.json")
	if os.IsNotExist(err) {
		return []Entry{}, nil // ファイルが存在しない場合は空のスライスを返す
	}
	if err != nil {
		return nil, fmt.Errorf("ファイルの読み込みに失敗しました: %w", err)
	}
	var entries []Entry
	err = json.Unmarshal(data, &entries) // JSONデータをEntry構造体に変換
	if err != nil {
		return nil, fmt.Errorf("構造体の変換に失敗しました: %w", err)
	}
	return entries, nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("使用方法: go run main.go [add|list]")
		return
	}

	command := os.Args[1]

	switch command {
	case "add":
		if len(os.Args) < 4 {
			fmt.Println("使用方法: go run main.go add \"タイトル\" \"内容\"")
			return
		}
		title := os.Args[2]
		content := os.Args[3]
		entry := Entry{
			Timestamp: time.Now(),
			Title:     title,
			Content:   content,
		}
		entries, err := loadDiary()
		if err != nil {
			fmt.Println("日記の読み込みに失敗しました:", err)
			return
		}
		entries = append(entries, entry)
		err = saveDiary(entries)
		if err != nil {
			fmt.Println("日記の保存に失敗しました:", err)
			return
		}
	case "list":
		entries, err := loadDiary()
		if err != nil {
			fmt.Println("日記の読み込みに失敗しました:", err)
			return
		}
		for _, entry := range entries {
			fmt.Printf("--- 日記データ ---\n")
			fmt.Printf("日時 :%s\n", entry.Timestamp.Format("2006-01-02 15:04:05"))
			fmt.Printf("タイトル :%s\n", entry.Title)
			fmt.Printf("本文 :\n%s\n", entry.Content)
		}
	default:
		fmt.Println("不明なコマンドです。使用方法: go run main.go [add|list]")
	}
}
