import { Suspense, useState } from 'react'
import { graphql, useLazyLoadQuery, useMutation } from 'react-relay'
import { Link, useNavigate, useParams } from 'react-router-dom'

const QuizDetailQuery = graphql`
  query QuizDetailQuery($id: ID!) {
    quiz(id: $id) {
      id
      title
      description
      createdAt
      tags {
        id
        name
      }
      questions {
        id
        type
        content
        difficulty
        options
      }
    }
  }
`

const DeleteQuizMutation = graphql`
  mutation QuizDetailDeleteMutation($id: ID!) {
    deleteQuiz(id: $id)
  }
`

const CreateQuestionMutation = graphql`
  mutation QuizDetailCreateQuestionMutation($input: CreateQuestionInput!) {
    createQuestion(input: $input) {
      id
      type
      content
      difficulty
      options
    }
  }
`

const DeleteQuestionMutation = graphql`
  mutation QuizDetailDeleteQuestionMutation($id: ID!) {
    deleteQuestion(id: $id)
  }
`

function QuizDetailContent() {
  const { id } = useParams<{ id: string }>()
  const navigate = useNavigate()
  const data = useLazyLoadQuery<any>(QuizDetailQuery, { id: id! })
  const [commitDeleteQuiz] = useMutation(DeleteQuizMutation)
  const [commitCreateQuestion] = useMutation(CreateQuestionMutation)
  const [commitDeleteQuestion] = useMutation(DeleteQuestionMutation)

  const [showQuestionForm, setShowQuestionForm] = useState(false)
  const [questionType, setQuestionType] = useState<
    'MULTIPLE_CHOICE' | 'TRUE_FALSE' | 'SHORT_ANSWER'
  >('MULTIPLE_CHOICE')
  const [questionContent, setQuestionContent] = useState('')
  const [questionOptions, setQuestionOptions] = useState(['', '', '', ''])
  const [correctAnswer, setCorrectAnswer] = useState('')
  const [explanation, setExplanation] = useState('')
  const [difficulty, setDifficulty] = useState<'EASY' | 'MEDIUM' | 'HARD'>('MEDIUM')

  const quiz = data.quiz

  if (!quiz) {
    return <div className="card">クイズが見つかりません</div>
  }

  const handleDeleteQuiz = () => {
    if (!confirm('このクイズを削除しますか？')) return

    commitDeleteQuiz({
      variables: { id: id! },
      onCompleted: () => {
        navigate('/')
      },
      onError: (error) => {
        alert(`削除に失敗しました: ${error.message}`)
      },
    })
  }

  const handleCreateQuestion = (e: React.FormEvent) => {
    e.preventDefault()

    const input: any = {
      quizID: id!,
      type: questionType,
      content: questionContent,
      correctAnswer,
      difficulty,
      explanation: explanation || undefined,
    }

    if (questionType === 'MULTIPLE_CHOICE') {
      input.options = questionOptions.filter((opt) => opt.trim())
    }

    commitCreateQuestion({
      variables: { input },
      onCompleted: () => {
        setShowQuestionForm(false)
        setQuestionContent('')
        setQuestionOptions(['', '', '', ''])
        setCorrectAnswer('')
        setExplanation('')
        setDifficulty('MEDIUM')
      },
      onError: (error) => {
        alert(`問題の作成に失敗しました: ${error.message}`)
      },
      updater: (store) => {
        const quizRecord = store.get(id!)
        if (quizRecord) {
          const newQuestion = store.getRootField('createQuestion')
          const questions = quizRecord.getLinkedRecords('questions') || []
          quizRecord.setLinkedRecords([...questions, newQuestion], 'questions')
        }
      },
    })
  }

  const handleDeleteQuestion = (questionId: string) => {
    if (!confirm('この問題を削除しますか？')) return

    commitDeleteQuestion({
      variables: { id: questionId },
      onCompleted: () => {},
      onError: (error) => {
        alert(`削除に失敗しました: ${error.message}`)
      },
      updater: (store) => {
        const quizRecord = store.get(id!)
        if (quizRecord) {
          const questions = quizRecord.getLinkedRecords('questions') || []
          const filtered = questions.filter((q) => q.getDataID() !== questionId)
          quizRecord.setLinkedRecords(filtered, 'questions')
        }
      },
    })
  }

  const updateOption = (index: number, value: string) => {
    const newOptions = [...questionOptions]
    newOptions[index] = value
    setQuestionOptions(newOptions)
  }

  const getQuestionTypeLabel = (type: string) => {
    switch (type) {
      case 'MULTIPLE_CHOICE':
        return '選択式'
      case 'TRUE_FALSE':
        return '○×'
      case 'SHORT_ANSWER':
        return '記述式'
      default:
        return type
    }
  }

  const getDifficultyLabel = (diff: string) => {
    switch (diff) {
      case 'EASY':
        return '易'
      case 'MEDIUM':
        return '中'
      case 'HARD':
        return '難'
      default:
        return diff
    }
  }

  return (
    <div>
      <div className="card">
        <div className="quiz-header">
          <h2>{quiz.title}</h2>
          <div className="actions">
            <Link to={`/take/${id}`} className="btn btn-success">
              クイズを開始
            </Link>
            <button onClick={handleDeleteQuiz} className="btn btn-danger">
              削除
            </button>
          </div>
        </div>
        {quiz.description && <p>{quiz.description}</p>}
        <div className="tags">
          {quiz.tags.map((tag: any) => (
            <span key={tag.id} className="tag">
              {tag.name}
            </span>
          ))}
        </div>
      </div>

      <div className="card">
        <div className="section-header">
          <h3>問題一覧 ({quiz.questions.length}問)</h3>
          <button
            className="btn btn-primary"
            onClick={() => setShowQuestionForm(!showQuestionForm)}
          >
            {showQuestionForm ? 'キャンセル' : '問題を追加'}
          </button>
        </div>

        {showQuestionForm && (
          <form onSubmit={handleCreateQuestion} className="question-form">
            <div className="form-row">
              <div className="form-group">
                <label>問題タイプ</label>
                <select
                  value={questionType}
                  onChange={(e) => setQuestionType(e.target.value as any)}
                >
                  <option value="MULTIPLE_CHOICE">選択式</option>
                  <option value="TRUE_FALSE">○×</option>
                  <option value="SHORT_ANSWER">記述式</option>
                </select>
              </div>
              <div className="form-group">
                <label>難易度</label>
                <select value={difficulty} onChange={(e) => setDifficulty(e.target.value as any)}>
                  <option value="EASY">易</option>
                  <option value="MEDIUM">中</option>
                  <option value="HARD">難</option>
                </select>
              </div>
            </div>

            <div className="form-group">
              <label>問題文 *</label>
              <textarea
                value={questionContent}
                onChange={(e) => setQuestionContent(e.target.value)}
                required
                placeholder="問題文を入力"
              />
            </div>

            {questionType === 'MULTIPLE_CHOICE' && (
              <div className="form-group">
                <label>選択肢</label>
                {questionOptions.map((option, index) => (
                  <input
                    key={index}
                    type="text"
                    value={option}
                    onChange={(e) => updateOption(index, e.target.value)}
                    placeholder={`選択肢 ${index + 1}`}
                  />
                ))}
              </div>
            )}

            <div className="form-group">
              <label>正解 *</label>
              {questionType === 'TRUE_FALSE' ? (
                <select
                  value={correctAnswer}
                  onChange={(e) => setCorrectAnswer(e.target.value)}
                  required
                >
                  <option value="">選択してください</option>
                  <option value="true">○ (正しい)</option>
                  <option value="false">× (誤り)</option>
                </select>
              ) : (
                <input
                  type="text"
                  value={correctAnswer}
                  onChange={(e) => setCorrectAnswer(e.target.value)}
                  required
                  placeholder="正解を入力"
                />
              )}
            </div>

            <div className="form-group">
              <label>解説</label>
              <textarea
                value={explanation}
                onChange={(e) => setExplanation(e.target.value)}
                placeholder="解説を入力（任意）"
              />
            </div>

            <button type="submit" className="btn btn-primary">
              追加
            </button>
          </form>
        )}

        <div className="questions-list">
          {quiz.questions.length === 0 ? (
            <p className="empty-message">問題がまだありません</p>
          ) : (
            quiz.questions.map((question: any, index: number) => (
              <div key={question.id} className="question-item">
                <div className="question-header">
                  <span className="question-number">Q{index + 1}</span>
                  <span className="question-type">{getQuestionTypeLabel(question.type)}</span>
                  <span className={`difficulty difficulty-${question.difficulty.toLowerCase()}`}>
                    {getDifficultyLabel(question.difficulty)}
                  </span>
                  <button
                    onClick={() => handleDeleteQuestion(question.id)}
                    className="btn-icon btn-danger-icon"
                  >
                    削除
                  </button>
                </div>
                <p className="question-content">{question.content}</p>
                {question.options && question.options.length > 0 && (
                  <ul className="question-options">
                    {question.options.map((option: string, i: number) => (
                      <li key={i}>{option}</li>
                    ))}
                  </ul>
                )}
              </div>
            ))
          )}
        </div>
      </div>
    </div>
  )
}

export default function QuizDetail() {
  return (
    <Suspense fallback={<div className="card">読み込み中...</div>}>
      <QuizDetailContent />
    </Suspense>
  )
}
