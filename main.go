package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type FunctionNameResponse struct {
	Names []string `json:"names"`
}

type ClaudeResponse struct {
	Result string `json:"result"`
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "使用方法: naming-helper <関数の機能説明>\n")
		os.Exit(1)
	}
	
	description := strings.Join(os.Args[1:], " ")

	prompt := buildPrompt(description)
	
	output, err := runClaudeCode(prompt)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ClaudeCode実行エラー: %v\n", err)
		os.Exit(1)
	}

	names, err := parseClaudeCodeOutput(output)
	if err != nil {
		fmt.Fprintf(os.Stderr, "出力パースエラー: %v\n", err)
		os.Exit(1)
	}

	// JSON形式で出力
	result := FunctionNameResponse{Names: names}
	jsonOutput, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "JSON生成エラー: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(string(jsonOutput))
}

func buildPrompt(description string) string {
	return fmt.Sprintf(`Suggest 5 function names for the following functionality.
Function description: %s

Output JSON only. Do not include any explanations or markdown code blocks.
Output exactly in this format:
{"names": ["functionName1", "functionName2", "functionName3", "functionName4", "functionName5"]}

Naming conventions:
- Use camelCase
- Start with a verb
- Be clear and concise
- Follow common programming conventions`, description)
}

func runClaudeCode(prompt string) (string, error) {
	cmd := exec.Command("claude", "-p", "--output-format", "json", prompt)
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(output), nil
}

func parseClaudeCodeOutput(output string) ([]string, error) {
	// まずClaude APIのレスポンスをパース
	var claudeResp ClaudeResponse
	if err := json.Unmarshal([]byte(output), &claudeResp); err != nil {
		return nil, fmt.Errorf("Claude応答のパースエラー: %v", err)
	}
	
	// 関数名のレスポンスをパース
	var response FunctionNameResponse
	if err := json.Unmarshal([]byte(claudeResp.Result), &response); err != nil {
		return nil, fmt.Errorf("関数名JSONのパースエラー: %v", err)
	}
	
	if len(response.Names) == 0 {
		return nil, fmt.Errorf("関数名が生成されませんでした")
	}
	
	return response.Names, nil
}