# 🧩 Persian CAPTCHA Generator (Go)

A simple and customizable Persian CAPTCHA generator for Go, rendering random 4-digit numbers in Persian numerals (`۰-۹`) with visual noise. Built using [fogleman/gg](https://github.com/fogleman/gg) and supports Persian fonts like IranNastaliq.

---

## ✨ Features

- ✅ Persian (Farsi) number CAPTCHA: `۰۱۲۳۴۵۶۷۸۹`
- ✅ Adds noisy lines for bot resistance
- ✅ Easy-to-use API (`Generate`, `GenerateImage`)
- ✅ Embed custom Persian fonts (e.g., IranNastaliq)
- ✅ Usable in web apps, file exports, or CLI tools

---

## 📦 Installation

```bash
go get github.com/Alirzamehrzad/persiancaptcha
```
🛠 Usage
🔹 Generate a CAPTCHA and save as PNG file

```code
package main

import (
	"log"
	"os"

	"github.com/yourusername/persiancaptcha"
)

func main() {
	img, text, err := persiancaptcha.Generate()
	if err != nil {
		log.Fatal(err)
	}

	// Save the image to a file
	file, err := os.Create("captcha.png")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	err = img.EncodePNG(file)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("CAPTCHA text:", text)
}
```
This generates an image like captcha.png with a number such as ۳۴۵۲.

🔹 Serve CAPTCHA in a web browser
```code
package main

import (
	"net/http"
	"log"

	"github.com/yourusername/persiancaptcha"
)

func captchaHandler(w http.ResponseWriter, r *http.Request) {
	img, text, err := persiancaptcha.Generate()
	if err != nil {
		http.Error(w, "Failed to generate CAPTCHA", 500)
		return
	}
	log.Println("CAPTCHA text:", text)

	w.Header().Set("Content-Type", "image/png")
	img.EncodePNG(w)
}

func main() {
	http.HandleFunc("/captcha", captchaHandler)
	log.Println("Server running at http://localhost:8080/captcha")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
```

Visit: http://localhost:8080/captcha to view the image in your browser.
