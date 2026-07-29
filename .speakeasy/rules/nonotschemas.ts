import { Rule, registerRule, createValidationError } from '@speakeasy-api/openapi-linter-types';
import type { Context, DocumentInfo, RuleConfig, ValidationError } from '@speakeasy-api/openapi-linter-types';

class NoNotSchemas extends Rule {
  id(): string { return 'custom-no-not-schemas'; }
  category(): string { return 'structure'; }
  description(): string {
    return 'The `not` keyword is unsupported by the SDK generator; any operation reaching it is silently skipped.';
  }
  summary(): string { return 'Schemas must not use the `not` keyword'; }

  run(_ctx: Context, docInfo: DocumentInfo, config: RuleConfig): ValidationError[] {
    const errors: ValidationError[] = [];
    const idx: any = (docInfo as any).index;
    if (!idx) return errors;

    for (const group of ['componentSchemas', 'inlineSchemas']) {
      const arr = idx[group];
      if (!Array.isArray(arr)) continue;
      for (const n of arr) {
        let s: any = n && n.node;
        if (s && typeof s.getSchema === 'function') s = s.getSchema();   // ← indexed nodes are JSONSchema wrappers
        if (!s || typeof s.getNot !== 'function') continue;
        const not = s.getNot();
        if (not !== null && not !== undefined) {
          errors.push(createValidationError(
            config.getSeverity(this.defaultSeverity()),
            this.id(),
            'schema uses `not`, which the SDK generator does not support - every operation reaching this schema will be silently dropped from the SDK',
            typeof s.getRootNode === 'function' ? s.getRootNode() : null
          ));
        }
      }
    }
    return errors;
  }
}
registerRule(new NoNotSchemas());
