// Apply only Kong Identity preview selection and Go name disambiguation.
// Run after platform-api/tools/gen_go_sdk.sh and before SDK generation.
const fs = require('node:fs');
const path = require('node:path');
const yaml = require('../../platform-api/node_modules/js-yaml');

const input = process.argv[2] || path.resolve(__dirname, '../openapi.yaml');
const spec = yaml.load(fs.readFileSync(input, 'utf8'));
const names = new Map([
  ['Algorithm', 'AuthServerSigningAlgorithm'],
  ['Client', 'AuthServerClient'],
  ['Scope', 'AuthServerScope'],
  ['GrantType', 'AuthServerGrantType'],
  ['TokenEndpointAuthMethod', 'AuthServerTokenEndpointAuthMethod'],
]);
const schemas = spec.components.schemas;
for (const [previous, next] of names) {
  if (schemas[previous]) {
    if (schemas[next]) throw new Error('Ambiguous preview schema: ' + next);
    schemas[next] = schemas[previous];
    delete schemas[previous];
  } else if (!schemas[next]) {
    throw new Error('Missing preview schema: ' + previous);
  }
}

function rewrite(node) {
  if (!node || typeof node !== 'object') return;
  for (const [key, value] of Object.entries(node)) {
    if (key === 'description' && typeof value === 'string') {
      node[key] = value.replace(/[ \t]+$/gm, '');
    } else if (typeof value === 'string' && value.startsWith('#/components/schemas/')) {
      const name = value.slice('#/components/schemas/'.length);
      if (names.has(name)) node[key] = '#/components/schemas/' + names.get(name);
    } else {
      rewrite(value);
    }
  }
}
rewrite(spec);

for (const variant of ['Shared', 'User']) {
  const credential = schemas['VaultProviderCredential' + variant].allOf.find(
    item => item.properties?.credential_type,
  );
  if (!credential) throw new Error('Missing credential discriminator: ' + variant);
  if (!credential.properties.credential_type.$ref) {
    const name = 'VaultProvider' + variant + 'CredentialType';
    schemas[name] = credential.properties.credential_type;
    credential.properties.credential_type = {
      $ref: '#/components/schemas/' + name,
    };
  }
}

let count = 0;
for (const [route, item] of Object.entries(spec.paths)) {
  if (route.startsWith('/v1/connected-apps/')) {
    delete spec.paths[route];
    continue;
  }
  if (!['/v1/auth-servers', '/v2/directories', '/v2/credential-provider-templates']
    .some(prefix => route === prefix || route.startsWith(prefix + '/'))) continue;
  for (const method of ['get', 'post', 'put', 'patch', 'delete']) {
    const operation = item[method];
    if (!operation) continue;
    count++;
    for (const parameter of operation.parameters || []) {
      if (!parameter.schema || parameter.schema.$ref) continue;
      const name = operation.operationId.replace(/(^.|[-_]\w)/g,
        match => match.replace(/[-_]/g, '').toUpperCase());
      if (parameter.name === 'filter') {
        parameter.schema['x-speakeasy-name-override'] = name + 'Filter';
      } else if (parameter.name === 'force') {
        parameter.schema['x-speakeasy-name-override'] = name + 'Force';
      }
    }
  }
}
if (count !== 112) throw new Error('Unexpected Identity operation count: ' + count);
fs.writeFileSync(input, yaml.dump(spec, {lineWidth: 100, noRefs: true}).replace(/[ \t]+$/gm, ''));
const workflowPath = path.resolve(__dirname, '../.speakeasy/workflow.yaml');
const workflow = yaml.load(fs.readFileSync(workflowPath, 'utf8'));
delete workflow.targets.go.codeSamples;
fs.writeFileSync(workflowPath, yaml.dump(workflow, {indent: 4, lineWidth: 120, noRefs: true}));
console.log('Selected 112 Identity operations with distinct Go schema names.');
