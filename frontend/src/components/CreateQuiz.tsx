import { Suspense, useState } from 'react'
import { useLazyLoadQuery, useMutation, graphql } from 'react-relay'
import { useNavigate } from 'react-router-dom'
import type { CreateQuizMutation as CreateQuizMutationType } from './__generated__/CreateQuizMutation.graphql'

const TagsQuery = graphql`
  query CreateQuizTagsQuery {
    tags {
      id
      name
    }
  }
`

const CreateQuizMutation = graphql`
  mutation CreateQuizMutation($input: CreateQuizInput!) {
    createQuiz(input: $input) {
      id
      title
      description
    }
  }
`

const CreateTagMutation = graphql`
  mutation CreateQuizCreateTagMutation($name: String!) {
    createTag(name: $name) {
      id
      name
    }
  }
`

function CreateQuizForm() {
  const data = useLazyLoadQuery<any>(TagsQuery, {})
  const [commitQuiz, isQuizInFlight] = useMutation<CreateQuizMutationType>(CreateQuizMutation)
  const [commitTag, isTagInFlight] = useMutation(CreateTagMutation)
  const navigate = useNavigate()

  const [title, setTitle] = useState('')
  const [description, setDescription] = useState('')
  const [selectedTags, setSelectedTags] = useState<string[]>([])
  const [newTagName, setNewTagName] = useState('')

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault()

    commitQuiz({
      variables: {
        input: {
          title,
          description: description || undefined,
          tagIDs: selectedTags.length > 0 ? selectedTags : undefined,
        },
      },
      onCompleted: (response) => {
        navigate(`/quiz/${response.createQuiz.id}`)
      },
      onError: (error) => {
        alert('クイズの作成に失敗しました: ' + error.message)
      },
    })
  }

  const handleCreateTag = () => {
    if (!newTagName.trim()) return

    commitTag({
      variables: { name: newTagName },
      onCompleted: () => {
        setNewTagName('')
      },
      onError: (error) => {
        alert('タグの作成に失敗しました: ' + error.message)
      },
      updater: (store) => {
        const root = store.getRoot()
        const newTag = store.getRootField('createTag')
        const tags = root.getLinkedRecords('tags') || []
        root.setLinkedRecords([...tags, newTag], 'tags')
      },
    })
  }

  const toggleTag = (tagId: string) => {
    setSelectedTags(prev =>
      prev.includes(tagId)
        ? prev.filter(id => id !== tagId)
        : [...prev, tagId]
    )
  }

  return (
    <div className="card">
      <h2>新規クイズ作成</h2>
      <form onSubmit={handleSubmit}>
        <div className="form-group">
          <label htmlFor="title">タイトル *</label>
          <input
            id="title"
            type="text"
            value={title}
            onChange={(e) => setTitle(e.target.value)}
            required
            placeholder="クイズのタイトルを入力"
          />
        </div>

        <div className="form-group">
          <label htmlFor="description">説明</label>
          <textarea
            id="description"
            value={description}
            onChange={(e) => setDescription(e.target.value)}
            placeholder="クイズの説明を入力（任意）"
          />
        </div>

        <div className="form-group">
          <label>タグ</label>
          <div className="tags-selection">
            {data.tags.map((tag: any) => (
              <button
                key={tag.id}
                type="button"
                className={`tag-button ${selectedTags.includes(tag.id) ? 'selected' : ''}`}
                onClick={() => toggleTag(tag.id)}
              >
                {tag.name}
              </button>
            ))}
          </div>
          <div className="new-tag-form">
            <input
              type="text"
              value={newTagName}
              onChange={(e) => setNewTagName(e.target.value)}
              placeholder="新しいタグを追加"
            />
            <button
              type="button"
              className="btn btn-secondary"
              onClick={handleCreateTag}
              disabled={isTagInFlight || !newTagName.trim()}
            >
              追加
            </button>
          </div>
        </div>

        <div className="form-actions">
          <button
            type="submit"
            className="btn btn-primary"
            disabled={isQuizInFlight || !title.trim()}
          >
            作成
          </button>
          <button
            type="button"
            className="btn btn-secondary"
            onClick={() => navigate('/')}
          >
            キャンセル
          </button>
        </div>
      </form>
    </div>
  )
}

export default function CreateQuiz() {
  return (
    <Suspense fallback={<div className="card">読み込み中...</div>}>
      <CreateQuizForm />
    </Suspense>
  )
}