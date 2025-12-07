#!/bin/bash

# GraphQL endpoint
ENDPOINT="http://localhost:8080/query"

# First, create a quiz about Rust
echo "Creating Rust quiz..."
QUIZ_RESPONSE=$(curl -s -X POST "$ENDPOINT" \
  -H "Content-Type: application/json" \
  -d '{
    "query": "mutation CreateQuiz($input: CreateQuizInput!) { createQuiz(input: $input) { id title description } }",
    "variables": {
      "input": {
        "title": "Rust Programming Fundamentals",
        "description": "A comprehensive quiz covering Rust programming language fundamentals including ownership, borrowing, lifetimes, and more."
      }
    }
  }')

echo "Quiz created: $QUIZ_RESPONSE"

# Extract quiz ID from response
QUIZ_ID=$(echo "$QUIZ_RESPONSE" | grep -o '"id":"[0-9]*"' | head -1 | grep -o '[0-9]*')
echo "Quiz ID: $QUIZ_ID"

if [ -z "$QUIZ_ID" ]; then
  echo "Failed to create quiz"
  exit 1
fi

# Create questions
echo "Creating questions..."

# Question 1: Ownership
curl -s -X POST "$ENDPOINT" \
  -H "Content-Type: application/json" \
  -d "{
    \"query\": \"mutation CreateQuestion(\$input: CreateQuestionInput!) { createQuestion(input: \$input) { id content } }\",
    \"variables\": {
      \"input\": {
        \"quizID\": \"$QUIZ_ID\",
        \"type\": \"MULTIPLE_CHOICE\",
        \"content\": \"What is the primary purpose of Rust's ownership system?\",
        \"options\": [
          \"To prevent memory leaks and data races at compile time\",
          \"To make the code run faster\",
          \"To enable dynamic typing\",
          \"To support object-oriented programming\"
        ],
        \"correctAnswer\": \"To prevent memory leaks and data races at compile time\",
        \"explanation\": \"Rust's ownership system is designed to guarantee memory safety and prevent data races without needing a garbage collector. It enforces these rules at compile time.\",
        \"difficulty\": \"MEDIUM\"
      }
    }
  }"
echo "Question 1 created"

# Question 2: Borrowing
curl -s -X POST "$ENDPOINT" \
  -H "Content-Type: application/json" \
  -d "{
    \"query\": \"mutation CreateQuestion(\$input: CreateQuestionInput!) { createQuestion(input: \$input) { id content } }\",
    \"variables\": {
      \"input\": {
        \"quizID\": \"$QUIZ_ID\",
        \"type\": \"MULTIPLE_CHOICE\",
        \"content\": \"Which of the following is true about mutable and immutable references in Rust?\",
        \"options\": [
          \"You can have multiple mutable references to the same data\",
          \"You can have either one mutable reference or any number of immutable references\",
          \"You can have multiple mutable and immutable references at the same time\",
          \"References in Rust are always mutable\"
        ],
        \"correctAnswer\": \"You can have either one mutable reference or any number of immutable references\",
        \"explanation\": \"Rust's borrowing rules state that at any given time, you can have either one mutable reference or any number of immutable references to a piece of data, but not both.\",
        \"difficulty\": \"MEDIUM\"
      }
    }
  }"
echo "Question 2 created"

# Question 3: Lifetimes
curl -s -X POST "$ENDPOINT" \
  -H "Content-Type: application/json" \
  -d "{
    \"query\": \"mutation CreateQuestion(\$input: CreateQuestionInput!) { createQuestion(input: \$input) { id content } }\",
    \"variables\": {
      \"input\": {
        \"quizID\": \"$QUIZ_ID\",
        \"type\": \"MULTIPLE_CHOICE\",
        \"content\": \"What do lifetime annotations in Rust specify?\",
        \"options\": [
          \"How long a variable will live in memory\",
          \"The relationship between the lifetimes of references\",
          \"When the garbage collector should free memory\",
          \"The execution time of a function\"
        ],
        \"correctAnswer\": \"The relationship between the lifetimes of references\",
        \"explanation\": \"Lifetime annotations don't change how long references live. Instead, they describe the relationships between the lifetimes of multiple references to help the compiler verify that references are valid.\",
        \"difficulty\": \"HARD\"
      }
    }
  }"
echo "Question 3 created"

# Question 4: Option and Result
curl -s -X POST "$ENDPOINT" \
  -H "Content-Type: application/json" \
  -d "{
    \"query\": \"mutation CreateQuestion(\$input: CreateQuestionInput!) { createQuestion(input: \$input) { id content } }\",
    \"variables\": {
      \"input\": {
        \"quizID\": \"$QUIZ_ID\",
        \"type\": \"MULTIPLE_CHOICE\",
        \"content\": \"What is the difference between Option<T> and Result<T, E> in Rust?\",
        \"options\": [
          \"Option is for values that might be absent, Result is for operations that might fail\",
          \"Option is faster than Result\",
          \"Result can only be used with I/O operations\",
          \"They are exactly the same\"
        ],
        \"correctAnswer\": \"Option is for values that might be absent, Result is for operations that might fail\",
        \"explanation\": \"Option<T> represents an optional value that can be Some(T) or None. Result<T, E> represents either success (Ok(T)) or failure (Err(E)), making it ideal for error handling.\",
        \"difficulty\": \"EASY\"
      }
    }
  }"
echo "Question 4 created"

