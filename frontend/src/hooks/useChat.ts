import { useEffect, useRef, useState } from 'react'
import type { ChatMessage } from '../types'

const MAX_MESSAGES = 200

export function useChat(path: string): ChatMessage[] {
  const [messages, setMessages] = useState<ChatMessage[]>([])
  const wsRef = useRef<WebSocket | null>(null)

  useEffect(() => {
    const protocol = window.location.protocol === 'https:' ? 'wss' : 'ws'
    const url = `${protocol}://${window.location.host}${path}`
    const ws = new WebSocket(url)
    wsRef.current = ws

    ws.onmessage = (event) => {
      const msg: ChatMessage = JSON.parse(event.data)
      setMessages((prev) => [msg, ...prev].slice(0, MAX_MESSAGES))
    }

    return () => ws.close()
  }, [path])

  return messages
}
