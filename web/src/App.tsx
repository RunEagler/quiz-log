import { BrowserRouter, Link, Route, Routes } from 'react-router-dom'
import CreateQuiz from './pages/CreateQuiz'
import QuizDetail from './pages/QuizDetail'
import QuizList from './pages/QuizList'
import Statistics from './pages/Statistics'
import TakeQuiz from './pages/TakeQuiz'

function App() {
  return (
    <BrowserRouter>
      <div className="app">
        <nav className="navbar">
          <h1>Quiz Log</h1>
          <ul>
            <li>
              <Link to="/">クイズ一覧</Link>
            </li>
            <li>
              <Link to="/create">クイズ作成</Link>
            </li>
            <li>
              <Link to="/statistics">統計</Link>
            </li>
          </ul>
        </nav>
        <main className="container">
          <Routes>
            <Route path="/" element={<QuizList />} />
            <Route path="/quiz/:id" element={<QuizDetail />} />
            <Route path="/create" element={<CreateQuiz />} />
            <Route path="/take/:id" element={<TakeQuiz />} />
            <Route path="/statistics" element={<Statistics />} />
          </Routes>
        </main>
      </div>
    </BrowserRouter>
  )
}

export default App
