/**
 * @generated SignedSource<<8185bb1c0feb21009fb87fc68739c75f>>
 * @lightSyntaxTransform
 * @nogrep
 */

/* tslint:disable */
/* eslint-disable */
// @ts-nocheck

import { ConcreteRequest } from 'relay-runtime';
export type SubmitAttemptInput = {
  answers: ReadonlyArray<AnswerInput>;
  quizID: string;
};
export type AnswerInput = {
  questionID: string;
  userAnswer: string;
};
export type TakeQuizSubmitMutation$variables = {
  input: SubmitAttemptInput;
};
export type TakeQuizSubmitMutation$data = {
  readonly submitAttempt: {
    readonly attempt: {
      readonly id: string;
      readonly score: number;
      readonly totalQuestions: number;
    };
    readonly correctCount: number;
    readonly score: number;
    readonly totalQuestions: number;
    readonly wrongQuestions: ReadonlyArray<{
      readonly content: string;
      readonly correctAnswer: string;
      readonly explanation: string | null | undefined;
      readonly id: string;
    }>;
  };
};
export type TakeQuizSubmitMutation = {
  response: TakeQuizSubmitMutation$data;
  variables: TakeQuizSubmitMutation$variables;
};

const node: ConcreteRequest = (function(){
var v0 = [
  {
    "defaultValue": null,
    "kind": "LocalArgument",
    "name": "input"
  }
],
v1 = {
  "alias": null,
  "args": null,
  "kind": "ScalarField",
  "name": "id",
  "storageKey": null
},
v2 = {
  "alias": null,
  "args": null,
  "kind": "ScalarField",
  "name": "score",
  "storageKey": null
},
v3 = {
  "alias": null,
  "args": null,
  "kind": "ScalarField",
  "name": "totalQuestions",
  "storageKey": null
},
v4 = [
  {
    "alias": null,
    "args": [
      {
        "kind": "Variable",
        "name": "input",
        "variableName": "input"
      }
    ],
    "concreteType": "AttemptResult",
    "kind": "LinkedField",
    "name": "submitAttempt",
    "plural": false,
    "selections": [
      {
        "alias": null,
        "args": null,
        "concreteType": "Attempt",
        "kind": "LinkedField",
        "name": "attempt",
        "plural": false,
        "selections": [
          (v1/*: any*/),
          (v2/*: any*/),
          (v3/*: any*/)
        ],
        "storageKey": null
      },
      (v2/*: any*/),
      (v3/*: any*/),
      {
        "alias": null,
        "args": null,
        "kind": "ScalarField",
        "name": "correctCount",
        "storageKey": null
      },
      {
        "alias": null,
        "args": null,
        "concreteType": "Question",
        "kind": "LinkedField",
        "name": "wrongQuestions",
        "plural": true,
        "selections": [
          (v1/*: any*/),
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
            "name": "correctAnswer",
            "storageKey": null
          },
          {
            "alias": null,
            "args": null,
            "kind": "ScalarField",
            "name": "explanation",
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
    "name": "TakeQuizSubmitMutation",
    "selections": (v4/*: any*/),
    "type": "Mutation",
    "abstractKey": null
  },
  "kind": "Request",
  "operation": {
    "argumentDefinitions": (v0/*: any*/),
    "kind": "Operation",
    "name": "TakeQuizSubmitMutation",
    "selections": (v4/*: any*/)
  },
  "params": {
    "cacheID": "73dcbc63f36115dffd3fcecdbc044c22",
    "id": null,
    "metadata": {},
    "name": "TakeQuizSubmitMutation",
    "operationKind": "mutation",
    "text": "mutation TakeQuizSubmitMutation(\n  $input: SubmitAttemptInput!\n) {\n  submitAttempt(input: $input) {\n    attempt {\n      id\n      score\n      totalQuestions\n    }\n    score\n    totalQuestions\n    correctCount\n    wrongQuestions {\n      id\n      content\n      correctAnswer\n      explanation\n    }\n  }\n}\n"
  }
};
})();

(node as any).hash = "988501d6a54407fc534d152efde9a1e5";

export default node;
