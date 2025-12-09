import { Suspense } from 'react'
import { graphql, useLazyLoadQuery } from 'react-relay'
import { Link } from 'react-router-dom'

const StatisticsQuery = graphql`
  query StatisticsQuery {
    statistics {
      totalAttempts
      averageScore
      categoryStats {
        tagName
        correctRate
        totalQuestions
      }
      recentAttempts {
        id
        quizID
        score
        totalQuestions
        completedAt
      }
    }
  }
`

function StatisticsContent() {
  const data = useLazyLoadQuery<any>(StatisticsQuery, {})
  const stats = data.statistics

  const formatDate = (dateString: string) => {
    const date = new Date(dateString)
    return date.toLocaleDateString('ja-JP', {
      year: 'numeric',
      month: '2-digit',
      day: '2-digit',
      hour: '2-digit',
      minute: '2-digit',
    })
  }

  const getScoreClass = (score: number, total: number) => {
    const percentage = (score / total) * 100
    if (percentage >= 80) return 'score-excellent'
    if (percentage >= 60) return 'score-good'
    return 'score-poor'
  }

  return (
    <div>
      <div className="card">
        <h2>学習統計</h2>
        <div className="stats-overview">
          <div className="stat-item">
            <div className="stat-label">総挑戦回数</div>
            <div className="stat-value">{stats.totalAttempts}回</div>
          </div>
          <div className="stat-item">
            <div className="stat-label">平均スコア</div>
            <div className="stat-value">{stats.averageScore.toFixed(1)}点</div>
          </div>
        </div>
      </div>

      {stats.categoryStats.length > 0 && (
        <div className="card">
          <h3>カテゴリー別統計</h3>
          <div className="category-stats">
            {stats.categoryStats.map((cat: any, index: number) => {
              const correctRate = (cat.correctRate * 100).toFixed(1)
              return (
                <div key={index} className="category-stat-item">
                  <div className="category-header">
                    <span className="tag">{cat.tagName}</span>
                    <span className="category-questions">{cat.totalQuestions}問</span>
                  </div>
                  <div className="category-rate">
                    <div className="rate-bar">
                      <div className="rate-fill" style={{ width: `${correctRate}%` }} />
                    </div>
                    <span className="rate-text">{correctRate}%</span>
                  </div>
                </div>
              )
            })}
          </div>
        </div>
      )}

      {stats.recentAttempts.length > 0 && (
        <div className="card">
          <h3>最近の挑戦</h3>
          <div className="recent-attempts">
            {stats.recentAttempts.map((attempt: any) => {
              const percentage = Math.round((attempt.score / attempt.totalQuestions) * 100)
              return (
                <div key={attempt.id} className="attempt-item">
                  <div className="attempt-info">
                    <Link to={`/quiz/${attempt.quizID}`} className="attempt-quiz-link">
                      クイズ #{attempt.quizID}
                    </Link>
                    <span className="attempt-date">{formatDate(attempt.completedAt)}</span>
                  </div>
                  <div
                    className={`attempt-score ${getScoreClass(attempt.score, attempt.totalQuestions)}`}
                  >
                    <span className="score-percentage">{percentage}点</span>
                    <span className="score-fraction">
                      {attempt.score} / {attempt.totalQuestions}
                    </span>
                  </div>
                </div>
              )
            })}
          </div>
        </div>
      )}

      {stats.totalAttempts === 0 && (
        <div className="card empty-state">
          <p>まだクイズに挑戦していません</p>
          <Link to="/" className="btn btn-primary">
            クイズ一覧へ
          </Link>
        </div>
      )}
    </div>
  )
}

export default function Statistics() {
  return (
    <Suspense fallback={<div className="card">読み込み中...</div>}>
      <StatisticsContent />
    </Suspense>
  )
}
