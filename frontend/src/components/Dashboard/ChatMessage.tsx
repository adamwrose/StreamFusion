import type { ChatMessage as ChatMessageType } from '../../types'

interface Props {
  message: ChatMessageType
}

const PLATFORM_COLORS: Record<string, string> = {
  twitch:  '#6441a5',
  youtube: '#ff0000',
  kick:    '#53fc18',
}

export default function ChatMessage({ message }: Props) {
  const platformColor = PLATFORM_COLORS[message.platform] ?? 'var(--text-muted)'

  return (
    <div className="flex items-start gap-2 rounded-[var(--border-radius)] px-3 py-1 hover:bg-[var(--surface)] transition-colors">
      <span
        className="text-xs font-bold uppercase mt-0.5"
        style={{ color: platformColor }}
      >
        {message.platform}
      </span>
      <span className="font-semibold" style={{ color: message.color || 'var(--accent)' }}>
        {message.display_name}:
      </span>
      <span className="text-[var(--text)]">{message.message}</span>
    </div>
  )
}
