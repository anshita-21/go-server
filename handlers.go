package main

import (
	"fmt"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "ParseForm() err: "+err.Error(), http.StatusBadRequest)
		return
	}
	name := r.FormValue("name")
	profession := r.FormValue("profession")
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Submission Result</title>
    <style>
        * { margin: 0; padding: 0; box-sizing: border-box; }
        body {
            font-family: 'Segoe UI', sans-serif;
            background: #0f172a;
            color: #e2e8f0;
            min-height: 100vh;
            display: flex;
            align-items: center;
            justify-content: center;
        }
        .card {
            background: #1e293b;
            border: 1px solid #334155;
            border-radius: 16px;
            padding: 2.5rem;
            width: 100%%;
            max-width: 420px;
            box-shadow: 0 25px 50px rgba(0,0,0,0.4);
            text-align: center;
        }
        .icon { font-size: 2.5rem; margin-bottom: 1rem; }
        h2 {
            font-size: 1.6rem;
            font-weight: 700;
            background: linear-gradient(135deg, #22d3ee, #818cf8);
            -webkit-background-clip: text;
            -webkit-text-fill-color: transparent;
            margin-bottom: 1.5rem;
        }
        .info {
            background: #0f172a;
            border: 1px solid #334155;
            border-radius: 10px;
            padding: 1rem 1.5rem;
            margin-bottom: 0.8rem;
            text-align: left;
        }
        .info span { color: #64748b; font-size: 0.8rem; display: block; margin-bottom: 0.2rem; }
        .info strong { font-size: 1rem; color: #e2e8f0; }
        .links { display: flex; gap: 0.8rem; margin-top: 1.5rem; }
        a {
            flex: 1;
            padding: 0.7rem;
            border-radius: 8px;
            text-decoration: none;
            font-weight: 600;
            font-size: 0.9rem;
            transition: all 0.2s;
            text-align: center;
        }
        .btn-primary { background: linear-gradient(135deg, #22d3ee, #818cf8); color: #0f172a; }
        .btn-primary:hover { opacity: 0.85; }
        .btn-secondary { border: 1px solid #334155; color: #94a3b8; }
        .btn-secondary:hover { border-color: #22d3ee; color: #22d3ee; }
    </style>
</head>
<body>
    <div class="card">
        <div class="icon">✅</div>
        <h2>Submission Successful</h2>
        <div class="info"><span>Name</span><strong>%s</strong></div>
        <div class="info"><span>Profession</span><strong>%s</strong></div>
        <div class="links">
            <a href="/form.html" class="btn-primary">Submit Again</a>
            <a href="/" class="btn-secondary">Home</a>
        </div>
    </div>
</body>
</html>`, name, profession)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "method is not supported", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Hello</title>
    <style>
        * { margin: 0; padding: 0; box-sizing: border-box; }
        body {
            font-family: 'Segoe UI', sans-serif;
            background: #0f172a;
            color: #e2e8f0;
            min-height: 100vh;
            display: flex;
            align-items: center;
            justify-content: center;
        }
        .hero { text-align: center; padding: 2rem; }
        .wave { font-size: 4rem; display: inline-block; animation: wave 1.5s infinite; }
        @keyframes wave {
            0%%, 100%% { transform: rotate(0deg); }
            25%% { transform: rotate(20deg); }
            75%% { transform: rotate(-10deg); }
        }
        h1 {
            font-size: 3rem;
            font-weight: 700;
            background: linear-gradient(135deg, #22d3ee, #818cf8);
            -webkit-background-clip: text;
            -webkit-text-fill-color: transparent;
            margin: 1rem 0 0.5rem;
        }
        p { color: #64748b; margin-bottom: 2rem; }
        a {
            padding: 0.7rem 1.8rem;
            border-radius: 8px;
            text-decoration: none;
            font-weight: 600;
            border: 1px solid #334155;
            color: #94a3b8;
            transition: all 0.2s;
        }
        a:hover { border-color: #22d3ee; color: #22d3ee; }
    </style>
</head>
<body>
    <div class="hero">
        <span class="wave">👋</span>
        <h1>Hello!</h1>
        <p>The Go server is up and running.</p>
        <a href="/">← Back to Home</a>
    </div>
</body>
</html>`)
}
