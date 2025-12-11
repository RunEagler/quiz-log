/**
 * @generated SignedSource<<cf85f860e0ad3fe46a7616bedb9da689>>
 * @lightSyntaxTransform
 * @nogrep
 */

/* tslint:disable */
/* eslint-disable */
// @ts-nocheck

import type { ConcreteRequest } from 'relay-runtime'
export type StatisticsQuery$variables = Record<PropertyKey, never>
export type StatisticsQuery$data = {
  readonly statistics: {
    readonly averageScore: number
    readonly categoryStats: ReadonlyArray<{
      readonly correctRate: number
      readonly tagName: string
      readonly totalQuestions: number
    }>
    readonly recentAttempts: ReadonlyArray<{
      readonly completedAt: any | null | undefined
      readonly id: string
      readonly quizID: string
      readonly score: number
      readonly totalQuestions: number
    }>
    readonly totalAttempts: number
  }
}
export type StatisticsQuery = {
  response: StatisticsQuery$data
  variables: StatisticsQuery$variables
}

const node: ConcreteRequest = (() => {
  var v0 = {
      alias: null,
      args: null,
      kind: 'ScalarField',
      name: 'totalQuestions',
      storageKey: null,
    },
    v1 = [
      {
        alias: null,
        args: null,
        concreteType: 'Statistics',
        kind: 'LinkedField',
        name: 'statistics',
        plural: false,
        selections: [
          {
            alias: null,
            args: null,
            kind: 'ScalarField',
            name: 'totalAttempts',
            storageKey: null,
          },
          {
            alias: null,
            args: null,
            kind: 'ScalarField',
            name: 'averageScore',
            storageKey: null,
          },
          {
            alias: null,
            args: null,
            concreteType: 'CategoryStat',
            kind: 'LinkedField',
            name: 'categoryStats',
            plural: true,
            selections: [
              {
                alias: null,
                args: null,
                kind: 'ScalarField',
                name: 'tagName',
                storageKey: null,
              },
              {
                alias: null,
                args: null,
                kind: 'ScalarField',
                name: 'correctRate',
                storageKey: null,
              },
              v0 /*: any*/,
            ],
            storageKey: null,
          },
          {
            alias: null,
            args: null,
            concreteType: 'Attempt',
            kind: 'LinkedField',
            name: 'recentAttempts',
            plural: true,
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
                name: 'quizID',
                storageKey: null,
              },
              {
                alias: null,
                args: null,
                kind: 'ScalarField',
                name: 'score',
                storageKey: null,
              },
              v0 /*: any*/,
              {
                alias: null,
                args: null,
                kind: 'ScalarField',
                name: 'completedAt',
                storageKey: null,
              },
            ],
            storageKey: null,
          },
        ],
        storageKey: null,
      },
    ]
  return {
    fragment: {
      argumentDefinitions: [],
      kind: 'Fragment',
      metadata: null,
      name: 'StatisticsQuery',
      selections: v1 /*: any*/,
      type: 'Query',
      abstractKey: null,
    },
    kind: 'Request',
    operation: {
      argumentDefinitions: [],
      kind: 'Operation',
      name: 'StatisticsQuery',
      selections: v1 /*: any*/,
    },
    params: {
      cacheID: '6511fc6c38e2b6c16a4cc2e710553bfa',
      id: null,
      metadata: {},
      name: 'StatisticsQuery',
      operationKind: 'query',
      text: 'query StatisticsQuery {\n  statistics {\n    totalAttempts\n    averageScore\n    categoryStats {\n      tagName\n      correctRate\n      totalQuestions\n    }\n    recentAttempts {\n      id\n      quizID\n      score\n      totalQuestions\n      completedAt\n    }\n  }\n}\n',
    },
  }
})()

;(node as any).hash = 'df4aac5b53cd4029e0baad8543806b46'

export default node
