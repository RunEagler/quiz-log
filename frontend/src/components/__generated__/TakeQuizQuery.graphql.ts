/**
 * @generated SignedSource<<145d6df6a08c8670e8aa85b109350544>>
 * @lightSyntaxTransform
 * @nogrep
 */

/* tslint:disable */
/* eslint-disable */
// @ts-nocheck

import { ConcreteRequest } from 'relay-runtime';
export type Difficulty = "EASY" | "HARD" | "MEDIUM" | "%future added value";
export type QuestionType = "MULTIPLE_CHOICE" | "SHORT_ANSWER" | "TRUE_FALSE" | "%future added value";
export type TakeQuizQuery$variables = {
  id: string;
};
export type TakeQuizQuery$data = {
  readonly quiz: {
    readonly description: string | null | undefined;
    readonly id: string;
    readonly questions: ReadonlyArray<{
      readonly content: string;
      readonly difficulty: Difficulty;
      readonly explanation: string | null | undefined;
      readonly id: string;
      readonly options: ReadonlyArray<string> | null | undefined;
      readonly type: QuestionType;
    }>;
    readonly title: string;
  } | null | undefined;
};
export type TakeQuizQuery = {
  response: TakeQuizQuery$data;
  variables: TakeQuizQuery$variables;
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
            "name": "options",
            "storageKey": null
          },
          {
            "alias": null,
            "args": null,
            "kind": "ScalarField",
            "name": "explanation",
            "storageKey": null
          },
          {
            "alias": null,
            "args": null,
            "kind": "ScalarField",
            "name": "difficulty",
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
    "name": "TakeQuizQuery",
    "selections": (v2/*: any*/),
    "type": "Query",
    "abstractKey": null
  },
  "kind": "Request",
  "operation": {
    "argumentDefinitions": (v0/*: any*/),
    "kind": "Operation",
    "name": "TakeQuizQuery",
    "selections": (v2/*: any*/)
  },
  "params": {
    "cacheID": "e5f900f4dbc4c9a69dc194a3edd3ef38",
    "id": null,
    "metadata": {},
    "name": "TakeQuizQuery",
    "operationKind": "query",
    "text": "query TakeQuizQuery(\n  $id: ID!\n) {\n  quiz(id: $id) {\n    id\n    title\n    description\n    questions {\n      id\n      type\n      content\n      options\n      explanation\n      difficulty\n    }\n  }\n}\n"
  }
};
})();

(node as any).hash = "908bbc1671f9a2b66ea3fc898816f17d";

export default node;
