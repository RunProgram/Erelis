# Erelis

**An AI-powered vulnerability scanner written in Go.**

Erelis is a CLI security tool that leverages **OpenAI's Codex CLI** to perform structured, category-focused vulnerability analysis on source code repositories.

It combines:

- Go performance and strong process control  
- Codex-powered AI reasoning  
- Structured multi-pass scanning logic  
- Clean Markdown reporting  
- Live streaming scan progress  

---

## Features

- Targeted vulnerability scanning by category  
- AI-assisted data-flow tracing  
- Real-time streaming updates during scans  
- Structured Markdown vulnerability reports  
- Simple and intuitive CLI interface  
- Designed to evolve into full multi-pass security auditing  

---

## Installation

### 1. Install Go

```bash
sudo apt install golang-go
git clone https://github.com/YOUR_USERNAME/erelis.git
cd erelis
go install
