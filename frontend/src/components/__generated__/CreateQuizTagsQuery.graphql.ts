/**
 * @generated SignedSource<<7a12961dc0796b4a03f3108df07e6e46>>
 * @lightSyntaxTransform
 * @nogrep
 */

/* tslint:disable */
/* eslint-disable */
// @ts-nocheck

import { ConcreteRequest } from 'relay-runtime';
export type CreateQuizTagsQuery$variables = Record<PropertyKey, never>;
export type CreateQuizTagsQuery$data = {
  readonly tags: ReadonlyArray<{
    readonly id: string;
    readonly name: string;
  }>;
};
export type CreateQuizTagsQuery = {
  response: CreateQuizTagsQuery$data;
  variables: CreateQuizTagsQuery$variables;
};

const node: ConcreteRequest = (function(){
var v0 = [
  {
    "alias": null,
    "args": null,
    "concreteType": "Tag",
    "kind": "LinkedField",
    "name": "tags",
    "plural": true,
    "selections": [
      {
        "alias": null,
        "args": null,
        "kind": "ScalarField",
        "name": "id",
        "storageKey": null
      },
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
];
return {
  "fragment": {
    "argumentDefinitions": [],
    "kind": "Fragment",
    "metadata": null,
    "name": "CreateQuizTagsQuery",
    "selections": (v0/*: any*/),
    "type": "Query",
    "abstractKey": null
  },
  "kind": "Request",
  "operation": {
    "argumentDefinitions": [],
    "kind": "Operation",
    "name": "CreateQuizTagsQuery",
    "selections": (v0/*: any*/)
  },
  "params": {
    "cacheID": "cbe2cafc7929b45d8acbad0edc057c12",
    "id": null,
    "metadata": {},
    "name": "CreateQuizTagsQuery",
    "operationKind": "query",
    "text": "query CreateQuizTagsQuery {\n  tags {\n    id\n    name\n  }\n}\n"
  }
};
})();

(node as any).hash = "ee1f551899cdad223296f35f249d526f";

export default node;
