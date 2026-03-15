🚀 StreamFusion (Working Title)

**The Ultimate Self-Hosted, Open-Source Streaming Command Center.**

StreamFusion is a high-performance, Go-based dashboard that aggregates chat, moderation, and real-time analytics from **Twitch, YouTube, and Kick** into a single, private instance. Stop relying on third-party cloud services—host your own data, customize your own overlays, and own your stream.

---

✨ Features

📡 Unified Multi-Platform Chat

- **Real-time Aggregation:** Combines chat from Twitch, YouTube, and Kick into one high-speed WebSocket feed.
- **Moderator Mode:** A dedicated, low-latency UI to ban, timeout, or delete messages across all platforms simultaneously.
- **Platform-Specific Actions:** Support for platform-unique features like Gifting Subs or SuperChat tracking.

📊 Real-Time Analytics Dashboard

- **Combined Stats:** View your total reach across all platforms in one "Total Live Viewers" metric.
- **Time-Series Insights:** Powered by **InfluxDB**, track viewer growth, chat velocity, and engagement trends over time.
- **Historical Data:** Unlike platform dashboards that reset, you keep your data forever on your own hardware.

🎨 Fully Themeable Overlays

- **OBS Integration:** Clean, dedicated browser source URLs for your chat and alerts.
- **CSS Customization:** Use a built-in web editor to tweak themes with CSS variables—no coding required to make it look professional.
- **Ultra-Lightweight:** Optimized for zero CPU impact on your streaming rig.

🤖 Open Bot Integration

- **Custom Commands:** Create global commands (e.g., `!socials`) that post to all active chats at once.
- **Extensible:** Simple Go-based plugin system to add your own automated moderation logic.

---

🏗️ The Tech Stack

- **Backend:** Go (Golang) — Engineered for high-concurrency and rock-solid stability.
- **Real-time:** Gorilla WebSockets — Sub-millisecond message delivery.
- **Database:**
    - **InfluxDB:** For high-performance time-series stream metrics.
    - **SQLite:** For lightweight, portable configuration and user settings.
- **Frontend:** **React + Tailwind CSS** — Modern, responsive, and easy to theme.
- **Infrastructure:** **Docker & Docker Compose** — Deploy anywhere (VPS, Raspberry Pi, or Local) with one command.

---

🚀 Quick Start (Development)

1. **Clone the repo:**
    
    bash
    
    ```
    git clone https://github.com
    cd streamfusion
    ```
    
    Use code with caution.
    
2. **Configure your Environment:**  
    Copy `.env.example` to `.env` and add your Twitch/YouTube/Kick API credentials.
3. **Spin up the stack:**
    
    bash
    
    ```
    docker-compose up -d
    ```
    
    Use code with caution.
    
4. **Access the Dashboard:**  
    Open `http://localhost:3000` to start your first session.

---

🤝 Contributing

This is an **Open Source** project. We welcome contributors who want to help build the future of independent streaming. Check out our `CONTRIBUTING.md` to get started with the Go backend or the React frontend.

---

📄 License

Distributed under the **MIT License**. See `LICENSE` for more information.
