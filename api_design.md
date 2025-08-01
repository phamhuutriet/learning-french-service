# Apprendre.ai REST API Design

## Base URL

```
https://api.apprendre.ai/v1
```

## Authentication

All endpoints require JWT Bearer token unless specified otherwise.

```
Authorization: Bearer <jwt_token>
```

---

## 🔐 Authentication & Users

### POST /auth/register

Register a new user account.

**Request:**

```json
{
    "email": "john@example.com",
    "password": "securepassword123",
    "username": "john_learner",
    "first_name": "John",
    "last_name": "Doe",
    "current_level": "A2",
    "target_level": "B2",
    "daily_goal": 25
}
```

**Response (201):**

```json
{
    "user": {
        "id": "uuid-123",
        "email": "john@example.com",
        "username": "john_learner",
        "current_level": "A2",
        "target_level": "B2",
        "daily_goal": 25,
        "created_at": "2024-01-15T10:00:00Z"
    },
    "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "refresh_token": "uuid-refresh-token"
}
```

### POST /auth/login

Login with email and password.

**Request:**

```json
{
    "email": "john@example.com",
    "password": "securepassword123"
}
```

**Response (200):**

```json
{
    "user": {
        "id": "uuid-123",
        "email": "john@example.com",
        "username": "john_learner",
        "current_level": "A2",
        "daily_goal": 25
    },
    "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "refresh_token": "uuid-refresh-token"
}
```

### GET /users/me

Get current user profile.

**Response (200):**

```json
{
    "id": "uuid-123",
    "email": "john@example.com",
    "username": "john_learner",
    "first_name": "John",
    "last_name": "Doe",
    "current_level": "A2",
    "target_level": "B2",
    "daily_goal": 25,
    "timezone": "UTC",
    "created_at": "2024-01-15T10:00:00Z",
    "last_active_at": "2024-01-20T15:30:00Z"
}
```

---

## 📚 Deck Management

### GET /decks

Get all decks for the authenticated user.

**Query Parameters:**

-   `category` (optional): Filter by category
-   `level` (optional): Filter by target level
-   `active` (optional): Filter by active status (true/false)

**Response (200):**

```json
{
    "decks": [
        {
            "id": "deck-uuid-1",
            "name": "Travel Vocabulary",
            "description": "Essential words for traveling in France",
            "category": "Travel",
            "target_level": "A2",
            "word_count": 45,
            "is_active": true,
            "created_at": "2024-01-10T09:00:00Z",
            "updated_at": "2024-01-20T14:15:00Z"
        }
    ],
    "total": 1
}
```

### POST /decks

Create a new deck.

**Request:**

```json
{
    "name": "Business French",
    "description": "Professional vocabulary for work",
    "category": "Business",
    "target_level": "B1"
}
```

**Response (201):**

```json
{
    "id": "deck-uuid-2",
    "name": "Business French",
    "description": "Professional vocabulary for work",
    "category": "Business",
    "target_level": "B1",
    "word_count": 0,
    "is_active": true,
    "created_at": "2024-01-20T16:00:00Z",
    "updated_at": "2024-01-20T16:00:00Z"
}
```

### GET /decks/:id

Get specific deck details.

**Response (200):**

```json
{
    "id": "deck-uuid-1",
    "name": "Travel Vocabulary",
    "description": "Essential words for traveling in France",
    "category": "Travel",
    "target_level": "A2",
    "word_count": 45,
    "is_active": true,
    "created_at": "2024-01-10T09:00:00Z",
    "updated_at": "2024-01-20T14:15:00Z"
}
```

### PUT /decks/:id

Update deck information.

**Request:**

```json
{
    "name": "Travel & Tourism",
    "description": "Updated description",
    "category": "Travel"
}
```

**Response (200):**

```json
{
    "id": "deck-uuid-1",
    "name": "Travel & Tourism",
    "description": "Updated description",
    "category": "Travel",
    "target_level": "A2",
    "word_count": 45,
    "is_active": true,
    "updated_at": "2024-01-20T16:30:00Z"
}
```

### DELETE /decks/:id

Delete a deck and all its words.

**Response (204):** No content

---

## 📝 Word Management

### GET /decks/:id/words

Get all words in a specific deck.

**Query Parameters:**

-   `page` (optional): Page number (default: 1)
-   `limit` (optional): Items per page (default: 20)
-   `search` (optional): Search in french_word or english_translation
-   `difficulty` (optional): Filter by difficulty level
-   `tags` (optional): Filter by tags (comma-separated)

**Response (200):**

