/**
 * @generated SignedSource<<4b94169e6e4964ed64105572abca6d2a>>
 * @lightSyntaxTransform
 * @nogrep
 */

/* tslint:disable */
/* eslint-disable */
// @ts-nocheck

import type { ConcreteRequest } from 'relay-runtime'
export type QuizDetailDeleteMutation$variables = {
  id: string
}
export type QuizDetailDeleteMutation$data = {
  readonly deleteQuiz: boolean
}
export type QuizDetailDeleteMutation = {
  response: QuizDetailDeleteMutation$data
  variables: QuizDetailDeleteMutation$variables
}

const node: ConcreteRequest = (() => {
  var v0 = [
      {
        defaultValue: null,
        kind: 'LocalArgument',
        name: 'id',
      },
    ],
    v1 = [
      {
        alias: null,
        args: [
          {
            kind: 'Variable',
            name: 'id',
            variableName: 'id',
          },
        ],
        kind: 'ScalarField',
        name: 'deleteQuiz',
        storageKey: null,
      },
    ]
  return {
    fragment: {
      argumentDefinitions: v0 /*: any*/,
      kind: 'Fragment',
      metadata: null,
      name: 'QuizDetailDeleteMutation',
      selections: v1 /*: any*/,
      type: 'Mutation',
      abstractKey: null,
    },
    kind: 'Request',
    operation: {
      argumentDefinitions: v0 /*: any*/,
      kind: 'Operation',
      name: 'QuizDetailDeleteMutation',
      selections: v1 /*: any*/,
    },
    params: {
      cacheID: '4df6d4c87755935d36ba86535aace8fd',
      id: null,
      metadata: {},
      name: 'QuizDetailDeleteMutation',
      operationKind: 'mutation',
      text: 'mutation QuizDetailDeleteMutation(\n  $id: ID!\n) {\n  deleteQuiz(id: $id)\n}\n',
    },
  }
})()

;(node as any).hash = 'a72a058b744c08b9f5d2cfbd109b6b76'

export default node
