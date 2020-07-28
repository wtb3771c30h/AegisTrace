# Architecture

## Execution pipeline

```text
YAML source -> strict parser -> normalizer -> verifier -> sealed artifact
                                                        |
                                                        v
input -> runtime -> provider boundary -> external model/tool
             |              |
             v              v
       checkpoint       hash-chained journal -> replay / proof
```

The compiler parses one YAML document with unknown-field rejection. Normalization supplies deterministic defaults and sequential edges. The verifier admits only finite, reachable graphs with valid budgets and explicit tool permissions. The artifact stores the normalized workflow and its SHA-256 digest.

