import { useChat } from '../../hooks/useChat'
import ChatMessage from './ChatMessage'

export default function Dashboard() {
  const messages = useChat('/ws/dashboard')

  return (
    <div className="min-h-screen bg-[var(--background)] text-[var(--text)] font-[var(--font-family)]">
      <header className="bg-[var(--surface)] border-b border-[var(--primary)] px-6 py-3">
        <h1 className="text-xl font-bold text-[var(--accent)]">StreamFusion — Dashboard</h1>
      </header>
      <main className="flex flex-col-reverse gap-1 p-4 overflow-y-auto h-[calc(100vh-56px)]">
        {messages.map((msg) => (
          <ChatMessage key={msg.id} message={msg} />
        ))}
      </main>
    </div>
  )
}
