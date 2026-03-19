#!/bin/bash
# Redact or restore sensitive values in examples/.
#
# Usage:
#   ./scripts/redact_examples.sh redact    # replace real values with placeholders
#   ./scripts/redact_examples.sh inject    # restore real values from .cache/sensitive
#
# 'redact' auto-detects current real values from the examples, saves them to
# .cache/sensitive (gitignored), then replaces them with placeholders.
# 'inject' reads .cache/sensitive and puts the real values back.
#
# Sensitive values handled:
#   PROJECT_ID   projectId fields
#   ORG_ID       organizationId / parentContainerId fields
#   SA_EMAIL     subject / ownerEmail fields (*.sa.stackit.cloud)
#   DNS_ZONE     dnsName / contactEmail domain
#   SSH_KEY      publicKey field

set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
EXAMPLES_DIR="$(cd "$SCRIPT_DIR/../examples" && pwd)"
SENSITIVE_FILE="$(cd "$SCRIPT_DIR/.." && pwd)/.cache/sensitive"

PH_PROJECT_ID="<PROJECT-ID>"
PH_ORG_ID="<ORGANIZATION-ID>"
PH_SA_EMAIL="<SA-EMAIL>"
PH_DNS_ZONE="<DNS-ZONE>"
PH_SSH_KEY="<SSH-PUBLIC-KEY>"

redact() {
    # Auto-detect real values from examples (ignore placeholders)
    PROJECT_ID=$(grep -rh "projectId:" "$EXAMPLES_DIR" --include="*.yaml" \
        | awk '{print $2}' | grep -v '<' | sort -u | head -1)
    ORG_ID=$(grep -rh -E "organizationId:|parentContainerId:" "$EXAMPLES_DIR" --include="*.yaml" \
        | awk '{print $2}' | grep -E '^[0-9a-f-]{36}$' | grep -v '<' | sort -u | head -1)
    SA_EMAIL=$(grep -rh -E "subject:|ownerEmail:" "$EXAMPLES_DIR" --include="*.yaml" \
        | awk '{print $2}' | grep "@" | grep -v '<' | sort -u | head -1)
    DNS_ZONE=$(grep -rh "dnsName:" "$EXAMPLES_DIR" --include="*.yaml" \
        | awk '{print $2}' | grep -v '<' | sort -u | head -1)
    SSH_KEY=$(grep -rh "publicKey:" "$EXAMPLES_DIR" --include="*.yaml" \
        | sed 's/.*publicKey: *"\?//' | sed 's/" *$//' | grep "^ssh-" | sort -u | head -1)

    if [[ -z "$PROJECT_ID" ]]; then
        echo "Nothing to redact (already redacted or no examples found)."
        exit 0
    fi

    mkdir -p "$(dirname "$SENSITIVE_FILE")"
    printf 'PROJECT_ID="%s"\n' "$PROJECT_ID"  > "$SENSITIVE_FILE"
    printf 'ORG_ID="%s"\n'     "$ORG_ID"     >> "$SENSITIVE_FILE"
    printf 'SA_EMAIL="%s"\n'   "$SA_EMAIL"   >> "$SENSITIVE_FILE"
    printf 'DNS_ZONE="%s"\n'   "$DNS_ZONE"   >> "$SENSITIVE_FILE"
    printf 'SSH_KEY="%s"\n'    "$SSH_KEY"    >> "$SENSITIVE_FILE"
    echo "Saved sensitive values to $SENSITIVE_FILE"

    _replace_in_examples \
        "$PROJECT_ID" "$PH_PROJECT_ID" \
        "$ORG_ID"     "$PH_ORG_ID" \
        "$SA_EMAIL"   "$PH_SA_EMAIL" \
        "$DNS_ZONE"   "$PH_DNS_ZONE" \
        "$SSH_KEY"    "$PH_SSH_KEY"

    echo "Redacted examples in $EXAMPLES_DIR"
}

inject() {
    if [[ ! -f "$SENSITIVE_FILE" ]]; then
        echo "No saved values found at $SENSITIVE_FILE"
        echo "Run '$0 redact' first, or create $SENSITIVE_FILE manually."
        exit 1
    fi

    # shellcheck source=/dev/null
    source "$SENSITIVE_FILE"

    _replace_in_examples \
        "$PH_PROJECT_ID" "${PROJECT_ID:-}" \
        "$PH_ORG_ID"     "${ORG_ID:-}" \
        "$PH_SA_EMAIL"   "${SA_EMAIL:-}" \
        "$PH_DNS_ZONE"   "${DNS_ZONE:-}" \
        "$PH_SSH_KEY"    "${SSH_KEY:-}"

    echo "Injected real values into $EXAMPLES_DIR"
}

# _replace_in_examples FROM1 TO1 FROM2 TO2 ...
_replace_in_examples() {
    local args=("$@")
    while IFS= read -r -d '' file; do
        for ((i = 0; i < ${#args[@]}; i += 2)); do
            from="${args[$i]}"
            to="${args[$((i + 1))]}"
            [[ -z "$from" || -z "$to" || "$from" == "$to" ]] && continue
            sed -i "s|${from}|${to}|g" "$file"
        done
    done < <(find "$EXAMPLES_DIR" -type f \( -name "*.yaml" -o -name "*.yml" \) -print0)
}

case "${1:-}" in
    redact) redact ;;
    inject) inject ;;
    *)
        echo "Usage: $0 {redact|inject}"
        echo "  redact  Replace real values with placeholders (safe to commit)"
        echo "  inject  Restore real values from .cache/sensitive (for testing)"
        exit 1
        ;;
esac
