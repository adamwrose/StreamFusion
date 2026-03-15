// Mirrors the Go models.ChatMessage struct exactly.
export interface ChatMessage {
  id:           string
  platform:     string
  username:     string
  display_name: string
  message:      string
  color:        string
  badges:       string[]
  is_mod:       boolean
  timestamp:    string
}
