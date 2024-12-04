package handlers

type CreatePromptRequest struct {
    Title    string `json:"title"`
    Content  string `json:"content"`
    Category string `json:"category"`
}

type RunPromptRequest struct {
    Input string `json:"input"`
}

type VotePromptRequest struct {
    Vote int `json:"vote"` // 1 for upvote, -1 for downvote
}

type ErrorResponse struct {
    Error string `json:"error"`
}

type PromptResponse struct {
    ID        int    `json:"id"`
    Title     string `json:"title"`
    Content   string `json:"content"`
    Category  string `json:"category"`
    Votes     int    `json:"votes"`
    CreatedAt string `json:"created_at"`
    CreatedBy string `json:"created_by"`
}

type ConversationResponse struct {
    ID        int    `json:"id"`
    Messages  string `json:"messages"`
    CreatedAt string `json:"created_at"`
    UserID    string `json:"user_id"`
    PromptID  int    `json:"prompt_id"`
}

type RunPromptResponse struct {
    ConversationID int    `json:"conversation_id"`
    Response       string `json:"response"`
}