# Question 5: Traits
curl -s -X POST "$ENDPOINT" \
  -H "Content-Type: application/json" \
  -d "{
    \"query\": \"mutation CreateQuestion(\$input: CreateQuestionInput!) { createQuestion(input: \$input) { id content } }\",
    \"variables\": {
      \"input\": {
        \"quizID\": \"$QUIZ_ID\",
        \"type\": \"MULTIPLE_CHOICE\",
        \"content\": \"What are traits in Rust?\",
        \"options\": [
          \"A way to define shared behavior across types\",
          \"Special variables that store metadata\",
          \"Functions that can only be called once\",
          \"Rust's version of classes\"
        ],
        \"correctAnswer\": \"A way to define shared behavior across types\",
        \"explanation\": \"Traits are Rust's way of defining shared behavior. They are similar to interfaces in other languages and allow you to define method signatures that types must implement.\",
        \"difficulty\": \"EASY\"
      }
    }
  }"
echo "Question 5 created"

# Question 6: Match
curl -s -X POST "$ENDPOINT" \
  -H "Content-Type: application/json" \
  -d "{
    \"query\": \"mutation CreateQuestion(\$input: CreateQuestionInput!) { createQuestion(input: \$input) { id content } }\",
    \"variables\": {
      \"input\": {
        \"quizID\": \"$QUIZ_ID\",
        \"type\": \"TRUE_FALSE\",
        \"content\": \"In Rust, match expressions must be exhaustive, covering all possible cases.\",
        \"correctAnswer\": \"true\",
        \"explanation\": \"Match expressions in Rust must be exhaustive. The compiler checks that all possible cases are handled, which helps prevent bugs. You can use _ as a catch-all pattern if needed.\",
        \"difficulty\": \"EASY\"
      }
    }
  }"
echo "Question 6 created"

# Question 7: Cargo
curl -s -X POST "$ENDPOINT" \
  -H "Content-Type: application/json" \
  -d "{
    \"query\": \"mutation CreateQuestion(\$input: CreateQuestionInput!) { createQuestion(input: \$input) { id content } }\",
    \"variables\": {
      \"input\": {
        \"quizID\": \"$QUIZ_ID\",
        \"type\": \"MULTIPLE_CHOICE\",
        \"content\": \"What is Cargo in the Rust ecosystem?\",
        \"options\": [
          \"Rust's package manager and build system\",
          \"A web framework for Rust\",
          \"A testing framework\",
          \"Rust's compiler\"
        ],
        \"correctAnswer\": \"Rust's package manager and build system\",
        \"explanation\": \"Cargo is Rust's build system and package manager. It handles building code, downloading dependencies, and building those dependencies.\",
        \"difficulty\": \"EASY\"
      }
    }
  }"
echo "Question 7 created"

# Question 8: Macros
curl -s -X POST "$ENDPOINT" \
  -H "Content-Type: application/json" \
  -d "{
    \"query\": \"mutation CreateQuestion(\$input: CreateQuestionInput!) { createQuestion(input: \$input) { id content } }\",
    \"variables\": {
      \"input\": {
        \"quizID\": \"$QUIZ_ID\",
        \"type\": \"MULTIPLE_CHOICE\",
        \"content\": \"What distinguishes a macro call from a function call in Rust?\",
        \"options\": [
          \"Macros end with an exclamation mark (!)\",
          \"Macros are always faster\",
          \"Macros can only be defined in the standard library\",
          \"There is no difference\"
        ],
        \"correctAnswer\": \"Macros end with an exclamation mark (!)\",
        \"explanation\": \"In Rust, macro calls are distinguished by the exclamation mark. For example, println! and vec! are macros. Macros are expanded at compile time and can take a variable number of arguments.\",
        \"difficulty\": \"EASY\"
      }
    }
  }"
echo "Question 8 created"

# Question 9: String vs &str
curl -s -X POST "$ENDPOINT" \
  -H "Content-Type: application/json" \
  -d "{
    \"query\": \"mutation CreateQuestion(\$input: CreateQuestionInput!) { createQuestion(input: \$input) { id content } }\",
    \"variables\": {
      \"input\": {
        \"quizID\": \"$QUIZ_ID\",
        \"type\": \"MULTIPLE_CHOICE\",
        \"content\": \"What is the main difference between String and &str in Rust?\",
        \"options\": [
          \"String is heap-allocated and growable, &str is a string slice reference\",
          \"String is faster than &str\",
          \"&str can only contain ASCII characters\",
          \"There is no difference\"
        ],
        \"correctAnswer\": \"String is heap-allocated and growable, &str is a string slice reference\",
        \"explanation\": \"String is an owned, heap-allocated, growable string type. &str is a string slice, which is a reference to a sequence of UTF-8 bytes, and is typically used for string views.\",
        \"difficulty\": \"MEDIUM\"
      }
    }
  }"
echo "Question 9 created"

# Question 10: Concurrency
curl -s -X POST "$ENDPOINT" \
  -H "Content-Type: application/json" \
  -d "{
    \"query\": \"mutation CreateQuestion(\$input: CreateQuestionInput!) { createQuestion(input: \$input) { id content } }\",
    \"variables\": {
      \"input\": {
        \"quizID\": \"$QUIZ_ID\",
        \"type\": \"MULTIPLE_CHOICE\",
        \"content\": \"How does Rust prevent data races in concurrent programming?\",
        \"options\": [
          \"Through its ownership and type system at compile time\",
          \"By using a global interpreter lock (GIL)\",
          \"By preventing the use of multiple threads\",
          \"By requiring all shared data to be immutable\"
        ],
        \"correctAnswer\": \"Through its ownership and type system at compile time\",
        \"explanation\": \"Rust's ownership system and type system prevent data races at compile time. The Send and Sync traits, combined with borrowing rules, ensure thread safety without runtime overhead.\",
        \"difficulty\": \"HARD\"
      }
    }
  }"
echo "Question 10 created"

echo "✅ Successfully created Rust quiz with 10 questions!"
