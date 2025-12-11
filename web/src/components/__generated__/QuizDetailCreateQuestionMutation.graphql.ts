/**
 * @generated SignedSource<<7e9af30faf02284ad77153c0ff0641a0>>
 * @lightSyntaxTransform
 * @nogrep
 */

/* tslint:disable */
/* eslint-disable */
// @ts-nocheck

import type { ConcreteRequest } from 'relay-runtime'
export type Difficulty = 'EASY' | 'HARD' | 'MEDIUM' | '%future added value'
export type QuestionType = 'MULTIPLE_CHOICE' | 'SHORT_ANSWER' | 'TRUE_FALSE' | '%future added value'
export type CreateQuestionInput = {
  content: string
  correctAnswer: string
  difficulty: Difficulty
  explanation?: string | null | undefined
  options?: ReadonlyArray<string> | null | undefined
  quizID: string
  tagIDs?: ReadonlyArray<string> | null | undefined
  type: QuestionType
}
export type QuizDetailCreateQuestionMutation$variables = {
  input: CreateQuestionInput
}
export type QuizDetailCreateQuestionMutation$data = {
  readonly createQuestion: {
    readonly content: string
    readonly difficulty: Difficulty
    readonly id: string
    readonly options: ReadonlyArray<string> | null | undefined
    readonly type: QuestionType
  }
}
export type QuizDetailCreateQuestionMutation = {
  response: QuizDetailCreateQuestionMutation$data
  variables: QuizDetailCreateQuestionMutation$variables
}

const node: ConcreteRequest = (() => {
  var v0 = [
      {
        defaultValue: null,
        kind: 'LocalArgument',
        name: 'input',
      },
    ],
    v1 = [
      {
        alias: null,
        args: [
          {
            kind: 'Variable',
            name: 'input',
            variableName: 'input',
          },
        ],
        concreteType: 'Question',
        kind: 'LinkedField',
        name: 'createQuestion',
        plural: false,
        selections: [
          {
            alias: null,
            args: null,
            kind: 'ScalarField',
            name: 'id',
            storageKey: null,
          },
          {
            alias: null,
            args: null,
            kind: 'ScalarField',
            name: 'type',
            storageKey: null,
          },
          {
            alias: null,
            args: null,
            kind: 'ScalarField',
            name: 'content',
            storageKey: null,
          },
          {
            alias: null,
            args: null,
            kind: 'ScalarField',
            name: 'difficulty',
            storageKey: null,
          },
          {
            alias: null,
            args: null,
            kind: 'ScalarField',
            name: 'options',
            storageKey: null,
          },
        ],
        storageKey: null,
      },
    ]
  return {
    fragment: {
      argumentDefinitions: v0 /*: any*/,
      kind: 'Fragment',
      metadata: null,
      name: 'QuizDetailCreateQuestionMutation',
      selections: v1 /*: any*/,
      type: 'Mutation',
      abstractKey: null,
    },
    kind: 'Request',
    operation: {
      argumentDefinitions: v0 /*: any*/,
      kind: 'Operation',
      name: 'QuizDetailCreateQuestionMutation',
      selections: v1 /*: any*/,
    },
    params: {
      cacheID: '759ec60298d5ea52be9f9aa036a41782',
      id: null,
      metadata: {},
      name: 'QuizDetailCreateQuestionMutation',
      operationKind: 'mutation',
      text: 'mutation QuizDetailCreateQuestionMutation(\n  $input: CreateQuestionInput!\n) {\n  createQuestion(input: $input) {\n    id\n    type\n    content\n    difficulty\n    options\n  }\n}\n',
    },
  }
})()

;(node as any).hash = '075dd2eefff767e0ab837b46272fff9d'

export default node
