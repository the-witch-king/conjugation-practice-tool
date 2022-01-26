package main

import (
    "encoding/json"
)

type ToolConfig struct {
    WordsPerReview  int  `json:"wordsPerReview"`
}

func ParseConfig(configData []byte) ToolConfig {
    var toolConfig ToolConfig
    json.Unmarshal(configData, &toolConfig)
    return toolConfig
}
