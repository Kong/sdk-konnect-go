// Add only the two Registry inventory reads needed by Identity retirement guards.
const assert = require('node:assert/strict');
const fs = require('node:fs');
const path = require('node:path');

const selectedPaths = ['/agents', '/agents/{agentId}/principals'];
const prefix = 'RegistryReference';
const tag = 'Registry Agent References';
const marker = 'x-kongctl-registry-reference-preview';
const escape = value => value.replace(/~/g, '~0').replace(/\//g, '~1');
const unescape = value => value.replace(/~1/g, '/').replace(/~0/g, '~');

function lookup(document, reference) {
  if (!reference.startsWith('#/')) throw new Error('Registry input has a non-local reference');
  let value = document;
  for (const part of reference.slice(2).split('/')) value = value?.[unescape(part)];
  if (value === undefined) throw new Error('Unresolved Registry reference: ' + reference);
  return value;
}

function prepareRegistryReferences(baseline, input) {
  assert.deepEqual(Object.keys(input.paths).sort(), [...selectedPaths].sort());
  const originalInput = input;
  input = structuredClone(input);
  const anchored = new Map();
  // Bundlers can point shared schemas into inline response allOf entries.
  // Generation overlays flatten those entries, so use stable components.
  function anchor(node) {
    if (!node || typeof node !== 'object') return;
    if (typeof node.$ref === 'string' && node.$ref.startsWith('#/paths/')) {
      const reference = node.$ref;
      if (!anchored.has(reference)) {
        const value = structuredClone(lookup(originalInput, reference));
        const section = value.in && value.name ? 'parameters' : value.content ? 'responses' : 'schemas';
        const name = 'ReadAnchor' + (anchored.size + 1);
        const target = '#/components/' + section + '/' + name;
        if (input.components[section]?.[name]) throw new Error('Reserved Registry anchor name already exists');
        anchored.set(reference, target);
        (input.components[section] ||= {})[name] = value;
        anchor(value);
      }
      node.$ref = anchored.get(reference);
    }
    for (const value of Object.values(node)) anchor(value);
  }
  anchor(input);
  const lifted = new Map();
  // Inline "name" and "meta" models otherwise rename existing public Go
  // types. Give only these new Registry shapes explicit component identities.
  for (const [parent, property, name] of [
    ['AgentFilterParameters', 'name', 'AgentNameFilter'],
    ['AgentList', 'meta', 'PaginationMeta'],
  ]) {
    const properties = input.components.schemas[parent]?.properties;
    if (!properties?.[property]) continue;
    if (input.components.schemas[name]) throw new Error('Registry lifted schema already exists: ' + name);
    input.components.schemas[name] = structuredClone(properties[property]);
    properties[property] = {$ref: '#/components/schemas/' + name};
    lifted.set('#/components/schemas/' + parent + '/properties/' + property,
      '#/components/schemas/' + name);
  }
  const result = structuredClone(baseline);
  const additions = {paths: [], components: {}};
  const existingOperationIDs = new Set();
  for (const [route, item] of Object.entries(baseline.paths)) {
    if (selectedPaths.some(p => route === '/v1' + p)) continue;
    for (const method of ['get', 'post', 'put', 'patch', 'delete']) {
      if (item[method]?.operationId) existingOperationIDs.add(item[method].operationId);
    }
  }
  for (const alternative of input.security || []) {
    for (const name of Object.keys(alternative)) {
      let source = input.components.securitySchemes[name];
      let target = baseline.components.securitySchemes?.[name];
      if (!source || !target) throw new Error('Unsupported Registry security scheme: ' + name);
      if (source.$ref) source = lookup(input, source.$ref);
      if (target.$ref) target = lookup(baseline, target.$ref);
      for (const field of ['type', 'scheme', 'in', 'name']) {
        assert.equal(source[field], target[field], 'Registry auth shape differs: ' + name);
      }
    }
  }
  function rewrite(node) {
    if (!node || typeof node !== 'object') return;
    for (const field of ['description', 'summary']) {
      if (typeof node[field] === 'string') node[field] = node[field].replace(/[ \t]+$/gm, '');
    }
    delete node['x-internal'];
    delete node['x-unstable'];
    if (typeof node.$ref === 'string') {
      const reference = lifted.get(node.$ref) || node.$ref;
      if (reference.startsWith('#/components/')) {
        const parts = reference.slice(2).split('/');
        const section = parts[1], name = unescape(parts[2]);
        if (!input.components[section]?.[name]) throw new Error('Unknown Registry component: ' + name);
        if (section !== 'securitySchemes') parts[2] = escape(prefix + name);
        node.$ref = '#/' + parts.join('/');
      } else if (reference.startsWith('#/paths/')) {
        const parts = reference.slice(2).split('/');
        const original = unescape(parts[1]);
        if (!selectedPaths.includes(original)) throw new Error('Reference escapes selected Registry reads');
        parts[1] = escape('/v1' + original);
        node.$ref = '#/' + parts.join('/');
      } else {
        throw new Error('Unsupported Registry reference: ' + reference);
      }
    }
    for (const value of Object.values(node)) rewrite(value);
  }
  function install(section, name, value) {
    result.components[section] ||= {};
    if (result.components[section][name]) {
      assert.deepEqual(result.components[section][name], value, 'Registry component collision: ' + name);
    } else {
      result.components[section][name] = value;
      (additions.components[section] ||= []).push(name);
    }
  }
  for (const [section, values] of Object.entries(input.components)) {
    if (section === 'securitySchemes') continue;
    for (const [name, original] of Object.entries(values)) {
      const value = structuredClone(original);
      rewrite(value);
      install(section, prefix + name, value);
    }
  }
  for (const original of selectedPaths) {
    const item = structuredClone(input.paths[original]);
    if (Object.keys(item).some(key => !['get', 'parameters'].includes(key))) {
      throw new Error('Registry selection contains an operation other than GET');
    }
    if (!item.get?.operationId || existingOperationIDs.has(item.get.operationId)) {
      throw new Error('Registry operation ID is missing or collides with the public SDK');
    }
    rewrite(item);
    item.get.tags = [tag];
    item.get['x-speakeasy-group'] = 'RegistryAgentReferences';
    item.get.security = structuredClone(input.security);
    item.get.servers = input.servers.map(server => {
      const url = new URL(server.url);
      if (url.protocol !== 'https:' || url.pathname !== '/v1') throw new Error('Unexpected Registry server');
      return {...server, url: url.origin + '/'};
    });
    const route = '/v1' + original;
    if (result.paths[route]) assert.deepEqual(result.paths[route], item, 'Registry path already exists');
    else {
      result.paths[route] = item;
      additions.paths.push(route);
    }
  }
  const declared = {name: tag, description: 'Read Registry Agent references before Identity retirement.'};
  result.tags ||= [];
  const previous = result.tags.find(value => value.name === tag);
  if (previous) assert.deepEqual(previous, declared);
  else result.tags.push(declared);
  const audit = {operations: 2, methods: ['GET'], componentPrefix: prefix};
  if (result[marker]) assert.deepEqual(result[marker], audit);
  else result[marker] = audit;

  // This transform is additive. Existing public definitions may not be renamed
  // or changed to accommodate the gated read-only dependency.
  for (const [route, item] of Object.entries(baseline.paths)) assert.deepEqual(result.paths[route], item);
  for (const [section, values] of Object.entries(baseline.components)) {
    for (const [name, value] of Object.entries(values)) assert.deepEqual(result.components[section][name], value);
  }
  function verify(node) {
    if (!node || typeof node !== 'object') return;
    if (node.$ref) lookup(result, node.$ref);
    for (const value of Object.values(node)) verify(value);
  }
  for (const route of selectedPaths.map(value => '/v1' + value)) verify(result.paths[route]);
  for (const [section, names] of Object.entries(input.components)) {
    if (section === 'securitySchemes') continue;
    for (const name of Object.keys(names)) verify(result.components[section][prefix + name]);
  }
  return {spec: result, audit, additions};
}

if (require.main === module) {
  const yaml = require('../../platform-api/node_modules/js-yaml');
  const filename = process.argv[2] || path.resolve(__dirname, '../openapi.yaml');
  const input = JSON.parse(fs.readFileSync(path.join(__dirname, 'registry-reference-input.json'), 'utf8'));
  const {spec, audit} = prepareRegistryReferences(yaml.load(fs.readFileSync(filename, 'utf8')), input);
  fs.writeFileSync(filename, yaml.dump(spec, {lineWidth: 100, noRefs: true}).replace(/[ \t]+$/gm, ''));
  fs.writeFileSync(path.join(__dirname, 'registry-reference-audit.json'), JSON.stringify(audit, null, 2) + '\n');
  console.log('Added two read-only Registry operations; existing public definitions are unchanged.');
}

module.exports = {prepareRegistryReferences};
