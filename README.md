# 🩺 AI-Powered WhatsApp Health Assistant (Health Buddy)

A lightweight, production-ready **AI health companion** built with **Go**, **WhatsApp Cloud API**, and **OpenAI LLMs**.
Users can describe their health discomforts in natural language, and the bot responds with **empathetic, safe, non-prescriptive home-care guidance** — directly on WhatsApp.

This project is designed as a **real-world backend system**, not a toy demo, and showcases end-to-end cloud deployment, webhook handling, and LLM integration.

---

## 🚀 Features

* 💬 WhatsApp-based chat interface (no app installation needed)
* 🤖 LLM-powered understanding of free-text symptoms
* 🧠 Context-aware replies using recent chat history
* ❤️ Empathetic, reassuring responses with safe home-care tips
* ⚖️ Strict medical guardrails (no diagnosis, no prescriptions)
* 🔐 Secure secret management using environment variables
* 🌐 Cloud-native deployment on Render
* 🧯 Defensive webhook parsing (handles all WhatsApp event types)
* 💰 Low-cost architecture (free tiers + minimal API usage)

---

## 🏗️ Architecture Overview

```
User (WhatsApp)
      ↓
Meta WhatsApp Cloud API
      ↓
Webhook (Gin / Go)
      ↓
Context Memory (in-memory / DB)
      ↓
OpenAI LLM (gpt-4o-mini)
      ↓
Response Builder (safety + empathy)
      ↓
WhatsApp Send API
```

The LLM is used **only for intelligence**, while all control flow remains deterministic and safe.

---

## 🛠️ Tech Stack

| Layer     | Technology                            |
| --------- | ------------------------------------- |
| Backend   | Go (Gin)                              |
| AI        | OpenAI (gpt-4o-mini)                  |
| Messaging | WhatsApp Cloud API                    |
| Hosting   | Render                                |

---

## 🔐 Safety Design

This bot is **not a doctor** and never replaces professional care.

Guardrails enforced in the system prompt:

* ❌ No diagnosis
* ❌ No medication recommendations
* ❌ No dosage advice
* ✅ Only safe, general home-care tips
* ⚠️ Always suggests seeing a doctor for severe symptoms

This makes the bot legally safe and ethically responsible.

---


## ⚠️ Disclaimer

This project is for **educational and demonstration purposes only**.
It does not provide medical advice and should not be used as a replacement for professional healthcare.

---

## 🙌 Acknowledgements

* OpenAI
* Meta WhatsApp Cloud API
* Gin Web Framework
* Render

---

## ⭐ If you like this project

Give it a star ⭐ and feel free to fork or extend it!