```json
{
    "words": [
        {
            "id": "word-uuid-1",
            "french_word": "bonjour",
            "english_translation": "hello, good morning",
            "part_of_speech": "interjection",
            "gender": null,
            "french_example": "Bonjour, comment allez-vous?",
            "english_example_translation": "Hello, how are you?",
            "phonetic_transcription": "/bon.ˈʒuʁ/",
            "pronunciation_audio_url": "https://audio.example.com/bonjour.mp3",
            "difficulty_level": "A1",
            "tags": ["greetings", "basic"],
            "usage_context": "Used to greet someone during the day",
            "created_at": "2024-01-15T11:00:00Z",
            "question_count": 3
        }
    ],
    "pagination": {
        "page": 1,
        "limit": 20,
        "total": 45,
        "pages": 3
    }
}
```

### POST /decks/:id/words

Add a new word to a deck (with synchronous AI enrichment).

**Request:**

```json
{
    "french_word": "merci",
    "tags": ["gratitude", "basic"],
    "difficulty_level": "A1"
}
```

**Response (201):**

```json
{
    "id": "word-uuid-2",
    "french_word": "merci",
    "english_translation": "thank you",
    "part_of_speech": "interjection",
    "gender": null,
    "french_example": "Merci beaucoup pour votre aide.",
    "english_example_translation": "Thank you very much for your help.",
    "phonetic_transcription": "/mɛʁ.si/",
    "pronunciation_audio_url": "https://audio.example.com/merci.mp3",
    "difficulty_level": "A1",
    "tags": ["gratitude", "basic"],
    "usage_context": "Used to express gratitude",
    "created_at": "2024-01-20T17:00:00Z",
    "questions": [
        {
            "id": "question-uuid-4",
            "question_type": "french_to_english",
            "question_text": "What does 'merci' mean in English?",
            "correct_answer": "thank you"
        },
        {
            "id": "question-uuid-5",
            "question_type": "english_to_french",
            "question_text": "How do you say 'thank you' in French?",
            "correct_answer": "merci"
        },
        {
            "id": "question-uuid-6",
            "question_type": "pronunciation",
            "question_text": "Pronounce the word 'merci'",
            "correct_answer": "/mɛʁ.si/"
        }
    ]
}
```

### GET /words/:id

Get detailed information about a specific word.

**Response (200):**

```json
{
    "id": "word-uuid-1",
    "deck_id": "deck-uuid-1",
    "french_word": "bonjour",
    "english_translation": "hello, good morning",
    "part_of_speech": "interjection",
    "gender": null,
    "french_example": "Bonjour, comment allez-vous?",
    "english_example_translation": "Hello, how are you?",
    "phonetic_transcription": "/bon.ˈʒuʁ/",
    "pronunciation_audio_url": "https://audio.example.com/bonjour.mp3",
    "difficulty_level": "A1",
    "tags": ["greetings", "basic"],
    "usage_context": "Used to greet someone during the day",
    "created_at": "2024-01-15T11:00:00Z",
    "updated_at": "2024-01-15T11:05:00Z",
    "questions": [
        {
            "id": "question-uuid-1",
            "question_type": "french_to_english",
            "question_text": "What does 'bonjour' mean in English?",
            "correct_answer": "hello"
        },
        {
            "id": "question-uuid-2",
            "question_type": "english_to_french",
            "question_text": "How do you say 'hello' in French?",
            "correct_answer": "bonjour"
        },
        {
            "id": "question-uuid-3",
            "question_type": "pronunciation",
            "question_text": "Pronounce the word 'bonjour'",
            "correct_answer": "/bon.ˈʒuʁ/"
        }
    ]
}
```

### PUT /words/:id

Update word information manually.

**Request:**

```json
{
    "english_translation": "hello, good day",
    "tags": ["greetings", "basic", "formal"],
    "usage_context": "Formal greeting used during daytime"
}
```

**Response (200):** Updated word object

### DELETE /words/:id

Delete a word and all its questions.

**Response (204):** No content

---

## 🎓 Learning & Questions

### GET /questions/due

Get questions due for review (spaced repetition).

**Query Parameters:**

-   `limit` (optional): Max questions to return (default: 20)
-   `deck_id` (optional): Filter by specific deck
-   `question_type` (optional): Filter by question type

**Response (200):**

```json
{
    "due_questions": [
        {
            "id": "question-uuid-1",
            "word_id": "word-uuid-1",
            "question_type": "french_to_english",
            "question_text": "What does 'bonjour' mean in English?",
            "difficulty_level": "A1",
            "deck_name": "Travel Vocabulary",
            "french_word": "bonjour",
            "hints": "This is a common greeting",
            "review_info": {
                "ease_factor": 2.5,
                "interval_days": 1,
                "repetition_count": 0,
                "last_reviewed_at": null,
                "accuracy_rate": 0.0
            }
        }
    ],
    "total_due": 15,
    "next_review_time": "2024-01-21T09:00:00Z"
}
```

