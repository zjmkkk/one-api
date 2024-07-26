package relaymode

import "strings"

func GetByPath(path string) int {
    relayMode := Unknown
    
    // 检查是否是 /hf/v1 路径
    isHuggingFace := strings.HasPrefix(path, "/hf/v1")
    
    // 移除 /v1 或 /hf/v1 前缀以统一处理
    path = strings.TrimPrefix(strings.TrimPrefix(path, "/v1"), "/hf/v1")

    switch {
    case strings.HasPrefix(path, "/chat/completions"):
        relayMode = ChatCompletions
    case strings.HasPrefix(path, "/completions"):
        relayMode = Completions
    case strings.HasPrefix(path, "/embeddings") || strings.HasSuffix(path, "embeddings"):
        relayMode = Embeddings
    case strings.HasPrefix(path, "/moderations"):
        relayMode = Moderations
    case strings.HasPrefix(path, "/images/generations"):
        relayMode = ImagesGenerations
    case strings.HasPrefix(path, "/edits"):
        relayMode = Edits
    case strings.HasPrefix(path, "/audio/speech"):
        relayMode = AudioSpeech
    case strings.HasPrefix(path, "/audio/transcriptions"):
        relayMode = AudioTranscription
    case strings.HasPrefix(path, "/audio/translations"):
        relayMode = AudioTranslation
    case strings.HasPrefix(path, "/oneapi/proxy"):
        relayMode = Proxy
    }

    // 如果是 Hugging Face 路径，可以在这里做额外处理
    if isHuggingFace {
        // 可能需要对 Hugging Face 特定的模式进行额外处理
        // 例如：relayMode = HuggingFaceMode(relayMode)
    }

    return relayMode
}
