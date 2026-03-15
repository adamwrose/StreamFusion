import { useChat } from '../../hooks/useChat'
import type { ChatMessage } from '../../types'

// Overlay is designed to be used as an OBS browser source.
// Keep it transparent-background-friendly.
export default function Overlay() {
  const messages = useChat('/ws/overlay')
  const recent = messages.slice(0, 8)

  return (
    <div className="flex flex-col-reverse gap-1 p-2 w-full">
      {recent.map((msg: ChatMessage) => (
        <div key={msg.id} className="flex items-center gap-2 text-sm drop-shadow-md">
          <span className="font-bold" style={{ color: msg.color || 'var(--accent)' }}>
            {msg.display_name}:
          </span>
          <span className="text-white">{msg.message}</span>
        </div>
      ))}
    </div>
  )
}
