/**
 * @generated SignedSource<<fffff185ca04bbbd009cb74230a3c62c>>
 * @lightSyntaxTransform
 * @nogrep
 */

/* tslint:disable */
/* eslint-disable */
// @ts-nocheck

import { ConcreteRequest } from 'relay-runtime';
export type QuizListQuery$variables = Record<PropertyKey, never>;
export type QuizListQuery$data = {
  readonly quizzes: ReadonlyArray<{
    readonly createdAt: any;
    readonly description: string | null | undefined;
    readonly id: string;
    readonly tags: ReadonlyArray<{
      readonly id: string;
      readonly name: string;
    }>;
    readonly title: string;
  }>;
};
export type QuizListQuery = {
  response: QuizListQuery$data;
  variables: QuizListQuery$variables;
};

const node: ConcreteRequest = (function(){
var v0 = {
  "alias": null,
  "args": null,
  "kind": "ScalarField",
  "name": "id",
  "storageKey": null
},
v1 = [
  {
    "alias": null,
    "args": null,
    "concreteType": "Quiz",
    "kind": "LinkedField",
    "name": "quizzes",
    "plural": true,
    "selections": [
      (v0/*: any*/),
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
          (v0/*: any*/),
          {
            "alias": null,
            "args": null,
            "kind": "ScalarField",
            "name": "name",
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
    "argumentDefinitions": [],
    "kind": "Fragment",
    "metadata": null,
    "name": "QuizListQuery",
    "selections": (v1/*: any*/),
    "type": "Query",
    "abstractKey": null
  },
  "kind": "Request",
  "operation": {
    "argumentDefinitions": [],
    "kind": "Operation",
    "name": "QuizListQuery",
    "selections": (v1/*: any*/)
  },
  "params": {
    "cacheID": "5690bcb50e6e326972a5574072c22564",
    "id": null,
    "metadata": {},
    "name": "QuizListQuery",
    "operationKind": "query",
    "text": "query QuizListQuery {\n  quizzes {\n    id\n    title\n    description\n    createdAt\n    tags {\n      id\n      name\n    }\n  }\n}\n"
  }
};
})();

(node as any).hash = "83228be2a3f3a60d1c010437c77a6ab1";

export default node;
