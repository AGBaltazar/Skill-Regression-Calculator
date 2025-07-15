# 🧠 Skill Regression Predictor

A backend-focused project built with **Go**, **PostgreSQL**, and **Docker** that estimates how much a user’s skill level has decayed over time and provides practice recommendations from various sources.

## 🚀 Project Purpose

Track skills over time and predict regression using a half-life decay model. This app allows users to:

- Log when they last practiced a skill
- Calculate current skill proficiency
- Receive personalized learning recommendations from YouTube
- View data in a simple dashboard

## 🛠️ Tech Stack

- **Go** – REST API backend
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
