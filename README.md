# Real-Time Forum Application

## Overview
This project is a single-page web application for a real-time forum built with Go, SQLite, Ja vaScript, HTML, and CSS. It supports user registration/login, post creation, commenting, and private messaging with real-time updates via WebSockets. The application ensures secure authentication, efficient data handling, and a responsive user interface without using frontend frameworks.

## Features
- **Registration and Login**:
  - Users register with nickname, age, gender, first name, last name, email, and password.
  - Login using nickname or email with password.
  - Logout available from any page.
- **Posts and Comments**:
  - View posts in a feed.
  - Comment on posts, visible only when viewing a specific post.
- **Private Messages**:
  - Real-time chat with online/offline user status.
  - Chat list sorted by last message (or alphabetically for new users).
  - Load last 10 messages, with scroll-up to fetch 10 more (throttled scrolling).
  - Messages display sender's nickname and timestamp.
- **Real-Time Updates**:
  - WebSockets for instant message notifications and chat updates.
- **Single Page Application**:
  - One HTML file with dynamic content managed by JavaScript.
- **Security**:
  - Passwords hashed with bcrypt.
  - Session management with cookies.

## Technologies
- **Backend**:
  - Go (standard library + allowed packages: gorilla/websocket, sqlite3, bcrypt, google/uuid).
  - SQLite for data storage.
  - WebSockets for real-time communication.
- **Frontend**:
  - Vanilla JavaScript for DOM manipulation and WebSocket handling.
  - HTML (single file for SPA).
  - CSS for styling.
- **No frontend frameworks** (e.g., React, Angular, Vue) are used.

## Project Structure
```
forum/
├── backend/
│   ├── main.go              # Entry point, server setup
│   ├── handlers/            # HTTP and WebSocket handlers
│   ├── models/              # Data structures (User, Post, Comment, Message)
│   ├── database/            # SQLite setup and queries
│   └── sessions/            # Session management
├── frontend/
│   ├── index.html           # Single HTML file
│   ├── css/                 # Stylesheets
│   │   └── styles.css
│   └── js/                  # JavaScript files
│       ├── main.js          # Core SPA logic
│       ├── websocket.js     # WebSocket client
│       └── utils.js         # Throttle/debounce and helpers
├── database/
│   └── forum.db             # SQLite database
├── go.mod                   # Go module dependencies
└── README.md                # This file
```



## Setup Instructions
### Prerequisites
- Go (version 1.21 or higher)
- SQLite3
- A modern web browser

### Installation
1. **Clone the Repository**:
   ```bash
   git clone <repository-url>
   cd forum
   ```

2. **Install Go Dependencies**:
   ```bash
   go mod init forum
   go get github.com/gorilla/websocket
   go get golang.org/x/crypto/bcrypt
   go get github.com/google/uuid
   ```

3. **Initialize the Database**:
   - Create the SQLite database (`forum.db`) in the `database/` directory.
   - Run the SQL schema (provided in `database/schema.sql`) to create tables for users, posts, comments, messages, and sessions:
     ```bash
     sqlite3 database/forum.db < database/schema.sql
     ```

4. **Run the Server**:
   ```bash
   go run backend/main.go
   ```
   The server runs on `http://localhost:8080` by default.

5. **Access the Application**:
   - Open a browser and navigate to `http://localhost:8080`.
   - Register a new account or log in to access the forum.

## Usage
- **Registration/Login**:
  - On the landing page, choose to register or log in.
  - Fill in the registration form with required fields (nickname, age, gender, first name, last name, email, password).
  - Log in using nickname or email and password.
- **Posts**:
  - Create a post with a title, content, and category from the feed page.
  - Click a post to view its content and comments.
  - Add comments to posts.
 —, **Private Messages**:
  - View the chat list (online/offline users) on the sidebar, sorted by last message or alphabetically.
  - Click a user to load the last 10 messages.
  - Scroll up to load more messages (throttled to prevent spam).
  - Send messages, which appear instantly for both users via WebSockets.
- **Logout**:
  - Click the logout button (available on all pages) to end the session.

## Development Notes
- **Backend**:
  - Uses Go routines and channels for concurrent WebSocket handling.
  - SQLite stores all data (users, posts, comments, messages).
  - Sessions are managed with cookies and UUIDs.
- **Frontend**:
  - JavaScript handles dynamic page updates (SPA) without reloading.
  - WebSocket client listens for real-time message updates.
  - Throttle/debounce used for scroll events to optimize message loading.
- **Database**:
  - Tables: `users`, `posts`, `comments`, `messages`, `sessions`.
  - Foreign keys ensure data integrity.
- **Security**:
  - Passwords are hashed with bcrypt.
  - Input validation prevents SQL injection and XSS.

## Allowed Packages
- Go standard library
- github.com/gorilla/websocket
- golang.org/x/crypto/bcrypt
- github.com/google/uuid
- SQLite3 (via database/sql)

## Learning Outcomes
This project demonstrates:
- Web basics (HTML, HTTP, CSS, DOM).
- Backend development (Go, WebSockets, SQLite).
- Frontend development (vanilla JavaScript, SPA).
- Real-time communication (WebSockets).
- Database management (SQL).
- Security practices (password hashing, session management).

## Troubleshooting
- **Server not starting**:
  - Ensure `forum.db` exists and is writable.
  - Check for port conflicts on `:8080`.
- **WebSocket connection fails**:
  - Verify the WebSocket endpoint (`ws://localhost:8080/ws`).
  - Check browser console for errors.
- **Messages not loading**:
  - Ensure the database has the correct schema.
  - Verify WebSocket connection is active.

## Future Improvements
- Add post filtering by category.
- Implement message search functionality.
- Enhance UI with responsive design for mobile devices.

## License
This project is for educational purposes and not licensed for commercial use.