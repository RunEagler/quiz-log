/**
 * @generated SignedSource<<f97bc8f54582d2f8c090de7288a96d16>>
 * @lightSyntaxTransform
 * @nogrep
 */

/* tslint:disable */
/* eslint-disable */
// @ts-nocheck

import { ConcreteRequest } from 'relay-runtime';
export type Difficulty = "EASY" | "HARD" | "MEDIUM" | "%future added value";
export type QuestionType = "MULTIPLE_CHOICE" | "SHORT_ANSWER" | "TRUE_FALSE" | "%future added value";
export type QuizDetailQuery$variables = {
  id: string;
};
export type QuizDetailQuery$data = {
  readonly quiz: {
    readonly createdAt: any;
    readonly description: string | null | undefined;
    readonly id: string;
    readonly questions: ReadonlyArray<{
      readonly content: string;
      readonly difficulty: Difficulty;
      readonly id: string;
      readonly options: ReadonlyArray<string> | null | undefined;
      readonly type: QuestionType;
    }>;
    readonly tags: ReadonlyArray<{
      readonly id: string;
      readonly name: string;
    }>;
    readonly title: string;
  } | null | undefined;
};
export type QuizDetailQuery = {
  response: QuizDetailQuery$data;
  variables: QuizDetailQuery$variables;
};

const node: ConcreteRequest = (function(){
var v0 = [
  {
    "defaultValue": null,
    "kind": "LocalArgument",
    "name": "id"
  }
],
v1 = {
  "alias": null,
  "args": null,
  "kind": "ScalarField",
  "name": "id",
  "storageKey": null
},
v2 = [
  {
    "alias": null,
    "args": [
      {
        "kind": "Variable",
        "name": "id",
        "variableName": "id"
      }
    ],
    "concreteType": "Quiz",
    "kind": "LinkedField",
    "name": "quiz",
    "plural": false,
    "selections": [
      (v1/*: any*/),
      {
        "alias": null,
        "args": null,
        "kind": "ScalarField",
        "name": "title",
        "storageKey": null
      },
      {
        "alias": null,
        "args": null,
        "kind": "ScalarField",
        "name": "description",
        "storageKey": null
      },
      {
        "alias": null,
        "args": null,
        "kind": "ScalarField",
        "name": "createdAt",
        "storageKey": null
      },
      {
        "alias": null,
        "args": null,
        "concreteType": "Tag",
        "kind": "LinkedField",
        "name": "tags",
        "plural": true,
        "selections": [
          (v1/*: any*/),
          {
            "alias": null,
            "args": null,
            "kind": "ScalarField",
            "name": "name",
            "storageKey": null
          }
        ],
        "storageKey": null
      },
      {
        "alias": null,
        "args": null,
        "concreteType": "Question",
        "kind": "LinkedField",
        "name": "questions",
        "plural": true,
        "selections": [
          (v1/*: any*/),
          {
            "alias": null,
            "args": null,
            "kind": "ScalarField",
            "name": "type",
            "storageKey": null
          },
          {
            "alias": null,
            "args": null,
            "kind": "ScalarField",
            "name": "content",
            "storageKey": null
          },
          {
            "alias": null,
            "args": null,
            "kind": "ScalarField",
            "name": "difficulty",
            "storageKey": null
          },
          {
            "alias": null,
            "args": null,
            "kind": "ScalarField",
            "name": "options",
            "storageKey": null
          }
        ],
        "storageKey": null
      }
    ],
    "storageKey": null
  }
];
return {
  "fragment": {
    "argumentDefinitions": (v0/*: any*/),
    "kind": "Fragment",
    "metadata": null,
    "name": "QuizDetailQuery",
    "selections": (v2/*: any*/),
    "type": "Query",
    "abstractKey": null
  },
  "kind": "Request",
  "operation": {
    "argumentDefinitions": (v0/*: any*/),
    "kind": "Operation",
    "name": "QuizDetailQuery",
    "selections": (v2/*: any*/)
  },
  "params": {
    "cacheID": "c0ba7326acaed67447079c8edb7aa323",
    "id": null,
    "metadata": {},
    "name": "QuizDetailQuery",
    "operationKind": "query",
    "text": "query QuizDetailQuery(\n  $id: ID!\n) {\n  quiz(id: $id) {\n    id\n    title\n    description\n    createdAt\n    tags {\n      id\n      name\n    }\n    questions {\n      id\n      type\n      content\n      difficulty\n      options\n    }\n  }\n}\n"
  }
};
})();

(node as any).hash = "2793c75f054b8e2adba761a93e47fdf3";

export default node;
