// Keep credential request expiration separate from the read-only IAM schema.
function prepareCredentialExpiration(schemas, writableExpiration) {
  if (writableExpiration.type !== 'string' || writableExpiration.format !== 'date-time' ||
      writableExpiration.nullable !== true || writableExpiration.readOnly) {
    throw new Error('Unexpected canonical writable expiration schema');
  }
  const expirationName = 'KongIdentityCredentialExpiresAt';
  if (schemas[expirationName] && JSON.stringify(schemas[expirationName]) !== JSON.stringify(writableExpiration)) {
    throw new Error('Unexpected existing credential expiration schema');
  }
  schemas[expirationName] = structuredClone(writableExpiration);
  for (const name of ['CreatePrincipalKeyBodyImported', 'CreatePrincipalKeyBodyv1', 'UpdatePrincipalKeyBody',
    'CreatePasswordBody', 'UpdatePasswordBody', 'ReplacePasswordBody']) {
    const property = schemas[name]?.properties?.expires_at;
    if (!property || !['#/components/schemas/ExpiresAt', '#/components/schemas/' + expirationName].includes(property.$ref)) {
      throw new Error('Unexpected credential expiration reference: ' + name);
    }
    schemas[name].properties.expires_at = {$ref: '#/components/schemas/' + expirationName};
  }
}

module.exports = {prepareCredentialExpiration};
