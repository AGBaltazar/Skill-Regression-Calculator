# 🧠 Skill Regression Predictor

A full-stack application built with **Go**, **JavaScript**, and **Bootstrap**, designed to estimate how much a technical skill has decayed over time using a mathematical decay model. It also provides targeted YouTube video recommendations to help users refresh their knowledge.

## 🚀 Project Purpose

The goal of this project is to help users reflect on their current proficiency with a skill based on how recently they practiced it and how difficult it is. By using a skill regression model inspired by Ebbinghaus’ Forgetting Curve, the app:

- Calculates a "retention score" between 0 and 10
- Interprets that score into levels like "Fresh", "Rusty", "Weak", or "Forgotten"
- Provides curated YouTube video recommendations for continued learning

## 🛠️ Tech Stack

- **Go** – Core backend logic and HTTP server
- **Docker & Docker Compose** – Containerized development
- **YouTube API** – For video recommendations
- **HTML/CSS/JS** – Simple frontend interface

## Flow
- User (HTML Form)

  ↓
- JavaScript (Captures Data)

   ↓
- HTTP Request (JSON Payload)

   ↓
- Go Backend API

   ↓
- API Calls (Recomendations)

   ↓
- HTTP Response (JSON Result)

   ↓
- JavaScript (Handles Response)

   ↓
- Updated DOM (Results Displayed)
