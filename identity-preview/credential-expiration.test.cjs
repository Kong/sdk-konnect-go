const assert = require('node:assert/strict');
const fs = require('node:fs');
const path = require('node:path');
const test = require('node:test');
const yaml = require('../../platform-api/node_modules/js-yaml');
const {prepareCredentialExpiration} = require('./credential-expiration.cjs');

const bodies = ['CreatePrincipalKeyBodyImported', 'CreatePrincipalKeyBodyv1', 'UpdatePrincipalKeyBody',
  'CreatePasswordBody', 'UpdatePasswordBody', 'ReplacePasswordBody'];
const canonical = yaml.load(fs.readFileSync(path.resolve(__dirname,
  '../../platform-api/src/common/definitions/properties.yaml'), 'utf8')).components.schemas.ExpiresAt;

function fixture() {
  const spec = yaml.load(fs.readFileSync(path.resolve(__dirname, '../openapi.yaml'), 'utf8'));
  delete spec.components.schemas.KongIdentityCredentialExpiresAt;
  for (const name of bodies) {
    spec.components.schemas[name].properties.expires_at = {$ref: '#/components/schemas/ExpiresAt'};
  }
  return spec;
}

test('only credential expiration references change; repeated preparation is stable', () => {
  const spec = fixture();
  const original = structuredClone(spec);
  assert.equal(spec.components.schemas.ExpiresAt.readOnly, true);
  prepareCredentialExpiration(spec.components.schemas, canonical);
  for (const name of bodies) {
    assert.deepEqual(spec.components.schemas[name].properties.expires_at,
      {$ref: '#/components/schemas/KongIdentityCredentialExpiresAt'});
  }
  assert.deepEqual(spec.components.schemas.KongIdentityCredentialExpiresAt, canonical);
  assert.ok(spec.components.schemas.ReplacePasswordBody.required.includes('expires_at'));
  const prepared = structuredClone(spec);
  prepareCredentialExpiration(spec.components.schemas, canonical);
  assert.deepEqual(spec, prepared);
  delete spec.components.schemas.KongIdentityCredentialExpiresAt;
  for (const name of bodies) {
    spec.components.schemas[name].properties.expires_at = original.components.schemas[name].properties.expires_at;
  }
  assert.deepEqual(spec, original);
});

test('unexpected source schemas and request references stop preparation', () => {
  for (const changed of [{readOnly: true}, {nullable: false}, {format: 'date'}, {type: 'number'}]) {
    assert.throws(() => prepareCredentialExpiration(fixture().components.schemas,
      {...canonical, ...changed}), /Unexpected canonical/);
  }
  const missing = fixture();
  delete missing.components.schemas.UpdatePasswordBody.properties.expires_at;
  assert.throws(() => prepareCredentialExpiration(missing.components.schemas, canonical), /reference/);
  const conflicting = fixture();
  conflicting.components.schemas.KongIdentityCredentialExpiresAt = {type: 'string'};
  assert.throws(() => prepareCredentialExpiration(conflicting.components.schemas, canonical), /existing/);
});
