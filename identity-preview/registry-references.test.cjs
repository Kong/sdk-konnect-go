const assert = require('node:assert/strict');
const test = require('node:test');
const {prepareRegistryReferences} = require('./registry-references.cjs');

function fixture() {
  const scheme = {type: 'http', scheme: 'bearer'};
  const baseline = {openapi: '3.0.3', info: {title: 'Public', version: '1'},
    paths: {'/existing': {get: {operationId: 'existing'}}}, tags: [{name: 'Existing'}],
    components: {schemas: {Agent: {type: 'string'}}, securitySchemes: {personalAccessToken: scheme}}};
  const input = {servers: [{url: 'https://us.api.konghq.com/v1'}], security: [{personalAccessToken: []}],
    paths: {
      '/agents': {get: {operationId: 'list-agents', 'x-internal': true, responses: {
        200: {content: {'application/json': {schema: {$ref: '#/components/schemas/Agent'}}}},
      }}},
      '/agents/{agentId}/principals': {get: {operationId: 'list-agent-principals', 'x-unstable': true}},
    }, components: {schemas: {Agent: {type: 'object', properties: {id: {type: 'string'}}}},
      securitySchemes: {personalAccessToken: structuredClone(scheme)}}};
  return {baseline, input};
}

test('Registry reads preserve colliding public schemas and are idempotent', () => {
  const {baseline, input} = fixture();
  input.components.schemas.Agent.description = 'A Registry agent.  \nDocumentation.\n';
  const before = structuredClone(baseline), source = structuredClone(input);
  const first = prepareRegistryReferences(baseline, input);
  assert.deepEqual(baseline, before);
  assert.deepEqual(input, source);
  assert.deepEqual(first.spec.components.schemas.Agent, {type: 'string'});
  assert.equal(first.spec.components.schemas.RegistryReferenceAgent.description,
    'A Registry agent.\nDocumentation.\n');
  assert.equal(first.spec.paths['/v1/agents'].get['x-internal'], undefined);
  assert.equal(first.spec.paths['/v1/agents'].get.servers[0].url, 'https://us.api.konghq.com/');
  assert.equal(first.spec.paths['/v1/agents'].get.responses[200].content['application/json'].schema.$ref,
    '#/components/schemas/RegistryReferenceAgent');
  assert.deepEqual(prepareRegistryReferences(first.spec, input).spec, first.spec);
});

test('Registry selection rejects writes, auth drift and existing component/path collisions', () => {
  for (const failure of ['write', 'auth', 'component', 'path', 'unknown-ref']) {
    const {baseline, input} = fixture();
    if (failure === 'write') input.paths['/agents'].post = {operationId: 'create-agent'};
    if (failure === 'auth') input.components.securitySchemes.personalAccessToken.scheme = 'basic';
    if (failure === 'component') baseline.components.schemas.RegistryReferenceAgent = {type: 'boolean'};
    if (failure === 'path') baseline.paths['/v1/agents'] = {get: {operationId: 'existing-agent'}};
    if (failure === 'unknown-ref') input.components.schemas.Agent.properties.link = {$ref: '#/paths/~1unselected/get'};
    assert.throws(() => prepareRegistryReferences(baseline, input), undefined, failure);
  }
});

test('Inline bundled references survive response allOf flattening through stable anchors', () => {
  const {baseline, input} = fixture();
  input.paths['/agents'].get.responses[400] = {content: {'application/problem+json': {
    schema: {allOf: [{type: 'object', properties: {status: {type: 'integer'}}}, {type: 'object'}]},
  }}};
  input.components.schemas.Agent.properties.error = {
    $ref: '#/paths/~1agents/get/responses/400/content/application~1problem+json/schema/allOf/0',
  };
  const output = prepareRegistryReferences(baseline, input).spec;
  const reference = output.components.schemas.RegistryReferenceAgent.properties.error.$ref;
  assert.match(reference, /^#\/components\/schemas\/RegistryReferenceReadAnchor/);
  assert.deepEqual(output.components.schemas.RegistryReferenceReadAnchor1,
    {type: 'object', properties: {status: {type: 'integer'}}});
  assert.deepEqual(prepareRegistryReferences(output, input).spec, output);
});

test('Inline Registry filter and pagination shapes do not consume public Name/Meta identities', () => {
  const {baseline, input} = fixture();
  baseline.components.schemas.Meta = {type: 'string'};
  input.components.schemas.AgentFilterParameters = {properties: {name: {type: 'object', properties: {eq: {type: 'string'}}}}};
  input.components.schemas.AgentList = {properties: {meta: {type: 'object', properties: {page: {type: 'object'}}}}};
  input.components.schemas.AgentPrincipalList = {properties: {
    meta: {$ref: '#/components/schemas/AgentList/properties/meta'},
  }};
  input.components.schemas.AgentFilterParameters.properties.slug = {
    $ref: '#/components/schemas/AgentFilterParameters/properties/name',
  };
  const output = prepareRegistryReferences(baseline, input).spec;
  assert.deepEqual(output.components.schemas.Meta, {type: 'string'});
  assert.equal(output.components.schemas.RegistryReferenceAgentFilterParameters.properties.name.$ref,
    '#/components/schemas/RegistryReferenceAgentNameFilter');
  assert.equal(output.components.schemas.RegistryReferenceAgentList.properties.meta.$ref,
    '#/components/schemas/RegistryReferencePaginationMeta');
  assert.equal(output.components.schemas.RegistryReferenceAgentPrincipalList.properties.meta.$ref,
    '#/components/schemas/RegistryReferencePaginationMeta');
  assert.equal(output.components.schemas.RegistryReferenceAgentFilterParameters.properties.slug.$ref,
    '#/components/schemas/RegistryReferenceAgentNameFilter');
});
