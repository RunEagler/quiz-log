import { Suspense, useState } from 'react'
import { useLazyLoadQuery, useMutation, graphql } from 'react-relay'
import { useParams, useNavigate } from 'react-router-dom'
import type { TakeQuizSubmitMutation as TakeQuizSubmitMutationType } from './__generated__/TakeQuizSubmitMutation.graphql'

const TakeQuizQuery = graphql`
  query TakeQuizQuery($id: ID!) {
    quiz(id: $id) {
      id
      title
      description
      questions {
        id
        type
        content
        options
        explanation
        difficulty
      }
    }
  }
`

const SubmitAttemptMutation = graphql`
  mutation TakeQuizSubmitMutation($input: SubmitAttemptInput!) {
    submitAttempt(input: $input) {
      attempt {
        id
        score
        totalQuestions
      }
      score
      totalQuestions
      correctCount
      wrongQuestions {
        id
        content
        correctAnswer
        explanation
      }
    }
  }
`

interface UserAnswer {
  questionID: string
  userAnswer: string
}

function TakeQuizContent() {
  const { id } = useParams<{ id: string }>()
  const navigate = useNavigate()
  const data = useLazyLoadQuery<any>(TakeQuizQuery, { id: id! })
  const [commitSubmit, isSubmitting] = useMutation<TakeQuizSubmitMutationType>(SubmitAttemptMutation)

  const [currentQuestionIndex, setCurrentQuestionIndex] = useState(0)
  const [answers, setAnswers] = useState<Map<string, string>>(new Map())
  const [showResult, setShowResult] = useState(false)
  const [result, setResult] = useState<any>(null)

  const quiz = data.quiz

  if (!quiz) {
    return <div className="card">クイズが見つかりません</div>
  }

  if (quiz.questions.length === 0) {
    return (
      <div className="card">
        <h2>{quiz.title}</h2>
        <p>このクイズにはまだ問題がありません</p>
        <button onClick={() => navigate(`/quiz/${id}`)} className="btn btn-primary">
          問題を追加
        </button>
      </div>
    )
  }

  const currentQuestion = quiz.questions[currentQuestionIndex]
  const isLastQuestion = currentQuestionIndex === quiz.questions.length - 1
  const currentAnswer = answers.get(currentQuestion.id) || ''

  const handleAnswerChange = (value: string) => {
    const newAnswers = new Map(answers)
    newAnswers.set(currentQuestion.id, value)
    setAnswers(newAnswers)
  }

  const handleNext = () => {
    if (currentQuestionIndex < quiz.questions.length - 1) {
      setCurrentQuestionIndex(currentQuestionIndex + 1)
    }
  }

  const handlePrevious = () => {
    if (currentQuestionIndex > 0) {
      setCurrentQuestionIndex(currentQuestionIndex - 1)
    }
  }

  const handleSubmit = () => {
    const answersArray: UserAnswer[] = quiz.questions.map((q: any) => ({
      questionID: q.id,
      userAnswer: answers.get(q.id) || '',
    }))

    commitSubmit({
      variables: {
        input: {
          quizID: id!,
          answers: answersArray,
        },
      },
      onCompleted: (response) => {
        setResult(response.submitAttempt)
        setShowResult(true)
      },
      onError: (error) => {
        alert('提出に失敗しました: ' + error.message)
      },
    })
  }

  const getQuestionTypeLabel = (type: string) => {
    switch (type) {
      case 'MULTIPLE_CHOICE': return '選択式'
      case 'TRUE_FALSE': return '○×'
      case 'SHORT_ANSWER': return '記述式'
      default: return type
    }
  }

  const getDifficultyLabel = (diff: string) => {
    switch (diff) {
      case 'EASY': return '易'
      case 'MEDIUM': return '中'
      case 'HARD': return '難'
      default: return diff
    }
  }

  const answeredCount = answers.size

  if (showResult && result) {
    const percentage = Math.round((result.correctCount / result.totalQuestions) * 100)

    return (
      <div>
        <div className="card result-card">
          <h2>結果</h2>
          <div className="result-summary">
            <div className="score-display">
              <div className="score-number">{percentage}点</div>
              <div className="score-detail">
                {result.correctCount} / {result.totalQuestions} 問正解
              </div>
            </div>
          </div>

          <div className="result-actions">
            <button onClick={() => navigate(`/quiz/${id}`)} className="btn btn-primary">
              クイズ詳細に戻る
            </button>
            <button onClick={() => navigate('/')} className="btn btn-secondary">
              クイズ一覧へ
            </button>
            <button
              onClick={() => {
                setShowResult(false)
                setAnswers(new Map())
                setCurrentQuestionIndex(0)
              }}
              className="btn btn-success"
            >
              もう一度挑戦
            </button>
          </div>
        </div>

        {result.wrongQuestions.length > 0 && (
          <div className="card">
            <h3>間違えた問題</h3>
            <div className="wrong-questions">
              {result.wrongQuestions.map((question: any, index: number) => (
                <div key={question.id} className="wrong-question-item">
                  <h4>問題 {index + 1}</h4>
                  <p className="question-content">{question.content}</p>
                  <div className="answer-info">
                    <div className="correct-answer">
                      <strong>正解:</strong> {question.correctAnswer}
                    </div>
                    {question.explanation && (
                      <div className="explanation">
                        <strong>解説:</strong> {question.explanation}
                      </div>
                    )}
                  </div>
                </div>
              ))}
            </div>
          </div>
        )}
      </div>
    )
  }

  return (
    <div>
      <div className="card">
        <h2>{quiz.title}</h2>
        {quiz.description && <p>{quiz.description}</p>}

        <div className="progress-bar">
          <div className="progress-info">
            <span>問題 {currentQuestionIndex + 1} / {quiz.questions.length}</span>
            <span>回答済み: {answeredCount} / {quiz.questions.length}</span>
          </div>
          <div className="progress-track">
            <div
              className="progress-fill"
              style={{ width: `${((currentQuestionIndex + 1) / quiz.questions.length) * 100}%` }}
            />
          </div>
        </div>
      </div>

      <div className="card question-card">
        <div className="question-meta">
          <span className="question-type">{getQuestionTypeLabel(currentQuestion.type)}</span>
          <span className={`difficulty difficulty-${currentQuestion.difficulty.toLowerCase()}`}>
            {getDifficultyLabel(currentQuestion.difficulty)}
          </span>
        </div>

        <h3 className="question-content">{currentQuestion.content}</h3>

        <div className="answer-section">
          {currentQuestion.type === 'MULTIPLE_CHOICE' && currentQuestion.options && (
            <div className="options-list">
              {currentQuestion.options.map((option: string, index: number) => (
                <label key={index} className="option-item">
                  <input
                    type="radio"
                    name="answer"
                    value={option}
                    checked={currentAnswer === option}
                    onChange={(e) => handleAnswerChange(e.target.value)}
                  />
                  <span className="option-text">{option}</span>
                </label>
              ))}
            </div>
          )}

          {currentQuestion.type === 'TRUE_FALSE' && (
            <div className="options-list">
              <label className="option-item">
                <input
                  type="radio"
                  name="answer"
                  value="true"
                  checked={currentAnswer === 'true'}
                  onChange={(e) => handleAnswerChange(e.target.value)}
                />
                <span className="option-text">○ (正しい)</span>
              </label>
              <label className="option-item">
                <input
                  type="radio"
                  name="answer"
                  value="false"
                  checked={currentAnswer === 'false'}
                  onChange={(e) => handleAnswerChange(e.target.value)}
                />
                <span className="option-text">× (誤り)</span>
              </label>
            </div>
          )}

          {currentQuestion.type === 'SHORT_ANSWER' && (
            <div className="form-group">
              <textarea
                value={currentAnswer}
                onChange={(e) => handleAnswerChange(e.target.value)}
                placeholder="回答を入力してください"
                rows={4}
              />
            </div>
          )}
        </div>

        <div className="navigation-buttons">
          <button
            onClick={handlePrevious}
            disabled={currentQuestionIndex === 0}
            className="btn btn-secondary"
          >
            前へ
          </button>

          {isLastQuestion ? (
            <button
              onClick={handleSubmit}
              disabled={isSubmitting}
              className="btn btn-success"
            >
              提出する
            </button>
          ) : (
            <button
              onClick={handleNext}
              className="btn btn-primary"
            >
              次へ
            </button>
          )}
        </div>

        <div className="question-navigation">
          {quiz.questions.map((q: any, index: number) => (
            <button
              key={q.id}
              onClick={() => setCurrentQuestionIndex(index)}
              className={`question-nav-button ${
                index === currentQuestionIndex ? 'active' : ''
              } ${answers.has(q.id) ? 'answered' : ''}`}
            >
              {index + 1}
            </button>
          ))}
        </div>
      </div>
    </div>
  )
}

export default function TakeQuiz() {
  return (
    <Suspense fallback={<div className="card">読み込み中...</div>}>
      <TakeQuizContent />
    </Suspense>
  )
}
