// Freeze only the two Registry reads and their transitive specification inputs.
const fs = require('node:fs');
const path = require('node:path');
const yaml = require('../../platform-api/node_modules/js-yaml');
const Parser = require('../../platform-api/node_modules/@apidevtools/swagger-parser');
const sourcePath = process.argv[2];
if (!sourcePath) throw new Error('Pass the platform-api agent-registry/v1/openapi.yaml source path');
const source = yaml.load(fs.readFileSync(sourcePath, 'utf8'));
const mini = {openapi: source.openapi, info: source.info, servers: source.servers,
  security: source.security, paths: {}, components: {}};
for (const name of ['/agents', '/agents/{agentId}/principals']) {
  const item = source.paths[name];
  mini.paths[name] = {get: structuredClone(item.get)};
  if (item.parameters) mini.paths[name].parameters = structuredClone(item.parameters);
}
const seen = new Set();
function add(section, name) {
  const key = section + '/' + name;
  if (seen.has(key)) return;
  seen.add(key);
  const original = source.components[section]?.[name];
  if (!original) throw new Error('Missing Registry source component ' + key);
  const node = structuredClone(original);
  (mini.components[section] ||= {})[name] = node;
  visit(node);
}
function visit(node) {
  if (!node || typeof node !== 'object') return;
  if (typeof node.$ref === 'string' && node.$ref.startsWith('#/components/')) {
    const [section, name] = node.$ref.slice('#/components/'.length).split('/');
    add(section, name.replace(/~1/g, '/').replace(/~0/g, '~'));
  }
  for (const value of Object.values(node)) visit(value);
}
visit(mini.paths);
for (const alternative of mini.security || []) {
  for (const name of Object.keys(alternative)) add('securitySchemes', name);
}
(async () => {
  const bundled = await Parser.bundle(path.resolve(sourcePath), mini, {resolve: {http: false}});
  fs.writeFileSync(path.join(__dirname, 'registry-reference-input.json'), JSON.stringify(bundled, null, 2) + '\n');
  console.log('Bundled two Registry inventory reads and their local component dependencies.');
})().catch(error => { console.error(error.message); process.exitCode = 1; });
