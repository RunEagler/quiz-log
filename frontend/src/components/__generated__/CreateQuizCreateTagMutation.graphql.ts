/**
 * @generated SignedSource<<7b0aca34f80e54420c36d85f4ab74e57>>
 * @lightSyntaxTransform
 * @nogrep
 */

/* tslint:disable */
/* eslint-disable */
// @ts-nocheck

import { ConcreteRequest } from 'relay-runtime';
export type CreateQuizCreateTagMutation$variables = {
  name: string;
};
export type CreateQuizCreateTagMutation$data = {
  readonly createTag: {
    readonly id: string;
    readonly name: string;
  };
};
export type CreateQuizCreateTagMutation = {
  response: CreateQuizCreateTagMutation$data;
  variables: CreateQuizCreateTagMutation$variables;
};

const node: ConcreteRequest = (function(){
var v0 = [
  {
    "defaultValue": null,
    "kind": "LocalArgument",
    "name": "name"
  }
],
v1 = [
  {
    "alias": null,
    "args": [
      {
        "kind": "Variable",
        "name": "name",
        "variableName": "name"
      }
    ],
    "concreteType": "Tag",
    "kind": "LinkedField",
    "name": "createTag",
    "plural": false,
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
    "argumentDefinitions": (v0/*: any*/),
    "kind": "Fragment",
    "metadata": null,
    "name": "CreateQuizCreateTagMutation",
    "selections": (v1/*: any*/),
    "type": "Mutation",
    "abstractKey": null
  },
  "kind": "Request",
  "operation": {
    "argumentDefinitions": (v0/*: any*/),
    "kind": "Operation",
    "name": "CreateQuizCreateTagMutation",
    "selections": (v1/*: any*/)
  },
  "params": {
    "cacheID": "dea04c44d870cc9006636ef7bab337ee",
    "id": null,
    "metadata": {},
    "name": "CreateQuizCreateTagMutation",
    "operationKind": "mutation",
    "text": "mutation CreateQuizCreateTagMutation(\n  $name: String!\n) {\n  createTag(name: $name) {\n    id\n    name\n  }\n}\n"
  }
};
})();

(node as any).hash = "f3e5ee0e1b9b75970efb59a6988dd764";

export default node;
