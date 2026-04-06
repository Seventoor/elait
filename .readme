# Project Name

A real-time audio streaming application with live transcription and translation powered by Deepgram, Google Cloud, and OpenAI.

---

## Prerequisites

- [Go](https://golang.org/dl/)
- [Node.js](https://nodejs.org/) + [pnpm](https://pnpm.io/) or npm
- A [Deepgram](https://deepgram.com/) account
- A [Google Cloud](https://console.cloud.google.com/) account
- An [OpenAI](https://platform.openai.com/) account

---

## Backend Setup

**1. Install dependencies**
```bash
cd backend
go mod tidy
```

**2. Create your `.env` file**
```bash
cp .env.sample .env
```

---

## API Keys Setup

### Deepgram
1. Sign up at [deepgram.com](https://deepgram.com/)
2. Go to **API Keys** in the dashboard
3. Click **Create a New API Key**
4. Copy the key into your `.env` as `DEEPGRAM_API_KEY`

---

### Google Cloud
1. Go to [console.cloud.google.com](https://console.cloud.google.com/)
2. Create a new project or select an existing one
3. Enable the following APIs:
   - **Cloud Speech-to-Text API**
   - **Cloud Translation API**
4. Create a Service Account:
   - Go to **IAM & Admin → Service Accounts**
   - Click **Create Service Account**
   - Give it a name and click **Create and Continue**
   - Assign the following roles:
     - `Cloud Speech Administrator`
     - `Cloud Translation API User`
   - Click **Done**
5. Generate a JSON key:
   - Click on the created service account
   - Go to the **Keys** tab
   - Click **Add Key → Create new key → JSON**
   - Download the file and place it in your backend directory
6. Set the path in your `.env` as `GOOGLE_APPLICATION_CREDENTIALS=./your-key-file.json`

---

### OpenAI
1. Go to [platform.openai.com](https://platform.openai.com/)
2. Navigate to **API Keys**
3. Click **Create new secret key**
4. Copy the key into your `.env` as `OPENAI_API_KEY`

---

## Frontend Setup

**1. Install dependencies**
```bash
cd frontend
pnpm install
# or
npm install
```

**2. Create your `.env` file**
```bash
cp .env.sample .env
```

---

## Running the App

Start the backend:
```bash
cd backend
go run .
```

Start the frontend:
```bash
cd frontend
pnpm dev
```

---

## Testing

Open two browser tabs:

| URL | Description |
|-----|-------------|
| `http://localhost:8080/` | Audio source — press **Start** to begin streaming |
| `http://localhost:5173/streams/616` | Audio player — select a language and press **Play** |

Start the source first, then connect the player.