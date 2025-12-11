import { Suspense } from 'react'
import { graphql, useLazyLoadQuery } from 'react-relay'
import { Link } from 'react-router-dom'

const QuizListQuery = graphql`
  query QuizListQuery {
    quizzes {
      id
      title
      description
      createdAt
      tags {
        id
        name
      }
    }
  }
`

function QuizListContent() {
  const data = useLazyLoadQuery<any>(QuizListQuery, {})

  return (
    <div>
      <h2>クイズ一覧</h2>
      {data.quizzes.length === 0 ? (
        <p>
          クイズがまだありません。<Link to="/create">新しく作成</Link>してください。
        </p>
      ) : (
        <div className="quiz-grid">
          {data.quizzes.map((quiz: any) => (
            <div key={quiz.id} className="card">
              <h3>{quiz.title}</h3>
              <p>{quiz.description}</p>
              <div className="tags">
                {quiz.tags.map((tag: any) => (
                  <span key={tag.id} className="tag">
                    {tag.name}
                  </span>
                ))}
              </div>
              <div className="actions">
                <Link to={`/quiz/${quiz.id}`} className="btn btn-primary">
                  詳細
                </Link>
                <Link to={`/take/${quiz.id}`} className="btn btn-success">
                  開始
                </Link>
              </div>
            </div>
          ))}
        </div>
      )}
    </div>
  )
}

export default function QuizList() {
  return (
    <Suspense fallback={<div>読み込み中...</div>}>
      <QuizListContent />
    </Suspense>
  )
}
