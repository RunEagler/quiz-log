import { BrowserRouter, Routes, Route, Link } from 'react-router-dom'
import QuizList from './components/QuizList'
import QuizDetail from './components/QuizDetail'
import CreateQuiz from './components/CreateQuiz'
import TakeQuiz from './components/TakeQuiz'
import Statistics from './components/Statistics'

function App() {
  return (
    <BrowserRouter>
      <div className="app">
        <nav className="navbar">
          <h1>Quiz Log</h1>
          <ul>
            <li><Link to="/">クイズ一覧</Link></li>
            <li><Link to="/create">クイズ作成</Link></li>
            <li><Link to="/statistics">統計</Link></li>
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
