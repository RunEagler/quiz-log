/**
 * @generated SignedSource<<28bed408527308365d9cf67b9179bc6a>>
 * @lightSyntaxTransform
 * @nogrep
 */

/* tslint:disable */
/* eslint-disable */
// @ts-nocheck

import { ConcreteRequest } from 'relay-runtime';
export type QuizDetailDeleteQuestionMutation$variables = {
  id: string;
};
export type QuizDetailDeleteQuestionMutation$data = {
  readonly deleteQuestion: boolean;
};
export type QuizDetailDeleteQuestionMutation = {
  response: QuizDetailDeleteQuestionMutation$data;
  variables: QuizDetailDeleteQuestionMutation$variables;
};

const node: ConcreteRequest = (function(){
var v0 = [
  {
    "defaultValue": null,
    "kind": "LocalArgument",
    "name": "id"
  }
],
v1 = [
  {
    "alias": null,
    "args": [
      {
        "kind": "Variable",
        "name": "id",
        "variableName": "id"
      }
    ],
    "kind": "ScalarField",
    "name": "deleteQuestion",
    "storageKey": null
  }
];
return {
  "fragment": {
    "argumentDefinitions": (v0/*: any*/),
    "kind": "Fragment",
    "metadata": null,
    "name": "QuizDetailDeleteQuestionMutation",
    "selections": (v1/*: any*/),
    "type": "Mutation",
    "abstractKey": null
  },
  "kind": "Request",
  "operation": {
    "argumentDefinitions": (v0/*: any*/),
    "kind": "Operation",
    "name": "QuizDetailDeleteQuestionMutation",
    "selections": (v1/*: any*/)
  },
  "params": {
    "cacheID": "86150b4c964d9c5f1558fc7f524b9012",
    "id": null,
    "metadata": {},
    "name": "QuizDetailDeleteQuestionMutation",
    "operationKind": "mutation",
    "text": "mutation QuizDetailDeleteQuestionMutation(\n  $id: ID!\n) {\n  deleteQuestion(id: $id)\n}\n"
  }
};
})();

(node as any).hash = "55428fcb7a6d71b49efd59970b0d35a6";

export default node;