### POST /questions/:id/attempt

Submit an answer attempt and get AI grading.

**Request:**

```json
{
    "user_answer": "hello",
    "response_time_ms": 3500
}
```

**Response (200):**

```json
{
    "attempt_id": "attempt-uuid-1",
    "question_id": "question-uuid-1",
    "user_answer": "hello",
    "correct_answer": "hello, good morning",
    "was_correct": true,
    "ai_grade": "correct",
    "ai_feedback": "Perfect! You got the exact translation.",
    "similarity_score": 1.0,
    "response_time_ms": 3500,
    "sm2_options": {
        "again": {
            "quality_rating": 0,
            "next_review": "2024-01-20T18:00:00Z",
            "new_interval_days": 1
        },
        "hard": {
            "quality_rating": 1,
            "next_review": "2024-01-21T17:00:00Z",
            "new_interval_days": 1
        },
        "good": {
            "quality_rating": 2,
            "next_review": "2024-01-23T17:00:00Z",
            "new_interval_days": 3
        },
        "easy": {
            "quality_rating": 3,
            "next_review": "2024-01-26T17:00:00Z",
            "new_interval_days": 6
        }
    }
}
```

### POST /questions/:id/rate

Rate the difficulty after seeing the answer (SM-2 algorithm).

**Request:**

```json
{
    "attempt_id": "attempt-uuid-1",
    "quality_rating": 2,
    "rating": "good"
}
```

**Response (200):**

```json
{
    "message": "Rating recorded successfully",
    "next_review_date": "2024-01-23T17:00:00Z",
    "new_interval_days": 3,
    "updated_ease_factor": 2.5,
    "repetition_count": 1
}
```

### POST /questions/:id/pronunciation-attempt

Submit pronunciation attempt for analysis.

**Request (multipart/form-data):**

```
audio_file: [audio file]
question_id: question-uuid-3
```

**Response (200):**

```json
{
    "attempt_id": "pronunciation-attempt-uuid-1",
    "user_audio_url": "https://audio.example.com/user-recordings/attempt-1.wav",
    "reference_audio_url": "https://audio.example.com/bonjour.mp3",
    "similarity_score": 0.85,
    "pronunciation_feedback": "Very good! Try to emphasize the 'r' sound more at the end.",
    "phonetic_accuracy": {
        "/b/": 0.95,
        "/o/": 0.9,
        "/n/": 0.85,
        "/ʒ/": 0.8,
        "/u/": 0.9,
        "/ʁ/": 0.7
    },
    "overall_grade": "good"
}
```

### GET /questions/:id

Get question details.

**Response (200):**

```json
{
    "id": "question-uuid-1",
    "word_id": "word-uuid-1",
    "question_type": "french_to_english",
    "question_text": "What does 'bonjour' mean in English?",
    "correct_answer": "hello, good morning",
    "difficulty_level": "A1",
    "hints": "This is a common greeting",
    "options": null,
    "word": {
        "french_word": "bonjour",
        "phonetic_transcription": "/bon.ˈʒuʁ/",
        "pronunciation_audio_url": "https://audio.example.com/bonjour.mp3"
    }
}
```

---

## 📊 Statistics & Analytics

### GET /users/me/stats

Get comprehensive user statistics.

**Query Parameters:**

-   `period` (optional): daily, weekly, monthly, all (default: all)
-   `start_date` (optional): Start date for period
-   `end_date` (optional): End date for period

**Response (200):**

```json
{
    "overview": {
        "total_words": 125,
        "words_learned": 87,
        "total_questions_answered": 650,
        "overall_accuracy": 0.78,
        "current_streak_days": 12,
        "longest_streak_days": 23,
        "total_study_time_minutes": 1450
    },
    "by_question_type": {
        "french_to_english": {
            "accuracy": 0.82,
            "total_attempts": 220,
            "average_response_time_ms": 3200
        },
        "english_to_french": {
            "accuracy": 0.75,
            "total_attempts": 215,
            "average_response_time_ms": 4500
        },
        "pronunciation": {
            "accuracy": 0.68,
            "total_attempts": 215,
            "average_response_time_ms": 8000
        }
    },
    "recent_performance": [
        {
            "date": "2024-01-20",
            "questions_answered": 25,
            "questions_correct": 20,
            "study_time_minutes": 45,
            "words_learned": 3
        }
    ],
    "ai_insights": {
        "strengths": ["vocabulary_recognition", "basic_grammar"],
        "weaknesses": ["pronunciation", "complex_sentences"],
        "recommendations": "Focus more on pronunciation practice. Try the pronunciation questions for words you already know well."
    }
}
```

