package main

import (
	"log"
	"regexp"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type CodeRequest struct {
	Code string `json:"code"`
}

type HelpRequest struct {
	Query string `json:"query"`
}

type RunResponse struct {
	Output string `json:"output"`
	Error  string `json:"error,omitempty"`
}

type AutoFixResponse struct {
	OriginalCode string `json:"original_code"`
	FixedCode    string `json:"fixed_code"`
}

type HelpResponse struct {
	Query string   `json:"query"`
	Tips  []string `json:"tips"`
}

func runHandler(c *fiber.Ctx) error {
	var body CodeRequest
	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid JSON"})
	}

	code := strings.ToLower(body.Code)
	output := ""
	errMsg := ""

	if strings.Contains(code, "error") {
		errMsg = "Simulated compilation error: found keyword 'error' in code"
	} else if strings.Contains(code, "print") || strings.Contains(code, "fmt.println") {
		output = "Simulated output: Hello from your code!"
	} else {
		output = "Code executed successfully (simulated). No output."
	}

	return c.JSON(RunResponse{
		Output: output,
		Error:  errMsg,
	})
}


func autoFixCode(code string) string {
	lines := strings.Split(code, "\n")
	var out []string
	indent := 0

	spaceRegex := regexp.MustCompile(`\s+`)

	for _, raw := range lines {
		line := strings.TrimSpace(raw)

		if line == "" {
			out = append(out, "")
			continue
		}

		
		if strings.HasPrefix(line, "}") {
			indent--
		}

		line = spaceRegex.ReplaceAllString(line, " ")

		if !strings.HasSuffix(line, "{") &&
			!strings.HasSuffix(line, "}") &&
			!strings.HasSuffix(line, ";") &&
			!strings.HasSuffix(line, ":") {
			line += ";"
		}

		indentedLine := strings.Repeat("  ", indent) + line
		out = append(out, indentedLine)

		if strings.HasSuffix(line, "{") {
			indent++
		}
	}

	for indent > 0 {
		indent--
		out = append(out, strings.Repeat("  ", indent)+"}")
	}

	return strings.Join(out, "\n")
}


func autofixHandler(c *fiber.Ctx) error {
	var body CodeRequest
	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid JSON"})
	}

	fixed := autoFixCode(body.Code)

	return c.JSON(fiber.Map{
		"original_code": body.Code,
		"fixed_code":    fixed,
	})
}


func helpHandler(c *fiber.Ctx) error {
	var body HelpRequest
	if err := c.BodyParser(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid JSON"})
	}

	q := strings.ToLower(body.Query)
	var tips []string

	if strings.Contains(q, "loop") || strings.Contains(q, "for") {
		tips = append(tips, "Ensure your loop variable updates to avoid infinite loops.")
	}
	if strings.Contains(q, "if") || strings.Contains(q, "condition") {
		tips = append(tips, "Check your conditional logic and consider edge cases.")
	}
	if strings.Contains(q, "error") || strings.Contains(q, "bug") {
		tips = append(tips, "Read error logs carefully and reproduce the bug in a small test.")
	}
	if strings.Contains(q, "syntax") || strings.Contains(q, "semicolon") {
		tips = append(tips, "Check brackets and ensure all statements end with semicolons.")
	}
	if len(tips) == 0 {
		tips = append(tips, "No direct tip found. Try asking about loops, syntax, or errors.")
	}

	return c.JSON(HelpResponse{
		Query: body.Query,
		Tips:  tips,
	})
}


const htmlPage = `
<!DOCTYPE html>
<html>
<head>
  <title>Code Helper Backend Test</title>
  <style>
    body { font-family: Arial; margin: 20px; background: #f5f5f5; }
    textarea { width: 100%; height: 100px; margin-bottom: 10px; }
    input { width: 100%; margin-bottom: 10px; }
    button { padding: 6px 14px; margin-bottom: 10px; }
    pre { background: #222; color:#eee; padding: 10px; border-radius: 6px; }
    .box { background: #fff; padding: 15px; border-radius: 8px; margin-bottom: 20px; }
  </style>
</head>
<body>

<h1>Code Helper Backend Test</h1>

<div class="box">
  <h2>/run</h2>
  <textarea id="runBox"></textarea>
  <button onclick="runCode()">Run Code</button>
  <pre id="runOut"></pre>
</div>

<div class="box">
  <h2>/autofix</h2>
  <textarea id="fixBox"></textarea>
  <button onclick="fixCode()">Auto Fix</button>
  <pre id="fixOut"></pre>
</div>

<div class="box">
  <h2>/help</h2>
  <input id="helpBox">
  <button onclick="getHelp()">Get Help</button>
  <pre id="helpOut"></pre>
</div>

<script>
async function runCode() {
  let res = await fetch('/run',{
    method:'POST',
    headers:{'Content-Type':'application/json'},
    body:JSON.stringify({code:runBox.value})
  });
  runOut.textContent = JSON.stringify(await res.json(),null,2);
}

async function fixCode() {
  let res = await fetch('/autofix',{
    method:'POST',
    headers:{'Content-Type':'application/json'},
    body:JSON.stringify({code:fixBox.value})
  });
  
  let data = await res.json();
  
  // ----- FIXED OUTPUT DISPLAY -----
  fixOut.textContent =
    "Original Code:\n\n" + data.original_code +
    "\n\nFixed Code:\n\n" + data.fixed_code;
}

async function getHelp() {
  let res = await fetch('/help',{
    method:'POST',
    headers:{'Content-Type':'application/json'},
    body:JSON.stringify({query:helpBox.value})
  });
  helpOut.textContent = JSON.stringify(await res.json(),null,2);
}
</script>


</body>
</html>
`

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Type("html").SendString(htmlPage)
	})

	app.Post("/run", runHandler)
	app.Post("/autofix", autofixHandler)
	app.Post("/help", helpHandler)

	log.Println("Running at http://localhost:3000")
	log.Fatal(app.Listen(":3000"))
}
