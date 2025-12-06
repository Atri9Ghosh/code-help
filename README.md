# Code Helper Backend (Go Fiber)

A simple backend service built using Go Fiber that provides code execution simulation, auto-formatting, and coding help features.

## ✨ Features

- **Running Code** — Simulates code execution with smart output handling
- **Auto-fixing Code** — Automatically formats and fixes basic code issues
- **Coding Help** — Provides intelligent tips based on coding keywords
- **Built-in HTML Interface** — Easy-to-use web page for testing all APIs

## 🛠 How to Run the Project

```bash
git clone https://github.com/<your-github-username>/code-helper-backend.git
cd code-helper-backend
go mod tidy
go run main.go
```

**Server starts at:**

👉 [http://localhost:3000](http://localhost:3000)

Open this in your browser — a small HTML page is included to test all APIs easily.

## 📡 API Routes

### 1️⃣ POST `/run`

Simulates executing code.

**Request:**
```json
{
  "code": "print('hello')"
}
```

**Behavior:**
- If code contains `print` → returns simulated output
- If code contains `error` → returns simulated compilation error
- Otherwise → returns a generic "executed successfully" message

**Example Response:**
```json
{
  "output": "hello",
  "status": "success"
}
```

---

### 2️⃣ POST `/autofix`

Auto-formats code using simple rules.

**Request:**
```json
{
  "code": "if(x==1){console.log('hi')}"
}
```

**Auto-fix Rules Applied:**
- Fix indentation (2 spaces per level)
- Add missing semicolons
- Remove extra spaces
- Normalize spacing inside lines
- Fix simple `{}` brace nesting
- Add closing braces if missing

**Example Response:**
```json
{
  "fixed_code": 
  "if (x==1){
     console.log("hi");
}"
```

---

### 3️⃣ POST `/help`

Returns predefined coding tips based on keyword matching.

**Request:**
```json
{
  "query": "syntax error in loop"
}
```

**Keyword → Tip Mapping:**

| Keyword | Help Tip |
|---------|----------|
| `loop`, `for` | Prevent infinite loops; ensure loop variable updates |
| `if`, `condition` | Check your condition logic and edge cases |
| `error`, `bug` | Read logs and reduce code to reproduce issues |
| `syntax`, `semicolon` | Check semicolons and bracket balancing |
| (no match) | Shows a generic guidance message |

**Example Response:**
```json
{
  "tip": "Prevent infinite loops; ensure loop variable updates"
}
```

---

## 🧪 Testing the APIs

### ✔ Option 1 — HTML Test Page (Recommended)

Simply open [http://localhost:3000](http://localhost:3000) in your browser.

**Features:**
- Textbox + button for `/run`
- Textbox + button for `/autofix`
- Textbox + button for `/help`
- Pretty-printed multiline results

### ✔ Option 2 — cURL Commands

**Test `/run` endpoint:**
```bash
curl -X POST http://localhost:3000/run \
  -H "Content-Type: application/json" \
  -d "{\"code\":\"print('hi')\"}"
```

**Test `/autofix` endpoint:**
```bash
curl -X POST http://localhost:3000/autofix \
  -H "Content-Type: application/json" \
  -d "{\"code\":\"if(x){console.log('a')}\"}"
```

**Test `/help` endpoint:**
```bash
curl -X POST http://localhost:3000/help \
  -H "Content-Type: application/json" \
  -d "{\"query\":\"loop\"}"
```

---

## 📦 Project Structure

```
code-helper-backend/
├── main.go          # Main application and route handlers
├── go.mod           # Go module file
└── README.md        # This file
```

---

## 🔧 Requirements

- **Go 1.16+**
- **Go Fiber** — Web framework for Go


---

## 💡 Notes

- The code execution is **simulated** — it doesn't actually compile or run code
- Auto-fix applies basic formatting rules and may not handle all edge cases
- Help tips are keyword-based and provide general guidance