### GET /users/me/stats/daily

Get daily statistics for a date range.

**Query Parameters:**

-   `start_date` (required): Start date (YYYY-MM-DD)
-   `end_date` (required): End date (YYYY-MM-DD)

**Response (200):**

```json
{
    "daily_stats": [
        {
            "date": "2024-01-20",
            "questions_answered": 25,
            "questions_correct": 20,
            "accuracy_rate": 0.8,
            "study_time_minutes": 45,
            "words_learned": 3,
            "streak_day": 12,
            "english_to_french_accuracy": 0.75,
            "french_to_english_accuracy": 0.85,
            "pronunciation_accuracy": 0.7
        }
    ],
    "summary": {
        "total_questions": 125,
        "total_correct": 98,
        "overall_accuracy": 0.78,
        "total_study_time": 180,
        "days_studied": 5
    }
}
```

### GET /users/me/streaks

Get learning streak information.

**Response (200):**

```json
{
    "current_streak": {
        "start_date": "2024-01-09",
        "length_days": 12,
        "is_active": true
    },
    "longest_streak": {
        "start_date": "2023-12-01",
        "end_date": "2023-12-23",
        "length_days": 23
    },
    "recent_streaks": [
        {
            "start_date": "2024-01-09",
            "end_date": null,
            "length_days": 12,
            "is_current": true
        },
        {
            "start_date": "2023-12-01",
            "end_date": "2023-12-23",
            "length_days": 23,
            "is_current": false
        }
    ]
}
```

### GET /decks/:id/stats

Get statistics for a specific deck.

**Response (200):**

```json
{
    "deck_id": "deck-uuid-1",
    "deck_name": "Travel Vocabulary",
    "total_words": 45,
    "words_learned": 32,
    "completion_rate": 0.71,
    "average_accuracy": 0.78,
    "total_reviews": 340,
    "by_question_type": {
        "french_to_english": { "accuracy": 0.82, "reviews": 115 },
        "english_to_french": { "accuracy": 0.75, "reviews": 112 },
        "pronunciation": { "accuracy": 0.68, "reviews": 113 }
    },
    "difficulty_breakdown": {
        "A1": { "words": 20, "learned": 18, "accuracy": 0.85 },
        "A2": { "words": 25, "learned": 14, "accuracy": 0.72 }
    }
}
```

---

## 🔍 Search & Filtering

### GET /search/words

Search across all user's words.

**Query Parameters:**

-   `q` (required): Search query
-   `deck_id` (optional): Filter by deck
-   `difficulty` (optional): Filter by difficulty
-   `tags` (optional): Filter by tags
-   `limit` (optional): Results limit (default: 20)

**Response (200):**

```json
{
    "words": [
        {
            "id": "word-uuid-1",
            "french_word": "bonjour",
            "english_translation": "hello, good morning",
            "deck_name": "Travel Vocabulary",
            "difficulty_level": "A1",
            "tags": ["greetings", "basic"],
            "match_type": "french_word"
        }
    ],
    "total": 1
}
```

---

## 📋 Error Responses

### Standard Error Format

```json
{
    "error": {
        "code": "INVALID_REQUEST",
        "message": "The request data is invalid",
        "details": {
            "field": "email",
            "issue": "Email format is invalid"
        }
    },
    "timestamp": "2024-01-20T17:00:00Z"
}
```

### Common Error Codes

-   `UNAUTHORIZED` (401): Invalid or missing auth token
-   `FORBIDDEN` (403): User doesn't have permission
-   `NOT_FOUND` (404): Resource not found
-   `INVALID_REQUEST` (400): Validation errors
-   `RATE_LIMITED` (429): Too many requests
-   `SERVER_ERROR` (500): Internal server error
-   `AI_SERVICE_ERROR` (503): AI enrichment service unavailable

---

## 🚀 Rate Limiting

-   **General API**: 1000 requests per hour per user
-   **AI Enrichment**: 50 requests per hour per user
-   **File Upload**: 20 uploads per hour per user
-   **Authentication**: 10 login attempts per hour per IP

## 📝 Response Headers

All responses include:

```
X-RateLimit-Limit: 1000
X-RateLimit-Remaining: 987
X-RateLimit-Reset: 1642694400
Content-Type: application/json
```
