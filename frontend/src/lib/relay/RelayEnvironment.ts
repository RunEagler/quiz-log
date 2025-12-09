import {
  Environment,
  Network,
  RecordSource,
  type RequestParameters,
  Store,
  type Variables,
} from 'relay-runtime'

async function fetchGraphQL(request: RequestParameters, variables: Variables) {
  const response = await fetch('/query', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({
      query: request.text,
      variables,
    }),
  })

  return await response.json()
}

const network = Network.create(fetchGraphQL)
const store = new Store(new RecordSource())

const RelayEnvironment = new Environment({
  network,
  store,
})

export default RelayEnvironment
