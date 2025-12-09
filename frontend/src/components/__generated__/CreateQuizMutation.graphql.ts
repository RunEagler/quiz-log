/**
 * @generated SignedSource<<dec8a78f5892ae43bbff5af32d9e8bc1>>
 * @lightSyntaxTransform
 * @nogrep
 */

/* tslint:disable */
/* eslint-disable */
// @ts-nocheck

import type { ConcreteRequest } from 'relay-runtime'
export type CreateQuizInput = {
  description?: string | null | undefined
  tagIDs?: ReadonlyArray<string> | null | undefined
  title: string
}
export type CreateQuizMutation$variables = {
  input: CreateQuizInput
}
export type CreateQuizMutation$data = {
  readonly createQuiz: {
    readonly description: string | null | undefined
    readonly id: string
    readonly title: string
  }
}
export type CreateQuizMutation = {
  response: CreateQuizMutation$data
  variables: CreateQuizMutation$variables
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
        concreteType: 'Quiz',
        kind: 'LinkedField',
        name: 'createQuiz',
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
            name: 'title',
            storageKey: null,
          },
          {
            alias: null,
            args: null,
            kind: 'ScalarField',
            name: 'description',
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
      name: 'CreateQuizMutation',
      selections: v1 /*: any*/,
      type: 'Mutation',
      abstractKey: null,
    },
    kind: 'Request',
    operation: {
      argumentDefinitions: v0 /*: any*/,
      kind: 'Operation',
      name: 'CreateQuizMutation',
      selections: v1 /*: any*/,
    },
    params: {
      cacheID: '819521afc10d1154e6384e2fb9d74462',
      id: null,
      metadata: {},
      name: 'CreateQuizMutation',
      operationKind: 'mutation',
      text: 'mutation CreateQuizMutation(\n  $input: CreateQuizInput!\n) {\n  createQuiz(input: $input) {\n    id\n    title\n    description\n  }\n}\n',
    },
  }
})()

;(node as any).hash = 'bfdaf6c23b0de4e86bdad2e3721f911a'

export default node
